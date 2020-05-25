#!/bin/bash

DATADIR_1=$HOME/.pls.dev
build/bin/geth --exec "web3.eth.sendTransaction({from: eth.accounts[0], to:eth.accounts[1], value: 1e18})" attach --datadir $DATADIR_1  


INIT_BALANCE=`build/bin/geth --exec "eth.getBalance(eth.accounts[1])" --datadir $DATADIR_1 attach`

bash send-tx1.sh &
bash send-tx2.sh &
bash send-tx3.sh &
bash send-tx4.sh &
bash send-tx5.sh &
bash send-tx6.sh &
bash send-tx7.sh &
bash send-tx8.sh &
bash send-tx9.sh &
bash send-tx10.sh &

wait
echo "Sending tx is done"


TOT_TX=0

TOTAL_BLOCK=``


CUR_BALANCE=`build/bin/geth --exec "eth.getBalance(eth.accounts[1])" --datadir $DATADIR_1 attach`


echo "Initial balance is $INIT_BALANCE"
echo "Currnet balance is $CUR_BALANCE"


