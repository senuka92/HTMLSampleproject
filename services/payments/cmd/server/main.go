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
	"github.com/example/pfm/services/payments"
	paymentsv1 "github.com/example/pfm/services/payments/gen"
)

func main() {
	addr := envOrDefault("PAYMENTS_GRPC_ADDR", ":50052")
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("payments listen: %v", err)
	}

	ctx := context.Background()
	pool, err := postgres.OpenPool(ctx, envOrDefault("PAYMENTS_DATABASE_URL", ""))
	if err != nil {
		log.Fatalf("payments db: %v", err)
	}
	migrationDir := envOrDefault("PAYMENTS_MIGRATIONS", filepath.Join("services", "payments", "migrations"))
	if err := migrate.Run(ctx, pool, migrationDir); err != nil {
		log.Fatalf("payments migrate: %v", err)
	}
	publisher, err := events.NewPublisher(envOrDefault("NATS_URL", ""))
	if err != nil {
		log.Fatalf("payments nats: %v", err)
	}

	store := payments.NewStore()
	repo := payments.NewRepository(pool)
	server := grpc.NewServer()
	paymentsv1.RegisterPaymentsServiceServer(server, payments.NewServer(store, repo, publisher))

	log.Printf("payments service listening on %s", addr)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("payments serve: %v", err)
	}
}

func envOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
