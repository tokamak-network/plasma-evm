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
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"strings"

	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind"
	"github.com/Onther-Tech/plasma-evm/accounts/keystore"
	"github.com/Onther-Tech/plasma-evm/cmd/utils"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/depositmanager"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/rootchain"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/rootchainregistry"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/seigmanager"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/ton"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/wton"
	"github.com/Onther-Tech/plasma-evm/core"
	"github.com/Onther-Tech/plasma-evm/core/rawdb"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/ethclient"
	"github.com/Onther-Tech/plasma-evm/ethdb"
	"github.com/Onther-Tech/plasma-evm/log"
	"github.com/Onther-Tech/plasma-evm/params"

	"gopkg.in/urfave/cli.v1"
)

// TODO: use other logger
// TODO: --rootchain.from flag unlock password..!
// TODO: return error -> utils.Fatalf
// TODO: refactor (init, ...)
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
					utils.RootChainGasPriceFlag,
				},
				Description: `
    geth staking deployManagers <withdrawalDelay> <seigPerBlock>

Deploy new manager contracts.

NOTE:
use --rootchain.ton, --rootchain.wton flags to use already deployed token contracts
`,
			},
			{
				Name:      "getManagers",
				Usage:     "Get staking managers addresses in database",
				Action:    utils.MigrateFlags(getManagers),
				ArgsUsage: " <configURL?>",
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
				},
				Description: `
    geth staking getManagers

Get staking contract addresses
`,
			},
			{
				Name:      "setManagers",
				Usage:     "Set staking managers addresses in database",
				ArgsUsage: " <configURL?>",
				Action:    utils.MigrateFlags(setManagers),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.RootChainTONFlag,
					utils.RootChainWTONFlag,
					utils.RootChainDepositManagerFlag,
					utils.RootChainRegistryFlag,
					utils.RootChainSeigManagerFlag,
					utils.RootChainGasPriceFlag,
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
					utils.RootChainGasPriceFlag,
				},
				Description: `
Register RootChain contract to RootChainRegistry`,
			},
			{
				Name:      "balances",
				Usage:     "Print balances of token and stake",
				ArgsUsage: " <to>",
				Action:    utils.MigrateFlags(getBalances),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.OperatorAddressFlag,
					utils.OperatorKeyFlag,
					utils.RootChainTONFlag,
					utils.RootChainGasPriceFlag,
				},
				Description: `
Mint TON to account`,
			},
			{
				Name:      "mintTON",
				Usage:     "Mint TON to account",
				ArgsUsage: " <to> <amount>",
				Action:    utils.MigrateFlags(mintTON),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.OperatorAddressFlag,
					utils.OperatorKeyFlag,
					utils.RootChainTONFlag,
					utils.RootChainGasPriceFlag,
				},
				Description: `
Mint TON to account`,
			},
			{
				Name:      "swapFromTON",
				Usage:     "Swap TON with WTON",
				ArgsUsage: "<tonAmount>",
				Action:    utils.MigrateFlags(swapFromTON),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.OperatorAddressFlag,
					utils.OperatorKeyFlag,
					utils.RootChainGasPriceFlag,
				},
				Description: `
Change TON to WTON

NOTE: <tonAmount> is in WAD, (decialms 18)
`,
			},
			{
				Name:      "swapToTON",
				Usage:     "Swap WTON with TON",
				ArgsUsage: "<wtonAmount>",
				Action:    utils.MigrateFlags(swapToTON),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.OperatorAddressFlag,
					utils.OperatorKeyFlag,
					utils.RootChainGasPriceFlag,
				},
				Description: `
Change WTON to TON

NOTE: <wtonAmount> is in RAY, (decialms 27)
`,
			},
			{
				Name:      "stake",
				Usage:     "Stake WTON",
				ArgsUsage: "<amount>",
				Action:    utils.MigrateFlags(stakeWTON),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.OperatorAddressFlag,
					utils.OperatorKeyFlag,
					utils.RootChainGasPriceFlag,
				},
				Description: `
Stake WTON`,
			},
			{
				Name:      "requestWithdrawal",
				Usage:     "Make a withdrawal request",
				ArgsUsage: "<amount?>",
				Action:    utils.MigrateFlags(requestWithdrawal),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.OperatorAddressFlag,
					utils.OperatorKeyFlag,
					utils.RootChainTONFlag,
				},
				Description: `
Make an unstaking request`,
			},
			{
				Name:      "processWithdrawal",
				Usage:     "Process pending withdrawals",
				ArgsUsage: "<numRequests?>",
				Action:    utils.MigrateFlags(processWithdrawal),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.OperatorAddressFlag,
					utils.OperatorKeyFlag,
					utils.RootChainTONFlag,
				},
				Description: `
Process unstaking requests`,
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

