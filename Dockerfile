# Build Geth in a stock Go builder container
FROM golang:1.11.6-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /go-ethereum

RUN cd /go-ethereum && make geth

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache jq ca-certificates
COPY --from=builder /go-ethereum/build/bin/geth /usr/local/bin/
COPY containers/docker/testnet/run.pls.sh /usr/local/bin/
WORKDIR /usr/local/bin/

EXPOSE 8547 8548 30305 30305/udp
# ENTRYPOINT ["geth"]
