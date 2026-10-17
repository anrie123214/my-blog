
# Stage 1: Build the Go application
FROM golang:1.26.5-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /out/site .

# Stage 2: Runtime image
FROM alpine:3.22

RUN apk add --no-cache ca-certificates \
    && adduser -D -u 10001 appuser

WORKDIR /app

COPY --from=builder /out/site ./site
COPY templates/ ./templates/
COPY static/ ./static/
COPY docs/ ./docs/

ENV PORT=8080

USER appuser

EXPOSE 8080

CMD ["./site"]
