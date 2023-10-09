FROM golang:1.21 as base

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app ./cmd/main.go

FROM alpine:latest

COPY --from=base /bin/app /bin/app

CMD ["/bin/app"]
