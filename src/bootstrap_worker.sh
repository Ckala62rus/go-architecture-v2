#!/bin/sh

go build -o worker ./worker/worker.go
./worker/worker
