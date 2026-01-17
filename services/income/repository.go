package income

import (
	context "context"
	fmt "fmt"

	pgxpool "github.com/jackc/pgx/v5/pgxpool"

	incomev1 "github.com/example/pfm/services/income/gen"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) UpsertIncome(ctx context.Context, income *incomev1.Income) error {
	_, err := r.pool.Exec(ctx, `
		INSERT INTO income.incomes (id, month, currency, amount, source, note)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO UPDATE SET amount = EXCLUDED.amount, source = EXCLUDED.source, note = EXCLUDED.note
	`, income.Id, income.Month, income.Currency, income.Amount, income.Source, income.Note)
	if err != nil {
		return fmt.Errorf("upsert income: %w", err)
	}
	return nil
}
