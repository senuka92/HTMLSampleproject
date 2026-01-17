FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
COPY services/alerts ./services/alerts
COPY services/alerts/gen ./services/alerts/gen
RUN go build -o /bin/alerts ./services/alerts/cmd/server

FROM alpine:3.19
COPY --from=builder /bin/alerts /bin/alerts
EXPOSE 50054
ENTRYPOINT ["/bin/alerts"]