func parseIntString(str string, decimals int) string {
	if decimals != 18 && decimals != 27 {
		utils.Fatalf("decimals should be 18 or 27, not %d", decimals)
	}
	i := strings.Index(str, ".")
	if i < 0 {
		return str
	}

	a := str[:i]
	b := str[i+1:]
	n := decimals - len(b)

	if n < 0 {
		utils.Fatalf("decimals out of precision: %d", decimals)
	}

	r := strings.Repeat("0", n)

	return a + b + r
}

func bigIntToString(v *big.Int, decimals int) string {
	if v.Cmp(big.NewInt(0)) == 0 {
		return "0"
	}

	if decimals != 18 && decimals != 27 {
		utils.Fatalf("decimals should be 18 or 27, not %d", decimals)
	}

	p := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)

	q, d := new(big.Int).DivMod(v, p, new(big.Int))

	return q.String() + "." + d.String()
}

func toWAD(v *big.Int) *big.Int {
	p := new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)
	q, _ := new(big.Int).DivMod(v, p, new(big.Int))
	return q
}

func toRAY(v *big.Int) *big.Int {
	p := new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)
	return new(big.Int).Mul(v, p)
}

func deployManagers(ctx *cli.Context) error {
	if len(ctx.Args()) != 2 {
		utils.Fatalf("Expected 2 parameters, not %d", len(ctx.Args()))
	}

	stack, cfg := makeConfigNode(ctx)

	chaindb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
		return err
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
	seigPerBlockStr = parseIntString(seigPerBlockStr, 27)

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
	configPath := ctx.Args().First()

	stack, _ := makeConfigNode(ctx)

	chaindb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
		return err
	}

	managers := getManagerConfig(chaindb)

	b, err := json.MarshalIndent(managers, "", "  ")
	data := string(b)

	if err != nil {
		return nil
	}

	if len(configPath) != 0 {
		log.Info("Exporting manager contracts", "path", configPath)

		fh, err := os.OpenFile(configPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return err
		}
		defer fh.Close()

		var writer io.Writer = fh

		if _, err := io.WriteString(writer, string(data)); err != nil {
			return err
		}

		log.Info("Exported manager contracts", "path", configPath)
	}

	fmt.Println(data)

	return nil
}

func setManagers(ctx *cli.Context) error {
	configPath := ctx.Args().First()
	managers := new(ManagerConfig)

	if len(configPath) != 0 {
		var r io.Reader

		if strings.HasPrefix(configPath, "http") {
			res, err := http.Get(configPath)
			if err != nil {
				utils.Fatalf("Failed to read response: %v", err)
			}
			r = res.Body
		} else {
			f, err := os.Open(configPath)
			if err != nil {
				utils.Fatalf("Failed to read managers file: %v", err)
			}
			defer f.Close()
			r = f
		}

		if err := json.NewDecoder(r).Decode(managers); err != nil {
			utils.Fatalf("invalid managers file: %v", err)
		}
	}

	if ctx.GlobalIsSet(utils.RootChainTONFlag.Name) {
		tonAddr := common.HexToAddress(ctx.GlobalString(utils.RootChainTONFlag.Name))
		if (managers.TON != common.Address{}) && managers.TON != tonAddr {
			log.Warn("Override TON address", "previous", managers.TON, "current", tonAddr)
		}
		managers.TON = tonAddr
	}
	if ctx.GlobalIsSet(utils.RootChainWTONFlag.Name) {
		wtonAddr := common.HexToAddress(ctx.GlobalString(utils.RootChainWTONFlag.Name))
		if (managers.WTON != common.Address{}) && managers.WTON != wtonAddr {
			log.Warn("Override WTON address", "previous", managers.WTON, "current", wtonAddr)
		}
		managers.WTON = wtonAddr
	}
	if ctx.GlobalIsSet(utils.RootChainDepositManagerFlag.Name) {
		depositManagerAddr := common.HexToAddress(ctx.GlobalString(utils.RootChainDepositManagerFlag.Name))
		if (managers.DepositManager != common.Address{}) && managers.DepositManager != depositManagerAddr {
			log.Warn("Override DepositManager address", "previous", managers.DepositManager, "current", depositManagerAddr)
		}
		managers.DepositManager = depositManagerAddr
	}
	if ctx.GlobalIsSet(utils.RootChainRegistryFlag.Name) {
		registryAddr := common.HexToAddress(ctx.GlobalString(utils.RootChainRegistryFlag.Name))
		if (managers.RootChainRegistry != common.Address{}) && managers.RootChainRegistry != registryAddr {
			log.Warn("Override RootChainRegistry address", "previous", managers.RootChainRegistry, "current", registryAddr)
		}
		managers.RootChainRegistry = registryAddr
	}
	if ctx.GlobalIsSet(utils.RootChainSeigManagerFlag.Name) {
		seigManagerAddr := common.HexToAddress(ctx.GlobalString(utils.RootChainSeigManagerFlag.Name))
		if (managers.SeigManager != common.Address{}) && managers.SeigManager != seigManagerAddr {
			log.Warn("Override SeigManager address", "previous", managers.SeigManager, "current", seigManagerAddr)
		}
		managers.SeigManager = seigManagerAddr
	}

	stack, _ := makeConfigNode(ctx)

	chaindb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
		return err
	}

	type w struct {
		name  string
		addr  common.Address
		read  func(ethdb.Reader) common.Address
		write func(ethdb.KeyValueWriter, common.Address)
	}

	targets := []w{
		{
			name:  "TON",
			addr:  managers.TON,
			read:  rawdb.ReadTON,
			write: rawdb.WriteTON,
		},
		{
			name:  "WTON",
			addr:  managers.WTON,
			read:  rawdb.ReadWTON,
			write: rawdb.WriteWTON,
		},
		{
			name:  "DepositManager",
			addr:  managers.DepositManager,
			read:  rawdb.ReadDepositManager,
			write: rawdb.WriteDepositManager,
		},
		{
			name:  "RootChainRegistry",
			addr:  managers.RootChainRegistry,
			read:  rawdb.ReadRegistry,
			write: rawdb.WriteRegistry,
		},
		{
			name:  "SeigManager",
			addr:  managers.SeigManager,
			read:  rawdb.ReadSeigManager,
			write: rawdb.WriteSeigManager,
		},
	}

	for _, target := range targets {
		// short circuit if no target address is provided
		if (target.addr == common.Address{}) {
			continue
		}

		addr := target.read(chaindb)

		switch addr {
		case common.Address{}:
			log.Info("Set address", "name", target.name, "addr", target.addr)
			target.write(chaindb, target.addr)
		case target.addr:
			log.Info("Address is already set", "name", target.name, "addr", target.addr)
		default:
			log.Error("Known address is different", "name", target.name, "known", addr, "target", target.addr)
		}
	}

	return nil
}

