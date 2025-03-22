FROM golang:latest AS builder
WORKDIR /worker
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o worker cmd/worker/main.go

FROM alpine:latest
WORKDIR /worker
COPY --from=builder /worker/worker .
CMD ["./worker"]
