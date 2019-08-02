FROM golang:1.12-alpine3.10 AS builder

RUN apk update && apk add git ca-certificates

WORKDIR /go/src/github.com/martinsirbe/prometheus-graphite-demo
COPY . .

ENV GO111MODULE=on
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o demo ./cmd/prometheus-graphite-demo

FROM alpine:3.10

COPY --from=builder /go/src/github.com/martinsirbe/prometheus-graphite-demo/demo /demo

ENTRYPOINT ["/demo"]
