// Copyright 2016 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind"
	"github.com/Onther-Tech/plasma-evm/accounts/keystore"
	"github.com/Onther-Tech/plasma-evm/cmd/utils"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/stakingmanager"
	"github.com/Onther-Tech/plasma-evm/core"
	"github.com/Onther-Tech/plasma-evm/core/rawdb"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/ethclient"
	"github.com/Onther-Tech/plasma-evm/ethdb"
	"github.com/Onther-Tech/plasma-evm/log"
	"strings"

	"github.com/Onther-Tech/plasma-evm/contracts/plasma"
	"gopkg.in/urfave/cli.v1"
	"math/big"
)

var (
	stakingCmd = cli.Command{
		Name:     "staking",
		Usage:    "Stake TON",
		Category: "TON STAKING COMMANDS",
		Description: `

Manage staking-related actions in the root chain.
`,
		Subcommands: []cli.Command{
			{
				Name:      "deployManagers",
				Usage:     "Deploy staking manager contract",
				ArgsUsage: "<withdrawalDelay> <seigPerBlock>",
				Action:    utils.MigrateFlags(deployManagers),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.OperatorAddressFlag,
					utils.OperatorKeyFlag,
					utils.RootChainTONFlag,
					utils.RootChainWTONFlag,
					utils.DeveloperKeyFlag,
				},
				Description: `
    geth staking deployManagers <withdrawalDelay> <seigPerBlock>

Deploy new manager contracts.

NOTE:
use --rootchain.ton, --rootchain.wton flags to use already deployed token contracts
`,
			},
			{
				Name:     "getManagers",
				Usage:    "Get staking managers addresses in database",
				Action:   utils.MigrateFlags(getManagers),
				Category: "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
				},
				Description: `
    geth staking getManagers

Get staking contract addresses
`,
			},
			{
				Name:     "setManagers",
				Usage:    "Set staking managers addresses in database",
				Action:   utils.MigrateFlags(setManagers),
				Category: "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.RootChainTONFlag,
					utils.RootChainWTONFlag,
					utils.RootChainDepositManagerFlag,
					utils.RootChainRegistryFlag,
					utils.RootChainSeigManagerFlag,
				},
				Description: `
    geth staking deployManagers <withdrawalDelay> <seigPerBlock>

Deploy new manager contracts.

NOTE:
use --rootchain.ton, --rootchain.wton flags to use already deployed token contracts
`,
			},
			{
				Name:     "register",
				Usage:    "Register RootChain contract",
				Action:   utils.MigrateFlags(registerRootChain),
				Category: "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.OperatorAddressFlag,
					utils.OperatorKeyFlag,
				},
				Description: `
Register RootChain contract to RootChainRegistry`,
			},
			{
				Name:      "swapFromTON",
				Usage:     "Swap TON with WTON",
				ArgsUsage: "<amount>",
				Action:    utils.MigrateFlags(deployManagers),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.OperatorAddressFlag,
					utils.OperatorKeyFlag,
					utils.RootChainTONFlag,
				},
				Description: `
Change TON to WTON`,
			},
			{
				Name:      "swapToTON",
				Usage:     "Swap WTON with TON",
				ArgsUsage: "<amount>",
				Action:    utils.MigrateFlags(deployManagers),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.OperatorAddressFlag,
					utils.OperatorKeyFlag,
					utils.RootChainTONFlag,
				},
				Description: `
Change WTON to TON`,
			},
			{
				Name:      "stake",
				Usage:     "Stake WTON",
				ArgsUsage: "<amount>",
				Action:    utils.MigrateFlags(deployManagers),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.OperatorAddressFlag,
					utils.OperatorKeyFlag,
					utils.RootChainTONFlag,
				},
				Description: `
Stake WTON`,
			},
			{
				Name:      "unstake",
				Usage:     "Untake WTON",
				ArgsUsage: "<amount>",
				Action:    utils.MigrateFlags(deployManagers),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.OperatorAddressFlag,
					utils.OperatorKeyFlag,
					utils.RootChainTONFlag,
				},
				Description: `
Stake WTON`,
			},
			{
				Name:      "requestWithdrawal",
				Usage:     "Make a withdrawal request",
				ArgsUsage: "<amount>",
				Action:    utils.MigrateFlags(deployManagers),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.OperatorAddressFlag,
					utils.OperatorKeyFlag,
					utils.RootChainTONFlag,
				},
				Description: `
Stake WTON`,
			},
			{
				Name:      "processWithdrawal",
				Usage:     "Process pending withdrawals",
				ArgsUsage: "<numRequests>",
				Action:    utils.MigrateFlags(deployManagers),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.OperatorAddressFlag,
					utils.OperatorKeyFlag,
					utils.RootChainTONFlag,
				},
				Description: `
Stake WTON`,
			},
			{
				Name:      "stats",
				Usage:     "Untake WTON",
				ArgsUsage: "<amount>",
				Action:    utils.MigrateFlags(deployManagers),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.OperatorAddressFlag,
					utils.OperatorKeyFlag,
					utils.RootChainTONFlag,
				},
				Description: `
Stake WTON`,
			},
		},
	}
)

