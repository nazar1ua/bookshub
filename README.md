# BooksHub API
Repository for books library API

## Запуск

1. Для запуску у вас має бути встановлений [Docker](https://www.docker.com/products/docker-desktop/) та [Go](https://go.dev/dl/)
2. Клонуйте цей репозиторій
3. Оновіть дані в файлі `docker-compose.yml` та `cmd/api/main.go`
4. В папці проєкту запустіть команду (перший запуск може бути довгим)
```shell
docker compose up
```
5. Після повідомлення `database system is ready to accept connections` зупиніть команду і запустіть її в фоні
```shell
docker compose start
```
6. Запустіть сервер Go
```shell
go run ./cmd/api
```
7. Ваш сервер доступний за посиланням `localhost:8081` :tada:
