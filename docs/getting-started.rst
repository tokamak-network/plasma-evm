===============================
Getting Started
===============================

| If feels complexity yourself or Could have time to do step-by-step,
| There is Dockerized section at the end of this document. You can go directly goto.
|
| Let considering Two environments. One is Private Environment, cannot connected outside of host, For plasma-chain test.
| Another one is Publically opened Environment almost production level.
| You can Mixing those two environments if fully used to.

---------------------------------
Setup Private Environment
---------------------------------

There are two essential steps.

1. Run **go-ethereum** as RootChain.
2. Run **Plasma-evm** as ChildChain with RootChain infomation.

Follow Instructions are tested on MacOSX.

1. Run RootChain
=================
Use forked version of go-ethereum v1.8.20 as RootChain.

Building geth requires both a GO (version 1.11 or later) and C compiler.


1.1. Clone go-ethereum Repository
----------------------------------

::

    git clone http://github.com/onther-tech/go-ethereum

| You dont need generate genesis file for eth balance allocation. It is going to use fake ethash as dev mode.
| It forked from geth v1.8.20 then added some features pre-founded, other things via flags.
| Use can use `run.rootchain.sh` running script instead geth run command like as below 1.3.

1.2. Build the source
----------------------------------

::

    make geth

If you want do running script, can skip this command.


1.3. Run ge-ethereum with flags.
----------------------------------

| Plasma-evm subscribe RootChain Events via rootchain's websocket, must be open to ChildChain.

::

    bash run.rootchain.sh

Or You can directly run with this commands

::

    build/bin/geth --dev --dev.period 1 --dev.faucetkey
    "b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291,78ae75d1cd5960d87e76a69760cb451a58928eee7890780c352186d23094a115,bfaa65473b85b3c33b2f5ddb511f0f4ef8459213ada2920765aaac25b4fe38c5,067394195895a82e685b000e592f771f7899d77e87cc8c79110e53a2f0b0b8fc,ae03e057a5b117295db86079ba4c8505df6074cdc54eec62f2050e677e5d4e66" --miner.gastarget 7500000 --miner.gasprice "10" --rpc --rpcport 8545 --rpcapi eth,debug,net --ws --wsport 8546

In this case, 5 privateKeys generate accounts files under datadir path.

2. Run ChildChain
==================
We currently working on Plasma-evm running stable.
Suggest, Clone master branch instead develop which is default.


2.1. Clone Plamsa-evm Repository
----------------------------------

::

    git clone http://github.com/onther-tech/plasma-evm


2.2. Build the source
----------------------------------

Do as same as go-ethereum

::

    make geth

2.3. Run Plasma-evm with flags
----------------------------------

| There are additional params to run Plasma-chain through flags. No need genesis file. It going to automatically generated.
| We added some flags for get params to run plasma-evm.
| You can get some information about the flags for plasma-evm as like below, using ``geth --help``.

::

 MISC OPTIONS:
  --operator.minether value                  Plasma operator minimum balance (default = 0.5 ether) (default: "0.5")
  --operator value                           Plasma operator address as hex. The account should be unlock by using --unlock
  --operator.key value                       Plasma operator key as hex(for dev)

  --dev.key value                            Comma seperated developer account key as hex(for dev)

  --rootchain.url value                      JSONRPC endpoint of rootchain provider (default: "ws://localhost:8546")
  --rootchain.contract value                 Address of the RootChain contract
  --rootchain.challenger value               Address of challenger account

  --tx.mingasprice "1000000000"              Minimum gas price for submitting a block (default = 1 Gwei)
  --tx.maxgasprice "100000000000"            Maximum gas price for submitting a block (default = 100 Gwei)
  --tx.interval value                        Pending interval time after submitting a block (default = 10s). If block submit transaction is not mined in 2 intervals, gas price will be adjusted. See https://golang.org/pkg/time/#ParseDuration (default: 10s)

  --stamina.mindeposit "500000000000000000"  MinDeposit variable state of stamina contract
  --stamina.recoverepochlength "10080"       RecoverEpochLength variable state of stamina contract
  --stamina.withdrawaldelay "30240"          WithdrawalDelay variable state of stamina contract