type ManagerConfig struct {
	TON               common.Address `json:TON`
	WTON              common.Address `json:WTON`
	DepositManager    common.Address `json:DepositManager`
	RootChainRegistry common.Address `json:RootChainRegistry`
	SeigManager       common.Address `json:SeigManager`
}

func getManagerConfig(reader ethdb.Reader) *ManagerConfig {
	return &ManagerConfig{
		TON:               rawdb.ReadTON(reader),
		WTON:              rawdb.ReadWTON(reader),
		DepositManager:    rawdb.ReadDepositManager(reader),
		RootChainRegistry: rawdb.ReadRegistry(reader),
		SeigManager:       rawdb.ReadSeigManager(reader),
	}
}

func deployManagers(ctx *cli.Context) error {
	stack, cfg := makeConfigNode(ctx)

	// TODO: other database name for light client?
	chaindb, err := stack.OpenDatabase("chaindata", 0, 0, "")
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
		return err
	}

	if len(ctx.Args()) != 2 {
		utils.Fatalf("Expected 2 parameters, not %d", len(ctx.Args()))
	}

	var (
		withdrawalDelay *big.Int
		seigPerBlock    *big.Int
	)

	withdrawalDelayStr := ctx.Args()[0]
	seigPerBlockStr := ctx.Args()[1]

	// parse int string
	withdrawalDelay, ok := big.NewInt(0).SetString(withdrawalDelayStr, 10)
	if !ok {
		return errors.New(fmt.Sprintf("Failed to parse integer: %s", withdrawalDelayStr))
	}

	// parse float string e.g., 12.4 to RAY value
	if i := strings.Index(seigPerBlockStr, "."); i >= 0 {
		a := seigPerBlockStr[:i]
		b := seigPerBlockStr[i+1:]

		if 27-len(b) < 0 {
			return errors.New("RAY value out of precision")
		}

		r := strings.Repeat("0", 27-len(b))

		seigPerBlockStr = a + b + r
	}
	seigPerBlock, ok = big.NewInt(0).SetString(seigPerBlockStr, 10)

	if !ok {
		return errors.New(fmt.Sprintf("Failed to parse integer: %s", seigPerBlockStr))
	}

	_tonAddr := common.HexToAddress(ctx.GlobalString(utils.RootChainTONFlag.Name))
	_wtonAddr := common.HexToAddress(ctx.GlobalString(utils.RootChainWTONFlag.Name))

	ks := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)

	opt := bind.NewAccountTransactor(ks, cfg.Pls.Operator)
	opt.GasLimit = 7000000
	opt.GasPrice = utils.GlobalBig(ctx, utils.RootChainGasPriceFlag.Name)

	backend, err := ethclient.Dial(cfg.Pls.RootChainURL)

	if err != nil {
		utils.Fatalf("Failed to connect rootchain: %v", err)
	}

	tonAddr, wtonAddr, registryAddr, depositManagerAddr, seigManagerAddr, err := plasma.DeployManagers(opt, backend, withdrawalDelay, seigPerBlock, _tonAddr, _wtonAddr)

	if err != nil {
		return err
	}

	log.Info("Staking manager contract deployed", "TON", tonAddr, "WTON", wtonAddr, "RootChainRegistry", registryAddr, "DepositManager", depositManagerAddr, "SeigManager", seigManagerAddr)

	rawdb.WriteTON(chaindb, tonAddr)
	rawdb.WriteWTON(chaindb, wtonAddr)
	rawdb.WriteRegistry(chaindb, registryAddr)
	rawdb.WriteDepositManager(chaindb, depositManagerAddr)
	rawdb.WriteSeigManager(chaindb, seigManagerAddr)

	return nil
}

func getManagers(ctx *cli.Context) error {
	stack, _ := makeConfigNode(ctx)

	// TODO: other database name for light client?
	chaindb, err := stack.OpenDatabase("chaindata", 0, 0, "")
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
		return err
	}

	config := getManagerConfig(chaindb)

	b, err := json.MarshalIndent(config, "", "  ")

	if err != nil {
		return nil
	}

	fmt.Println(string(b))

	return nil
}

