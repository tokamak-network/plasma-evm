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
$ git clone https://github.com/Onther-Tech/go-ethereum.git root-chain-geth && cd root-chain-geth
$ bash run.rootchain.sh
```

Then you can run  or test plasma-evm client.

```bash
$ git clone https://github.com/Onther-Tech/plasma-evm.git && cd plasma-evm
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
  --rootchain.deploygasprice value    Transaction gas price to deploy rootchain in GWei (default: 10000000000). This flag applies only to deploy command.

PLASMA EVM - STAKING OPTIONS:
  --unlock value                      Comma separated list of accounts to unlock
  --password value                    Password file to use for non-interactive password input
  --rootchain.sender value            Address of root chain transaction sender account. it MUST be unlocked by --unlock, --password flags (CAVEAT: To set plasma operator, use --operator flag)
  --rootchain.gasprice value          Transaction gas price to root chain in GWei (default: 10000000000)
  --rootchain.ton value               Address of TON token contract
  --rootchain.wton value              Address of WTON token contract
  --rootchain.registry value          Address of RootChainRegistry contract
  --rootchain.depositmanager value    Address of Deposit Manager contract
  --rootchain.seigmanager value       Address of SeigManager contract
  --rootchain.powerton value          Address of PowerTON contract
  
PLASMA EVM - CHILDCHAIN OPTIONS:
  --childchain.url value              JSONRPC endpoint of child chain provider.
  --childchain.sender value           Sender address of child chain transaction
  --childchain.gasprice value         Gas price for child chain transaction in GWei (default: 0)
```

## Additional Commands
For more information, run below command (and sub-command) with `--help` flag.

### account

```bash
$ geth account import-key <privateKey>    # Import a private key from hex key into a new account

ETHEREUM OPTIONS:
  --datadir value       Data directory for the databases and keystore (default: "/Users/thomashin/Library/Ethereum")
  --keystore value      Directory for the keystore (default = inside the datadir)
  --lightkdf            Reduce key-derivation RAM & CPU usage at some expense of KDF strength

ACCOUNT OPTIONS:
  --password value      Password file to use for non-interactive password input
```

```bash
$ geth account import-hdwallet <mnemonic> <path>    # Import a mnemonic into a new account

ETHEREUM OPTIONS:
  --datadir value       Data directory for the databases and keystore (default: "/Users/thomashin/Library/Ethereum")
  --keystore value      Directory for the keystore (default = inside the datadir)
  --lightkdf            Reduce key-derivation RAM & CPU usage at some expense of KDF strength

ACCOUNT OPTIONS:
  --password value      Password file to use for non-interactive password input
```

### deploy
```bash
$ geth deploy <genesisPath> <chainId> <withPETH> <NRELength>    # Deploy RootChain contract and make genesis file

ETHEREUM OPTIONS:
  --datadir value                       Data directory for the databases and keystore (default: "/Users/thomashin/Library/Ethereum")

ACCOUNT OPTIONS:
  --unlock value                        Comma separated list of accounts to unlock
  --password value                      Password file to use for non-interactive password input

PLASMA EVM - DEVELOPMENT MODE OPTIONS:
  --dev.key value                       Comma seperated developer account key as hex(for dev)

PLASMA EVM - STAMINA OPTIONS:
  --stamina.operatoramount value        Operator stamina amount at genesis block in ETH (default: 1)
  --stamina.mindeposit value            Minimum deposit amount in ETH (default: 0.5)
  --stamina.recoverepochlength value    The length of recovery epoch in block (default: 120960)
  --stamina.withdrawaldelay value       Withdrawal delay in block (default: 362880)

PLASMA EVM - ROOTCHAIN CONTRACT OPTIONS:
  --rootchain.url value                 JSONRPC endpoint of rootchain provider. If URL is empty, ignore the provider.
  --rootchain.deploygasprice value      Transaction gas price to deploy rootchain in GWei (default: 10000000000). This flag applies only to deploy command.

PLASMA EVM - STAKING OPTIONS OPTIONS:
  --rootchain.sender value              Address of root chain transaction sender account. it MUST be unlocked by --unlock, --password flags (CAVEAT: To set plasma operator, use --operator flag)
```

### stamina
```bash
$ geth stamina get-delegatee <address>    # Get delegatee of account

PLASMA EVM - CHILDCHAIN OPTIONS OPTIONS:
  --childchain.url value      JSONRPC endpoint of child chain provider.
