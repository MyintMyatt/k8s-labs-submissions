FROM golang:1.27 AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o log_reader .

FROM alpine:3.22
WORKDIR /app
COPY --from=builder /app/log_reader .

CMD ["./log_reader"]
