# Book Store Microservice in GO

## Dependencies

- `go get -u github.com/gin-gonic/gin`
- `go get -u gorm.io/gorm`
- `go get gorm.io/driver/sqlite`

## Build Commands

- `go build .`
- `go run main.go`
- `go install github.com/ashishsingh4u/bookmicroservice`

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
