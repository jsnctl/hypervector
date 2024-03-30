build:
	go build -o hypervector-binary ./cmd/hypervector/hypervector.go

build-docker:
	go build -o hypervector-binary-docker ./cmd/hypervector/hypervector.go

test:
	go test ./...


all: build test
	./hypervector-binary

docker:
	docker build -t hypervector .
	docker run -p 8000:8000 hypervector:latest