#!/bin/bash

echo "Exemplos de uso do Load Test CLI"
echo "=================================="
echo ""

echo "1. Teste básico com 100 requests e 10 workers:"
echo "docker run loadtest --url=https://example.com --requests=100 --concurrency=10"
echo ""

echo "2. Teste de carga intensivo:"
echo "docker run loadtest --url=https://api.example.com/health --requests=10000 --concurrency=100"
echo ""

echo "3. Teste rápido com baixa concorrência:"
echo "docker run loadtest --url=http://google.com --requests=50 --concurrency=5"
echo ""

echo "4. Teste de stress com alta concorrência:"
echo "docker run loadtest --url=https://jsonplaceholder.typicode.com/posts --requests=5000 --concurrency=200"
echo ""

