package main

import (
	"log/slog"
	"net/http"
	"os"

	routes "github.com/rohitvpatil0810/go-url-shortener-api/api/v1"
	"github.com/rohitvpatil0810/go-url-shortener-api/internal/config"
	"github.com/rohitvpatil0810/go-url-shortener-api/internal/middleware"
	"github.com/rohitvpatil0810/go-url-shortener-api/pkg/logger"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		slog.Error("Error loading config", "error", err)
		os.Exit(1)
	}

	router := routes.RegisterV1Routes()

	middlewareChain := middleware.CreateChain(middleware.Logging)

	server := http.Server{
		Addr:    ":" + config.Port,
		Handler: middlewareChain(router),
	}

	logger.Logger.Info("Starting server", slog.String("port", config.Port))
	server.ListenAndServe()
}
