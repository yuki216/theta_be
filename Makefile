.PHONY: all build run test tidy lint clean docker-up docker-down

all: test build

tidy:
	go mod tidy

build:
	go build -v -o bin/app.exe ./cmd/app

run:
	go run ./cmd/app

test:
	go test -v -race ./...

lint:
	go vet ./...

clean:
	rm -rf bin/

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down