```

```bash
$ geth stamina get-stamina <address>    # Get stamina of account

PLASMA EVM - CHILDCHAIN OPTIONS OPTIONS:
  --childchain.url value      JSONRPC endpoint of child chain provider.
```

```bash
$ geth stamina get-totaldeposit <address>    # Get total deposit of account

PLASMA EVM - CHILDCHAIN OPTIONS OPTIONS:
  --childchain.url value      JSONRPC endpoint of child chain provider.
```

```bash
$ geth stamina get-deposit <depositor> <delegatee>    # Get deposit of account from the depositor

PLASMA EVM - CHILDCHAIN OPTIONS OPTIONS:
  --childchain.url value      JSONRPC endpoint of child chain provider.
```

```bash
$ geth stamina get-lastrecoveryblock <delegatee>    # Get last recovery block of the delegatee

PLASMA EVM - CHILDCHAIN OPTIONS OPTIONS:
  --childchain.url value      JSONRPC endpoint of child chain provider.
```

```bash
$ geth stamina get-withdrawal <depositor>    # Get withdrawal requests

PLASMA EVM - CHILDCHAIN OPTIONS OPTIONS:
  --childchain.url value      JSONRPC endpoint of child chain provider.
```

```bash
$ geth stamina set-delegator <delegator>    # Set delegator

ETHEREUM OPTIONS:
--datadir value                    Data directory for the databases and keystore (default: "/Users/thomashin/Library/Ethereum")

ACCOUNT OPTIONS:
  --unlock value                   Comma separated list of accounts to unlock
  --password value                 Password file to use for non-interactive password input

PLASMA EVM - CHILDCHAIN OPTIONS OPTIONS:
  --childchain.url value           JSONRPC endpoint of child chain provider.
  --childchain.sender value        Sender address of child chain transaction
  --childchain.gasprice value      Gas price for child chain transaction in GWei (default: 0)
```

```bash
$ geth stamina deposit <delegatee> <value>    # Deposit PETH to gain stamina

ETHEREUM OPTIONS:
  --datadir value                  Data directory for the databases and keystore (default: "/Users/thomashin/Library/Ethereum")

ACCOUNT OPTIONS:
  --unlock value                   Comma separated list of accounts to unlock
  --password value                 Password file to use for non-interactive password input

PLASMA EVM - CHILDCHAIN OPTIONS OPTIONS:
  --childchain.url value           JSONRPC endpoint of child chain provider.
  --childchain.sender value        Sender address of child chain transaction
  --childchain.gasprice value      Gas price for child chain transaction in GWei (default: 0)
```

```bash
$ geth stamina request-withdrawal <delegatee> <value>    # Request withdraw

ETHEREUM OPTIONS:
  --datadir value                  Data directory for the databases and keystore (default: "/Users/thomashin/Library/Ethereum")

ACCOUNT OPTIONS:
  --unlock value                   Comma separated list of accounts to unlock
  --password value                 Password file to use for non-interactive password input

PLASMA EVM - CHILDCHAIN OPTIONS OPTIONS:
  --childchain.url value           JSONRPC endpoint of child chain provider.
  --childchain.sender value        Sender address of child chain transaction
  --childchain.gasprice value      Gas price for child chain transaction in GWei (default: 0)
```

```bash
$ geth stamina withdraw    # Process withdraw

ETHEREUM OPTIONS:
  --datadir value                  Data directory for the databases and keystore (default: "/Users/thomashin/Library/Ethereum")

ACCOUNT OPTIONS:
  --unlock value                   Comma separated list of accounts to unlock
  --password value                 Password file to use for non-interactive password input

PLASMA EVM - CHILDCHAIN OPTIONS OPTIONS:
  --childchain.url value           JSONRPC endpoint of child chain provider.
  --childchain.sender value        Sender address of child chain transaction
  --childchain.gasprice value      Gas price for child chain transaction in GWei (default: 0)                              
```

### manage-staking

```bash
$ geth manage-staking deploy-managers <withdrawalDelay> <seigPerBlock>    # Deploy staking manager contracts (except PowerTON)

ETHEREUM OPTIONS:
  --datadir value                       Data directory for the databases and keystore (default: "/Users/thomashin/Library/Ethereum")

ACCOUNT OPTIONS:
  --unlock value                        Comma separated list of accounts to unlock
  --password value                      Password file to use for non-interactive password input

PLASMA EVM - DEVELOPMENT MODE OPTIONS:
  --dev.key value                       Comma seperated developer account key as hex(for dev)