`dev.key`, `operator.key`, `rootchain.challenger` Are for testing.

| If there is already deployed `rootchain` contract on RootChain Network,
| you can use like ``--rootchain.contract 0x123456789aa`` instead ``--dev`` mode. Cannot use `--dev` and `--rootchain.contract` at the same time.
| In plasma-evm `dev` mode has additional features, which automatically deploying `rootchain` contract if no have rootchain address via `--roothchain.contract`.

| In this case for testing. use dev mode.

[Important Notice] Operator Account must have some ether balance at RootChain.
If Does not have, Could not start ChildChain.

::

    geth --miner.etherbase 0x71562b71999873DB5b286dF957af199Ec94617F7 --dev --rpc --rpcaddr 0.0.0.0 --rpcport 8547 --port 30307 --dev.key b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291 --operator 0x71562b71999873DB5b286dF957af199Ec94617F7 --tx.interval "300ms" --rootchain.url "ws://127.0.0.1:8546"


If you consider to run in production level, Recommand raise `tx.interval` time, at least 20s.

| Finally, You can see like Plasma-chain running!

| Remember deployed contract addresses! It can be use for later.
| **RootChain contract** address is most Important contract addresses on below. In here, address start with 0x880EC...

.. code::

  INFO [04-04|01:02:11.897] Persisted trie from memory database      nodes=13 size=5.80kB time=66.292µs gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
  INFO [04-04|01:02:11.897] Deploying contracts for development mode
  INFO [04-04|01:02:11.908] Deploy MintableToken contract            hash=2febea…3b8f02 address=0x3A220f351252089D385b29beca14e27F204c296A
  INFO [04-04|01:02:11.908] Wait until deploy transaction is mined
  INFO [04-04|01:02:12.920] Deploy EtherToken contract               hash=b48a9b…0e82f3 address=0xdB7d6AB1f17c6b31909aE466702703dAEf9269Cf
  INFO [04-04|01:02:12.921] Wait until deploy transaction is mined
  INFO [04-04|01:02:19.953] Deploy EpochHandler contract             hash=f30f84…337aa6 address=0x537e697c7AB75A26f9ECF0Ce810e3154dFcaaf44
  INFO [04-04|01:02:19.953] Wait until deploy transaction is mined
  INFO [04-04|01:02:22.983] Deploy RootChain contract                hash=175321…ff616f address=0x880EC53Af800b5Cd051531672EF4fc4De233bD5d
  INFO [04-04|01:02:30.012] Initialize EtherToken                    hash=584c83…0a41e1
  INFO [04-04|01:02:32.019] Set options for submitting a block       mingaspirce=1000000000 maxgasprice=300000000000 interval=10s
  INFO [04-04|01:02:32.019] Starting peer-to-peer node               instance=Geth/v1.8.20-stable-3a343606/darwin-amd64/go1.9.5
  INFO [04-04|01:02:32.019] Allocated cache and file handles         database=/Users/jins/.pls.dev/geth/chaindata cache=512 handles=4611686018427387903
  INFO [04-04|01:02:32.026] Writing custom genesis block             rootChainContract=0x880EC53Af800b5Cd051531672EF4fc4De233bD5d
  INFO [04-04|01:02:32.027] Persisted trie from memory database      nodes=13 size=5.80kB time=124.834µs gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
  INFO [04-04|01:02:32.027] Initialised chain configuration          config="{ChainID: 16 Homestead: 0 DAO: <nil> DAOSupport: false EIP150: 0 EIP155: 0 EIP158: 0 Byzantium: 0 Constantinople: <nil> Engine: ethash}"
  WARN [04-04|01:02:32.027] Ethash used in fake mode
  INFO [04-04|01:02:32.027] Initialising Plasma protocol             versions="[63 62]" network=1337
  INFO [04-04|01:02:32.048] Loaded most recent local header          number=0 hash=e413e8…e44af1 td=1 age=49y11mo2w
  INFO [04-04|01:02:32.048] Loaded most recent local full block      number=0 hash=e413e8…e44af1 td=1 age=49y11mo2w
  INFO [04-04|01:02:32.048] Loaded most recent local fast block      number=0 hash=e413e8…e44af1 td=1 age=49y11mo2w
  INFO [04-04|01:02:32.049] Regenerated local transaction journal    transactions=0 accounts=0
  INFO [04-04|01:02:32.051] Rootchain provider connected             url=ws://localhost:8546
  INFO [04-04|01:02:32.061] New local node record                    seq=1 id=df4cc248d21c5db6 ip=127.0.0.1 udp=0 tcp=55563
  INFO [04-04|01:02:32.061] Started P2P networking                   self="enode://6f7ff81c34959c797e96704e5082fab0550ba603c5dec6825fc1b31f85f1a441303eb94af46ca2ab36165bd0f9738b3337e5c8fee4b51b22bafad08fb201fe6e@127.0.0.1:55563?discport=0"
  INFO [04-04|01:02:32.063] Iterating epoch prepared event
  INFO [04-04|01:02:32.063] RootChain epoch prepared                 epochNumber=1 epochLength=2 isRequest=false userActivated=false isEmpty=false ForkNumber=0 isRebase=false
  INFO [04-04|01:02:32.063] NRB epoch is prepared, NRB epoch is started NRBepochLength=2
  INFO [04-04|01:02:32.064] Iterating block finalized event
  INFO [04-04|01:02:32.064] RootChain block finalized                forkNumber=0 blockNubmer=0
  INFO [04-04|01:02:32.064] Watching epoch prepared event            start block number=0
  INFO [04-04|01:02:32.065] Watching block finalized event           start block number=0
  INFO [04-04|01:02:32.065] Updated mining threads                   threads=8
  INFO [04-04|01:02:32.065] started whisper v.6.0
  INFO [04-04|01:02:32.068] IPC endpoint opened                      url=/Users/jins/.pls.dev/geth.ipc
  INFO [04-04|01:02:32.068] HTTP endpoint opened                     url=http://127.0.0.1:8547         cors= vhosts=localhost
  INFO [04-04|01:02:34.312] Mapped network port                      proto=tcp extport=55563 intport=55563 interface="UPNP IGDv2-IP1"


