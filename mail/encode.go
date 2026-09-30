package mail

import (
	"crypto/tls"
	"mime"
)

func mimeQEncode(value string) string {
	return mime.QEncoding.Encode("utf-8", value)
}

func tlsConfig(host string) *tls.Config {
	return &tls.Config{ServerName: host, MinVersion: tls.VersionTLS12}
}
