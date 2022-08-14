# go-gin-boilerplate

A boilerplate for a go-gin project.

A Full-Stack Web Application built with Go and Gin.

## Installation

Clone it from GitHub,

```bash
git clone https://github.com/mrinjamul/go-gin-boilerplate.git
```

And start development,

```bash
    cd go-gin-boilerplate
    go mod download
    go run main.go
```

Generate OpenAPI spec:

```bash
swag init --parseDependency --parseInternal
```

Initialize database,

```bash
http POST http://localhost:3000/api/v1/quote data:=@data.json
```

## Technology Stack

Technologies used:

- [Golang](https://golang.org/): [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [Gorm](https://gorm.io/)
- [Sqlite](https://www.sqlite.org/): [go-sqlite3](https://github.com/mattn/go-sqlite3)
- [Swagger](https://swagger.io/): [Go Swagger](https://github.com/swaggo/gin-swagger)

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) to learn how to contribute to this project.

Also see [CODE OF CONDUCT](CODE_OF_CONDUCT.md).

## License

open-sourced under MIT license [MIT](LICENSE)
