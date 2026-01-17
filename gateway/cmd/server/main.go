package main

import (
	context "context"
	encodingjson "encoding/json"
	log "log"
	nethttp "net/http"
	os "os"
	time "time"

	grpc "google.golang.org/grpc"
	credentials "google.golang.org/grpc/credentials/insecure"

	alertsv1 "github.com/example/pfm/services/alerts/gen"
	cardsv1 "github.com/example/pfm/services/cards/gen"
	paymentsv1 "github.com/example/pfm/services/payments/gen"
)

const (
	defaultCardsAddr    = "localhost:50051"
	defaultPaymentsAddr = "localhost:50052"
	defaultAlertsAddr   = "localhost:50054"
)

type server struct {
	cardsClient     cardsv1.CardsServiceClient
	paymentsClient  paymentsv1.PaymentsServiceClient
	alertsClient    alertsv1.AlertsServiceClient
	dashboardClient alertsv1.DashboardServiceClient
}

func main() {
	cardsConn := dialGRPC(envOrDefault("CARDS_GRPC_ADDR", defaultCardsAddr))
	paymentsConn := dialGRPC(envOrDefault("PAYMENTS_GRPC_ADDR", defaultPaymentsAddr))
	alertsConn := dialGRPC(envOrDefault("ALERTS_GRPC_ADDR", defaultAlertsAddr))
	defer cardsConn.Close()
	defer paymentsConn.Close()
	defer alertsConn.Close()

	svc := &server{
		cardsClient:     cardsv1.NewCardsServiceClient(cardsConn),
		paymentsClient:  paymentsv1.NewPaymentsServiceClient(paymentsConn),
		alertsClient:    alertsv1.NewAlertsServiceClient(alertsConn),
		dashboardClient: alertsv1.NewDashboardServiceClient(alertsConn),
	}

	mux := nethttp.NewServeMux()
	mux.HandleFunc("/healthz", svc.handleHealth)
	mux.HandleFunc("/banks", svc.handleCreateBank)
	mux.HandleFunc("/cards", svc.handleCreateCard)
	mux.HandleFunc("/statements", svc.handleCreateStatement)
	mux.HandleFunc("/payments", svc.handleCreatePayment)
	mux.HandleFunc("/alerts/schedule", svc.handleScheduleAlert)
	mux.HandleFunc("/dashboard/summary", svc.handleSummary)

	addr := envOrDefault("GATEWAY_ADDR", ":8081")
	log.Printf("gateway listening on %s", addr)
	if err := nethttp.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("gateway serve: %v", err)
	}
}

func dialGRPC(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(credentials.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc dial %s: %v", addr, err)
	}
	return conn
}

func (s *server) handleHealth(w nethttp.ResponseWriter, r *nethttp.Request) {
	writeJSON(w, nethttp.StatusOK, map[string]string{"status": "ok"})
}

