# Testes do Loadtest Go

O objetivo deste documento é orientar como validar o projeto tanto via testes automatizados quanto via execuções manuais.

## Pré-requisitos

- Go 1.21 (conforme `go.mod`)
- `git` e acesso à internet para baixar dependências
- Opcional: Docker (para validar a imagem gerada pelo `Dockerfile`)

## Testes automatizados

1. A partir do diretório raiz, execute a suíte de testes:

   ```bash
   go test ./...
   ```

   Isso roda os testes unitários definidos em `internal/usecase/loadtest_test.go`, que usam um cliente HTTP simulado para validar métricas de carga e limites de concorrência.

2. Para gerar um relatório de cobertura:

   ```bash
   go test ./... -coverprofile=coverage.out
   go tool cover -func=coverage.out
   ```

   O arquivo `coverage.out` pode ser revisado ou transformado em um relatório HTML (`go tool cover -html=coverage.out`).

## Testes manuais / Smoke tests

1. Compile o binário:

   ```bash
   go build -o loadtest ./cmd/loadtest
   ```

2. Execute o binário com a URL e os parâmetros desejados:

   ```bash
   ./loadtest --url=https://example.com --requests=50 --concurrency=5
   ```

3. Para validar o fluxo apresentado no README, também é possível usar o script de exemplos:

   ```bash
   chmod +x examples.sh
   ./examples.sh
   ```

   Esse script mostra comandos `docker run` preparados. Para executá-los, primeiro crie a imagem:

   ```bash
   docker build -t loadtest .
   docker run --rm loadtest --url=https://example.com --requests=20 --concurrency=2
   ```

## Verificações adicionais

- Garanta que o `GOOS` e o `GOPATH` estejam configurados conforme o padrão da sua máquina (`go env`).
- Sempre que modificar lógica de negócios, execute novamente `go test ./...` antes de commitar.

Com esses passos você deve conseguir validar o comportamento esperado do projeto.

