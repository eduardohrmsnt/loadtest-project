package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/eduardohermesneto/loadtest-go/internal/infrastructure"
	"github.com/eduardohermesneto/loadtest-go/internal/usecase"
)

func main() {
	var url string
	var requests int
	var concurrency int

	// Normalizar argumentos (substituir diferentes tipos de hífens Unicode por hífen duplo padrão)
	args := os.Args[1:]
	for i, arg := range args {
		// Substituir diferentes tipos de hífens Unicode por hífen duplo padrão (--)
		arg = strings.ReplaceAll(arg, "—", "--") // em-dash (U+2014)
		arg = strings.ReplaceAll(arg, "–", "--") // en-dash (U+2013)
		arg = strings.ReplaceAll(arg, "―", "--") // horizontal bar (U+2015)
		args[i] = arg
	}
	os.Args = append([]string{os.Args[0]}, args...)

	flag.StringVar(&url, "url", "", "URL para testar (obrigatório)")
	flag.StringVar(&url, "u", "", "URL para testar (obrigatório, forma curta)")
	flag.IntVar(&requests, "requests", 100, "Número total de requests")
	flag.IntVar(&requests, "r", 100, "Número total de requests (forma curta)")
	flag.IntVar(&concurrency, "concurrency", 10, "Número de workers concorrentes")
	flag.IntVar(&concurrency, "c", 10, "Número de workers concorrentes (forma curta)")
	flag.Parse()

	if url == "" {
		fmt.Println("Erro: --url é obrigatório")
		flag.Usage()
		os.Exit(1)
	}

	if requests <= 0 {
		fmt.Println("Erro: --requests deve ser maior que 0")
		os.Exit(1)
	}

	if concurrency <= 0 {
		fmt.Println("Erro: --concurrency deve ser maior que 0")
		os.Exit(1)
	}

	// Inicializar dependências
	httpClient := infrastructure.NewHTTPClient(30 * time.Second)
	loadTester := usecase.NewLoadTestUseCase(httpClient)
	reporter := infrastructure.NewConsoleReporter()

	// Executar teste de carga
	fmt.Printf("Iniciando teste de carga...\n")
	fmt.Printf("URL: %s\n", url)
	fmt.Printf("Total de requests: %d\n", requests)
	fmt.Printf("Concorrência: %d\n\n", concurrency)

	result := loadTester.Execute(url, requests, concurrency)

	// Gerar relatório
	reporter.Generate(result)
}
