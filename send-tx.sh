#!/bin/bash

DATADIR_1=$HOME/.pls.dev

INIT_TIME=`date +%Y%m%d%H%M%S`
INIT_BALANCE=`build/bin/geth --exec "eth.getBalance(eth.accounts[0])" --datadir $DATADIR_1 attach`

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
# for((i=0;i<50;i++))
# do
#   # send empty transaction
#   build/bin/geth --exec "web3.eth.sendTransaction({from: eth.accounts[0], to:eth.accounts[0], value: 0})" attach --datadir $DATADIR_1
#   build/bin/geth --exec "web3.eth.sendTransaction({from: eth.accounts[0], to:eth.accounts[0], value: 0})" attach --datadir $DATADIR_1
  
# done
CUR_TIME=`date +%Y%m%d%H%M%S`
CUR_BALANCE=`build/bin/geth --exec "eth.getBalance(eth.accounts[0])" --datadir $DATADIR_1 attach`

echo "Start Time is $INIT_TIME"
echo "Current Time is $CUR_TIME"
TIME_DIFF=`expr $CUR_TIME - $INIT_TIME`
echo "It takes $TIME_DIFF seconds."

echo "Initial balance is $INIT_BALANCE"
echo "Currnet balance is $CUR_BALANCE"

BALANCE_DIFF=`expr $INIT_BALANCE - $CUR_BALANCE`
echo "It use $BALANCE_DIFF"

echo "10000 tx done"
