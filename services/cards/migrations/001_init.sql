CREATE SCHEMA IF NOT EXISTS cards;

CREATE TABLE IF NOT EXISTS cards.banks (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS cards.cards (
    id TEXT PRIMARY KEY,
    bank_id TEXT NOT NULL REFERENCES cards.banks(id),
    name TEXT NOT NULL,
    last_four TEXT,
    due_day TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS cards.statements (
    id TEXT PRIMARY KEY,
    card_id TEXT NOT NULL REFERENCES cards.cards(id),
    statement_month TEXT NOT NULL,
    due_date TEXT,
    due_amount NUMERIC(12, 2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
