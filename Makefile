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
	go test ./... -count=1 -v

testd:
	docker build -t redocker_test -f Dockerfile.tests .
	docker run --rm redocker_test

push-image:
	./scripts/push_image