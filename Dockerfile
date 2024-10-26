FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . /app
RUN go mod download && go mod verify

RUN go build -o gnsagrent ./cmd/

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/gnsagrent /app/gnsagrent

CMD ["/app/gnsagrent"]