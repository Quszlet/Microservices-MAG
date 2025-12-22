package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Quszlet/booking_service/internal/api/consumer"
	"github.com/Quszlet/booking_service/internal/api/handler"
	"github.com/Quszlet/booking_service/internal/kafka"
	"github.com/Quszlet/booking_service/internal/repository"
	"github.com/Quszlet/booking_service/internal/service"
)

func main() {
	slog.Info("Booking service is start...")

	slog.SetLogLoggerLevel(slog.LevelDebug)

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "postgres",
		Port:     "5432",
		Username: "postgres",
		DBName:   "booking_db",
		SSLMode:  "disable",
		Password: "postgres",
	})

	if err != nil {
		slog.Error("Failed to initialize db", "error", err.Error())
	}

	repo := repository.NewRepository(db)

	service := service.NewService(repo)

	consumers, err := consumer.PrepareConsumers(service)
	if err != nil {
		slog.Error("failed to initialize kafka consumers", "error", err.Error())
	}
	defer func() {
		for _, c := range consumers {
			if c.Consumer != nil {
				c.Consumer.Close()
			}
		}
	}()
	for _, c := range consumers {
		go func(r consumer.Runner) {
			if err := r.Run(context.Background()); err != nil {
				slog.Error("kafka consumer error", "error", err.Error())
			}
		}(c)
	}

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
