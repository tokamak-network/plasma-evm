#!/bin/bash

DATADIR_1=$HOME/.pls.dev

INIT_TIME=`date +%Y%m%d%H%M%S`
INIT_BALANCE=`build/bin/geth --exec "eth.getBalance(eth.accounts[0])" --datadir $DATADIR_1 attach`
INIT_BN=`build/bin/geth --exec "eth.getBlock('latest')).number" attach --datadir $DATADIR_1`
INIT_TIME=`build/bin/geth --exec "eth.getBlock('latest')).timestamp" attach --datadir $DATADIR_1`

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

TOT_TX=0

LAST_BN=`build/bin/geth --exec "eth.getBlock('latest')).number" attach --datadir $DATADIR_1`
LAST_TIME=`build/bin/geth --exec "eth.getBlock('latest')).timestamp" attach --datadir $DATADIR_1`

for((i=$INIT_BN;i<$LAST_BN;i++))
do
  # send empty transaction
  TX=`build/bin/geth --exec "eth.getBlock('latest').transactions" attach --datadir $DATADIR_1`
  NUM_TX=${#TX[@]}
  TOT_TX=`expr $TOT_TX + $NUM_TX`  
done
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

DURATION=`expr $LAST_TIME - $INIT_TIME`
echo "Duration is $DURATION"

TPS=`expr $TOT_TX / $DURATION`
echo "TPS is $TPS"

