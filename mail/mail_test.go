package mail

import (
	"bufio"
	"context"
	"net"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComposeRejectsHeaderInjection(t *testing.T) {
	_, err := compose("a@example.com", Message{
		To:      "b@example.com",
		Subject: "hi\r\nBcc: c@example.com",
	})
	require.ErrorIs(t, err, errInvalidMessage)
}

func TestComposeEncodesSubjectAndLineEndings(t *testing.T) {
	body, err := compose("a@example.com", Message{
		To:      "b@example.com",
		Subject: "Sign in to TodoSet ✓",
		Text:    "line one\nline two",
	})
	require.NoError(t, err)

	text := string(body)
	assert.Contains(t, text, "Subject: =?utf-8?q?")
	assert.Contains(t, text, "\r\n\r\nline one\r\nline two")
}

func TestConfigValidate(t *testing.T) {
	require.NoError(t, Config{From: "a@example.com"}.Validate())
	require.Error(t, Config{From: "nope"}.Validate())
	require.Error(t, Config{From: "a@example.com", Host: "smtp", Port: 0}.Validate())
	require.Error(
		t,
		Config{From: "a@example.com", Host: "smtp", Port: 25, Password: "x"}.Validate(),
	)
}

// TestSMTPSendsToAPlainServer talks to a minimal in-process SMTP server, the way Mailpit behaves
// in development: no STARTTLS and no authentication.
func TestSMTPSendsToAPlainServer(t *testing.T) {
	var config net.ListenConfig

	listener, err := config.Listen(t.Context(), "tcp", "127.0.0.1:0")
	require.NoError(t, err)

	t.Cleanup(func() { _ = listener.Close() })

	received := make(chan string, 1)

	go serveOneMessage(t, listener, received)

	host, port, err := net.SplitHostPort(listener.Addr().String())
	require.NoError(t, err)

	sender := &SMTP{config: Config{Host: host, Port: mustAtoi(t, port), From: "a@example.com"}}

	err = sender.Send(context.Background(), Message{
		To:      "b@example.com",
		Subject: "Hello",
		Text:    "Body",
	})
	require.NoError(t, err)

	assert.Contains(t, <-received, "Subject: Hello")
}

func TestRecorderLast(t *testing.T) {
	var recorder Recorder

	require.NoError(t, recorder.Send(context.Background(), Message{To: "a@example.com", Text: "1"}))
	require.NoError(t, recorder.Send(context.Background(), Message{To: "b@example.com", Text: "2"}))
	require.NoError(t, recorder.Send(context.Background(), Message{To: "A@example.com", Text: "3"}))

	last, ok := recorder.Last("a@example.com")
	require.True(t, ok)
	assert.Equal(t, "3", last.Text)

	_, ok = recorder.Last("c@example.com")
	assert.False(t, ok)
}

func serveOneMessage(t *testing.T, listener net.Listener, received chan<- string) {
	t.Helper()

	conn, err := listener.Accept()
	if err != nil {
		return
	}
	defer func() { _ = conn.Close() }()

	reader := bufio.NewReader(conn)
	write := func(line string) { _, _ = conn.Write([]byte(line + "\r\n")) }

	write("220 test ready")

	var data strings.Builder

	inData := false

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		if inData {
			if line == ".\r\n" {
				inData = false

				write("250 queued")

				received <- data.String()

				continue
			}

			data.WriteString(line)

			continue
		}

		switch command := strings.ToUpper(strings.TrimSpace(line)); {
		case strings.HasPrefix(command, "EHLO"), strings.HasPrefix(command, "HELO"):
			write("250 test")
		case strings.HasPrefix(command, "DATA"):
			inData = true

			write("354 go ahead")
		case strings.HasPrefix(command, "QUIT"):
			write("221 bye")

			return
		default:
			write("250 ok")
		}
	}
}

func mustAtoi(t *testing.T, value string) int {
	t.Helper()

	n, err := strconv.Atoi(value)
	require.NoError(t, err)

	return n
}
