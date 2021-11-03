FROM golang:1.17.1 as builder

RUN mkdir /build
WORKDIR /build

COPY go.mod .
COPY go.sum .


ADD . .

RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -gcflags=-trimpath=$PWD -o main .

FROM alpine:3.12.0

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

ENV GOROOT /usr/local/go
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip

# RUN mkdir -p /app && adduser -S -D -H -h /app appuser && chown -R appuser /app
COPY --from=builder /build/main /app/
# COPY --from=builder /build/config.toml  /app/config/
# USER appuser
EXPOSE 8081
WORKDIR /app
CMD ["./main"]
