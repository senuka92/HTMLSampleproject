package main

import (
	"context"
	"log"
	"net"
	"os"
	"path/filepath"

	"google.golang.org/grpc"

	"github.com/example/pfm/pkg/events"
	"github.com/example/pfm/pkg/migrate"
	"github.com/example/pfm/pkg/postgres"
	"github.com/example/pfm/services/income"
	incomev1 "github.com/example/pfm/services/income/gen"
)

func main() {
	addr := envOrDefault("INCOME_GRPC_ADDR", ":50053")
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("income listen: %v", err)
	}

	ctx := context.Background()
	pool, err := postgres.OpenPool(ctx, envOrDefault("INCOME_DATABASE_URL", ""))
	if err != nil {
		log.Fatalf("income db: %v", err)
	}
	migrationDir := envOrDefault("INCOME_MIGRATIONS", filepath.Join("services", "income", "migrations"))
	if err := migrate.Run(ctx, pool, migrationDir); err != nil {
		log.Fatalf("income migrate: %v", err)
	}
	publisher, err := events.NewPublisher(envOrDefault("NATS_URL", ""))
	if err != nil {
		log.Fatalf("income nats: %v", err)
	}

	store := income.NewStore()
	repo := income.NewRepository(pool)
	server := grpc.NewServer()
	incomev1.RegisterIncomeServiceServer(server, income.NewServer(store, repo, publisher))

	log.Printf("income service listening on %s", addr)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("income serve: %v", err)
	}
}

func envOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