PLASMA EVM - ROOTCHAIN CONTRACT OPTIONS:
  --rootchain.url value                 JSONRPC endpoint of rootchain provider. If URL is empty, ignore the provider.

PLASMA EVM - STAKING OPTIONS OPTIONS:
  --rootchain.sender value              Address of root chain transaction sender account. it MUST be unlocked by --unlock, --password flags (CAVEAT: To set plasma operator, use --operator flag)
  --rootchain.ton value                 Address of TON token contract
  --rootchain.wton value                Address of WTON token contract
  --rootchain.gasprice value            Transaction gas price to root chain in GWei (default: 10000000000)
```

```bash
$ geth manage-staking deploy-powerton <roundDuration>    # Deploy PowerTON contract

ETHEREUM OPTIONS:
  --datadir value                       Data directory for the databases and keystore (default: "/Users/thomashin/Library/Ethereum")

ACCOUNT OPTIONS:
  --unlock value                        Comma separated list of accounts to unlock
  --password value                      Password file to use for non-interactive password input

PLASMA EVM - ROOTCHAIN CONTRACT OPTIONS:
  --rootchain.url value                 JSONRPC endpoint of rootchain provider. If URL is empty, ignore the provider.

PLASMA EVM - STAKING OPTIONS OPTIONS:
  --rootchain.sender value              Address of root chain transaction sender account. it MUST be unlocked by --unlock, --password flags (CAVEAT: To set plasma operator, use --operator flag)
  --rootchain.gasprice value            Transaction gas price to root chain in GWei (default: 10000000000)
  --rootchain.wton value                Address of WTON token contract
  --rootchain.seigmanager value         Address of SeigManager contract
```

```bash
$ geth manage-staking start-powerton    # Start PowerTON first round

ETHEREUM OPTIONS:
  --datadir value                       Data directory for the databases and keystore (default: "/Users/thomashin/Library/Ethereum")

ACCOUNT OPTIONS:
  --unlock value                        Comma separated list of accounts to unlock
  --password value                      Password file to use for non-interactive password input

PLASMA EVM - DEVELOPMENT MODE OPTIONS:
  --dev.key value                       Comma seperated developer account key as hex(for dev)

PLASMA EVM - ROOTCHAIN CONTRACT OPTIONS:
  --rootchain.url value                 JSONRPC endpoint of rootchain provider. If URL is empty, ignore the provider.

PLASMA EVM - STAKING OPTIONS OPTIONS:
  --rootchain.sender value              Address of root chain transaction sender account. it MUST be unlocked by --unlock, --password flags (CAVEAT: To set plasma operator, use --operator flag)
  --rootchain.powerton value            Address of PowerTON contract
  --rootchain.gasprice value            Transaction gas price to root chain in GWei (default: 10000000000)
```

```bash
$ geth manage-staking register    # Register RootChain contract

ETHEREUM OPTIONS:
  --datadir value                       Data directory for the databases and keystore (default: "/Users/thomashin/Library/Ethereum")

ACCOUNT OPTIONS:
  --unlock value                        Comma separated list of accounts to unlock
  --password value                      Password file to use for non-interactive password input

PLASMA EVM - DEVELOPMENT MODE OPTIONS:
  --dev.key value                       Comma seperated developer account key as hex(for dev)

PLASMA EVM - ROOTCHAIN CONTRACT OPTIONS:
  --rootchain.url value                 JSONRPC endpoint of rootchain provider. If URL is empty, ignore the provider.

PLASMA EVM - STAKING OPTIONS OPTIONS:
  --rootchain.sender value              Address of root chain transaction sender account. it MUST be unlocked by --unlock, --password flags (CAVEAT: To set plasma operator, use --operator flag)
  --rootchain.gasprice value            Transaction gas price to root chain in GWei (default: 10000000000)
```

```bash
$ geth manage-staking get-managers <path>    # Get staking managers addresses in database

ETHEREUM OPTIONS:
  --datadir value                       Data directory for the databases and keystore (default: "/Users/thomashin/Library/Ethereum")
```

```bash
$ geth manage-staking set-managers <uri>    # Set staking managers addresses in database

ETHEREUM OPTIONS:
  --datadir value                       Data directory for the databases and keystore (default: "/Users/thomashin/Library/Ethereum")

PLASMA EVM - ROOTCHAIN CONTRACT OPTIONS:
  --rootchain.url value                 JSONRPC endpoint of rootchain provider. If URL is empty, ignore the provider.

