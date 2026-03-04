FROM golang:1.25-alpine AS builder
LABEL authors="aaron"
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main ./main.go

FROM alpine:latest
RUN apk add --no-cache tzdata
WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]