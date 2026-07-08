FROM golang:1.26.4-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /bin/api ./cmd/api

FROM alpine:3.22

RUN addgroup -S app && adduser -S -G app app
WORKDIR /app

COPY --from=builder /bin/api /app/api

USER app
EXPOSE 8080

ENTRYPOINT ["/app/api"]
