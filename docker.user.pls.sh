#!/usr/bin/env bash
DATADIR=$HOME/.pls.dev

rm -rf $DATADIR

geth \
  --datadir $DATADIR \
  --rpc \
  --rpcaddr 0.0.0.0 \
  --rpcport 8549 \
  --rootchain.url "ws://172.30.0.3:8546" \
  --rootchain.contract "0x880EC53Af800b5Cd051531672EF4fc4De233bD5d" \
  --syncmode "full" \
  --bootnodes "enode://401bd6383fe11a5224d5b4277b53d7c0278efed3ca685b6593935751ad1fe734a8e35d2b3ebd9d7fc6da6cff12e72cfcfca8db408bf2e49f1fad4c503956d07f@172.30.0.6:30301"


