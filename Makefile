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
	go test ./internal/configuration/... -cover

test-infrastructure:
	go test ./internal/infrastructure/... -cover

test-adapters:
	go test ./internal/adapters/... -cover

test-application:
	go test ./internal/application/... -cover

test-domain:
	go test ./internal/domain/... -cover
