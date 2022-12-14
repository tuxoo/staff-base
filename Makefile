.PHONY: all lint test build
.SILENT:

lint:
	golangci-lint run

swagger:
	swag init -g ./internal/app/app.go

tidy:
	go mod tidy

clean:
	go clean -modcache

build:
	go build -o ./.bin/app ./cmd/main.go

docker:
	docker compose up