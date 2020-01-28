package plasma

//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/RootChain.sol --pkg rootchain --out rootchain/rootchain.go
//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/handlers/EpochHandler.sol --pkg epochhandler --out epochhandler/epochhandler.go
//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/handlers/SubmitHandler.sol --pkg submithandler --out submithandler/submithandler.go

//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/stake/TON.sol --pkg ton --out ton/ton.go
//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/stake/WTON.sol --pkg wton --out wton/wton.go
//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/stake/SeigManager.sol --pkg stakingmanager --out stakingmanager/stakingmanager.go

//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/RequestableSimpleToken.sol --pkg token --out token/token.go
//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/node_modules/openzeppelin-solidity/contracts/token/ERC20/ERC20Mintable.sol --pkg mintabletoken --out mintabletoken/mintabletoken.go

//go:generate ../../build/bin/abigen --sol plasma-evm-cotracts/contracts/EtherToken.sol --pkg ethertoken --out ethertoken/ethertoken.go

import (
	"bytes"
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
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/stakingmanager"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/submithandler"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/ton"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/wton"
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

	// 1. deploy MintableToken
	mintableTokenContract, tx, _, err := mintabletoken.DeployERC20Mintable(opt, backend)
	if err != nil {
		return common.Address{}, errors.New(fmt.Sprintf("Failed to deploy MintableToken contract: %v", err))
	}
	log.Info("Deploy MintableToken contract", "hash", tx.Hash(), "address", mintableTokenContract)

	log.Info("Wait until deploy transaction is mined")
	wait(tx.Hash())

	// 2. deploy EtherToken
	etherTokenContract, tx, etherToken, err := ethertoken.DeployEtherToken(opt, backend, development, mintableTokenContract, false)
	if err != nil {
		return common.Address{}, errors.New(fmt.Sprintf("Failed to deploy EtherToken contract: %v", err))
	}
	log.Info("Deploy EtherToken contract", "hash", tx.Hash(), "address", etherTokenContract)

	log.Info("Wait until deploy transaction is mined")
	wait(tx.Hash())

	// 3. deploy EpochHandler
	epochHandlerContract, tx, _, err := epochhandler.DeployEpochHandler(opt, backend)
	if err != nil {
		return common.Address{}, errors.New(fmt.Sprintf("Failed to deploy EpochHandler contract: %v", err))
	}
	log.Info("Deploy EpochHandler contract", "hash", tx.Hash(), "address", epochHandlerContract)

	log.Info("Wait until deploy transaction is mined")
	wait(tx.Hash())

	// 4. deploy SubmitHandler
	submitHandlerContract, tx, _, err := submithandler.DeploySubmitHandler(opt, backend, epochHandlerContract)
	if err != nil {
		return common.Address{}, errors.New(fmt.Sprintf("Failed to deploy SubmitHandler contract: %v", err))
	}
	log.Info("Deploy EpochHandler contract", "hash", tx.Hash(), "address", epochHandlerContract)

	log.Info("Wait until deploy transaction is mined")
	wait(tx.Hash())

	// 5. deploy RootChain
	rootchainContract, tx, _, err := rootchain.DeployRootChain(opt, backend, epochHandlerContract, submitHandlerContract, etherTokenContract, development, NRELength, dummyBlock.Root(), dummyBlock.TxHash(), dummyBlock.ReceiptHash())
	if err != nil {
		return common.Address{}, errors.New(fmt.Sprintf("Failed to deploy RootChain contract: %v", err))
	}
	log.Info("Deploy RootChain contract", "hash", tx.Hash(), "address", rootchainContract)
	wait(tx.Hash())

	// 6. initialize EtherToken
	tx, err = etherToken.Init(opt, rootchainContract)
	if err != nil {
		return common.Address{}, errors.New(fmt.Sprintf("Failed to initialize EtherToken: %v", err))
	}
	log.Info("Initialize EtherToken", "hash", tx.Hash())
	wait(tx.Hash())

	return rootchainContract, nil
}

type seigManagerSetter interface {
	SetSeigManager(opts *bind.TransactOpts, _seigManager common.Address) (*types.Transaction, error)
	SeigManager(opts *bind.CallOpts) (common.Address, error)
}

