# Architecture Blueprint

## Goals

- Free/low-cost services with minimal operational overhead.
- Clear service boundaries with gRPC internal contracts.
- REST/JSON public API for broad client compatibility.
- Event-driven workflow for alerts and dashboards.

## Services (initial)

- **cards**: banks, cards, statements.
- **payments**: payments and allocations.
- **income**: income records (salary + other sources).
- **alerts**: reminders and alert scheduling.

## Event topics (NATS JetStream)

- `card.created`
- `bank.created`
- `statement.created`
- `payment.created`
- `income.upserted`
- `alert.scheduled`
- `alert.sent`

## Data strategy

- Shared Postgres instance.
- One schema per service, owned by that service.
- Migrate to per-service databases later if desired.

## Public API (REST)

- `POST /banks`
- `POST /cards`
- `POST /statements`
- `POST /payments`
- `POST /alerts/schedule`
- `GET /dashboard/summary`
- `POST /auth/refresh`

## Internal API (gRPC)

Each service owns a gRPC API defined in `/proto/*.proto`.
