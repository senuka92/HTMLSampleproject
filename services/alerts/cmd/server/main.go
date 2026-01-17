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
	"github.com/example/pfm/services/alerts"
	alertsv1 "github.com/example/pfm/services/alerts/gen"
)

func main() {
	addr := envOrDefault("ALERTS_GRPC_ADDR", ":50054")
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("alerts listen: %v", err)
	}

	ctx := context.Background()
	pool, err := postgres.OpenPool(ctx, envOrDefault("ALERTS_DATABASE_URL", ""))
	if err != nil {
		log.Fatalf("alerts db: %v", err)
	}
	migrationDir := envOrDefault("ALERTS_MIGRATIONS", filepath.Join("services", "alerts", "migrations"))
	if err := migrate.Run(ctx, pool, migrationDir); err != nil {
		log.Fatalf("alerts migrate: %v", err)
	}
	publisher, err := events.NewPublisher(envOrDefault("NATS_URL", ""))
	if err != nil {
		log.Fatalf("alerts nats: %v", err)
	}

	store := alerts.NewStore()
	repo := alerts.NewRepository(pool)
	handler := alerts.NewServer(store, repo, publisher)
	server := grpc.NewServer()
	alertsv1.RegisterAlertsServiceServer(server, handler)
	alertsv1.RegisterDashboardServiceServer(server, handler)

	log.Printf("alerts service listening on %s", addr)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("alerts serve: %v", err)
	}
}

func envOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