func (s *server) handleCreateBank(w nethttp.ResponseWriter, r *nethttp.Request) {
	if r.Method != nethttp.MethodPost {
		writeError(w, nethttp.StatusMethodNotAllowed, "method not allowed")
		return
	}
	var payload struct {
		Name string `json:"name"`
	}
	if err := decodeJSON(r, &payload); err != nil {
		writeError(w, nethttp.StatusBadRequest, err.Error())
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	res, err := s.cardsClient.CreateBank(ctx, &cardsv1.CreateBankRequest{Name: payload.Name})
	if err != nil {
		writeError(w, nethttp.StatusBadGateway, err.Error())
		return
	}
	writeJSON(w, nethttp.StatusCreated, res)
}

func (s *server) handleCreateCard(w nethttp.ResponseWriter, r *nethttp.Request) {
	if r.Method != nethttp.MethodPost {
		writeError(w, nethttp.StatusMethodNotAllowed, "method not allowed")
		return
	}
	var payload struct {
		BankID   string `json:"bank_id"`
		Name     string `json:"name"`
		LastFour string `json:"last_four"`
		DueDay   string `json:"due_day"`
	}
	if err := decodeJSON(r, &payload); err != nil {
		writeError(w, nethttp.StatusBadRequest, err.Error())
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	res, err := s.cardsClient.CreateCard(ctx, &cardsv1.CreateCardRequest{
		BankId:   payload.BankID,
		Name:     payload.Name,
		LastFour: payload.LastFour,
		DueDay:   payload.DueDay,
	})
	if err != nil {
		writeError(w, nethttp.StatusBadGateway, err.Error())
		return
	}
	writeJSON(w, nethttp.StatusCreated, res)
}

func (s *server) handleCreateStatement(w nethttp.ResponseWriter, r *nethttp.Request) {
	if r.Method != nethttp.MethodPost {
		writeError(w, nethttp.StatusMethodNotAllowed, "method not allowed")
		return
	}
	var payload struct {
		CardID         string  `json:"card_id"`
		StatementMonth string  `json:"statement_month"`
		DueDate        string  `json:"due_date"`
		DueAmount      float64 `json:"due_amount"`
	}
	if err := decodeJSON(r, &payload); err != nil {
		writeError(w, nethttp.StatusBadRequest, err.Error())
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	res, err := s.cardsClient.CreateStatement(ctx, &cardsv1.CreateStatementRequest{
		CardId:         payload.CardID,
		StatementMonth: payload.StatementMonth,
		DueDate:        payload.DueDate,
		DueAmount:      payload.DueAmount,
	})
	if err != nil {
		writeError(w, nethttp.StatusBadGateway, err.Error())
		return
	}
	writeJSON(w, nethttp.StatusCreated, res)
}

func (s *server) handleCreatePayment(w nethttp.ResponseWriter, r *nethttp.Request) {
	if r.Method != nethttp.MethodPost {
		writeError(w, nethttp.StatusMethodNotAllowed, "method not allowed")
		return
	}
	var payload struct {
		CardID      string  `json:"card_id"`
		StatementID string  `json:"statement_id"`
		Amount      float64 `json:"amount"`
		PaidAt      string  `json:"paid_at"`
		Note        string  `json:"note"`
	}
	if err := decodeJSON(r, &payload); err != nil {
		writeError(w, nethttp.StatusBadRequest, err.Error())
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	res, err := s.paymentsClient.CreatePayment(ctx, &paymentsv1.CreatePaymentRequest{
		CardId:      payload.CardID,
		StatementId: payload.StatementID,
		Amount:      payload.Amount,
		PaidAt:      payload.PaidAt,
		Note:        payload.Note,
	})
	if err != nil {
		writeError(w, nethttp.StatusBadGateway, err.Error())
		return
	}
	writeJSON(w, nethttp.StatusCreated, res)
}

func (s *server) handleScheduleAlert(w nethttp.ResponseWriter, r *nethttp.Request) {
	if r.Method != nethttp.MethodPost {
		writeError(w, nethttp.StatusMethodNotAllowed, "method not allowed")
		return
	}
	var payload struct {
		CardID      string  `json:"card_id"`
		StatementID string  `json:"statement_id"`
		DueDate     string  `json:"due_date"`
		DueAmount   float64 `json:"due_amount"`
	}
	if err := decodeJSON(r, &payload); err != nil {
		writeError(w, nethttp.StatusBadRequest, err.Error())
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	res, err := s.alertsClient.ScheduleAlert(ctx, &alertsv1.ScheduleAlertRequest{
		CardId:      payload.CardID,
		StatementId: payload.StatementID,
		DueDate:     payload.DueDate,
		DueAmount:   payload.DueAmount,
	})
	if err != nil {
		writeError(w, nethttp.StatusBadGateway, err.Error())
		return
	}
	writeJSON(w, nethttp.StatusCreated, res)
}

func (s *server) handleSummary(w nethttp.ResponseWriter, r *nethttp.Request) {
	if r.Method != nethttp.MethodGet {
		writeError(w, nethttp.StatusMethodNotAllowed, "method not allowed")
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	res, err := s.dashboardClient.GetSummary(ctx, &alertsv1.GetSummaryRequest{})
	if err != nil {
		writeError(w, nethttp.StatusBadGateway, err.Error())
		return
	}
	writeJSON(w, nethttp.StatusOK, res)
}

func decodeJSON(r *nethttp.Request, target interface{}) error {
	decoder := encodingjson.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(target)
}

func writeJSON(w nethttp.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := encodingjson.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("json encode: %v", err)
	}
}

func writeError(w nethttp.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func envOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
