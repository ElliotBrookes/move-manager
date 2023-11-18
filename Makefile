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

test:
	go test ./...