func getRootChainAddr(datadir string) (rootchainAddr common.Address, err error) {
	path := filepath.Join(datadir, "geth", "genesis.json")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	genesis := new(core.Genesis)
	if err = json.Unmarshal(data, genesis); err != nil {
		return
	}

	rootchainAddr = common.BytesToAddress(genesis.ExtraData)
	if (rootchainAddr == common.Address{}) {
		utils.Fatalf("RootChain address is NULL ADDRESS")
	}
	return
}

func registerRootChain(ctx *cli.Context) error {
	stack, cfg := makeConfigNode(ctx)

	chaindb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}

	managers := getManagerConfig(chaindb)

	if (managers.RootChainRegistry == common.Address{}) || (managers.SeigManager == common.Address{}) {
		return errors.New("manager contract addresses is empty. please write contracts before register using `geth staking setManagers`")
	}

	rootchainAddr, err := getRootChainAddr(cfg.Node.DataDir)
	if err != nil {
		return err
	}

	ks := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)

	opt := bind.NewAccountTransactor(ks, cfg.Pls.Operator)
	opt.GasLimit = 7000000
	opt.GasPrice = utils.GlobalBig(ctx, utils.RootChainGasPriceFlag.Name)

	backend, err := ethclient.Dial(cfg.Pls.RootChainURL)

	if err != nil {
		utils.Fatalf("Failed to connect rootchain: %v", err)
	}

	log.Info("Using manager contracts", "TON", managers.TON, "WTON", managers.WTON, "DepositManager", managers.DepositManager, "RootChainRegistry", managers.RootChainRegistry, "SeigManager", managers.SeigManager)

	// load contract instances
	registry, err := rootchainregistry.NewRootChainRegistry(managers.RootChainRegistry, backend)
	if err != nil {
		utils.Fatalf("Failed to load RootChainRegistry contract: %v", err)
	}
	rootchainCtr, err := rootchain.NewRootChain(rootchainAddr, backend)
	if err != nil {
		utils.Fatalf("Failed to load RootChain contract: %v", err)

	}

	// send transactions

	// 1. register SeigManager to RootChain
	f1 := func() error {
		var tx *types.Transaction
		var err error

		seigManagerAddr, err := rootchainCtr.SeigManager(&bind.CallOpts{Pending: false})
		if err != nil {
			return err
		}

		if seigManagerAddr == managers.SeigManager {
			log.Warn("SeigManager already registered to RootChain")
			return nil
		}

		if (seigManagerAddr != common.Address{}) && (seigManagerAddr != managers.SeigManager) {
			return errors.New("RootChain already write SeigManager to another contract: " + seigManagerAddr.String())
		}

		if tx, err = rootchainCtr.SetSeigManager(opt, managers.SeigManager); err != nil {
			return err
		}

		if err = plasma.WaitTx(backend, tx.Hash()); err != nil {
			return err
		}
		log.Info("Registered SeigManager to RootChain", "registry", managers.RootChainRegistry, "rootchain", rootchainAddr, "seigManager", managers.SeigManager, "tx", tx.Hash())

		return nil
	}

	// 2. register RootChain to SeigManager
	f2 := func() error {
		var tx *types.Transaction
		var err error

		registered, err := registry.Rootchains(&bind.CallOpts{Pending: false}, rootchainAddr)
		if err != nil {
			return err
		}

		if registered {
			log.Warn("RootChain already registered to SeigManager")
			return nil
		}

		if tx, err = registry.RegisterAndDeployCoinage(opt, rootchainAddr, managers.SeigManager); err != nil {
			return err
		}

		if err = plasma.WaitTx(backend, tx.Hash()); err != nil {
			return err
		}
		log.Info("Registered RootChain to SeigManager", "registry", managers.RootChainRegistry, "rootchain", rootchainAddr, "seigManager", managers.SeigManager, "tx", tx.Hash())

		return nil
	}

	if err = f1(); err != nil {
		return err
	}
	if err = f2(); err != nil {
		return err
	}

	return nil
}

