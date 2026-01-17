package main

import (
	log "log"
	net "net"
	os "os"

	grpc "google.golang.org/grpc"

	cards "github.com/example/pfm/services/cards"
	cardsv1 "github.com/example/pfm/services/cards/gen"
)

func main() {
	addr := envOrDefault("CARDS_GRPC_ADDR", ":50051")
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("cards listen: %v", err)
	}

	store := cards.NewStore()
	server := grpc.NewServer()
	cardsv1.RegisterCardsServiceServer(server, cards.NewServer(store))

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
