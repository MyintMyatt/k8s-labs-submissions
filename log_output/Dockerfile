FROM golang:1.27 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o log-output .

FROM alpine:3.22
WORKDIR /app
COPY --from=builder /app/log-output .

CMD ["./log-output"]