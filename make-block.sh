#!/bin/bash

DATADIR_1=$HOME/.pls.dev


INITIAL_BLOCK=`build/bin/geth --exec "eth.getBlock('latest').number" --datadir $DATADIR_1 attach`
INIT_BN=`build/bin/geth --exec "web3.eth.getBlock('latest').number" --datadir $DATADIR_1 attach`
echo "Initial block number is $INITIAL_BLOCK"
UNCLE=(`build/bin/geth --exec "eth.getBlock('latest').uncles" --datadir /Users/hwangjaeseung/.pls.dev attach`)
LEN=${#UNCLE[@]}
echo $UNCLE
echo "uncleblock length is $LEN"

GASLIMIT=`build/bin/geth --exec "eth.getBlock('latest').gasLimit" --datadir $DATADIR_1 attach`
echo $GASLIMIT

BLOCK=`expr $INITIAL_BLOCK + 10`
MINED_BLOCK=0

while [ $MINED_BLOCK -lt $BLOCK ]
do
  # send empty transaction
  build/bin/geth --exec "web3.eth.sendTransaction({from: eth.accounts[0], to:eth.accounts[0], value: 0})" attach --datadir $DATADIR_1 &
  
  MINED_BLOCK=`build/bin/geth --exec "eth.getBlock('latest').number" --datadir $DATADIR_1 attach`
  echo $MINED_BLOCK
done

TOTAL_GASLIMIT=0
UNCLE_BLOCK=0

for((i=$INITIAL_BLOCK;i<$BLOCK;i++))
do
  # Calculate total block size
  GASLIMIT=`build/bin/geth --exec "eth.getBlock($i).gasLimit" --datadir $DATADIR_1 attach`
  # echo $GASLIMIT
  TOTAL_GASLIMIT=`expr $TOTAL_GASLIMIT + $GASLIMIT`

  # Calculate Uncle ratio
  UNCLE=(`build/bin/geth --exec "eth.getBlock($i).uncles" --datadir $DATADIR_1 attach`)
  LEN=${#UNCLE[@]}
  UNCLE_BLOCK=`expr $UNCLE_BLOCK + $LEN`
done



AVG_GASLIMIT=`expr $TOTAL_GASLIMIT / 1000`
echo "Average block size is $AVG_GASLIMIT"

UNCLE_RATIO=`expr $UNCLE_BLOCK / 1000`
echo "Total uncle block is $UNCLE_BLOCK"
echo "Uncle ratio is $UNCLE_RATIO"

LAST_BT=`build/bin/geth --exec "eth.getBlock('latest').timestamp" --datadir $DATADIR_1 attach`
INIT_BT=`build/bin/geth --exec "eth.getBlock($INIT_BN).timestamp" --datadir $DATADIR_1 attach`

TOT_BT=`expr $LAST_BT - $INIT_BT`
AVG_BT=`expr $TOT_BT / 1000`

echo $TOT_BT
echo $AVG_BT