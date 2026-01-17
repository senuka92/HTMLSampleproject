CREATE SCHEMA IF NOT EXISTS income;

CREATE TABLE IF NOT EXISTS income.incomes (
    id TEXT PRIMARY KEY,
    month TEXT NOT NULL,
    currency TEXT NOT NULL,
    amount NUMERIC(12, 2) NOT NULL DEFAULT 0,
    source TEXT,
    note TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
