package serve

import (
	"log/slog"
	"net/http"
	"time"
	"uuid"

	"github.com/applicaset/buildset/pkg/reqid"
)

// WithRequestID honours an inbound identifier only when trustInbound is set. Set it for an internal
// API a sibling calls, never for anything a browser reaches: any client could pick its own.
func WithRequestID(trustInbound bool, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := ""
		if trustInbound {
			id = r.Header.Get(reqid.Header)
		}

		if !validRequestID(id) {
			id = uuid.NewV7().String()
		}

		w.Header().Set(reqid.Header, id)

		next.ServeHTTP(w, r.WithContext(reqid.NewContext(r.Context(), id)))
	})
}

// validRequestID keeps a forwarded identifier from carrying anything that would corrupt a log line
// or a response header.
func validRequestID(id string) bool {
	if len(id) == 0 || len(id) > 64 {
		return false
	}

	for _, r := range id {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || r == '-' ||
			r == '_' {
			continue
		}

		return false
	}

	return true
}

// RecoverPanics keeps one bad handler from taking the process down, and answers the client rather
// than dropping the connection.
func RecoverPanics(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		defer func() {
			recovered := recover()
			if recovered == nil {
				return
			}

			logger.ErrorContext(ctx, "handler panicked",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Any("panic", recovered),
			)

			// If the handler already started writing, this is a no-op; writing again would corrupt
			// the response.
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}()

		next.ServeHTTP(w, r)
	})
}

// LogRequests logs method, path, status and duration only. Identity service request bodies carry
// live session tokens, so logging bodies or headers leaks credentials. Check before adding a field.
func LogRequests(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		recorder := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(recorder, r)

		logger.InfoContext(r.Context(), "request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int("status", recorder.status),
			slog.Duration("duration", time.Since(start)),
		)
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func (s *statusRecorder) WriteHeader(status int) {
	if s.wroteHeader {
		return
	}

	s.status = status
	s.wroteHeader = true
	s.ResponseWriter.WriteHeader(status)
}

// Unwrap lets http.ResponseController reach the connection beneath, so a streaming handler can
// flush and lift its write deadline.
func (s *statusRecorder) Unwrap() http.ResponseWriter { return s.ResponseWriter }
