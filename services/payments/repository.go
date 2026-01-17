package payments

import (
	context "context"
	fmt "fmt"

	pgxpool "github.com/jackc/pgx/v5/pgxpool"

	paymentsv1 "github.com/example/pfm/services/payments/gen"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) CreatePayment(ctx context.Context, payment *paymentsv1.Payment) error {
	_, err := r.pool.Exec(ctx, `
		INSERT INTO payments.payments (id, card_id, statement_id, amount, paid_at, note)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, payment.Id, payment.CardId, payment.StatementId, payment.Amount, payment.PaidAt, payment.Note)
	if err != nil {
		return fmt.Errorf("insert payment: %w", err)
	}
	return nil
}
