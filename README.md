# Book Store Microservice in GO

## Create new project

- `mkdir -p bookmicroservice`
- `cd bookmicroservice`
- `go mod init bookmicroservice`

## Dependencies

- `go get -u github.com/gin-gonic/gin`
- `go get -u gorm.io/gorm`
- `go get gorm.io/driver/sqlite`
- `go get github.com/tkanos/gonfig`

## Build Commands

- `go build .`
- `go run main.go`
- `go install github.com/ashishsingh4u/bookmicroservice`

## Profiling Command

- `go tool pprof http://localhost:8080/debug/pprof/heap`
- `go tool pprof http://localhost:8080/debug/pprof/block`
- `go tool pprof http://localhost:8080/debug/pprof/mutex`
- `https://go.dev/blog/pprof`

## Upgrade dependencies

- `go get -u`
- `go mod tidy`

## Swagger Support

- `go get -u github.com/swaggo/swag/cmd/swag`
- `../../bin/swag init`
- `go get -u github.com/swaggo/gin-swagger`
- `go get -u github.com/swaggo/files`
- `http://localhost:8080/swagger/index.html`

## Docker Support

- `docker build -t bookmicroservice .`

## Debug Server

- Create launch.json with following content

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch Package",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "main.go"
    }
  ]
}
```
