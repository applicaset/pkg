package httpx

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

// maxRequestBytes bounds a request body, so a caller cannot use one to exhaust a service.
const maxRequestBytes = 1 << 20

// DecodeJSON answers invalid_input and reports false when the body cannot be read. The handler
// must then return.
func DecodeJSON(w http.ResponseWriter, r *http.Request, target any) bool {
	decoder := json.NewDecoder(io.LimitReader(r.Body, maxRequestBytes))
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(target); err != nil {
		// Never echo the error. A decode error can quote the body, and some bodies carry a live
		// session token.
		WriteError(w, CodeInvalidInput, "That request could not be read.")

		return false
	}

	return true
}

func WriteJSON(w http.ResponseWriter, value any) {
	body, err := json.Marshal(value)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)
}

// WriteError answers with an Envelope. The site shows message to the visitor, so write it for them.
func WriteError(w http.ResponseWriter, code Code, message string) {
	body, err := json.Marshal(Envelope{Code: code, Message: message})
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(code.Status())
	_, _ = w.Write(body)
}

// WriteInternal logs err and sends a bare 500, a failure the client cannot mistake for an answer.
func WriteInternal(
	ctx context.Context,
	w http.ResponseWriter,
	logger *slog.Logger,
	operation string,
	err error,
) {
	logger.ErrorContext(ctx, operation, slog.Any("error", err))

	http.Error(w, operation+" failed", http.StatusInternalServerError)
}
