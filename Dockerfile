FROM golang:1.26.4-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/gopportunities ./main.go

FROM alpine:3.22

RUN apk add --no-cache ca-certificates \
	&& addgroup -S app \
	&& adduser -S -G app app

WORKDIR /app

ENV GIN_MODE=release \
	DB_HOST=postgres \
	DB_PORT=5432 \
	DB_USER=postgres \
	DB_NAME=app \
	DB_SSLMODE=disable

COPY --from=builder /out/gopportunities ./gopportunities
COPY .env.example ./.env

USER app

EXPOSE 8080

ENTRYPOINT ["./gopportunities"]
