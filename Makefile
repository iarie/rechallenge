.DEFAULT_GOAL := run

.PHONY: fmt vet build run
fmt:
	go fmt ./...
vet: fmt
	go vet ./...
build: vet
	go build -o bin/run cmd/main.go
run: build
	./bin/run
test:
	go test ./...