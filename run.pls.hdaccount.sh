OPERATOR_ADDRESS="c49926c4124cee1cba0ea94ea31a6c12318df947"
KEY1="78ae75d1cd5960d87e76a69760cb451a58928eee7890780c352186d23094a115"
KEY2="bfaa65473b85b3c33b2f5ddb511f0f4ef8459213ada2920765aaac25b4fe38c5"
KEY3="067394195895a82e685b000e592f771f7899d77e87cc8c79110e53a2f0b0b8fc"
KEY4="ae03e057a5b117295db86079ba4c8505df6074cdc54eec62f2050e677e5d4e66"
KEY5="eda4515e1bc6c08e8606b51ffb6ffe70b3fe76781ed49872285e484064e3b634"

if [ -z "$1" ]
then
  echo "[Usage] run.pls.sh [RootChain address]"
else
  DATADIR=$HOME/.pls.data.dir

  rm -rf $DATADIR

  make geth && \
  build/bin/geth account importHDwallet "tag volcano eight thank tide danger coast health above argue embrace heavy" "m/44'/60'/0'/0/0"\
    --datadir $DATADIR
  build/bin/geth \
    --datadir $DATADIR \
    --miner.etherbase 0x71562b71999873DB5b286dF957af199Ec94617F7 \
    --rpc \
    --rpcport 8547 \
    --dev.key $KEY1,$KEY2,$KEY3,$KEY4,$KEY5 \
    --rootchain.operatorAddress $OPERATOR_ADDRESS \
    --rootchain.contract $1
fi
