FROM golang:1.11.5-alpine3.9

COPY service.go .

RUN go build -o /app

ENTRYPOINT ["/app"]