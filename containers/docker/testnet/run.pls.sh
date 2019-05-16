#!/usr/bin/env bash

OPERATOR_KEY="b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291"
KEY1="78ae75d1cd5960d87e76a69760cb451a58928eee7890780c352186d23094a115"
KEY2="bfaa65473b85b3c33b2f5ddb511f0f4ef8459213ada2920765aaac25b4fe38c5"
KEY3="067394195895a82e685b000e592f771f7899d77e87cc8c79110e53a2f0b0b8fc"
KEY4="ae03e057a5b117295db86079ba4c8505df6074cdc54eec62f2050e677e5d4e66"
KEY5="eda4515e1bc6c08e8606b51ffb6ffe70b3fe76781ed49872285e484064e3b634"
CHALLENGER_KEY="78ae75d1cd5960d87e76a69760cb451a58928eee7890780c352186d23094a114"

DATADIR=$HOME/.pls.dev
OPERATOR="0x71562b71999873DB5b286dF957af199Ec94617F7"
CHALLENGER="0x3616BE06D68dD22886505e9c2CaAa9EcA84564b8"

ROOTCHAIN_IP=172.31.0.3

# Only once running at first time to run
if [ ! -d "$DATADIR" ]; then
  sleep 15s
  # Deploy contracts at rootchain
  echo "Deploy rootchain contract and others"
  geth \
    --rootchain.url "ws://$ROOTCHAIN_IP:8546" \
    --operator.key $OPERATOR_KEY \
    --datadir $DATADIR \
    deploy "./genesis.json" 1663 true 2 
fi

if [ ! -d "$DATADIR" ]; then
  # Initialize with genesis file.
  echo "Plasma-evm Initialize with genesis file"
  geth \
    --datadir $DATADIR \
    --operator.key $OPERATOR_KEY \
    --rootchain.url "ws://ROOTCHAIN_IP:8546"\
    init genesis.json
fi

# get rootchain contract address from genesis
export ROOTADDR=$(cat genesis.json | jq .extraData)
echo "Rootchain contract : $ROOTADDR"

geth \
  --datadir $DATADIR \
  --rpc \
  --rpcaddr 0.0.0.0 \
  --rpcport 8547 \
  --rootchain.contract ${ROOTADDR:1:-1} \
  --rootchain.url "ws://$ROOTCHAIN_IP:8546" \
  --operator.key $OPERATOR_KEY
