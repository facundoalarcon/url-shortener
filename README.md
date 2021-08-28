## api usage example
With [Postman](https://www.postman.com/)

- `GET http://localhost:8080/hello`
- `POST http://localhost:8080/shorturl`
- `POST http://localhost:8080/mongoshorturl`

Body/rw/JSON
```
{
    "original_url" : "http://google.com",
    "short" : "g1",
    "short_url" : "http://localhost:8080/g1",
    "url_id" : 1
}
```

## Config

`go mod init <repo-name>`

`go mod init yourRepo`

`go get -u repo`

`go mod download`

`go mod tidy -v`

`go mod vendor`

## Environment file example

- add `.env` file in root folder
```
MYSQL_ROOT_PASSWORD = passw
MYSQL_DATABASE = urlshortener
MYSQL_USER = urlshortener
MYSQL_PASSWORD = passw
MYSQL_HOST = url-shortener_db_1
MYSQL_PORT = 3306

MONGO_INITDB_ROOT_USERNAME = root
MONGO_INITDB_ROOT_PASSWORD = urlshortener
MONGO_DATABASE = urlshortenerDB
MONGO_HOST = url-shortener_mongo_1
MONGO_PORT = 27017
```

## Docker

`docker build -t facundo-alarcon/url-shortener .`

Port (3000) defined on `main.go`

`docker run -p 8080:3000 -it facundo-alarcon/url-shortener`

Remove image

`docker rmi [image]`

List and Remove all docker volumes

`docker volume ls`

`docker volume prune`

Examples

`docker rmi url-shortener-2_api`

## docker-compose

`docker-compose up`

`docker-compose down`

## Middlewares
Funciones Go que estan dentro de las rutas

## Notes
- By default the DB name is `url-shortener_db_1 `
- DB Commands in /migrations. `name_table.up.sql`
- If you want use `[IF [NOT] EXISTS]` sintax
- Go Middlewares (Go functions for routes) uses [chi](https://github.com/go-chi/chi)