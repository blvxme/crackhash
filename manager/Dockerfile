FROM golang:latest AS builder
WORKDIR /manager
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o manager cmd/manager/main.go

FROM alpine:latest
WORKDIR /manager
COPY --from=builder /manager/manager .
CMD ["./manager"]
