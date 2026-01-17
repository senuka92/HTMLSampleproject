CREATE SCHEMA IF NOT EXISTS payments;

CREATE TABLE IF NOT EXISTS payments.payments (
    id TEXT PRIMARY KEY,
    card_id TEXT NOT NULL,
    statement_id TEXT NOT NULL,
    amount NUMERIC(12, 2) NOT NULL DEFAULT 0,
    paid_at TEXT,
    note TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
