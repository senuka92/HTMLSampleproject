package alerts

import (
	context "context"
	fmt "fmt"
	strconv "strconv"
	strings "strings"

	pgxpool "github.com/jackc/pgx/v5/pgxpool"

	alertsv1 "github.com/example/pfm/services/alerts/gen"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) CreateAlert(ctx context.Context, alert *alertsv1.Alert) error {
	_, err := r.pool.Exec(ctx, `
		INSERT INTO alerts.alerts (id, card_id, statement_id, due_date, due_amount, status)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, alert.Id, alert.CardId, alert.StatementId, alert.DueDate, alert.DueAmount, alert.Status)
	if err != nil {
		return fmt.Errorf("insert alert: %w", err)
	}
	return nil
}

func (r *Repository) Summary(ctx context.Context) (*alertsv1.DashboardSummary, error) {
	row := r.pool.QueryRow(ctx, `
		SELECT COALESCE(SUM(due_amount), 0) AS total_due,
		       SUM(CASE WHEN due_date IS NOT NULL AND due_date <> '' AND due_date < TO_CHAR(NOW(), 'YYYY-MM-DD') THEN 1 ELSE 0 END) AS overdue
		FROM alerts.alerts
	`)
	var totalDue float64
	var overdue int32
	if err := row.Scan(&totalDue, &overdue); err != nil {
		return nil, fmt.Errorf("summary: %w", err)
	}
	return &alertsv1.DashboardSummary{
		TotalDue:       round(totalDue),
		TotalPaid:      0,
		TotalRemaining: round(totalDue),
		OverdueCount:   overdue,
	}, nil
}

func round(value float64) float64 {
	formatted := strconv.FormatFloat(value, 'f', 2, 64)
	parsed, err := strconv.ParseFloat(strings.TrimSpace(formatted), 64)
	if err != nil {
		return value
	}
	return parsed
}
