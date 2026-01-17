package cards

import (
	context "context"
	fmt "fmt"

	pgxpool "github.com/jackc/pgx/v5/pgxpool"

	cardsv1 "github.com/example/pfm/services/cards/gen"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) CreateBank(ctx context.Context, bank *cardsv1.Bank) error {
	_, err := r.pool.Exec(ctx, `INSERT INTO cards.banks (id, name) VALUES ($1, $2)`, bank.Id, bank.Name)
	if err != nil {
		return fmt.Errorf("insert bank: %w", err)
	}
	return nil
}

func (r *Repository) CreateCard(ctx context.Context, card *cardsv1.Card) error {
	_, err := r.pool.Exec(ctx, `
		INSERT INTO cards.cards (id, bank_id, name, last_four, due_day)
		VALUES ($1, $2, $3, $4, $5)
	`, card.Id, card.BankId, card.Name, card.LastFour, card.DueDay)
	if err != nil {
		return fmt.Errorf("insert card: %w", err)
	}
	return nil
}

func (r *Repository) CreateStatement(ctx context.Context, statement *cardsv1.Statement) error {
	_, err := r.pool.Exec(ctx, `
		INSERT INTO cards.statements (id, card_id, statement_month, due_date, due_amount)
		VALUES ($1, $2, $3, $4, $5)
	`, statement.Id, statement.CardId, statement.StatementMonth, statement.DueDate, statement.DueAmount)
	if err != nil {
		return fmt.Errorf("insert statement: %w", err)
	}
	return nil
}
