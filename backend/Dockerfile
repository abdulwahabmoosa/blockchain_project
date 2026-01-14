# Use Go 1.25 for Ziren library compatibility
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
# Install certificates for HTTPS connections to blockchain nodes
RUN apk --no-cache add ca-certificates
EXPOSE 3000
CMD ["./main"]