// TODO: pending withdrawal amount
func getBalances(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		utils.Fatalf("Expected 2 parameters, not %d", len(ctx.Args()))
	}

	var (
		depositor common.Address
	)

	depositor = common.HexToAddress(ctx.Args()[0])

	stack, cfg := makeConfigNode(ctx)

	log.Info("cfg.Node.DataDir", "v", filepath.Join(cfg.Node.DataDir, "geth", "genesis.json"))

	chaindb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}

	managers := getManagerConfig(chaindb)

	if (managers.TON == common.Address{}) ||
		(managers.WTON == common.Address{}) ||
		(managers.DepositManager == common.Address{}) ||
		(managers.RootChainRegistry == common.Address{}) ||
		(managers.SeigManager == common.Address{}) {
		return errors.New("manager contract addresses is empty. please write contracts before register using `geth staking setManagers`")
	}

	backend, err := ethclient.Dial(cfg.Pls.RootChainURL)

	if err != nil {
		utils.Fatalf("Failed depositor connect rootchain: %v", err)
	}

	log.Info("Using manager contracts", "TON", managers.TON, "WTON", managers.WTON, "DepositManager", managers.DepositManager, "RootChainRegistry", managers.RootChainRegistry, "SeigManager", managers.SeigManager)

	opt := &bind.CallOpts{Pending: false}

	rootchainAddr, err := getRootChainAddr(cfg.Node.DataDir)
	if err != nil {
		return err
	}

	var (
		TON            *ton.TON
		WTON           *wton.WTON
		depositManager *depositmanager.DepositManager
		seigManager    *seigmanager.SeigManager

		tot     *seigmanager.ERC20
		coinage *seigmanager.ERC20

		tonBalance  *big.Int
		wtonBalance *big.Int

		accStaked   *big.Int
		accUnstaked *big.Int
		deposit     *big.Int

		numPendingRequests *big.Int
		pendingUnstaked    *big.Int

		totalStake          *big.Int
		totalStakeRootChain *big.Int

		uncomittedStakeOf *big.Int
		stakeOf           *big.Int
	)

	// load contract instances
	if TON, err = ton.NewTON(managers.TON, backend); err != nil {
		utils.Fatalf("Failed depositor load TON contract: %v", err)
	}
	if WTON, err = wton.NewWTON(managers.WTON, backend); err != nil {
		utils.Fatalf("Failed depositor load WTON contract: %v", err)
	}
	if depositManager, err = depositmanager.NewDepositManager(managers.DepositManager, backend); err != nil {
		utils.Fatalf("Failed depositor load DepositManager contract: %v", err)
	}
	if seigManager, err = seigmanager.NewSeigManager(managers.SeigManager, backend); err != nil {
		utils.Fatalf("Failed depositor load SeigManager contract: %v", err)
	}

	totAddr, err := seigManager.Tot(opt)
	if err != nil {
		utils.Fatalf("Failed depositor load tot address: %v", err)
	}
	coinageAddr, err := seigManager.Coinages(opt, rootchainAddr)
	if err != nil {
		utils.Fatalf("Failed depositor load coinage address: %v", err)
	}

	if tot, err = seigmanager.NewERC20(totAddr, backend); err != nil {
		utils.Fatalf("Failed depositor load tot contract: %v", err)
	}
	if coinage, err = seigmanager.NewERC20(coinageAddr, backend); err != nil {
		utils.Fatalf("Failed depositor load tot contract: %v", err)
	}

	// read balances
	if tonBalance, err = TON.BalanceOf(opt, depositor); err != nil {
		utils.Fatalf("Failed depositor read TON balance: %v", err)
	}
	if wtonBalance, err = WTON.BalanceOf(opt, depositor); err != nil {
		utils.Fatalf("Failed depositor read WTON balance: %v", err)
	}
	if accStaked, err = depositManager.AccStaked(opt, rootchainAddr, depositor); err != nil {
		utils.Fatalf("Failed depositor read accumulated stake: %v", err)
	}
	if accUnstaked, err = depositManager.AccUnstaked(opt, rootchainAddr, depositor); err != nil {
		utils.Fatalf("Failed depositor read accumulated unstake: %v", err)
	}
	if numPendingRequests, err = depositManager.NumPendingRequests(opt, rootchainAddr, depositor); err != nil {
		utils.Fatalf("Failed depositor read num pending requests: %v", err)
	}
	if pendingUnstaked, err = depositManager.PendingUnstaked(opt, rootchainAddr, depositor); err != nil {
		utils.Fatalf("Failed depositor read pending withdrawal amount: %v", err)
	}

	if totalStake, err = tot.TotalSupply(opt); err != nil {
		log.Warn("Failed depositor read total stake", "err", err)
		totalStake = big.NewInt(0)
	}
	if totalStakeRootChain, err = coinage.TotalSupply(opt); err != nil {
		log.Warn("Failed depositor read total stake of root chain", "err", err)
		totalStakeRootChain = big.NewInt(0)
	}
	if uncomittedStakeOf, err = seigManager.UncomittedStakeOf(opt, rootchainAddr, depositor); err != nil {
		log.Warn("Failed depositor read uncomitted stake", "err", err)
		uncomittedStakeOf = big.NewInt(0)
	}
	if stakeOf, err = seigManager.StakeOf(opt, rootchainAddr, depositor); err != nil {
		log.Warn("Failed depositor read stake", "err", err)
		stakeOf = big.NewInt(0)
	}

	deposit = new(big.Int).Sub(accStaked, accUnstaked)

	// print balances
	log.Info("TON Balance", "amount", bigIntToString(tonBalance, params.TONDecimals)+" TON", "depositor", depositor)
	log.Info("WON Balance", "amount", bigIntToString(wtonBalance, params.WTONDecimals)+" WTON", "depositor", depositor)
	log.Info("Deposit", "amount", bigIntToString(deposit, params.WTONDecimals)+" WTON", "rootchain", rootchainAddr, "depositor", depositor)

	log.Info("Pending withdrawal requests", "num", numPendingRequests)
	log.Info("Pending withdrawal WTON", "amount", bigIntToString(pendingUnstaked, params.WTONDecimals)+" WTON", "rootchain", rootchainAddr, "depositor", depositor)

	log.Info("Total Stake", "amount", bigIntToString(totalStake, params.WTONDecimals)+" WTON")
	log.Info("Total Stake of Root Chain", "amount", bigIntToString(totalStakeRootChain, params.WTONDecimals)+" WTON", "rootchain", rootchainAddr)

	log.Info("Uncomitted Stake", "amount", bigIntToString(uncomittedStakeOf, params.WTONDecimals)+" WTON", "rootchain", rootchainAddr, "depositor", depositor)
	log.Info("Comitted Stake", "amount", bigIntToString(stakeOf, params.WTONDecimals)+" WTON", "rootchain", rootchainAddr, "depositor", depositor)

	return nil
}

