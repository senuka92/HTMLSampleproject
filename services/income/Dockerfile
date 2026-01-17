FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
COPY pkg ./pkg
COPY services/income ./services/income
COPY services/income/gen ./services/income/gen
RUN go build -o /bin/income ./services/income/cmd/server

FROM alpine:3.19
COPY --from=builder /bin/income /bin/income
EXPOSE 50053
ENTRYPOINT ["/bin/income"]
