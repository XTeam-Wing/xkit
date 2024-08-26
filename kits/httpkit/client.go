package httpkit

import (
	"net/http"
	"time"
)

// NewClient creates a new HTTP client with custom timeout settings.
func NewClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
	}
}
