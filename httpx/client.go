package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"net"
	"net/http"
	"time"

	"github.com/buildset/buildset/pkg/reqid"
)

const (
	// maxErrorBytes caps how much of a failed response is read. A huge error body from a dependency
	// cannot exhaust this process.
	maxErrorBytes = 8 << 10
	// maxDrainBytes is how much of a body is read back before closing, so the connection returns
	// to the pool instead of being dropped.
	maxDrainBytes = 4 << 10

	defaultTimeout             = 5 * time.Second
	defaultMaxIdleConnsPerHost = 32
)

var (
	errInvalidArgument  = errors.New("invalid argument")
	errUnexpectedAnswer = errors.New("unexpected answer")
)

type ClientOptions struct {
	// Timeout bounds one call. The 5s default sits under the 30s write timeout of a page-serving
	// process, leaving time to render an error page.
	Timeout time.Duration
	// MaxIdleConnsPerHost bounds the pool this client keeps to its one dependency.
	MaxIdleConnsPerHost int
}

// Client calls one service and never retries. Most calls are writes during a browser request, and
// a retry multiplies load on a dependency that is already slow. A caller retries at its own site.
type Client struct {
	service string
	baseURL string
	http    *http.Client
}

func NewClient(service, baseURL string, opts ClientOptions) (*Client, error) {
	if service == "" {
		return nil, fmt.Errorf("%w: service name must not be empty", errInvalidArgument)
	}

	if baseURL == "" {
		return nil, fmt.Errorf("%w: base URL of %s must not be empty", errInvalidArgument, service)
	}

	timeout := opts.Timeout
	if timeout <= 0 {
		timeout = defaultTimeout
	}

	idle := opts.MaxIdleConnsPerHost
	if idle <= 0 {
		idle = defaultMaxIdleConnsPerHost
	}

	transport := &http.Transport{
		DialContext:         (&net.Dialer{Timeout: 2 * time.Second}).DialContext,
		MaxIdleConns:        idle,
		MaxIdleConnsPerHost: idle,
		IdleConnTimeout:     90 * time.Second,
		// Plain HTTP/1.1 on a private network. Multiplexing gains nothing here and makes failures
		// harder to reason about.
		ForceAttemptHTTP2: false,
	}

	return &Client{
		service: service,
		baseURL: baseURL,
		http:    &http.Client{Timeout: timeout, Transport: transport},
	}, nil
}

// Call POSTs request as JSON to path and decodes the answer into response, which may be nil.
// A domain failure returns *Error. Every other failure is a plain wrapped error, never an *Error.
func (c *Client) Call(ctx context.Context, path string, request, response any) error {
	body, err := json.Marshal(request)
	if err != nil {
		return c.wrap(path, fmt.Errorf("encode request: %w", err))
	}

	httpRequest, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.baseURL+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return c.wrap(path, fmt.Errorf("build request: %w", err))
	}

	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Accept", "application/json")

	if id, ok := reqid.FromContext(ctx); ok {
		httpRequest.Header.Set(reqid.Header, id)
	}

	httpResponse, err := c.http.Do(httpRequest)
	if err != nil {
		return c.wrap(path, err)
	}
	defer func() {
		_, _ = io.Copy(io.Discard, io.LimitReader(httpResponse.Body, maxDrainBytes))
		_ = httpResponse.Body.Close()
	}()

	if httpResponse.StatusCode != http.StatusOK {
		return c.wrap(path, decodeError(httpResponse))
	}

	if response == nil {
		return nil
	}

	if err := json.NewDecoder(httpResponse.Body).Decode(response); err != nil {
		return c.wrap(path, fmt.Errorf("decode response: %w", err))
	}

	return nil
}

// The request body stays out of the error: an identity service request carries a session token.
func (c *Client) wrap(path string, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s %s: %w", c.service, path, err)
}

// Only a well-formed domain failure becomes *Error. Anything else stays an ordinary error, so "the
// identity service is down" is never read as "no such session".
func decodeError(response *http.Response) error {
	switch response.StatusCode {
	case http.StatusBadRequest, http.StatusForbidden, http.StatusNotFound, http.StatusConflict:
	default:
		return fmt.Errorf("%w: status %d", errUnexpectedAnswer, response.StatusCode)
	}

	mediaType, _, err := mime.ParseMediaType(response.Header.Get("Content-Type"))
	if err != nil || mediaType != "application/json" {
		return fmt.Errorf(
			"%w: status %d with content type %q",
			errUnexpectedAnswer,
			response.StatusCode,
			response.Header.Get("Content-Type"),
		)
	}

	var envelope Envelope
	if err := json.NewDecoder(io.LimitReader(response.Body, maxErrorBytes)).
		Decode(&envelope); err != nil {
		return fmt.Errorf("status %d with undecodable body: %w", response.StatusCode, err)
	}

	if !envelope.Code.Known() {
		return fmt.Errorf(
			"%w: status %d with unknown code %q",
			errUnexpectedAnswer,
			response.StatusCode,
			envelope.Code,
		)
	}

	return &Error{Code: envelope.Code, Message: envelope.Message}
}
