FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .

# Build for Linux AMD64 explicitly
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o aggregator ./aggregator/cmd/aggregator

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/aggregator .

CMD ["./aggregator"]