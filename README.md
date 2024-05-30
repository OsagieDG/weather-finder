# weather-finder
gRPC and HTTP endpoints that efficiently serve weather data.

## Generate your own `cert.pem` and `key.pem` files with this simple command
```
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

## Installing and Running grpcurl command in the terminal
```
          go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
          go get github.com/fullstorydev/grpcurl/cmd/grpcurl 
```

```
                  grpcurl -d '{"city": "Lisbon", "country": "Portugal"}' -import-path . -proto api/v1/weather.proto -cacert tls/cert.pem -servername localhost localhost:50051 weather.v1.WeatherService/GetWeather
```

## Automating Program Compilation with a Makefile
- To generate code from weather.proto file simply use:
make compile

- Build and Run Target with:
```
make build
make run 
```

## Project environment variables
```
GRPC_SERVER_ADDRESS=
HTTP_SERVER_PORT=
OPENWEATHERMAP_API_KEY=
OPENWEATHERMAP_API_BASE_URL=
```

### gRPC request and response
![weather1](https://github.com/osag1e/weather-finder/blob/main/images/weather1.png)

### http request and response
![weather2](https://github.com/osag1e/weather-finder/blob/main/images/weather2.png)


