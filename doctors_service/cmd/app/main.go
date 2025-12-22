package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Quszlet/doctors_service/internal/api/handler"
	"github.com/Quszlet/doctors_service/internal/kafka"
	"github.com/Quszlet/doctors_service/internal/repository"
	"github.com/Quszlet/doctors_service/internal/service"
)

func main() {
	slog.Info("Doctors service is start...")

	slog.SetLogLoggerLevel(slog.LevelDebug)

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "postgres",
		Port:     "5432",
		Username: "postgres",
		DBName:   "doctors_db",
		SSLMode:  "disable",
		Password: "postgres",
	})

	if err != nil {
		slog.Error("failed to initialize db", "error", err)
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)

	producers, cleanupProducers, err := kafka.InitKafkaProducers()
	if err != nil {
		slog.Error("failed to initialize kafka producers", "error", err)
	}
	defer cleanupProducers()

	router := handler.NewHandler(service, producers)
	routes := router.InitRoutes()

	srv := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	slog.Error(srv.ListenAndServe().Error())
}