func mintTON(ctx *cli.Context) error {
	if len(ctx.Args()) != 2 {
		utils.Fatalf("Expected 2 parameters, not %d", len(ctx.Args()))
	}

	decimals := params.TONDecimals

	var (
		to     common.Address
		amount *big.Int

		ok bool
	)

	to = common.HexToAddress(ctx.Args()[0])
	amountStr := parseIntString(ctx.Args()[1], decimals)
	if amount, ok = big.NewInt(0).SetString(amountStr, 10); !ok {
		return errors.New(fmt.Sprintf("Failed to parse integer: %s", amountStr))
	}

	stack, cfg := makeConfigNode(ctx)

	chaindb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}

	managers := getManagerConfig(chaindb)

	if (managers.TON == common.Address{}) {
		return errors.New("manager contract addresses is empty. please write contracts before register using `geth staking setManagers`")
	}

	ks := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)

	opt := bind.NewAccountTransactor(ks, cfg.Pls.Operator)
	opt.GasPrice = utils.GlobalBig(ctx, utils.RootChainGasPriceFlag.Name)

	backend, err := ethclient.Dial(cfg.Pls.RootChainURL)

	if err != nil {
		utils.Fatalf("Failed to connect rootchain: %v", err)
	}

	log.Info("Using manager contracts", "TON", managers.TON, "WTON", managers.WTON, "DepositManager", managers.DepositManager, "RootChainRegistry", managers.RootChainRegistry, "SeigManager", managers.SeigManager)

	TON, err := ton.NewTON(managers.TON, backend)
	if err != nil {
		return err
	}

	var tx *types.Transaction
	if tx, err = TON.Mint(opt, to, amount); err != nil {
		return err
	}
	log.Info("Minting TON", "to", to, "amount", bigIntToString(amount, decimals)+" TON", "tx", tx.Hash())

	if err = plasma.WaitTx(backend, tx.Hash()); err != nil {
		return err
	}

	return nil
}

type approvable interface {
	Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error)
	Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error)
}

