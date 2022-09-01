## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## run: run the cmd/app application
run:
	go run ./cmd/app

## migrate: run the cmd/app application with migrations
migrate:
	go run -tags migrate ./cmd/app

## run: run the cmd/app application
doc:
	swag init -g ./internal/controller/http/v1/router.go

## build: build the cmd/app application
build:
	@echo 'Building cmd/api...'
	go build -ldflags='-s' -o=./bin/api ./cmd/api
