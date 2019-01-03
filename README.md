## Plasma EVM PoC Implementation

Proof-of-concept Implementation of [Plasma EVM](https://hackmd.io/s/HyZ2ms8EX), forked from [Onther-Tech/plasma-evm](https://github.com/Onther-Tech/plasma-evm). You can check the RootChain smart contract [here](https://github.com/Onther-Tech/plasma-evm-contracts).


## Development Status
- [x] Make enter / exit requests
- [x] Submit NRBs / ORBs
- [x] Finalize block and requests
- [ ] Challenge on Null Address Transaction in NRBs

Implementing TrueBit's verification game is under research as well as challenges that requires verification game.


## Ethereum client
Building geth requires both a Go (version 1.9 or later) and a C compiler.
You can install them using your favourite package manager.
Once the dependencies are installed, run

We're using [Onther-Tech/go-ethereum](https://github.com/Onther-Tech/go-ethereum) as an ethereum client instead of `ganache-cli` due to the support for WebSocket JSONRPC endpoint. Especially, the client is modified with respect to genesis and developer accounts setup like `ganache-cli`.

For the simple start of both clients(ethereum, plasma-evm), you can run `run.rootchain.sh` in [Onther-Tech/go-ethereum](https://github.com/Onther-Tech/go-ethereum) and `run.pls.sh` in [Onther-Tech/plasma-evm](https://github.com/Onther-Tech/plasma-evm) (this is recommended way to setup devlop and test plasma-evm).


## Build

Just run `make geth` in terminal.

## Command-line Options

Here is some additional command-line options for Plasma-EVM.

```bash
ROOTCHAIN OPTIONS:
  --rootchain.operatorKey   Specify operator key as hex
  --rootchain.contract      The address of RootChain contract
  --rootchain.url           JSONRPC endpoint of rootchain provider. Use WebSocket to subscribe events. (default: ws://localhost:8546)

DEVELOPER CHAIN OPTIONS:
  --dev.key                 Comma seperated keys as hex for developer accounts
```

## Test

Some original `geth` tests may fail. You can just test plasam-evm related feature by running
```bash
$ go test pls/rootchain_manager_test.go
```