func approveToken(
	name string,
	contract approvable,
	backend *ethclient.Client,
	opts *bind.TransactOpts,
	spender common.Address, target *big.Int,
	decimals int,
) {

	current, err := contract.Allowance(&bind.CallOpts{Pending: false}, opts.From, spender)

	if current.Cmp(target) >= 0 {
		return
	}

	diff := new(big.Int).Sub(target, current)

	log.Warn("Allowances is inefficient", "current", bigIntToString(current, decimals), "target", bigIntToString(target, decimals), "diff", bigIntToString(diff, decimals))
	log.Warn(fmt.Sprintf("Approve to deposit %s", name), "amount", bigIntToString(target, decimals))

	var tx *types.Transaction

	if tx, err = contract.Approve(opts, spender, target); err != nil {
		utils.Fatalf("Failed to send transaction: %v", err)
	}
	if err = plasma.WaitTx(backend, tx.Hash()); err != nil {
		utils.Fatalf("Failed to send transaction: %v", err)
	}

	log.Warn(fmt.Sprintf("Approved to deposit %s", name), "amount", bigIntToString(target, decimals), "tx", tx.Hash())

}

func swapFromTON(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		utils.Fatalf("Expected 1 parameters, not %d", len(ctx.Args()))
	}

	decimals := params.TONDecimals

	var (
		amount *big.Int

		ok bool
	)

	amountStr := parseIntString(ctx.Args()[0], decimals)
	if amount, ok = big.NewInt(0).SetString(amountStr, 10); !ok {
		return errors.New(fmt.Sprintf("Failed to parse integer: %s", amountStr))
	}

	stack, cfg := makeConfigNode(ctx)

	chaindb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}

	managers := getManagerConfig(chaindb)

	if (managers.WTON == common.Address{}) || (managers.TON == common.Address{}) {
		return errors.New("manager contract addresses is empty. please write contracts before register using `geth staking setManagers`")
	}

	ks := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)

	opt := bind.NewAccountTransactor(ks, cfg.Pls.Operator)
	opt.GasPrice = utils.GlobalBig(ctx, utils.RootChainGasPriceFlag.Name)

	backend, err := ethclient.Dial(cfg.Pls.RootChainURL)

	if err != nil {
		utils.Fatalf("Failed to connect rootchain: %v", err)
	}

	log.Info("Using manager contracts", "TON", managers.TON, "WTON", managers.WTON, "DepositManager", managers.DepositManager, "RootChainRegistry", managers.RootChainRegistry, "SeigManager", managers.SeigManager)

	// load contract instances
	TON, err := ton.NewTON(managers.TON, backend)
	if err != nil {
		utils.Fatalf("Failed to load TON contract: %v", err)
	}
	WTON, err := wton.NewWTON(managers.WTON, backend)
	if err != nil {
		utils.Fatalf("Failed to load WTON contract: %v", err)
	}

	// check TON balance
	tonBalance, err := TON.BalanceOf(&bind.CallOpts{Pending: false}, opt.From)
	if err != nil {
		utils.Fatalf("Failed to read TON balance: %v", err)
	}

	if tonBalance.Cmp(amount) < 0 {
		utils.Fatalf("Insufficient TON Balance (%s)", bigIntToString(tonBalance, params.TONDecimals))
	}

	// send transaction(s)
	approveToken("TON", TON, backend, opt, managers.WTON, amount, params.TONDecimals)

	var tx *types.Transaction
	if tx, err = WTON.SwapFromTON(opt, amount); err != nil {
		utils.Fatalf("Failed to send transaction: %v", err)
	}
	log.Info("Swap from TON to WTON", "amount", bigIntToString(amount, decimals)+" TON", "from", opt.From, "tx", tx.Hash())

	if err = plasma.WaitTx(backend, tx.Hash()); err != nil {
		return err
	}

	return nil
}

func swapToTON(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		utils.Fatalf("Expected 1 parameters, not %d", len(ctx.Args()))
	}

	decimals := params.WTONDecimals

	var (
		amount *big.Int

		ok bool
	)

	amountStr := parseIntString(ctx.Args()[0], decimals)
	if amount, ok = big.NewInt(0).SetString(amountStr, 10); !ok {
		return errors.New(fmt.Sprintf("Failed to parse integer: %s", amountStr))
	}

	stack, cfg := makeConfigNode(ctx)

	chaindb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}

	managers := getManagerConfig(chaindb)

	if (managers.WTON == common.Address{}) || (managers.TON == common.Address{}) {
		return errors.New("manager contract addresses is empty. please write contracts before register using `geth staking setManagers`")
	}

	ks := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)

	opt := bind.NewAccountTransactor(ks, cfg.Pls.Operator)
	opt.GasPrice = utils.GlobalBig(ctx, utils.RootChainGasPriceFlag.Name)

	backend, err := ethclient.Dial(cfg.Pls.RootChainURL)

	if err != nil {
		utils.Fatalf("Failed to connect rootchain: %v", err)
	}

	log.Info("Using manager contracts", "TON", managers.TON, "WTON", managers.WTON, "DepositManager", managers.DepositManager, "RootChainRegistry", managers.RootChainRegistry, "SeigManager", managers.SeigManager)

	// load contract instance
	WTON, err := wton.NewWTON(managers.WTON, backend)
	if err != nil {
		return err
	}

	// check WTON balance
	wtonBalance, err := WTON.BalanceOf(&bind.CallOpts{Pending: false}, opt.From)
	if err != nil {
		utils.Fatalf("Failed to read WTON balance: %v", err)
	}

	if wtonBalance.Cmp(amount) < 0 {
		utils.Fatalf("Insufficient WTON Balance (%s)", bigIntToString(wtonBalance, params.WTONDecimals))
	}

	// send transaction
	var tx *types.Transaction
	if tx, err = WTON.SwapToTON(opt, amount); err != nil {
		return err
	}
	log.Info("Swap from WTON to TON", "amount", bigIntToString(amount, decimals)+" WTON", "from", opt.From, "tx", tx.Hash())

	if err = plasma.WaitTx(backend, tx.Hash()); err != nil {
		return err
	}

	return nil
}

