package main

import (
	log "log"
	net "net"
	os "os"

	grpc "google.golang.org/grpc"

	alerts "github.com/example/pfm/services/alerts"
	alertsv1 "github.com/example/pfm/services/alerts/gen"
)

func main() {
	addr := envOrDefault("ALERTS_GRPC_ADDR", ":50054")
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("alerts listen: %v", err)
	}

	store := alerts.NewStore()
	handler := alerts.NewServer(store)
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
