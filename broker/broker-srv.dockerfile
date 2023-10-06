# The builder stage
FROM golang:1.21.1-alpine3.18 AS builder

WORKDIR /app

COPY go.mod .

RUN go get ./...

COPY . .

RUN go build -o broker ./cmd/main.go

# The runtime stage
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/broker .

COPY ./config /app/config

# default command once container is up 
CMD [ "/app/broker" ]