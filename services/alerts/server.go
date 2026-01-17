package alerts

import (
	context "context"
	fmt "fmt"
	strconv "strconv"
	sync "sync"
	atomic "sync/atomic"
	time "time"

	events "github.com/example/pfm/pkg/events"
	alertsv1 "github.com/example/pfm/services/alerts/gen"
)

type Store struct {
	mu        sync.RWMutex
	alerts    map[string]*alertsv1.Alert
	idCounter uint64
}

func NewStore() *Store {
	return &Store{alerts: make(map[string]*alertsv1.Alert)}
}

type Server struct {
	alertsv1.UnimplementedAlertsServiceServer
	alertsv1.UnimplementedDashboardServiceServer
	store *Store
	repo  *Repository
	pubs  *events.Publisher
}

func NewServer(store *Store, repo *Repository, pubs *events.Publisher) *Server {
	return &Server{store: store, repo: repo, pubs: pubs}
}

func (s *Server) ScheduleAlert(ctx context.Context, req *alertsv1.ScheduleAlertRequest) (*alertsv1.Alert, error) {
	if req.GetCardId() == "" || req.GetStatementId() == "" {
		return nil, fmt.Errorf("card_id and statement_id are required")
	}
	alert := &alertsv1.Alert{
		Id:          s.nextID(),
		CardId:      req.GetCardId(),
		StatementId: req.GetStatementId(),
		DueDate:     req.GetDueDate(),
		DueAmount:   req.GetDueAmount(),
		Status:      "scheduled",
	}
	s.store.mu.Lock()
	s.store.alerts[alert.Id] = alert
	s.store.mu.Unlock()
	if s.repo != nil {
		if err := s.repo.CreateAlert(ctx, alert); err != nil {
			return nil, err
		}
	}
	if s.pubs != nil {
		_ = s.pubs.Publish(ctx, "alert.scheduled", alert)
	}
	return alert, nil
}

func (s *Server) GetSummary(ctx context.Context, req *alertsv1.GetSummaryRequest) (*alertsv1.DashboardSummary, error) {
	if s.repo != nil {
		return s.repo.Summary(ctx)
	}
	var totalDue float64
	var overdueCount int32
	now := time.Now()

	s.store.mu.RLock()
	for _, alert := range s.store.alerts {
		totalDue += alert.GetDueAmount()
		if alert.GetDueDate() != "" {
			if due, err := time.Parse("2006-01-02", alert.GetDueDate()); err == nil && due.Before(now) {
				overdueCount++
			}
		}
	}
	s.store.mu.RUnlock()

	return &alertsv1.DashboardSummary{
		TotalDue:       totalDue,
		TotalPaid:      0,
		TotalRemaining: totalDue,
		OverdueCount:   overdueCount,
	}, nil
}

func (s *Server) nextID() string {
	id := atomic.AddUint64(&s.store.idCounter, 1)
	return strconv.FormatUint(id, 10)
}
