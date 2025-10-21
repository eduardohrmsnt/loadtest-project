package infrastructure

import (
	"fmt"
	"sort"

	"github.com/eduardohermesneto/loadtest-go/internal/domain"
)

type consoleReporter struct{}

func NewConsoleReporter() domain.Reporter {
	return &consoleReporter{}
}

func (c *consoleReporter) Generate(result domain.TestResult) {
	fmt.Println("\n========================================")
	fmt.Println("         RELATÓRIO DE TESTES")
	fmt.Println("========================================")
	fmt.Printf("Tempo total de execução: %v\n", result.TotalDuration)
	fmt.Printf("Total de requests realizados: %d\n", result.TotalRequests)
	fmt.Printf("Requests com status 200: %d\n", result.SuccessRequests)

	if result.Errors > 0 {
		fmt.Printf("Requests com erro: %d\n", result.Errors)
	}

	fmt.Println("\nDistribuição de status HTTP:")

	statusCodes := make([]int, 0, len(result.StatusCodes))
	for code := range result.StatusCodes {
		statusCodes = append(statusCodes, code)
	}
	sort.Ints(statusCodes)

	for _, code := range statusCodes {
		count := result.StatusCodes[code]
		percentage := float64(count) / float64(result.TotalRequests) * 100
		fmt.Printf("  Status %d: %d (%.2f%%)\n", code, count, percentage)
	}

	fmt.Println("========================================")
}
