build:
	go build -o bin/cli cmd/cli/main.go
	chmod +x bin/cli

execute:
	./bin/cli

cleanup:
	rm -r bin/*

clear-db:
	rm -r dataset.db

run: build execute cleanup

test-all:
	go test ./...

test-layers: test-config test-infrastructure test-adapters test-application test-domain

test-config:
	go test ./internal/configuration/...

test-infrastructure:
	go test ./internal/infrastructure/...

test-adapters:
	go test ./internal/adapters/...

test-application:
	go test ./internal/application/...

test-domain:
	go test ./internal/domain/...
