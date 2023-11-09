build:
	@go build -o bin/checker

run: build
	@./bin/checker

test:
	@go test ./...
