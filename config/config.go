// Package config holds settings blocks shared by more than one binary, so every binary reads PORT
// or DATABASE_DSN the same way. A binary embeds the blocks it needs and adds its own fields.
package config

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/nasermirzaei89/env"
)

// SessionCookieName is fixed rather than configurable, so the service that sets the cookie and the
// ones that read it cannot disagree.
const SessionCookieName = "ms_session"

var errInvalidConfig = errors.New("invalid configuration")

// Server keeps Port as the raw environment string. net.Listen takes a string, and PORT may name a
// service ("http") as well as a number.
type Server struct {
	Host            string
	Port            string
	ShutdownTimeout time.Duration
}

func LoadServer() Server {
	return Server{
		Host:            env.GetString("HOST", ""),
		Port:            env.GetString("PORT", "8080"),
		ShutdownTimeout: env.GetDuration("SHUTDOWN_TIMEOUT", 15*time.Second),
	}
}

// Address is what the HTTP server listens on. An empty host means every interface.
func (s Server) Address() string {
	return net.JoinHostPort(s.Host, s.Port)
}

// ResolvePort resolves a service name in PORT to a number. A URL authority may not contain a
// service name, so anything building one needs this rather than Port.
func (s Server) ResolvePort(ctx context.Context) (string, error) {
	port, err := net.DefaultResolver.LookupPort(ctx, "tcp", s.Port)
	if err != nil {
		return "", fmt.Errorf("resolve PORT %q: %w", s.Port, err)
	}

	return strconv.Itoa(port), nil
}

func (s Server) Validate(ctx context.Context) error {
	if _, err := s.ResolvePort(ctx); err != nil {
		return err
	}

	if s.ShutdownTimeout <= 0 {
		return fmt.Errorf(
			"%w: SHUTDOWN_TIMEOUT must be positive, got %s",
			errInvalidConfig,
			s.ShutdownTimeout,
		)
	}

	return nil
}

type Log struct {
	Level slog.Level
	JSON  bool
}

func LoadLog() (Log, error) {
	var level slog.Level
	if err := level.UnmarshalText([]byte(env.GetString("LOG_LEVEL", "info"))); err != nil {
		return Log{}, fmt.Errorf("parse LOG_LEVEL: %w", err)
	}

	return Log{Level: level, JSON: env.GetBool("LOG_JSON", false)}, nil
}

// Cookie is shared by the binary that sets the session cookie and the ones that read it. They must
// agree, or a signed-in visitor looks anonymous to half the system.
type Cookie struct {
	Name   string
	Secure bool
}

func LoadCookie() Cookie {
	return Cookie{
		Name:   SessionCookieName,
		Secure: env.GetBool("SECURE_COOKIES", true),
	}
}

// LoadURL reads an http or https URL from key. A missing or malformed value fails at start-up, not
// as a 500 on a visitor's first click.
func LoadURL(key string) (string, error) {
	raw := env.GetString(key, "")
	if raw == "" {
		return "", fmt.Errorf("%w: %s must be set", errInvalidConfig, key)
	}

	parsed, err := url.Parse(raw)
	if err != nil {
		return "", fmt.Errorf("parse %s: %w", key, err)
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return "", fmt.Errorf(
			"%w: %s must be an http or https URL, got %q",
			errInvalidConfig,
			key,
			raw,
		)
	}

	if parsed.Host == "" {
		return "", fmt.Errorf("%w: %s must include a host, got %q", errInvalidConfig, key, raw)
	}

	// A trailing slash would double up when a path is appended.
	return strings.TrimRight(raw, "/"), nil
}

// LoadBasePath reads BASE_PATH, where a gateway mounts a site: empty for the root, or a path such as
// "/blog". A trailing slash would double up when a route is appended.
func LoadBasePath() (string, error) {
	raw := env.GetString("BASE_PATH", "")
	if raw == "" {
		return "", nil
	}

	if !strings.HasPrefix(raw, "/") || strings.HasSuffix(raw, "/") ||
		strings.ContainsAny(raw, "?#") {
		return "", fmt.Errorf(
			"%w: BASE_PATH must start with / and not end with one, got %q",
			errInvalidConfig,
			raw,
		)
	}

	return raw, nil
}
