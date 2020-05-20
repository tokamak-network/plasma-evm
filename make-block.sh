#!/bin/bash

DATADIR_1=$HOME/.pls.dev


INITIAL_BLOCK=`build/bin/geth --exec "eth.getBlock('latest').number" --datadir $DATADIR_1 attach`
echo "Initial block number is $INITIAL_BLOCK"
UNCLE=`build/bin/geth --exec "eth.getBlock('latest').uncles" --datadir /Users/hwangjaeseung/.pls.dev attach`
LEN=${#UNCLE[@]}
echo $UNCLE
echo "uncleblock length is $LEN"

GASLIMIT=`build/bin/geth --exec "eth.getBlock(i).gasLimit" --datadir $DATADIR_1 attach`
echo $GASLIMIT

BLOCK=`expr $INITIAL_BLOCK + 1000`

while [ $MINED_BLOCK -lt $BLOCK ]
do
  # send empty transaction
  build/bin/geth --exec "web3.eth.sendTransaction({from: eth.accounts[0], to:eth.accounts[0], value: 0})" \
               attach --datadir $DATADIR_1
  
  MINED_BLOCK=`build/bin/geth --exec "eth.getBlock('latest').number" --datadir $DATADIR_1 attach`
done

TOTAL_GASLIMIT=0
UNCLE_BLOCK=0

# for((i=$INITIAL_BLOCK;i<$BLOCK;i++))
# do
#   # Calculate total block size
#   GASLIMIT=`build/bin/geth --exec "eth.getBlock(i).gasLimit" --datadir $DATADIR_1 attach`
#   TOTAL_GASLIMIT=`expr $TOTAL_GASLIMIT + $GASLIMIT`

#   # Calculate Uncle ratio
#   UNCLE=`build/bin/geth --exec "eth.getBlock('latest').uncles" --datadir /Users/hwangjaeseung/.pls.dev attach`
#   UNCLE_BLOCK=$UNCLE_BLOCK
# done

# AVG_GASLIMIT=`expr $TOTAL_GASLIMIT/1000`
# echo "Average block size is $AVG_GASLIMIT"