PLASMA EVM - STAKING OPTIONS OPTIONS:
  --rootchain.ton value                 Address of TON token contract
  --rootchain.wton value                Address of WTON token contract
  --rootchain.depositmanager value      Address of Deposit Manager contract
  --rootchain.registry value            Address of RootChainRegistry contract
  --rootchain.seigmanager value         Address of SeigManager contract
  --rootchain.powerton value            Address of PowerTON contract
  --rootchain.gasprice value            Transaction gas price to root chain in GWei (default: 10000000000)
```

```bash
$ geth manage-staking mint-ton <to> <amount>    # Mint TON to account (for dev)

ETHEREUM OPTIONS:
  --datadir value                       Data directory for the databases and keystore (default: "/Users/thomashin/Library/Ethereum")

ACCOUNT OPTIONS:
  --unlock value                        Comma separated list of accounts to unlock
  --password value                      Password file to use for non-interactive password input

PLASMA EVM - ROOTCHAIN CONTRACT OPTIONS:
  --rootchain.url value                 JSONRPC endpoint of rootchain provider. If URL is empty, ignore the provider.

PLASMA EVM - STAKING OPTIONS OPTIONS:
  --rootchain.sender value              Address of root chain transaction sender account. it MUST be unlocked by --unlock, --password flags (CAVEAT: To set plasma operator, use --operator flag)
  --rootchain.ton value                 Address of TON token contract
  --rootchain.gasprice value            Transaction gas price to root chain in GWei (default: 10000000000)
```

### staking

```bash
$ geth staking balances <address>    # Print balances of token and stake

ETHEREUM OPTIONS:
  --datadir value                       Data directory for the databases and keystore (default: "/Users/thomashin/Library/Ethereum")

PLASMA EVM - ROOTCHAIN CONTRACT OPTIONS:
  --rootchain.url value                 JSONRPC endpoint of rootchain provider. If URL is empty, ignore the provider.

PLASMA EVM - STAKING OPTIONS OPTIONS:
  --rootchain.sender value              Address of root chain transaction sender account. it MUST be unlocked by --unlock, --password flags (CAVEAT: To set plasma operator, use --operator flag)
  --rootchain.ton value                 Address of TON token contract
  --rootchain.wton value                Address of WTON token contract
  --rootchain.depositmanager value      Address of Deposit Manager contract
  --rootchain.seigmanager value         Address of SeigManager contract
```

```bash
$ geth staking swap-from-ton <tonAmount>    # Swap TON to WTON

ETHEREUM OPTIONS:
  --datadir value                       Data directory for the databases and keystore (default: "/Users/thomashin/Library/Ethereum")

ACCOUNT OPTIONS:
  --unlock value                        Comma separated list of accounts to unlock
  --password value                      Password file to use for non-interactive password input

PLASMA EVM - ROOTCHAIN CONTRACT OPTIONS:
  --rootchain.url value                 JSONRPC endpoint of rootchain provider. If URL is empty, ignore the provider.

PLASMA EVM - STAKING OPTIONS OPTIONS:
  --rootchain.sender value              Address of root chain transaction sender account. it MUST be unlocked by --unlock, --password flags (CAVEAT: To set plasma operator, use --operator flag)
  --rootchain.ton value                 Address of TON token contract
  --rootchain.wton value                Address of WTON token contract
  --rootchain.depositmanager value      Address of Deposit Manager contract
  --rootchain.seigmanager value         Address of SeigManager contract
  --rootchain.gasprice value            Transaction gas price to root chain in GWei (default: 10000000000)
```

```bash
$ geth staking swap-to-ton <wtonAmount>    # Swap WTON to TON

ETHEREUM OPTIONS:
  --datadir value                       Data directory for the databases and keystore (default: "/Users/thomashin/Library/Ethereum")

ACCOUNT OPTIONS:
  --unlock value                        Comma separated list of accounts to unlock
  --password value                      Password file to use for non-interactive password input

PLASMA EVM - ROOTCHAIN CONTRACT OPTIONS:
  --rootchain.url value                 JSONRPC endpoint of rootchain provider. If URL is empty, ignore the provider.

PLASMA EVM - STAKING OPTIONS OPTIONS:
  --rootchain.sender value              Address of root chain transaction sender account. it MUST be unlocked by --unlock, --password flags (CAVEAT: To set plasma operator, use --operator flag)
  --rootchain.ton value                 Address of TON token contract
  --rootchain.wton value                Address of WTON token contract
  --rootchain.depositmanager value      Address of Deposit Manager contract
  --rootchain.seigmanager value         Address of SeigManager contract
  --rootchain.gasprice value            Transaction gas price to root chain in GWei (default: 10000000000)
