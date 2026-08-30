FROM golang:1.27 AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o todo_app .

FROM alpine:3.22
WORKDIR /app

COPY --from=builder /app/todo_app .

CMD [ "./todo_app" ]