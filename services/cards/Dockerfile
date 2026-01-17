FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
COPY services/cards ./services/cards
COPY services/cards/gen ./services/cards/gen
RUN go build -o /bin/cards ./services/cards/cmd/server

FROM alpine:3.19
COPY --from=builder /bin/cards /bin/cards
EXPOSE 50051
ENTRYPOINT ["/bin/cards"]
