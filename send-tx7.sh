#!/bin/bash

DATADIR_1=$HOME/.pls.dev

for((i=0;i<100;i++))
do
  # send empty transaction
  build/bin/geth --exec "web3.eth.sendTransaction({from: eth.accounts[0], to:eth.accounts[0], value: 0})" attach --datadir $DATADIR_1  
done
CUR_TIME=`date +%Y%m%d%H%M%S`

echo "tx7 done"
