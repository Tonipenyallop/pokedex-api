FROM golang:1.24.2-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o docker-pokedex-api

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/docker-pokedex-api .
COPY .env .

EXPOSE 8080
CMD ["./docker-pokedex-api"]