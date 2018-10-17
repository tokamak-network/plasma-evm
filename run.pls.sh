if [ -z "$1" ]
then
  echo "[Usage] run.pls.sh [RootChain address]"
else
  DATADIR=$HOME/.pls.data.dir

  rm -rf $DATADIR

  make geth && \
  build/bin/geth \
    --datadir $DATADIR \
    --miner.etherbase 0x71562b71999873DB5b286dF957af199Ec94617F7 \
    --rpc \
    --rpcport 8547 \
    --operatorKey b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291 \
    --rootchain.contract $1
fi
