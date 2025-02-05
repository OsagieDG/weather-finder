package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	pb "github.com/OsagieDG/weather-finder/api/v1"
	service "github.com/OsagieDG/weather-finder/service/weather"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	grpcServerAddress := os.Getenv("GRPC_SERVER_ADDRESS")
	httpServerPort := os.Getenv("HTTP_SERVER_PORT")

	certFile := "tls/cert.pem"
	keyFile := "tls/key.pem"

	cert, err := os.ReadFile(certFile)
	if err != nil {
		log.Fatalf("Failed to load server certificate: %v", err)
	}

	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(cert); !ok {
		log.Fatalf("Failed to append server certificate to the certificate pool")
	}

	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}

	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to load TLS certificates: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterWeatherServiceServer(grpcServer, &service.WeatherServer{})

	go func() {
		lis, err := net.Listen("tcp", grpcServerAddress)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		defer lis.Close()

		fmt.Printf("gRPC server is running on %s...\n", grpcServerAddress)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	router := weatherRouter(grpcServerAddress, tlsConfig)

	fmt.Printf("HTTP server is running on port %s...\n", httpServerPort)
	if err := http.ListenAndServe(httpServerPort, router); err != nil {
		panic(err)
	}
}
