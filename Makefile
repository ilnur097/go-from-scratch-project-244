build:
	go build -o bin/gendiff cmd/gendiff/main.go
test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=cp.out ./...
	go tool cover -func=cp.out
	go tool cover -html=cp.out -o coverage.html

lint:
	golangci-lint run
