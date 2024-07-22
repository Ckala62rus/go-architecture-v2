1. generate swagger documentation

   go install github.com/swaggo/swag/cmd/swag@latest

   swag init -g ./cmd/main.go -o ./docs

   documentation swagger
   https://github.com/swaggo/swag/blob/master/README.md#use-multiple-path-params

2. run app and regenerate swagger docs
   swag init -g ./cmd/main.go -o ./docs; go run .\cmd\main.go

3. swagger documentation
   http://localhost:8081/swagger/index.html

4. если будет проблема при запуске скриптов sh,
   то заходим в контейнер и применяем на скриптах
   dos2unix bootstrap_app.sh
   (проблема связана с переводом коретки на новую строку windows и linux)

5. Docker
   docker exec -ti backend /bin/sh
   docker-compose down
   docker-compose build

   go build -o app ./cmd/main.go
   sh bootstrap.sh
