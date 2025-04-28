FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

RUN CGO_ENABLED=0 GOOS=linux go build -o seed src/cmd/seed/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/seed .

EXPOSE 8080

CMD sh -c "./seed && ./main" 