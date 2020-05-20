#!/bin/bash

DATADIR_1=$HOME/.pls.dev


INITIAL_BLOCK=`build/bin/geth --exec "eth.getBlock('latest').number" --datadir $DATADIR_1 attach`
echo "Initial block number is $INITIAL_BLOCK"

MINED_BLOCK=`expr $INITIAL_BLOCK + 10000`
echo $MINED_BLOCK
while :
do
  # send empty transaction
  build/bin/geth --exec "web3.eth.sendTransaction({from: eth.accounts[0], to:eth.accounts[0], value: 0})" \
               attach --datadir $DATADIR_1
  

done
