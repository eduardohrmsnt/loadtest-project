package usecase

import (
	"sync"
	"time"

	"github.com/eduardohermesneto/loadtest-go/internal/domain"
)

type loadTestUseCase struct {
	client domain.HTTPClient
}

func NewLoadTestUseCase(client domain.HTTPClient) domain.LoadTester {
	return &loadTestUseCase{
		client: client,
	}
}

func (l *loadTestUseCase) Execute(url string, totalRequests, concurrency int) domain.TestResult {
	startTime := time.Now()

	results := make(chan domain.Response, totalRequests)
	requests := make(chan domain.Request, totalRequests)

	var wg sync.WaitGroup

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go l.worker(&wg, requests, results)
	}

	go func() {
		for i := 0; i < totalRequests; i++ {
			requests <- domain.Request{
				URL:    url,
				Method: "GET",
			}
		}
		close(requests)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	return l.aggregateResults(results, totalRequests, time.Since(startTime))
}

func (l *loadTestUseCase) worker(wg *sync.WaitGroup, requests <-chan domain.Request, results chan<- domain.Response) {
	defer wg.Done()

	for req := range requests {
		response := l.client.Do(req)
		results <- response
	}
}

func (l *loadTestUseCase) aggregateResults(results <-chan domain.Response, totalRequests int, duration time.Duration) domain.TestResult {
	statusCodes := make(map[int]int)
	successCount := 0
	errorCount := 0

	for response := range results {
		if response.Error != nil {
			errorCount++
			continue
		}

		statusCodes[response.StatusCode]++

		if response.StatusCode == 200 {
			successCount++
		}
	}

	return domain.TestResult{
		TotalRequests:   totalRequests,
		SuccessRequests: successCount,
		StatusCodes:     statusCodes,
		TotalDuration:   duration,
		Errors:          errorCount,
	}
}
