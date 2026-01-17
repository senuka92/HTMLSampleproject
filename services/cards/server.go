package cards

import (
	context "context"
	fmt "fmt"
	strconv "strconv"
	sync "sync"
	atomic "sync/atomic"

	cardsv1 "github.com/example/pfm/services/cards/gen"
)

type Store struct {
	mu         sync.RWMutex
	banks      map[string]*cardsv1.Bank
	cards      map[string]*cardsv1.Card
	statements map[string]*cardsv1.Statement
	idCounter  uint64
}

func NewStore() *Store {
	return &Store{
		banks:      make(map[string]*cardsv1.Bank),
		cards:      make(map[string]*cardsv1.Card),
		statements: make(map[string]*cardsv1.Statement),
	}
}

type Server struct {
	cardsv1.UnimplementedCardsServiceServer
	store *Store
}

func NewServer(store *Store) *Server {
	return &Server{store: store}
}

func (s *Server) CreateBank(ctx context.Context, req *cardsv1.CreateBankRequest) (*cardsv1.Bank, error) {
	if req.GetName() == "" {
		return nil, fmt.Errorf("name is required")
	}
	bank := &cardsv1.Bank{
		Id:   s.nextID(),
		Name: req.GetName(),
	}
	s.store.mu.Lock()
	s.store.banks[bank.Id] = bank
	s.store.mu.Unlock()
	return bank, nil
}

func (s *Server) CreateCard(ctx context.Context, req *cardsv1.CreateCardRequest) (*cardsv1.Card, error) {
	if req.GetBankId() == "" || req.GetName() == "" {
		return nil, fmt.Errorf("bank_id and name are required")
	}
	card := &cardsv1.Card{
		Id:       s.nextID(),
		BankId:   req.GetBankId(),
		Name:     req.GetName(),
		LastFour: req.GetLastFour(),
		DueDay:   req.GetDueDay(),
	}
	s.store.mu.Lock()
	s.store.cards[card.Id] = card
	s.store.mu.Unlock()
	return card, nil
}

func (s *Server) CreateStatement(ctx context.Context, req *cardsv1.CreateStatementRequest) (*cardsv1.Statement, error) {
	if req.GetCardId() == "" || req.GetStatementMonth() == "" {
		return nil, fmt.Errorf("card_id and statement_month are required")
	}
	statement := &cardsv1.Statement{
		Id:             s.nextID(),
		CardId:         req.GetCardId(),
		StatementMonth: req.GetStatementMonth(),
		DueDate:        req.GetDueDate(),
		DueAmount:      req.GetDueAmount(),
	}
	s.store.mu.Lock()
	s.store.statements[statement.Id] = statement
	s.store.mu.Unlock()
	return statement, nil
}

func (s *Server) nextID() string {
	id := atomic.AddUint64(&s.store.idCounter, 1)
	return strconv.FormatUint(id, 10)
}
