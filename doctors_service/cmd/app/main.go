package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Quszlet/provider_service/internal/api/handler"
)

func main() {
	slog.Info("Doctors service is start...")

	slog.SetLogLoggerLevel(slog.LevelDebug)

	routes := handler.NewHandler().InitRoutes();

	srv := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	slog.Error(srv.ListenAndServe().Error())
}
