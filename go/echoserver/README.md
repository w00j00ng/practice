# Practice for echo server

## 참고자료

- https://echo.labstack.com/cookbook/crud/
- https://github.com/swaggo/swag
- https://github.com/swaggo/echo-swagger

## swagger UI

### swaggo 설치

```sh
$ go install github.com/swaggo/swag/cmd/swag@latest
```

### swagger 문서 최신화

```sh
$ swag init -g main.go -o docs
```

