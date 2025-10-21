package domain

type HTTPClient interface {
	Do(request Request) Response
}

type LoadTester interface {
	Execute(url string, totalRequests, concurrency int) TestResult
}

type Reporter interface {
	Generate(result TestResult)
}
