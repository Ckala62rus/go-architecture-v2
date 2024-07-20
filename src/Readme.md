1. generate swagger documentation
   swag init -g ./cmd/main.go -o ./docs

   documentation swagger
   https://github.com/swaggo/swag/blob/master/README.md#use-multiple-path-params

2. run app and regenerate swagger docs
   swag init -g ./cmd/main.go -o ./docs; go run .\cmd\main.go

3. swagger documentation
   http://localhost:8081/swagger/index.html

docker exec -ti backend /bin/sh
docker-compose down
docker-compose build
go build -o app ./cmd/main.go

run app in container
./cmd/app

