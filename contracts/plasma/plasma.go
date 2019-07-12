package plasma

//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/RootChain.sol --pkg rootchain --out rootchain/rootchain.go
//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/EpochHandler.sol --pkg epochhandler --out epochhandler/epochhandler.go
//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/RequestableSimpleToken.sol --pkg token --out token/token.go
//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/node_modules/openzeppelin-solidity/contracts/token/ERC20/ERC20Mintable.sol --pkg mintabletoken --out mintabletoken/mintabletoken.go
//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/EtherToken.sol --pkg ethertoken --out ethertoken/ethertoken.go

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/epochhandler"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/ethertoken"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/mintabletoken"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/rootchain"
	"github.com/Onther-Tech/plasma-evm/core"
	"github.com/Onther-Tech/plasma-evm/core/rawdb"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/ethclient"
	"github.com/Onther-Tech/plasma-evm/log"
	"github.com/Onther-Tech/plasma-evm/pls"
)

func DeployPlasmaContracts(opt *bind.TransactOpts, backend *ethclient.Client, cfg *pls.Config, withPETH bool, development bool, NRELength *big.Int) (common.Address, error) {
	opt.GasLimit = 7000000

	wait := func(hash common.Hash) {
		<-time.NewTimer(1 * time.Second).C

		for receipt, _ := backend.TransactionReceipt(context.Background(), hash); receipt == nil; {
			<-time.NewTimer(1 * time.Second).C

			receipt, _ = backend.TransactionReceipt(context.Background(), hash)
		}
	}

	dummyDB := rawdb.NewMemoryDatabase()
	defer dummyDB.Close()
	var dummyBlock *types.Block

	if withPETH {
		// give PETH to operator
		dummyBlock = core.DeveloperGenesisBlock(
			1,
			common.HexToAddress("0xdead"),
			opt.From,
			cfg.StaminaConfig,
		).ToBlock(dummyDB)
	} else {
		// Nobody has PETH in genesis block
		dummyBlock = core.DefaultGenesisBlock(
			common.HexToAddress("0xdead"),
			opt.From,
			cfg.StaminaConfig,
		).ToBlock(dummyDB)
	}

	// 1. deploy MintableToken in root chain
	mintableTokenContract, tx, _, err := mintabletoken.DeployERC20Mintable(opt, backend)
	if err != nil {
		return common.Address{}, errors.New(fmt.Sprintf("Failed to deploy MintableToken contract: %v", err))
	}
	log.Info("Deploy MintableToken contract", "hash", tx.Hash(), "address", mintableTokenContract)

	log.Info("Wait until deploy transaction is mined")
	wait(tx.Hash())

	// 2. deploy EtherToken in root chain
	etherTokenContract, tx, etherToken, err := ethertoken.DeployEtherToken(opt, backend, development, mintableTokenContract, false)
	if err != nil {
		return common.Address{}, errors.New(fmt.Sprintf("Failed to deploy EtherToken contract: %v", err))
	}
	log.Info("Deploy EtherToken contract", "hash", tx.Hash(), "address", etherTokenContract)

	log.Info("Wait until deploy transaction is mined")
	wait(tx.Hash())

	// 3. deploy EpochHandler in root chain
	epochHandlerContract, tx, _, err := epochhandler.DeployEpochHandler(opt, backend)
	if err != nil {
		return common.Address{}, errors.New(fmt.Sprintf("Failed to deploy EpochHandler contract: %v", err))
	}
	log.Info("Deploy EpochHandler contract", "hash", tx.Hash(), "address", epochHandlerContract)

	log.Info("Wait until deploy transaction is mined")
	wait(tx.Hash())

	// 4. deploy RootChain in root chain
	rootchainContract, tx, _, err := rootchain.DeployRootChain(opt, backend, epochHandlerContract, etherTokenContract, development, NRELength, dummyBlock.Root(), dummyBlock.TxHash(), dummyBlock.ReceiptHash())
	if err != nil {
		return common.Address{}, errors.New(fmt.Sprintf("Failed to deploy RootChain contract: %v", err))
	}
	log.Info("Deploy RootChain contract", "hash", tx.Hash(), "address", rootchainContract)
	wait(tx.Hash())

	// 5. initialize EtherToken
	tx, err = etherToken.Init(opt, rootchainContract)
	if err != nil {
		return common.Address{}, errors.New(fmt.Sprintf("Failed to initialize EtherToken: %v", err))
	}
	log.Info("Initialize EtherToken", "hash", tx.Hash())
	wait(tx.Hash())

	return rootchainContract, nil
}