func setManagers(ctx *cli.Context) error {
	stack, _ := makeConfigNode(ctx)

	// TODO: other database name for light client?
	chaindb, err := stack.OpenDatabase("chaindata", 0, 0, "")
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
		return err
	}

	tonAddr := common.HexToAddress(ctx.GlobalString(utils.RootChainTONFlag.Name))
	wtonAddr := common.HexToAddress(ctx.GlobalString(utils.RootChainWTONFlag.Name))
	depositManagerAddr := common.HexToAddress(ctx.GlobalString(utils.RootChainDepositManagerFlag.Name))
	registryAddr := common.HexToAddress(ctx.GlobalString(utils.RootChainRegistryFlag.Name))
	seigManagerAddr := common.HexToAddress(ctx.GlobalString(utils.RootChainSeigManagerFlag.Name))

	type w struct {
		name string
		addr common.Address
		get  func(ethdb.Reader) common.Address
		set  func(ethdb.KeyValueWriter, common.Address)
	}

	targets := []w{
		{
			name: "TON",
			addr: tonAddr,
			get:  rawdb.ReadTON,
			set:  rawdb.WriteTON,
		},
		{
			name: "WTON",
			addr: wtonAddr,
			get:  rawdb.ReadWTON,
			set:  rawdb.WriteWTON,
		},
		{
			name: "DepositManager",
			addr: depositManagerAddr,
			get:  rawdb.ReadDepositManager,
			set:  rawdb.WriteDepositManager,
		},
		{
			name: "RootChainRegistry",
			addr: registryAddr,
			get:  rawdb.ReadRegistry,
			set:  rawdb.WriteRegistry,
		},
		{
			name: "SeigManager",
			addr: seigManagerAddr,
			get:  rawdb.ReadSeigManager,
			set:  rawdb.WriteSeigManager,
		},
	}

	for _, target := range targets {
		addr := target.get(chaindb)

		switch addr {
		case common.Address{}:
			log.Info("Set $s address", target.name)
			target.set(chaindb, target.addr)
		case target.addr:
			log.Info("%s address is already set as %s", target.name, addr.String())
		default:
			log.Error("%s address is already set as %s, not same with %s", target.name, addr.String(), target.addr)
		}
	}

	return nil
}

func registerRootChain(ctx *cli.Context) error {
	stack, cfg := makeConfigNode(ctx)
	log.Info("cfg.Pls", "cfg.Pls", cfg.Pls)

	// TODO: other database name for light client?
	chaindb, err := stack.OpenDatabase("chaindata", 0, 0, "")
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
		return err
	}

	config := getManagerConfig(chaindb)

	if (config.TON == common.Address{}) ||
		(config.WTON == common.Address{}) ||
		(config.DepositManager == common.Address{}) ||
		(config.RootChainRegistry == common.Address{}) ||
		(config.SeigManager == common.Address{}) {
		return errors.New("manager contract addresses is empty. please set contracts before register using `geth staking setManagers`")
	}

	data := rawdb.ReadGenesis(chaindb)
	genesis := new(core.Genesis)
	if err = json.Unmarshal(data, genesis); err != nil {
		return err
	}

	rootchainAddr := common.BytesToAddress(genesis.ExtraData)

	ks := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)

	opt := bind.NewAccountTransactor(ks, cfg.Pls.Operator)
	opt.GasLimit = 7000000
	opt.GasPrice = utils.GlobalBig(ctx, utils.RootChainGasPriceFlag.Name)

	backend, err := ethclient.Dial(cfg.Pls.RootChainURL)

	if err != nil {
		utils.Fatalf("Failed to connect rootchain: %v", err)
	}

	log.Info("Using manager contracts", "TON", config.TON, "WTON", config.WTON, "DepositManager", config.DepositManager, "RootChainRegistry", config.RootChainRegistry, "SeigManager", config.SeigManager)

	registry, err := stakingmanager.NewRootChainRegistry(config.RootChainRegistry, backend)
	if err != nil {
		return err
	}

	var tx *types.Transaction
	if tx, err = registry.Register(opt, rootchainAddr); err != nil {
		return err
	}
	log.Info("Register RootChain", "registry", config.RootChainRegistry, "rootchain", rootchainAddr, "tx", tx.Hash())

	if err = plasma.WaitTx(backend, tx.Hash()); err != nil {
		return err
	}

	log.Info("RootChain registered", "registry", config.RootChainRegistry, "rootchain", rootchainAddr, "tx", tx.Hash())

	return nil
}
