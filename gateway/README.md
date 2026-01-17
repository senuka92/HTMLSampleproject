# Gateway

Placeholder for the REST/BFF gateway.

Planned responsibilities:

- Public REST endpoints.
- Token validation via Keycloak.
- Refresh token handling (BFF).
- Route requests to internal gRPC services.

## Local endpoints (initial)

- `POST /banks`
- `POST /cards`
- `POST /statements`
- `POST /payments`
- `POST /alerts/schedule`
- `GET /dashboard/summary`
