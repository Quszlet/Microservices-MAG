package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Quszlet/gateway_service/graph"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8070"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	resolvers := &graph.Resolver{
		DoctorsURL:  ensureHTTP(os.Getenv("DOCTORS_SERVICE_URL"), "http://localhost:80"),
		ScheduleURL: ensureHTTP(os.Getenv("SCHEDULE_SERVICE_URL"), "http://schedule_service:8090"),
		BookingURL:  ensureHTTP(os.Getenv("BOOKING_SERVICE_URL"), "http://booking_service:8100"),
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolvers}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func ensureHTTP(url, fallback string) string {
	if url == "" {
		return fallback
	}
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	}
	return "http://" + url
}
