package payments

import (
	context "context"
	fmt "fmt"
	strconv "strconv"
	sync "sync"
	atomic "sync/atomic"

	events "github.com/example/pfm/pkg/events"
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
	repo  *Repository
	pubs  *events.Publisher
}

func NewServer(store *Store, repo *Repository, pubs *events.Publisher) *Server {
	return &Server{store: store, repo: repo, pubs: pubs}
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
	if s.repo != nil {
		if err := s.repo.CreatePayment(ctx, payment); err != nil {
			return nil, err
		}
	}
	if s.pubs != nil {
		_ = s.pubs.Publish(ctx, "payment.created", payment)
	}
	return payment, nil
}

func (s *Server) nextID() string {
	id := atomic.AddUint64(&s.store.idCounter, 1)
	return strconv.FormatUint(id, 10)
}
