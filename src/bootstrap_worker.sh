#!/bin/sh

go build -o worker ./worker/worker.go
chmod 777 ./cmd/app
./worker/worker
