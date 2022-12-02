FROM --platform=linux/x86_64 golang:1.18.0-alpine AS builder

WORKDIR /app

COPY ./src/ .

RUN go mod tidy && \
    go build -o main main.go

FROM --platform=linux/x86_64 golang:1.18.0-alpine AS production
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 9090

CMD ["/app/main"]