Looks like stop, but It Just waiting Tx!
In dev mode, Start block mine when transaction has on txpool.


----------------------------
Setup Public Environment
----------------------------

1. Run RootChain
=================
Recommand to use go-ethereum v1.8.23 as RootChain.

Building geth requires both a GO (version 1.11 or later) and C compiler.


1.1. Clone go-ethereum Repository
----------------------------------


::

    git clone -b v1.8.23 http://github.com/ethereum/go-ethereum

You should generate genesis file via puppeth. Recommand consensus ethash, not Clique.


1.2. Build the source
----------------------------------

::

    make geth

1.3. Run ge-ethereum with flags.
----------------------------------

| If you want to run with Ropsten testnet in here, Add `--testnet` then this geth going to have chainId 3.
| Plasma-evm subscribe RootChain Events via rootchain's websocket, must be open to ChildChain.
| If you want to running on Ropsten network, use **--testnet** flag instead --dev flag.


| geth initialize with genesis file

::

  geth init --datadir data genesis.json

then

::

    geth --datadir data --mine --miner.etherbase 0x71562b71999873DB5b286dF957af199Ec94617F7 --miner.gastarget 7500000 --miner.gasprice "10" --rpc --rpcaddr 0.0.0.0 --rpcport 8545 --rpcapi web3,eth,personal,miner,net,txpool --ws --wsaddr 0.0.0.0 --wsport 8546 --wsorigins="*" --unlock 0x71562b71999873DB5b286dF957af199Ec94617F7,0x5df7107c960320b90a3d7ed9a83203d1f98a811d,0x3cd9f729c8d882b851f8c70fb36d22b391a288cd --password ./signer.pass


