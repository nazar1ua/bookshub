# BooksHub API
Repository for books library API

## Запуск

1. Для запуску у вас має бути встановлений [Docker](https://www.docker.com/products/docker-desktop/) та [Go](https://go.dev/dl/)
2. Клонуйте цей репозиторій
3. Оновіть дані в файлі `docker-compose.yml` та `cmd/api/main.go`
4. В папці проєкту запустіть команду (перший запуск може бути довгим, прапорець `d` запускає команду в фоні)
```shell
docker compose up -d
```
5. Запустіть сервер Go
```shell
make start # або go run ./cmd/api
```
6. Ваш сервер доступний за посиланням `localhost:8081` :tada:
