start:
	go run cmd/main.go postgres://postgres:3344@localhost:5432/go
build:
	go build -o test cmd/main.go