| In this case, Inserted 3 Keyfiles already in `data` path.

2. Run ChildChain
==================
We currently working on Plasma-evm running stable.
Suggest, Clone master branch instead develop which is Default.

Important Notice, Every User has own node in plasma-evm which syncing operator's one.
If user does not, It is securely vulnerable by Data Availability.

And also, Operator's node has own private key for commit transactions to RootChain.
Have to properly secure action via firewall etc.


2.1. Clone Plamsa-evm Repository
----------------------------------

::

    git clone -b master http://github.com/onther-tech/plasma-evm


2.2. Build the source
----------------------------------

Do as same as go-ethereum

::

    make geth

2.3. Run Plasma-evm as Operator
----------------------------------

| If there is already deployed `rootchain` contract on RootChain Network,
| you can use like ``--rootchain.contract 0x123456789aa`` instead ``--dev`` mode. Cannot use `--dev` and `--rootchain.contract` at the same time.
| In plasma-evm `dev` mode has additional features, which automatically deploying `rootchain` contract if no have rootchain address via `--roothchain.contract`.

| In this case for testing. use dev mode.

[Important Notice] Operator Account must have some ether balance at RootChain.
If Does not have, Could not start ChildChain.

::

    geth --miner.etherbase 0x71562b71999873DB5b286dF957af199Ec94617F7 --dev --port 30307 --dev.key b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291 --operator 0x71562b71999873DB5b286dF957af199Ec94617F7 --tx.interval "300ms" --rootchain.url "ws://127.0.0.1:8546"

If you consider to run in production level, Recommand raise `tx.interval` time, at least 10s.

| Remember deployed contract addresses! It can be use for later.

2.4. Run Plasma-evm as User
----------------------------------

| There are different between Operator Node and User Node.
| User Node Has no Own Private Key and No Mining.
| But Need same networkId, genesis hash at start.

| There is two way to connect Operator Node and User Node Each other.
| One is that Using bootnode, It is simpler than other way. run bootnode then insert bootnode address on flag `--bootnodes`.

| The Other is that add peer, Operator Node, manually at User Node.
|
| Important thing is that Exactly Match networkid, genesis block hash between User and Operator Node.
| If does not, Cannot find each other forever.

| There is one Problem, has different genesis block hash when run plasma-evm everytime. Must fix one of Two.
| So, If you want to run User node, use other plasma-evm branch, p2p-in-dev-mode, still not merged soon.


::

    git clone -b p2p-in-dev-mode http://github.com/onther-tech/plasma-evm



then you should check `rootchain`, `operator` addresses as same as Operator's.


::

    geth --dev --dev.p2p --networkid 1337 --rpc --rpcaddr 0.0.0.0 --rpcport 8549 --port 30307 --rootchain.url "ws://127.0.0.1:8546"  --dev.key b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291 --dev.rootchain 0x880EC53Af800b5Cd051531672EF4fc4De233bD5d --operator 0x71562b71999873DB5b286dF957af199Ec94617F7


In here, `--dev.p2p` mode make turn on p2p networking. so Please do not forgot.
And others, `dev.rootchain`, `dev.operator`, are need to generate same genesis block hash.

Add `--bootnode [bootnode key@ip:port]` flag if you using Bootnode. or Add peer using geth console.


Remember, have to same Networkid and Genesis hash between User & Operator Node.


------------------------
Quick Start with Docker
------------------------

It is Private Environment for test.


1. Clone dockerize branch Plasma-evm

::

  git clone -b dockerize http://github.com/onther-tech/plasma-evm

2. Update Submodules

::

  git submodule update --init --recursive

