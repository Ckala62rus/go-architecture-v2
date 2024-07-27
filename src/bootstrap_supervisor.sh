#!/bin/sh

go build -o ./cmd/app ./cmd/main.go
chmod 777 ./cmd/app
/usr/bin/supervisord -n -c /etc/supervisord.conf
