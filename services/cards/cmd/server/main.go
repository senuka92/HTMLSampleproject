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
	"github.com/example/pfm/services/cards"
	cardsv1 "github.com/example/pfm/services/cards/gen"
)

func main() {
	addr := envOrDefault("CARDS_GRPC_ADDR", ":50051")
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("cards listen: %v", err)
	}

	ctx := context.Background()
	pool, err := postgres.OpenPool(ctx, envOrDefault("CARDS_DATABASE_URL", ""))
	if err != nil {
		log.Fatalf("cards db: %v", err)
	}
	migrationDir := envOrDefault("CARDS_MIGRATIONS", filepath.Join("services", "cards", "migrations"))
	if err := migrate.Run(ctx, pool, migrationDir); err != nil {
		log.Fatalf("cards migrate: %v", err)
	}
	publisher, err := events.NewPublisher(envOrDefault("NATS_URL", ""))
	if err != nil {
		log.Fatalf("cards nats: %v", err)
	}

	store := cards.NewStore()
	repo := cards.NewRepository(pool)
	server := grpc.NewServer()
	cardsv1.RegisterCardsServiceServer(server, cards.NewServer(store, repo, publisher))

	log.Printf("cards service listening on %s", addr)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("cards serve: %v", err)
	}
}

func envOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
