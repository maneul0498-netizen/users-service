FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

RUN go build -o app .

FROM debian:stable-slim

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]