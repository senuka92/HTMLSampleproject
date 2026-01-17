package income

import (
	context "context"
	fmt "fmt"
	strconv "strconv"
	sync "sync"
	atomic "sync/atomic"

	events "github.com/example/pfm/pkg/events"
	incomev1 "github.com/example/pfm/services/income/gen"
)

type Store struct {
	mu        sync.RWMutex
	incomes   map[string]*incomev1.Income
	idCounter uint64
}

func NewStore() *Store {
	return &Store{incomes: make(map[string]*incomev1.Income)}
}

type Server struct {
	incomev1.UnimplementedIncomeServiceServer
	store *Store
	repo  *Repository
	pubs  *events.Publisher
}

func NewServer(store *Store, repo *Repository, pubs *events.Publisher) *Server {
	return &Server{store: store, repo: repo, pubs: pubs}
}

func (s *Server) UpsertIncome(ctx context.Context, req *incomev1.UpsertIncomeRequest) (*incomev1.Income, error) {
	if req.GetMonth() == "" || req.GetCurrency() == "" {
		return nil, fmt.Errorf("month and currency are required")
	}
	income := &incomev1.Income{
		Id:       s.nextID(),
		Month:    req.GetMonth(),
		Currency: req.GetCurrency(),
		Amount:   req.GetAmount(),
		Source:   req.GetSource(),
		Note:     req.GetNote(),
	}
	s.store.mu.Lock()
	s.store.incomes[income.Id] = income
	s.store.mu.Unlock()
	if s.repo != nil {
		if err := s.repo.UpsertIncome(ctx, income); err != nil {
			return nil, err
		}
	}
	if s.pubs != nil {
		_ = s.pubs.Publish(ctx, "income.upserted", income)
	}
	return income, nil
}

func (s *Server) nextID() string {
	id := atomic.AddUint64(&s.store.idCounter, 1)
	return strconv.FormatUint(id, 10)
}
