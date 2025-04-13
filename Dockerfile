# Build
FROM golang:1.24.1-alpine AS builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN GOOS=linux go build -o /entrypoint ./main.go

# Production
FROM ubuntu:24.04
WORKDIR /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /entrypoint /entrypoint
ENTRYPOINT [ "/entrypoint" ]