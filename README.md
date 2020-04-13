# Plasma EVM Implementation
[![CircleCI](https://circleci.com/gh/Onther-Tech/plasma-evm.svg?style=shield)](https://app.circleci.com/pipelines/github/Onther-Tech/plasma-evm) [![Discord](https://img.shields.io/badge/discord-join%20chat-blue.svg)](https://discord.gg/8wSpJKz)

Implementation of [Plasma EVM](https://tokamak.network). You can check smart contracts [here](https://github.com/Onther-Tech/plasma-evm-contracts). For more information, see [documentations](http://docs.tokamak.network/).

- [Plasma EVM Implementation](#plasma-evm-implementation)
  - [Development Status](#development-status)
  - [Ethereum client](#ethereum-client)
  - [Public Testnet](#public-testnet)
  - [Build](#build)
  - [Run](#run)
  - [Test](#test)
  - [Command-line Options](#command-line-options)
  - [Additional Commands](#additional-commands)
    - [account](#account)
    - [deploy](#deploy)
    - [manage-staking](#manage-staking)
    - [Staking](#staking)

## Development Status
- [x] Make enter / exit requests
- [x] Submit NRBs / ORBs
- [x] Finalize block and requests
- [ ] Challenge on Null Address Transaction in NRBs
- [ ] Continuous Rebase
- [ ] Integration Computation Challenge using [solevm](https://github.com/Onther-Tech/solEVM).

## Ethereum client

For the simple start of both clients(ethereum, plasma-evm), you can run `run.rootchain.sh` in [Onther-Tech/go-ethereum](https://github.com/Onther-Tech/go-ethereum) and `run.pls.sh` in [Onther-Tech/plasma-evm](https://github.com/Onther-Tech/plasma-evm) (this is recommended way to setup develop and test environment).

## Public Testnet

Faraday testnet is running based on Rinkeby network as root chain.

See [Onther-Tech/plasma-evm-networks](https://github.com/Onther-Tech/plasma-evm-networks)

## Build

```bash
$ make geth
```

## Run

Before running plasma-evm client, you **MUST** run another geth client as root chain client.

```bash
$ git clone https://github.con/Onther-Tech/go-ethereum.git root-chain-geth && cd root-chain-geth
$ bash run.rootchain.sh
```

Then you can run  or test plasma-evm client.

```bash
$ git clone https://github.con/Onther-Tech/plasma-evm.git && cd plasma-evm
$ bash run.pls.sh
```

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
  --dev                               Ephemeral proof-of-authority network with a pre-funded developer account, mining enabled
  --dev.period value                  Block period to use in developer mode (0 = mine only if transaction pending) (default: 0)
  --dev.key value                     Comma seperated developer account key as hex(for dev)

PLASMA EVM - OPERATOR OPTIONS:
  --operator value                    Plasma operator address as hex.
  --operator.key value                Plasma operator key as hex(for dev)
  --operator.password value           Operator password file to use for non-interactive password input
  --operator.minether value           Plasma operator minimum balance (default = 0.5 ether) (default: "0.5")
  --miner.recommit value              Time interval to recreate the block being mined (default: 3s)

PLASMA EVM - ROOTCHAIN TRANSACTION MANAGER OPTIONS:
  --tx.gasprice value                 Gas price for transaction (default = 10 Gwei) (default: 0)
  --tx.mingasprice value              Minimum gas price for submitting a block (default = 1 Gwei) (default: 1000000000)
  --tx.maxgasprice value              Maximum gas price for submitting a block (default = 100 Gwei) (default: 100000000000)
  --tx.interval value                 Pending interval time after submitting a block (default = 10s). If block submit transaction is not mined in 2 intervals, gas price will be adjusted. See https://golang.org/pkg/time/#ParseDuration (default: 10s)

PLASMA EVM - STAMINA OPTIONS:
  --stamina.operatoramount value      Operator stamina amount at genesis block in ETH (default: 1)
  --stamina.mindeposit value          Minimum deposit amount in ETH (default: 0.5)
  --stamina.recoverepochlength value  The length of recovery epoch in block (default: 120960)
  --stamina.withdrawaldelay value     Withdrawal delay in block (default: 362880)

PLASMA EVM - CHALLENGER OPTIONS:
  --rootchain.challenger value        Address of challenger account
  --challenger.password value         Challenger password file to use for non-interactive password input

PLASMA EVM - ROOTCHAIN CONTRACT OPTIONS:
  --rootchain.url value               JSONRPC endpoint of rootchain provider. If URL is empty, ignore the provider.
  --rootchain.contract value          Address of the RootChain contract

PLASMA EVM - STAKING OPTIONS OPTIONS:
  --unlock value                      Comma separated list of accounts to unlock
  --password value                    Password file to use for non-interactive password input
  --rootchain.sender value            Address of root chain transaction sender account. it MUST be unlocked by --unlock, --password flags (CAVEAT: To set plasma operator, use --operator flag)
  --rootchain.gasPrice value          Transaction gas price to root chain in GWei (default: 10000000000)
  --rootchain.ton value               Address of TON token contract
  --rootchain.wton value              Address of WTON token contract
  --rootchain.registry value          Address of RootChainRegistry contract
  --rootchain.depositManager value    Address of Deposit Manager contract
  --rootchain.seigManager value       Address of SeigManager contract
  --rootchain.powerton value          Address of PowerTON contract
```

## Additional Commands
For more information, run below command (and sub-command) with `--help` flag.

### account

```bash
$ geth account importKey <privateKey>            # Import a private key from hex key into a new account
$ geth account importHDwallet <mnemonic> <path>  # Import a mnemonic into a new account
```

### deploy
```bash
$ geth deploy <genesisPath> <chainId> <withPETH> <NRELength>  # Deploy RootChain contract and make genesis file
```

### manage-staking

```bash
$ geth manage-staking deployManagers <withdrawalDelay> <seigPerBlock>  # Deploy staking manager contracts (except PowerTON)
$ geth manage-staking deployPowerTON <roundDuration>                   # Deploy PowerTON contract
$ geth manage-staking startPowerTON                                    # Start PowerTON first round
$ geth manage-staking register                                         # Register RootChain contract
$ geth manage-staking getManagers <path?>                              # Get staking managers addresses in database
$ geth manage-staking setManagers <uri>                                # Set staking managers addresses in database
$ geth manage-staking mintTON <to> <amount>                            # Mint TON to account (for dev)
```

### Staking

```bash
$ geth staking balances <address>                               # Print balances of token and stake
$ geth staking swapFromTON <tonAmount>                          # Swap TON to WTON
$ geth staking swapToTON <wtonAmount>                           # Swap WTON to TON
$ geth staking stake <amount>                                   # Stake WTON
$ geth staking requestWithdrawal <amount?>                      # Make a withdrawal request
$ geth staking processWithdrawal <numRequests?>                 # Process pending withdrawals
```
