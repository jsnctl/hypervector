build:
	go build -o hypervector-binary ./cmd/hypervector/hypervector.go



test:
	go test ./...


all: build test
	./hypervector-binary