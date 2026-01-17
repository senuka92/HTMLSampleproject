FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
COPY pkg ./pkg
COPY services/payments ./services/payments
COPY services/payments/gen ./services/payments/gen
RUN go build -o /bin/payments ./services/payments/cmd/server

FROM alpine:3.19
COPY --from=builder /bin/payments /bin/payments
EXPOSE 50052
ENTRYPOINT ["/bin/payments"]
