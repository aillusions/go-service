FROM golang:1.9.2
WORKDIR /go/src/go-service/app/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM scratch
COPY --from=0 /go/src/go-service/app/app .
ENTRYPOINT ["/app"]