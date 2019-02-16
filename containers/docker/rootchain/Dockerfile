# Build Geth in a stock Go builder container
FROM golang:1.11-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git ca-certificates && \
    git clone --depth 1 https://github.com/Onther-Tech/go-ethereum.git && \
    cd go-ethereum && make geth

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

COPY run.rootchain.sh /usr/local/bin/
COPY --from=builder /go/go-ethereum/build/bin/geth /usr/local/bin/
WORKDIR /usr/local/bin
