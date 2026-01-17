package payments

import (
	context "context"
	fmt "fmt"
	strconv "strconv"
	sync "sync"
	atomic "sync/atomic"

	paymentsv1 "github.com/example/pfm/services/payments/gen"
)

type Store struct {
	mu        sync.RWMutex
	payments  map[string]*paymentsv1.Payment
	idCounter uint64
}

func NewStore() *Store {
	return &Store{payments: make(map[string]*paymentsv1.Payment)}
}

type Server struct {
	paymentsv1.UnimplementedPaymentsServiceServer
	store *Store
}

func NewServer(store *Store) *Server {
	return &Server{store: store}
}

func (s *Server) CreatePayment(ctx context.Context, req *paymentsv1.CreatePaymentRequest) (*paymentsv1.Payment, error) {
	if req.GetCardId() == "" || req.GetStatementId() == "" {
		return nil, fmt.Errorf("card_id and statement_id are required")
	}
	payment := &paymentsv1.Payment{
		Id:          s.nextID(),
		CardId:      req.GetCardId(),
		StatementId: req.GetStatementId(),
		Amount:      req.GetAmount(),
		PaidAt:      req.GetPaidAt(),
		Note:        req.GetNote(),
	}
	s.store.mu.Lock()
	s.store.payments[payment.Id] = payment
	s.store.mu.Unlock()
	return payment, nil
}

func (s *Server) nextID() string {
	id := atomic.AddUint64(&s.store.idCounter, 1)
	return strconv.FormatUint(id, 10)
}