3. Up docker-compose

::

  docker-compose up


- If you turn down containers `docker-compose down` on plasma-evm path.

------------------------
Quick Start with Truffle
------------------------

You can use the truffle framework to deploy contracts to the tokamak testnet, faraday. To use the faraday network, you need to specify the network in truffle-config.js as follows.

::

  module.exports = {
  networks: {
    development: {
      host: 'localhost',
      port: 8545,
      network_id: '*'
    },
    ropsten: {
      provider: ropstenProvider,
      network_id: 3
    },
    faraday: {
      host: "112.169.69.41",
      port: 48549,
      network_id: "*"
    }
  }

If you want to test on the faraday network, you can use faucet at `here <http://faucet.tokamak.network:8080/#/>`_.

.. note::
  Faraday is the Tokamak network used as a testing environment. Faraday is connected to the ropsten network as a root chain.

Enter and exit token at faraday network
=======================================

We will enter and exit token using the :ref:`RequestableSimpleToken<requestable-simple-token>` without ownership, `RequestableSimpleTokenWithoutOwnership <https://github.com/Onther-Tech/requestable-simple-token/blob/master/contracts/RequestableSimpleTokenWithoutOwnership.sol>`_. Anyone can mint the token because there is no ownership. 

RequestableSimpleTokenWithoutOwnership contract and RootChain contract are already deployed, so you can use it to test enter and exit. RequestableSimpleTokenWithoutOwnership contract address at Ropsten is ``0x6B27C38e3376C4E8B29cFbB3986f00676267D489`` and contract address at Faraday is
``0x1d93d7bd7d820ac7691109ace371e42d5004e1c1``. `RootChain` contract address at Ropsten is ``0x3122546c1544FD0F910A423A8c80fdCD48d742Fd``.

The scenario will work as follows:

1. Alice mint ``RequestableSimpleTokenWithoutOwnership`` at the Ropsten.
2. Alice get her trieKey by using ``RequestableSimpleTokenWithoutOwnership.getBalanceTrieKey()`` and trieValue.
3. Alice call ``RootChain.startEnter()`` method to start entering process to Faraday.

::

    truffle(ropsten)> alice = '0xb60e8dd61c5d32be8058bb8eb970870f07233155'
    truffle(ropsten)> rootchain = RootChain.at('0x3122546c1544FD0F910A423A8c80fdCD48d742Fd')
    truffle(ropsten)> token = RequestableSimpleTokenWithoutOwnership.at('0x6B27C38e3376C4E8B29cFbB3986f00676267D489')

    truffle(ropsten)> token.mint(alice, 10000000)
    truffle(ropsten)> key = token.getBalanceTrieKey(alice)
    truffle(ropsten)> value = '0x0000000000000000000000000000000000000000000000000000000000000100'
    truffle(ropsten)> rootchain.startEnter(alice, key, value)
    
4. After entering process is finished, you can check entered token balance by using ``RequestableSimpleTokenWithoutOwnership.balances()`` at Faraday.

::

    truffle(faraday)> alice = '0xb60e8dd61c5d32be8058bb8eb970870f07233155'
    truffle(faraday)> token = RequestableSimpleTokenWithoutOwnership.at('0x1d93d7bd7d820ac7691109ace371e42d5004e1c1')
    truffle(faraday)> token.balances(alice)

5. If Alice wants to use his token at rootchain, then start exit process to rootchain by using ``RootChain.startExit()``.

::  

    truffle(ropsten)> alice = '0xb60e8dd61c5d32be8058bb8eb970870f07233155'
    truffle(ropsten)> rootchain = RootChain.at('0x3122546c1544FD0F910A423A8c80fdCD48d742Fd')

    truffle(ropsten)> key = token.getBalanceTrieKey(alice)
    truffle(ropsten)> value = '0x0000000000000000000000000000000000000000000000000000000000000100'
    truffle(ropsten)> rootchain.startExit(alice, key, value)
