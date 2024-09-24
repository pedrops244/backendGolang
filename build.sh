#!/bin/bash

# Função para mostrar o uso
usage() {
    echo "Uso: $0 {build|test|run|migration|migrate-up|migrate-down} [args]"
    exit 1
}

# Verifica se algum argumento foi passado
if [ $# -eq 0 ]; then
    usage
fi

case "$1" in
    build)
        echo "Construindo o projeto..."
        go build -o bin/ecom cmd/main.go
        ;;
    test)
        echo "Executando os testes..."
        go test -v ./...
        ;;
    run)
        echo "Executando o projeto..."
        ./bin/ecom
        ;;
    migration)
        echo "Criando uma nova migração..."
        shift
        migrate create -ext sql -dir cmd/migrate/migrations "$@"
        ;;
    migrate-up)
        echo "Executando migrações para cima..."
        go run cmd/migrate/main.go up
        ;;
    migrate-down)
        echo "Executando migrações para baixo..."
        go run cmd/migrate/main.go down
        ;;
    *)
        usage
        ;;
esac
