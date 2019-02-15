package plasma

//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/RootChain.sol --pkg rootchain --out rootchain/rootchain.go
//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/EpochHandler.sol --pkg epochhandler --out epochhandler/epochhandler.go
//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/RequestableSimpleToken.sol --pkg token --out token/token.go
//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/node_modules/openzeppelin-solidity/contracts/token/ERC20/MintableToken.sol --pkg mintabletoken --out mintabletoken/mintabletoken.go
//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/EtherToken.sol --pkg ethertoken --out ethertoken/ethertoken.go
