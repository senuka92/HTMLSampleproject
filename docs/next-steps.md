# Next Steps Roadmap

This checklist tracks the recommended sequence to move the scaffold into a working vertical slice.

## 1) Generate gRPC code

- [ ] Install `protoc` and the Go plugins (`protoc-gen-go`, `protoc-gen-go-grpc`).
- [ ] Run `make proto` to regenerate stubs in `services/*/gen`.
- [ ] Remove placeholder stubs once generated code is in place.

## 2) Wire persistence into services

- [ ] Add database configuration (DSN/env vars) per service.
- [ ] Run SQL migrations from `services/*/migrations` at startup.
- [ ] Replace in-memory stores with Postgres-backed repositories.

## 3) Finish the REST/BFF layer

- [ ] Validate JWT access tokens against Keycloak.
- [ ] Add refresh-token handling (BFF pattern).
- [ ] Enforce auth/authorization per route.

## 4) Implement the full vertical slice

- [ ] Bank creation
- [ ] Card creation
- [ ] Statement creation
- [ ] Payment creation
- [ ] Dashboard summary
- [ ] Alert scheduling

## 5) Add event publishing

- [ ] Publish lifecycle events to NATS JetStream.
- [ ] Consume events for alert scheduling and dashboard projection.

## 6) Improve local developer experience

- [ ] Add health checks in `docker-compose`.
- [ ] Add seed data script.
- [ ] Add curl examples for smoke testing.