func DeployManagers(
	opt *bind.TransactOpts,
	backend *ethclient.Client,
	withdrawalDelay *big.Int,
	seigPerBlock *big.Int,
	_tonAddr common.Address,
	_wtonAddr common.Address,
) (
	tonAddr common.Address,
	wtonAddr common.Address,
	registryAddr common.Address,
	depositManagerAddr common.Address,
	seigManagerAddr common.Address,
	err error,
) {
	opt.GasLimit = 7000000

	var (
		TON            *ton.TON
		WTON           *wton.WTON
		registry       *stakingmanager.RootChainRegistry
		depositManager *stakingmanager.DepositManager
		seigManager    *stakingmanager.SeigManager

		tx *types.Transaction
	)

	// 1. deploy TON
	if (_tonAddr == common.Address{}) {
		log.Info("1. deploy TON contract")
		if tonAddr, tx, TON, err = ton.DeployTON(opt, backend); err != nil {
			err = errors.New(fmt.Sprintf("Failed to deploy TON: %v", err))
			return
		}

		if err = wait(backend, tx.Hash()); err != nil {
			err = errors.New(fmt.Sprintf("Failed to deploy TON: %v", err))
			return
		}
	} else {
		tonAddr = _tonAddr
		log.Info("use TON contract at %s", tonAddr.String())
		if TON, err = ton.NewTON(tonAddr, backend); err != nil {
			err = errors.New(fmt.Sprintf("Failed to instantiate TON: %v", err))
			return
		}
	}

	// 2. deploy WTON
	if (_wtonAddr == common.Address{}) {
		log.Info("2. deploy WTON contract")
		if wtonAddr, tx, WTON, err = wton.DeployWTON(opt, backend, tonAddr); err != nil {
			err = errors.New(fmt.Sprintf("Failed to deploy WTON: %v", err))
			return
		}

		if err = wait(backend, tx.Hash()); err != nil {
			err = errors.New(fmt.Sprintf("Failed to deploy WTON: %v", err))
			return
		}
	} else {
		wtonAddr = _wtonAddr
		log.Info("use WTON contract at %s", wtonAddr.String())
		if WTON, err = wton.NewWTON(wtonAddr, backend); err != nil {
			err = errors.New(fmt.Sprintf("Failed to instantiate WTON: %v", err))
			return
		}
	}

	if addr, _ := WTON.SeigManager(&bind.CallOpts{Pending: false}); (addr != common.Address{}) {
		err = errors.New("WTON already set SeigManager")
		return
	}

	// 3. deploy RootChainRegistry
	if registryAddr, tx, registry, err = stakingmanager.DeployRootChainRegistry(opt, backend); err != nil {
		err = errors.New(fmt.Sprintf("Failed to deploy RootChainRegistry: %v", err))
		return
	}

	if err = wait(backend, tx.Hash()); err != nil {
		err = errors.New(fmt.Sprintf("Failed to deploy RootChainRegistry: %v", err))
		return
	}

	// 4. deploy DepositManager
	if depositManagerAddr, tx, depositManager, err = stakingmanager.DeployDepositManager(opt, backend, wtonAddr, registryAddr, withdrawalDelay); err != nil {
		err = errors.New(fmt.Sprintf("Failed to deploy DepositManager: %v", err))
		return
	}

	if err = wait(backend, tx.Hash()); err != nil {
		err = errors.New(fmt.Sprintf("Failed to deploy DepositManager: %v", err))
		return
	}

	// 5. deploy SeigManager
	if seigManagerAddr, tx, seigManager, err = stakingmanager.DeploySeigManager(opt, backend, tonAddr, wtonAddr, registryAddr, depositManagerAddr, seigPerBlock); err != nil {
		err = errors.New(fmt.Sprintf("Failed to deploy SeigManager: %v", err))
		return
	}

	if err = wait(backend, tx.Hash()); err != nil {
		err = errors.New(fmt.Sprintf("Failed to deploy SeigManager: %v", err))
		return
	}

	// 6. add WTON minter role to SeigManager
	if tx, err = WTON.AddMinter(opt, seigManagerAddr); err != nil {
		err = errors.New(fmt.Sprintf("Failed to add WTON minter role to SeigManager: %v", err))
		return
	}
	if err = wait(backend, tx.Hash()); err != nil {
		err = errors.New(fmt.Sprintf("Failed to add WTON minter role to SeigManager: %v", err))
		return
	}

	// 7. set seig manager to contracts
	var contracts = []seigManagerSetter{depositManager, WTON}
	for _, c := range contracts {
		if tx, err = c.SetSeigManager(opt, seigManagerAddr); err != nil {
			err = errors.New(fmt.Sprintf("Failed to set SeigManager: %v", err))
			return
		}
		if err = wait(backend, tx.Hash()); err != nil {
			err = errors.New(fmt.Sprintf("Failed to set SeigManager: %v", err))
			return
		}
	}

	return
}

func wait(backend *ethclient.Client, hash common.Hash) error {
	var receipt *types.Receipt

	<-time.NewTimer(1 * time.Second).C

	for receipt, _ = backend.TransactionReceipt(context.Background(), hash); receipt == nil; {
		<-time.NewTimer(1 * time.Second).C

		receipt, _ = backend.TransactionReceipt(context.Background(), hash)
	}

	if receipt.Status == 0 {
		return errors.New(fmt.Sprintf("Transaction reverted: %s", receipt.TxHash.String()))
	}
	return nil
}
