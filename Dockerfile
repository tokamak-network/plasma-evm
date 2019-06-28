# Build Geth in a stock Go builder container
FROM golang:1.12-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

ADD . /go-ethereum
RUN cd /go-ethereum && make geth

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-ethereum/build/bin/geth /usr/local/bin/
WORKDIR /usr/local/bin/

EXPOSE 8547 8548 30305 30305/udp
# ENTRYPOINT ["geth"]
