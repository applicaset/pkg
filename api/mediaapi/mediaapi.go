// Package mediaapi is the wire shape of the media service: types and paths only, so the service,
// its client and its callers share one definition without importing each other.
package mediaapi

import "time"

// Only sibling services on the private network reach these paths. The service trusts its caller:
// it stores bytes and never decides who may read them.
const (
	// PathUpload takes the file's bytes as the raw request body, typed by its Content-Type header,
	// and answers with a FileResponse.
	PathUpload = "/v1/files"
	// PathDownload is PathUpload plus "/{id}". It answers GET with the bytes and honours Range and
	// If-None-Match, so a caller can pass a browser's request straight through.
	PathDownload = "/v1/files/"
	PathStat     = "/v1/stat"
	PathDelete   = "/v1/delete"
)

type File struct {
	ID          string    `json:"id"`
	SHA256      string    `json:"sha256"`
	Size        int64     `json:"size"`
	ContentType string    `json:"content_type"`
	CreatedAt   time.Time `json:"created_at"`
}

type FileRequest struct {
	FileID string `json:"file_id"`
}

type FileResponse struct {
	File File `json:"file"`
}