func stakeWTON(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		utils.Fatalf("Expected 1 parameters, not %d", len(ctx.Args()))
	}

	decimals := params.WTONDecimals

	var (
		amount *big.Int

		ok bool

		tx *types.Transaction
	)

	amountStr := parseIntString(ctx.Args()[0], decimals)
	if amount, ok = big.NewInt(0).SetString(amountStr, 10); !ok {
		return errors.New(fmt.Sprintf("Failed to parse integer: %s", amountStr))
	}

	stack, cfg := makeConfigNode(ctx)

	chaindb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}

	rootchainAddr, err := getRootChainAddr(cfg.Node.DataDir)
	if err != nil {
		return err
	}

	managers := getManagerConfig(chaindb)

	if (managers.TON == common.Address{}) ||
		(managers.WTON == common.Address{}) ||
		(managers.DepositManager == common.Address{}) ||
		(managers.RootChainRegistry == common.Address{}) ||
		(managers.SeigManager == common.Address{}) {
		return errors.New("manager contract addresses is empty. please write contracts before register using `geth staking setManagers`")
	}

	ks := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)

	opt := bind.NewAccountTransactor(ks, cfg.Pls.Operator)
	opt.GasPrice = utils.GlobalBig(ctx, utils.RootChainGasPriceFlag.Name)

	backend, err := ethclient.Dial(cfg.Pls.RootChainURL)

	if err != nil {
		utils.Fatalf("Failed to connect rootchain: %v", err)
	}

	log.Info("Using manager contracts", "TON", managers.TON, "WTON", managers.WTON, "DepositManager", managers.DepositManager, "RootChainRegistry", managers.RootChainRegistry, "SeigManager", managers.SeigManager)

	var (
		WTON           *wton.WTON
		depositManager *depositmanager.DepositManager
	)

	// load contract instances
	if WTON, err = wton.NewWTON(managers.WTON, backend); err != nil {
		utils.Fatalf("Failed to load WTON contract: %v", err)
	}
	if depositManager, err = depositmanager.NewDepositManager(managers.DepositManager, backend); err != nil {
		utils.Fatalf("Failed to load DepositManager contract: %v", err)
	}

	// check WTON balance
	wtonBalance, err := WTON.BalanceOf(&bind.CallOpts{Pending: false}, opt.From)
	if err != nil {
		utils.Fatalf("Failed to read WTON balance: %v", err)
	}

	if wtonBalance.Cmp(amount) < 0 {
		utils.Fatalf("Insufficient WTON Balance (%s)", bigIntToString(wtonBalance, params.WTONDecimals))
	}

	// send transaction(s)
	approveToken("WTON", WTON, backend, opt, managers.DepositManager, amount, params.WTONDecimals)

	if tx, err = depositManager.Deposit(opt, rootchainAddr, amount); err != nil {
		return err
	}

	if err = plasma.WaitTx(backend, tx.Hash()); err != nil {
		return err
	}

	log.Info("Deposit WTON to RootChain", "rootchain", rootchainAddr, "amount", bigIntToString(amount, decimals)+" WTON", "tx", tx.Hash())

	return nil
}

