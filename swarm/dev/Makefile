.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/Onther-Tech/plasma-evm/cmd/swarm
	go build -o bin/geth     github.com/Onther-Tech/plasma-evm/cmd/geth
	go build -o bin/bootnode github.com/Onther-Tech/plasma-evm/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/Onther-Tech/plasma-evm/swarm/...
