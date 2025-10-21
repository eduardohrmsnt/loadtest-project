package infrastructure

import (
	"net/http"
	"time"

	"github.com/eduardohermesneto/loadtest-go/internal/domain"
)

type httpClient struct {
	client *http.Client
}

func NewHTTPClient(timeout time.Duration) domain.HTTPClient {
	return &httpClient{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (h *httpClient) Do(request domain.Request) domain.Response {
	startTime := time.Now()

	req, err := http.NewRequest(request.Method, request.URL, nil)
	if err != nil {
		return domain.Response{
			Error:    err,
			Duration: time.Since(startTime),
		}
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return domain.Response{
			Error:    err,
			Duration: time.Since(startTime),
		}
	}
	defer resp.Body.Close()

	return domain.Response{
		StatusCode: resp.StatusCode,
		Duration:   time.Since(startTime),
		Error:      nil,
	}
}
