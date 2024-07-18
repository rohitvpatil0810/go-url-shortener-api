package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/rohitvpatil0810/go-url-shortener-api/internal/config"
	"github.com/rohitvpatil0810/go-url-shortener-api/pkg/logger"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		slog.Error("Error loading config", "error", err)
		os.Exit(1)
	}

	http.Handle("/", http.HandlerFunc(serveHome))

	logger.Logger.Info("Starting server", slog.String("port", config.Port))
	http.ListenAndServe(":"+config.Port, nil)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	slog.Info("Serving home")
	w.Write([]byte("<p>Welcome to Url Shortener API</p>"))
}
