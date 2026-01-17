package main

import (
	log "log"
	net "net"
	os "os"

	grpc "google.golang.org/grpc"

	income "github.com/example/pfm/services/income"
	incomev1 "github.com/example/pfm/services/income/gen"
)

func main() {
	addr := envOrDefault("INCOME_GRPC_ADDR", ":50053")
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("income listen: %v", err)
	}

	store := income.NewStore()
	server := grpc.NewServer()
	incomev1.RegisterIncomeServiceServer(server, income.NewServer(store))

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
