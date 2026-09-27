// Package asset serves an embedded static file under a URL that changes with its content, so
// browsers can cache it forever.
package asset

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
)

// versionLength is how many hex digits of the content hash go into the URL.
const versionLength = 8

type Asset struct {
	// Path is the route the file is served at, without the version.
	Path string
	// URL is Path with the content version appended. Pages link to this.
	URL         string
	contentType string
	content     []byte
}

func New(path, contentType string, content []byte) *Asset {
	sum := sha256.Sum256(content)

	return &Asset{
		Path:        path,
		URL:         path + "?v=" + hex.EncodeToString(sum[:])[:versionLength],
		contentType: contentType,
		content:     content,
	}
}

func (a *Asset) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", a.contentType)
	w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
	_, _ = w.Write(a.content)
}
