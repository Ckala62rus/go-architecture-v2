#!/bin/sh

go build -o ./cmd/app ./cmd/main.go
chmod 777 ./cmd/app
./cmd/app
