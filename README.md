# Personal Finance Manager (Microservices)

This repository scaffolds a hobby-grade, free-first microservices architecture for a personal finance manager.
The focus is on a simple local developer experience and clear service boundaries while keeping costs low.

## Stack (locked)

- **External API:** REST/JSON (gateway/BFF)
- **Internal API:** gRPC
- **Auth:** Keycloak (self-hosted)
- **Auth flow:** BFF handles refresh tokens
- **Database:** Shared Postgres with strict schemas per service
- **Event bus:** NATS JetStream
- **Dev/hosting:** Docker Compose locally → single VPS later

## Repository layout

```
/apps
  /flutter
/docs
/gateway
  /bff
  /rest
/infra
/proto
/services
  /alerts
  /cards
  /income
  /payments
```

## Quick start (infra + services)

> This spins up Postgres, NATS JetStream, Keycloak, the gateway, and all gRPC services.

```bash
docker compose -f infra/docker-compose.yml up
```

Keycloak admin UI (dev): http://localhost:8080
Gateway (REST): http://localhost:8081

## Local development notes

- gRPC stubs live under `services/*/gen`. These are placeholders until `protoc` and the Go plugins are available; regenerate them with `make proto` when tooling is installed.
- Service ports: cards `50051`, payments `50052`, income `50053`, alerts `50054`.

## Next steps (implementation)

- Add Go service skeletons and generated gRPC stubs.
- Implement the first vertical slice: bank → card → statement → payment → dashboard → alerts.
- Add the REST/BFF gateway with token management and service routing.

See `docs/architecture.md` for the full blueprint.
For the implementation checklist, see `docs/next-steps.md`.
For deployment configuration, see `docs/deployment.md`.
