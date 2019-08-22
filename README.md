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

## Run

Before running plasma-evm client, you **MUST** run another geth client as root chain client.

```bash
$ git clone https://github.con/Onther-Tech/go-ethereum.git root-chain-geth && cd root-chain-geth
$ bash run.rootchain.sh
```
Then you can run plasma-evm client or test.

### Development Network
Just run
```bash
bash run.pls.sh
```

### Custom network

#### 1. Deploy RootChain contract and make genesis file.
All paramaeters are based on ethereum test network run by `run.rootchain.sh`.

```bash
$ rm -rf ~/.plasma-evm

$ geth deploy ./genesis.json 1019 true 1024 \
  --datadir ~/.plasma-evm \
  --rootchain.url ws://127.0.0.1:8546 \
  --operator.key b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291
```

#### 2. Run Plasma EVM client as operator mode.
```bash
$ geth \
  --datadir ~/.plasma-evm \
  --rpc \
  --rpcport 8547 \
  --miner.etherbase 0x71562b71999873DB5b286dF957af199Ec94617F7 \
  --operator 0x71562b71999873DB5b286dF957af199Ec94617F7 \
  --nodekeyhex b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f49b\
  --dev \
  --dev.key 78ae75d1cd5960d87e76a69760cb451a58928eee7890780c352186d23094a114 \
  --rootchain.challenger 0x3616BE06D68dD22886505e9c2CaAa9EcA84564b8 \
  --rootchain.url ws://127.0.0.1:8546 \
  --tx.interval "300ms"
```
This is based on [`run.pls.sh`](./run.pls.sh)

##### CAVEAT: Challenger key is in step 2. So when restart node, you should not use `--dev.key` flag in step 2 not to import key again.

#### 3. (Optional) Run another node as user mode.
```bash
$ geth init genesis.json \
  --datadir ~/.plasma-evm-user

$ geth \
  --datadir ~/.plasma-evm-user \
  --rpc \
  --rpcport 8548 \
  --rootchain.url ws://127.0.0.1:8546 \
  --bootnodes enode://a8d63d4760dbcd8656b8a61f28eb246c2f990949cd7a5d2efe483caa05a5381a61cfb30eb8622f4cb0909a6e94a82b7505fb5b2cd737398860f9369b3a8522ca@127.0.0.1:30305?discport=30305
```

Then use JSONRPC console to connect two nodes.

## Test

Some original go-ethereum tests may fail.


To test plasam-evm related features,
```bash
$ go test github.com/Onther-Tech/plasma-evm/pls
$ go test github.com/Onther-Tech/plasma-evm/tx
```


## Command-line Options

`build/bin/geth help` returns available flags. Belows are plasma-evm-specific flags.

```bash
PLASMA EVM - DEVELOPMENT MODE OPTIONS:
  --dev               Ephemeral proof-of-authority network with a pre-funded developer account, mining enabled
  --dev.period value  Block period to use in developer mode (0 = mine only if transaction pending) (default: 0)
  --dev.key value     Comma seperated developer account key as hex(for dev)

PLASMA EVM - OPERATOR OPTIONS:
  --operator value           Plasma operator address as hex. The account should be unlock by using --unlock
  --operator.key value       Plasma operator key as hex(for dev)
  --operator.minether value  Plasma operator minimum balance (default = 0.5 ether) (default: "0.5")
  --miner.recommit value     Time interval to recreate the block being mined (default: 3s)

PLASMA EVM - ROOTCHAIN TRANSACTION MANAGER OPTIONS:
  --tx.gasprice "0"                Gas price for transaction (default = 10 Gwei)
  --tx.mingasprice "1000000000"    Minimum gas price for submitting a block (default = 1 Gwei)
  --tx.maxgasprice "100000000000"  Maximum gas price for submitting a block (default = 100 Gwei)
  --tx.interval value              Pending interval time after submitting a block (default = 10s). If block submit transaction is not mined in 2 intervals, gas price will be adjusted. See https://golang.org/pkg/time/#ParseDuration (default: 10s)

PLASMA EVM - STAMINA OPTIONS:
  --stamina.mindeposit "500000000000000000"  MinDeposit variable state of stamina contract
  --stamina.recoverepochlength "10080"       RecoverEpochLength variable state of stamina contract
  --stamina.withdrawaldelay "30240"          WithdrawalDelay variable state of stamina contract

PLASMA EVM - CHALLENGER OPTIONS:
  --rootchain.challenger value  Address of challenger account

PLASMA EVM - ROOTCHAIN CONTRACT OPTIONS:
  --rootchain.url value       JSONRPC endpoint of rootchain provider. If URL is empty, ignore the provider.
  --rootchain.contract value  Address of the RootChain contract
```

## Additional Commands
For more information, run below command with `--help` flag.

### Account

#### Import Private Key
```bash
geth account importKey <privateKey>
```

#### Import HDWallet
```bash
geth account importHDwallet <mnemonic> <path>
```

### Deploy
#### Deploy RootChain contract and make genesis.json file
```bash
geth deploy <genesisPath> <chainId> <withPETH> <NRELength>
```

