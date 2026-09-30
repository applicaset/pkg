package serve

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/SladkyCitron/slogcolor"
	"github.com/applicaset/buildset/pkg/config"
	"github.com/applicaset/buildset/pkg/reqid"
)

// NewLogger builds the logger every binary uses.
func NewLogger(cfg config.Log) *slog.Logger {
	opts := &slog.HandlerOptions{Level: cfg.Level}

	var handler slog.Handler
	if cfg.JSON {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slogcolor.NewHandler(os.Stdout, &slogcolor.Options{
			Level:      cfg.Level,
			TimeFormat: time.RFC3339,
			// Colour escape codes are noise in a file or a log collector.
			NoColor: !isTerminal(os.Stdout),
		})
	}

	return slog.New(requestIDHandler{Handler: handler})
}

// requestIDHandler adds the context's request identifier to every record. Services never touch the
// field, yet all lines about one request share it.
type requestIDHandler struct {
	slog.Handler
}

func (h requestIDHandler) Handle(ctx context.Context, record slog.Record) error {
	if id, ok := reqid.FromContext(ctx); ok {
		record.AddAttrs(slog.String("request_id", id))
	}

	return h.Handler.Handle(ctx, record)
}

func (h requestIDHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return requestIDHandler{Handler: h.Handler.WithAttrs(attrs)}
}

func (h requestIDHandler) WithGroup(name string) slog.Handler {
	return requestIDHandler{Handler: h.Handler.WithGroup(name)}
}

func isTerminal(file *os.File) bool {
	info, err := file.Stat()
	if err != nil {
		return false
	}

	return info.Mode()&os.ModeCharDevice != 0
}
