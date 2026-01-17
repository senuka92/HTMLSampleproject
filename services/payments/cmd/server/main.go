package main

import (
	log "log"
	net "net"
	os "os"

	grpc "google.golang.org/grpc"

	payments "github.com/example/pfm/services/payments"
	paymentsv1 "github.com/example/pfm/services/payments/gen"
)

func main() {
	addr := envOrDefault("PAYMENTS_GRPC_ADDR", ":50052")
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("payments listen: %v", err)
	}

	store := payments.NewStore()
	server := grpc.NewServer()
	paymentsv1.RegisterPaymentsServiceServer(server, payments.NewServer(store))

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
