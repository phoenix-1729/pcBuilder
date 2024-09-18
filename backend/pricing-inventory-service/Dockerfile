FROM golang:1.19

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o pricing-inventory-service main.go

CMD ["/app/pricing-inventory-service"]