func requestWithdrawal(ctx *cli.Context) error {
	if len(ctx.Args()) > 1 {
		utils.Fatalf("Expected 1 or 0 parameters, not %d", len(ctx.Args()))
	}

	decimals := params.WTONDecimals

	var (
		amount *big.Int

		ok bool

		tx *types.Transaction
	)

	if len(ctx.Args()) == 1 {
		amountStr := parseIntString(ctx.Args()[0], decimals)
		if amount, ok = big.NewInt(0).SetString(amountStr, 10); !ok {
			return errors.New(fmt.Sprintf("Failed to parse integer: %s", amountStr))
		}
	}

	stack, cfg := makeConfigNode(ctx)

	chaindb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}

	rootchainAddr, err := getRootChainAddr(cfg.Node.DataDir)
	if err != nil {
		return err
	}

	managers := getManagerConfig(chaindb)

	if (managers.TON == common.Address{}) ||
		(managers.WTON == common.Address{}) ||
		(managers.DepositManager == common.Address{}) ||
		(managers.RootChainRegistry == common.Address{}) ||
		(managers.SeigManager == common.Address{}) {
		return errors.New("manager contract addresses is empty. please write contracts before register using `geth staking setManagers`")
	}

	ks := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)

	opt := bind.NewAccountTransactor(ks, cfg.Pls.Operator)
	opt.GasPrice = utils.GlobalBig(ctx, utils.RootChainGasPriceFlag.Name)

	from := opt.From

	backend, err := ethclient.Dial(cfg.Pls.RootChainURL)

	if err != nil {
		utils.Fatalf("Failed to connect rootchain: %v", err)
	}

	log.Info("Using manager contracts", "TON", managers.TON, "WTON", managers.WTON, "DepositManager", managers.DepositManager, "RootChainRegistry", managers.RootChainRegistry, "SeigManager", managers.SeigManager)

	var (
		depositManager *depositmanager.DepositManager
		seigManager    *seigmanager.SeigManager
	)

	// load contract instances
	if depositManager, err = depositmanager.NewDepositManager(managers.DepositManager, backend); err != nil {
		utils.Fatalf("Failed to load DepositManager contract: %v", err)
	}
	if seigManager, err = seigmanager.NewSeigManager(managers.SeigManager, backend); err != nil {
		utils.Fatalf("Failed to load SeigManager contract: %v", err)
	}

	// check staked WTON balance
	staked, err := seigManager.StakeOf(&bind.CallOpts{Pending: false}, rootchainAddr, from)
	if err != nil {
		utils.Fatalf("Failed to read stake amount: %v", err)
	}

	// write all staked amount if no parameter given
	if amount == nil {
		amount = new(big.Int).Set(staked)
	}

	if staked.Cmp(amount) < 0 {
		utils.Fatalf("Insufficient Staked amount to withdraw (%s)", bigIntToString(staked, params.WTONDecimals))
	}

	if tx, err = depositManager.RequestWithdrawal(opt, rootchainAddr, amount); err != nil {
		return err
	}

	if err = plasma.WaitTx(backend, tx.Hash()); err != nil {
		return err
	}

	log.Info("Withdrawal requested", "rootchain", rootchainAddr, "amount", bigIntToString(amount, decimals)+" WTON", "tx", tx.Hash())

	return nil
}

func processWithdrawal(ctx *cli.Context) error {
	if len(ctx.Args()) > 1 {
		utils.Fatalf("Expected 1 or 0 parameters, not %d", len(ctx.Args()))
	}

	var (
		n int

		tx  *types.Transaction
		err error
	)

	if len(ctx.Args()) == 1 {
		if n, err = strconv.Atoi(ctx.Args()[0]); err != nil {
			return err
		}
	}

	stack, cfg := makeConfigNode(ctx)

	chaindb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}

	rootchainAddr, err := getRootChainAddr(cfg.Node.DataDir)
	if err != nil {
		return err
	}

	managers := getManagerConfig(chaindb)

	if (managers.TON == common.Address{}) ||
		(managers.WTON == common.Address{}) ||
		(managers.DepositManager == common.Address{}) ||
		(managers.RootChainRegistry == common.Address{}) ||
		(managers.SeigManager == common.Address{}) {
		return errors.New("manager contract addresses is empty. please write contracts before register using `geth staking setManagers`")
	}

	ks := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)

	opt := bind.NewAccountTransactor(ks, cfg.Pls.Operator)
	opt.GasPrice = utils.GlobalBig(ctx, utils.RootChainGasPriceFlag.Name)

	backend, err := ethclient.Dial(cfg.Pls.RootChainURL)

	if err != nil {
		utils.Fatalf("Failed to connect rootchain: %v", err)
	}

	log.Info("Using manager contracts", "TON", managers.TON, "WTON", managers.WTON, "DepositManager", managers.DepositManager, "RootChainRegistry", managers.RootChainRegistry, "SeigManager", managers.SeigManager)

	var (
		numPendingRequests *big.Int

		depositManager *depositmanager.DepositManager
	)

	// load contract instances
	if depositManager, err = depositmanager.NewDepositManager(managers.DepositManager, backend); err != nil {
		utils.Fatalf("Failed to load DepositManager contract: %v", err)
	}

	if numPendingRequests, err = depositManager.NumPendingRequests(&bind.CallOpts{Pending: false}, rootchainAddr, opt.From); err != nil {
		utils.Fatalf("Failed depositor read num pending requests: %v", err)
	}

	// check num pending requests
	if n == 0 || int(numPendingRequests.Int64()) > n {
		log.Warn("Set num requests", "n", numPendingRequests)
		n = int(numPendingRequests.Int64())
	}

	if n == 0 {
		utils.Fatalf("No request to process")
	}

	if tx, err = depositManager.ProcessRequests(opt, rootchainAddr, big.NewInt(int64((n))), false); err != nil {
		return err
	}

	if err = plasma.WaitTx(backend, tx.Hash()); err != nil {
		return err
	}

	// TODO: log request index and amount
	log.Info("Withdraw request processed", "rootchain", rootchainAddr, "tx", tx.Hash())

	return nil
}
