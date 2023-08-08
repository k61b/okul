build:
	@go build -o bin/okul

run: build
	@./bin/okul

test: 
	@go test -v ./...

coverage:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out