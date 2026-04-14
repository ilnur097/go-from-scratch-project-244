build:
	go build -o bin/gendiff cmd/gendiff/main.go
test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=cp.out ./...

lint:
	golangci-lint run
