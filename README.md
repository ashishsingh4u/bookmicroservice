# Book Store Microservice in GO

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

- `http://localhost:8080/debug/pprof/heap`
- `http://localhost:8080/debug/pprof/block`
- `http://localhost:8080/debug/pprof/mutex`
- `https://go.dev/blog/pprof`

## Upgrade dependencies

- `go get -u`
- `go mod tidy`

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
