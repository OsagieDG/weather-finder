package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"time"

	pb "github.com/OsagieDG/weather-finder/api/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func weatherRouter(grpcServerAddress string, tlsConfig *tls.Config) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/weather/{city}/{country}", func(w http.ResponseWriter, r *http.Request) {
		city := chi.URLParam(r, "city")
		country := chi.URLParam(r, "country")

		conn, err := grpc.Dial(
			grpcServerAddress,
			grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		client := pb.NewWeatherServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		req := &pb.WeatherRequest{
			City:    city,
			Country: country,
		}

		weather, err := client.GetWeather(ctx, req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(map[string]interface{}{
			"Location":    weather.Location,
			"Temperature": weather.Temperature,
			"Conditions":  weather.Conditions,
		}); err != nil {
			log.Printf("Error encoding JSON response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	return router
}
