#!/bin/bash

DATADIR_1=$HOME/.pls.dev

for((i=0;i<1000;i++))
do
  # send empty transaction
  build/bin/geth --exec "web3.eth.sendTransaction({from: eth.accounts[2], to:eth.accounts[2], value: 0})" --datadir $DATADIR_1 attach
done
CUR_TIME=`date +%Y%m%d%H%M%S`

echo "tx4 done"
