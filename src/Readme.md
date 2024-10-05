# Golang gin api template

### Enviroment variables
> Создать файл .env из .env.example

### Generate swagger documentation.

#### Swagger documentation
```Bash

# Package for create swagger documentation
go install github.com/swaggo/swag/cmd/swag@latest

# Command for generate documentation
swag init -g ./cmd/main.go -o ./docs

# alternative generate documentation with run application 
swag init -g ./cmd/main.go -o ./docs; go run .\cmd\main.go
```

### Official documentation for examples
> https://github.com/swaggo/swag/blob/master/README.md#use-multiple-path-params

### swagger documentation  address (by default)
>   http://localhost:8081/swagger/index.html

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

   docker-compose up --build --force-recreate --renew-anon-volumes
   
   docker-compose build --no-cache

6. Документация по Gin framework https://gin-gonic.com/

//todo
   1) Сделать DTO структуры для входных и выходных данных
   2) Сделать авторизацию по JWT или по BasicAuth
   3) Привести код в порядок
   4) Добаавить реализацию Redis PubSub
   5) Создать приложение блог с категориями, тегами и пользователем
   6) подумать, что сделать с транзакциями. есть мысль передавать GORM прямо в метод репозитория.
   7) сделать сохранение токена JWT в редисе с временем жизни.
   8) сделать хранение токенов JWT в Redis. 
   9) перенести инициализацию ллогера из main в сам логер.
   10) логировать коннекты к Redis и Postgres.

