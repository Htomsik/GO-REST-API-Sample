

<h2 align="center">GO REST API</h3>


<h3 align="center">

Just my GO playground

</h3>



## What can

- Base account operation (Authorize, create, deactivate)
- Store and process info in db
- Have swagger (url/swagger)
> All configs in configs/apiServer.toml


## You need

- GCC and CGO_ENABLED=1 (for go-sqlite3) `go env -w CGO_ENABLED=1`
- [Swag](https://github.com/swaggo/swag) (for generate swagger docs)

### RUN

#### With Make file

``` bash
    Make build  #  build
    Make Start  #  build and start
    Make fullBuild # build with swaggo
```
#### Manual

``` bash
    go mod tidy # if needed
    swag init -g .\cmd\apiServer\main.go # optional (swagger docs)
    go build -v ./cmd/apiServer
    ./apiServer
```

