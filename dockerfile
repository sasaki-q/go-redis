FROM golang:1.18.0-alpine

ENV GOPATH /go
ENV GO111MODULE on

RUN apk update && \
    apk --no-cache add git

WORKDIR /go/src

COPY ./src/ /go/src/

RUN go mod tidy && \
    go install github.com/cosmtrek/air@v1.27.3

ENV GIN_MODE=debug
ENV PORTS=9090
ENV DB_PORT=5432
ENV DB_HOST=cache-database
ENV DB_NAME=cache-database
ENV DB_USER=postgres
ENV DB_PASSWORD=postgres

EXPOSE 9090

CMD ["air", "-c", ".air.conf"]