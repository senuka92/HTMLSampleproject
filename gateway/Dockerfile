FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
COPY gateway ./gateway
COPY services/alerts/gen ./services/alerts/gen
COPY services/cards/gen ./services/cards/gen
COPY services/payments/gen ./services/payments/gen
RUN go build -o /bin/gateway ./gateway/cmd/server

FROM alpine:3.19
COPY --from=builder /bin/gateway /bin/gateway
EXPOSE 8081
ENTRYPOINT ["/bin/gateway"]
