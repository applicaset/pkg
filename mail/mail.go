// Package mail sends plain-text email. Every service that mails a person goes through Sender, so
// tests swap in a Recorder and development without a mail server swaps in a Logger.
package mail

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/mail"
	"net/smtp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/nasermirzaei89/env"
)

var (
	errInvalidConfig  = errors.New("invalid configuration")
	errInvalidMessage = errors.New("invalid message")
)

type Message struct {
	To      string
	Subject string
	Text    string
}

type Sender interface {
	Send(ctx context.Context, message Message) error
}

type Config struct {
	// Host is the SMTP server. Empty sends nothing and logs each message instead.
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

func Load() Config {
	return Config{
		Host:     env.GetString("SMTP_HOST", ""),
		Port:     env.GetInt("SMTP_PORT", 587),
		Username: env.GetString("SMTP_USERNAME", ""),
		Password: env.GetString("SMTP_PASSWORD", ""),
		From:     env.GetString("MAIL_FROM", "Buildset <no-reply@localhost>"),
	}
}

func (c Config) Validate() error {
	if _, err := mail.ParseAddress(c.From); err != nil {
		return fmt.Errorf("%w: MAIL_FROM must be an email address: %w", errInvalidConfig, err)
	}

	if c.Host == "" {
		return nil
	}

	if c.Port < 1 || c.Port > 65535 {
		return fmt.Errorf("%w: SMTP_PORT must be a port number, got %d", errInvalidConfig, c.Port)
	}

	if c.Password != "" && c.Username == "" {
		return fmt.Errorf("%w: SMTP_PASSWORD needs SMTP_USERNAME", errInvalidConfig)
	}

	return nil
}

// New returns the SMTP sender, or a Logger when no host is configured.
func New(cfg Config, logger *slog.Logger) Sender {
	if cfg.Host == "" {
		return Logger{Logger: logger}
	}

	return &SMTP{config: cfg}
}

type SMTP struct {
	config Config
}

// sendTimeout bounds one delivery, so a hanging server cannot hold a request open.
const sendTimeout = 15 * time.Second

// Send upgrades to TLS whenever the server offers STARTTLS. net/smtp refuses to send credentials
// over a plain connection to anything but localhost, so a server without STARTTLS gets none.
func (s *SMTP) Send(ctx context.Context, message Message) error {
	body, err := compose(s.config.From, message)
	if err != nil {
		return err
	}

	from, err := mail.ParseAddress(s.config.From)
	if err != nil {
		return fmt.Errorf("parse sender: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, sendTimeout)
	defer cancel()

	address := net.JoinHostPort(s.config.Host, strconv.Itoa(s.config.Port))

	var dialer net.Dialer

	conn, err := dialer.DialContext(ctx, "tcp", address)
	if err != nil {
		return fmt.Errorf("dial smtp server %s: %w", address, err)
	}

	if deadline, ok := ctx.Deadline(); ok {
		_ = conn.SetDeadline(deadline)
	}

	client, err := smtp.NewClient(conn, s.config.Host)
	if err != nil {
		_ = conn.Close()

		return fmt.Errorf("start smtp session: %w", err)
	}
	defer func() { _ = client.Close() }()

	if ok, _ := client.Extension("STARTTLS"); ok {
		if err := client.StartTLS(tlsConfig(s.config.Host)); err != nil {
			return fmt.Errorf("start tls: %w", err)
		}
	}

	if s.config.Username != "" {
		auth := smtp.PlainAuth("", s.config.Username, s.config.Password, s.config.Host)
		if err := client.Auth(auth); err != nil {
			return fmt.Errorf("authenticate to smtp server: %w", err)
		}
	}

	if err := client.Mail(from.Address); err != nil {
		return fmt.Errorf("smtp MAIL FROM: %w", err)
	}

	if err := client.Rcpt(message.To); err != nil {
		return fmt.Errorf("smtp RCPT TO: %w", err)
	}

	writer, err := client.Data()
	if err != nil {
		return fmt.Errorf("smtp DATA: %w", err)
	}

	if _, err := writer.Write(body); err != nil {
		_ = writer.Close()

		return fmt.Errorf("write message: %w", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("finish message: %w", err)
	}

	if err := client.Quit(); err != nil {
		return fmt.Errorf("smtp QUIT: %w", err)
	}

	return nil
}

// compose builds an RFC 5322 message. The subject is Q-encoded and header values are checked for
// line breaks, so a caller's string can never add a header.
func compose(from string, message Message) ([]byte, error) {
	if _, err := mail.ParseAddress(message.To); err != nil {
		return nil, fmt.Errorf("%w: recipient: %w", errInvalidMessage, err)
	}

	for _, value := range []string{message.To, message.Subject} {
		if strings.ContainsAny(value, "\r\n") {
			return nil, fmt.Errorf("%w: header value contains a line break", errInvalidMessage)
		}
	}

	var builder strings.Builder

	builder.WriteString("From: " + from + "\r\n")
	builder.WriteString("To: " + message.To + "\r\n")
	builder.WriteString("Subject: " + mimeQEncode(message.Subject) + "\r\n")
	builder.WriteString("Date: " + time.Now().UTC().Format(time.RFC1123Z) + "\r\n")
	builder.WriteString("MIME-Version: 1.0\r\n")
	builder.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
	builder.WriteString("Content-Transfer-Encoding: 8bit\r\n")
	builder.WriteString("\r\n")
	builder.WriteString(strings.ReplaceAll(normalizeNewlines(message.Text), "\n", "\r\n"))

	return []byte(builder.String()), nil
}

func normalizeNewlines(text string) string {
	return strings.ReplaceAll(text, "\r\n", "\n")
}

// Logger writes each message to the log instead of sending it, for development without a mail
// server. The body holds live links, so it must never be used where logs are shared.
type Logger struct {
	Logger *slog.Logger
}

func (l Logger) Send(ctx context.Context, message Message) error {
	l.Logger.WarnContext(ctx, "mail not sent: SMTP_HOST is empty",
		slog.String("to", message.To),
		slog.String("subject", message.Subject),
		slog.String("text", message.Text),
	)

	return nil
}

// Recorder keeps every message for a test to read back.
type Recorder struct {
	mu       sync.Mutex
	messages []Message
}

func (r *Recorder) Send(_ context.Context, message Message) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.messages = append(r.messages, message)

	return nil
}

func (r *Recorder) Messages() []Message {
	r.mu.Lock()
	defer r.mu.Unlock()

	return append([]Message(nil), r.messages...)
}

// Last returns the latest message sent to address, and false when there is none.
func (r *Recorder) Last(address string) (Message, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, v := range slices.Backward(r.messages) {
		if strings.EqualFold(v.To, address) {
			return v, true
		}
	}

	return Message{}, false
}
