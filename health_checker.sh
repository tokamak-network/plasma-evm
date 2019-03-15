#!/bin/bash

# check if the operator node is alive.
json=$(curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["latest", true],"id":1}' \
 -H "Content-Type: application/json" 127.0.0.1:8547 || echo "Operator node is not initialized" & exit 1)

blocknumber=$(echo $json | jq '.result.number')

echo $blocknumber >> result.log

second_to_last_block=$(tail -n 2 result.log | head -1)
last_block=$(tail -n 1 result.log)

if [ -n "${second_to_last_block}" ] && [ $second_to_last_block = $last_block ]
then
  echo "Operator is not mining block." & exit 1
elif [ -z "${second_to_last_block}" ] && [ -n "${last_block}" ]
then
  echo "Operator is initialized" & exit 0
else
  echo "Operator is mining block." & exit 0
fi
