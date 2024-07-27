#!/bin/sh

go install github.com/swaggo/swag/cmd/swag@latest
swag init -g ./cmd/main.go -o ./docs

go build -o ./cmd/app ./cmd/main.go
chmod 777 ./cmd/app
./cmd/app
