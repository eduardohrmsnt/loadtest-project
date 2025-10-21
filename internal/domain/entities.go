package domain

import "time"

type Request struct {
	URL    string
	Method string
}

type Response struct {
	StatusCode int
	Duration   time.Duration
	Error      error
}

type TestResult struct {
	TotalRequests   int
	SuccessRequests int
	StatusCodes     map[int]int
	TotalDuration   time.Duration
	Errors          int
}
