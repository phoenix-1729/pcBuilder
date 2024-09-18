# Stage 1: Build the Go binary
FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o component-service main.go

# Stage 2: Create the final image
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/component-service .

EXPOSE 50051

CMD ["./component-service"]
