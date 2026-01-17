# Deployment Guide

This guide explains how to run the stack locally and what environment variables are required for deployment.

## Required services

- Postgres
- NATS JetStream
- Keycloak (OIDC)
- Gateway
- Cards/Payments/Income/Alerts services

## Environment variables

### Gateway

- `GATEWAY_ADDR` (default `:8081`)
- `AUTH_REQUIRED` (`true`/`false`)
- `KEYCLOAK_JWKS_URL` (e.g. `http://localhost:8080/realms/pfm/protocol/openid-connect/certs`)
- `KEYCLOAK_TOKEN_URL` (e.g. `http://localhost:8080/realms/pfm/protocol/openid-connect/token`)
- `KEYCLOAK_CLIENT_ID`
- `KEYCLOAK_CLIENT_SECRET` (optional for public clients)

### Services (shared patterns)

- `*_GRPC_ADDR` for listen port
- `*_DATABASE_URL` for Postgres DSN
- `*_MIGRATIONS` for local migration directory (defaults to `services/<service>/migrations`)
- `NATS_URL` for JetStream publishing

## Local run

```bash
docker compose -f infra/docker-compose.yml up --build
```

## Deployment notes

- Run the gateway behind HTTPS and configure `AUTH_REQUIRED=true` for production.
- Use per-service DB users or schemas as isolation boundaries.
- Configure Keycloak realm/clients before enabling auth enforcement.
