# Multi-stage Dockerfile for Go backend with distroless
FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o tic-tac-toe-backend main.go

# Use distroless for minimal runtime image
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/tic-tac-toe-backend .
EXPOSE 9090
USER nonroot:nonroot
ENTRYPOINT ["/app/tic-tac-toe-backend"]