```

```bash
$ geth staking stake-wton <amount>    # Stake WTON

ETHEREUM OPTIONS:
  --datadir value                       Data directory for the databases and keystore (default: "/Users/thomashin/Library/Ethereum")

ACCOUNT OPTIONS:
  --unlock value                        Comma separated list of accounts to unlock
  --password value                      Password file to use for non-interactive password input

PLASMA EVM - ROOTCHAIN CONTRACT OPTIONS:
  --rootchain.url value                 JSONRPC endpoint of rootchain provider. If URL is empty, ignore the provider.

PLASMA EVM - STAKING OPTIONS OPTIONS:
  --rootchain.sender value              Address of root chain transaction sender account. it MUST be unlocked by --unlock, --password flags (CAVEAT: To set plasma operator, use --operator flag)
  --rootchain.ton value                 Address of TON token contract
  --rootchain.wton value                Address of WTON token contract
  --rootchain.depositmanager value      Address of Deposit Manager contract
  --rootchain.seigmanager value         Address of SeigManager contract
  --rootchain.gasprice value            Transaction gas price to root chain in GWei (default: 10000000000)
```

```bash
$ geth staking restake <amount>    # Restake pending withdrawal request

ETHEREUM OPTIONS:
  --datadir value                       Data directory for the databases and keystore (default: "/Users/thomashin/Library/Ethereum")

ACCOUNT OPTIONS:
  --unlock value                        Comma separated list of accounts to unlock
  --password value                      Password file to use for non-interactive password input

PLASMA EVM - ROOTCHAIN CONTRACT OPTIONS:
  --rootchain.url value                 JSONRPC endpoint of rootchain provider. If URL is empty, ignore the provider.

PLASMA EVM - STAKING OPTIONS OPTIONS:
  --rootchain.sender value              Address of root chain transaction sender account. it MUST be unlocked by --unlock, --password flags (CAVEAT: To set plasma operator, use --operator flag)
  --rootchain.ton value                 Address of TON token contract
  --rootchain.wton value                Address of WTON token contract
  --rootchain.depositmanager value      Address of Deposit Manager contract
  --rootchain.seigmanager value         Address of SeigManager contract
  --rootchain.gasprice value            Transaction gas price to root chain in GWei (default: 10000000000)
```

```bash
$ geth staking request-withdrawal <amount>    # Make a withdrawal request

ETHEREUM OPTIONS:
  --datadir value                       Data directory for the databases and keystore (default: "/Users/thomashin/Library/Ethereum")

ACCOUNT OPTIONS:
  --unlock value                        Comma separated list of accounts to unlock
  --password value                      Password file to use for non-interactive password input

PLASMA EVM - ROOTCHAIN CONTRACT OPTIONS:
  --rootchain.url value                 JSONRPC endpoint of rootchain provider. If URL is empty, ignore the provider.

PLASMA EVM - STAKING OPTIONS OPTIONS:
  --rootchain.sender value              Address of root chain transaction sender account. it MUST be unlocked by --unlock, --password flags (CAVEAT: To set plasma operator, use --operator flag)
  --rootchain.depositmanager value      Address of Deposit Manager contract
  --rootchain.gasprice value            Transaction gas price to root chain in GWei (default: 10000000000)
```

```bash
$ geth staking process-withdrawal <numRequests>    # Process pending withdrawals

ETHEREUM OPTIONS:
  --datadir value                       Data directory for the databases and keystore (default: "/Users/thomashin/Library/Ethereum")

ACCOUNT OPTIONS:
  --unlock value                        Comma separated list of accounts to unlock
  --password value                      Password file to use for non-interactive password input

PLASMA EVM - ROOTCHAIN CONTRACT OPTIONS:
  --rootchain.url value                 JSONRPC endpoint of rootchain provider. If URL is empty, ignore the provider.

PLASMA EVM - STAKING OPTIONS OPTIONS:
  --rootchain.sender value              Address of root chain transaction sender account. it MUST be unlocked by --unlock, --password flags (CAVEAT: To set plasma operator, use --operator flag)
  --rootchain.depositmanager value      Address of Deposit Manager contract
  --rootchain.gasprice value            Transaction gas price to root chain in GWei (default: 10000000000)
```
