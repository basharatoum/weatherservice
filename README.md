to run 
```go mod init main.go```
```go mod tidy```
```go run main.go```

Example url to call:
GET AT : ```http://localhost:8080/weather/34/-118.0```

Example result:
```{"shortForecast":"Sunny","temperature":"moderate"}```

USED https://transform.tools/json-to-go for model extraction from swagger

Rest it boiler plate http client code and gin for the http server!
