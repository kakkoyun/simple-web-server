FROM golang:1.12-alpine AS builder

RUN apk add --update git ca-certificates && update-ca-certificates

ADD . /build/
WORKDIR /build

RUN go mod vendor -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -a -ldflags '-s -w' -o main

FROM scratch

COPY --from=builder /build/main /app/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app
CMD ["./main"]
