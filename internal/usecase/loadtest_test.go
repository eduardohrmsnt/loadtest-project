package usecase

import (
	"testing"
	"time"

	"github.com/loadtest/internal/domain"
)

type mockHTTPClient struct {
	statusCode int
	shouldFail bool
}

func (m *mockHTTPClient) Do(request domain.Request) domain.Response {
	if m.shouldFail {
		return domain.Response{
			Error:    nil,
			Duration: time.Millisecond,
		}
	}

	return domain.Response{
		StatusCode: m.statusCode,
		Duration:   time.Millisecond,
		Error:      nil,
	}
}

func TestLoadTestUseCase_Execute(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		totalRequests  int
		concurrency    int
		statusCode     int
		expectedStatus int
		shouldFail     bool
	}{
		{
			name:           "successful requests with status 200",
			url:            "http://example.com",
			totalRequests:  10,
			concurrency:    2,
			statusCode:     200,
			expectedStatus: 200,
			shouldFail:     false,
		},
		{
			name:           "requests with status 404",
			url:            "http://example.com",
			totalRequests:  5,
			concurrency:    1,
			statusCode:     404,
			expectedStatus: 404,
			shouldFail:     false,
		},
		{
			name:           "high concurrency test",
			url:            "http://example.com",
			totalRequests:  100,
			concurrency:    10,
			statusCode:     200,
			expectedStatus: 200,
			shouldFail:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &mockHTTPClient{
				statusCode: tt.statusCode,
				shouldFail: tt.shouldFail,
			}

			useCase := NewLoadTestUseCase(mockClient)
			result := useCase.Execute(tt.url, tt.totalRequests, tt.concurrency)

			if result.TotalRequests != tt.totalRequests {
				t.Errorf("Expected %d total requests, got %d", tt.totalRequests, result.TotalRequests)
			}

			if tt.expectedStatus == 200 && result.SuccessRequests != tt.totalRequests {
				t.Errorf("Expected %d successful requests, got %d", tt.totalRequests, result.SuccessRequests)
			}

			if _, exists := result.StatusCodes[tt.statusCode]; !exists && !tt.shouldFail {
				t.Errorf("Expected status code %d to be present in results", tt.statusCode)
			}
		})
	}
}

func TestLoadTestUseCase_ConcurrencyLimit(t *testing.T) {
	mockClient := &mockHTTPClient{
		statusCode: 200,
		shouldFail: false,
	}

	useCase := NewLoadTestUseCase(mockClient)
	result := useCase.Execute("http://example.com", 50, 5)

	if result.TotalRequests != 50 {
		t.Errorf("Expected 50 total requests, got %d", result.TotalRequests)
	}

	if result.SuccessRequests != 50 {
		t.Errorf("Expected 50 successful requests, got %d", result.SuccessRequests)
	}
}
