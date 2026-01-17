CREATE SCHEMA IF NOT EXISTS alerts;

CREATE TABLE IF NOT EXISTS alerts.alerts (
    id TEXT PRIMARY KEY,
    card_id TEXT NOT NULL,
    statement_id TEXT NOT NULL,
    due_date TEXT,
    due_amount NUMERIC(12, 2) NOT NULL DEFAULT 0,
    status TEXT NOT NULL DEFAULT 'scheduled',
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
