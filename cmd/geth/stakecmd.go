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
	"time"

	"strings"

	"github.com/Onther-Tech/plasma-evm/accounts"
	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind"
	"github.com/Onther-Tech/plasma-evm/accounts/keystore"
	"github.com/Onther-Tech/plasma-evm/cmd/utils"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/depositmanager"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/powerton"
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
	"github.com/Onther-Tech/plasma-evm/node"
	"github.com/Onther-Tech/plasma-evm/params"
	"github.com/Onther-Tech/plasma-evm/pls"

	"gopkg.in/urfave/cli.v1"
)

// TODO: use other logger
// TODO: --rootchain.from flag unlock password..!
// TODO: return error -> utils.Fatalf
// TODO: refactor (init, ...)
var (
	manageStakingCmd = cli.Command{
		Name:     "manage-staking",
		Usage:    "Manage Staking Contracts",
		Category: "TON STAKING MANAGE COMMANDS",
		Description: `

The manage-staking command deploys and set up contracts in TON ecosystem.
`,
		Subcommands: []cli.Command{
			{
				Name:      "deploy-managers",
				Usage:     "Deploy staking manager contracts (except PowerTON)",
				ArgsUsage: "<withdrawalDelay> <seigPerBlock>",
				Action:    utils.MigrateFlags(deployManagers),
				Category:  "TON STAKING MANAGE COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.UnlockedAccountFlag,
					utils.PasswordFileFlag,
					utils.RootChainSenderFlag,
					utils.RootChainTONFlag,
					utils.RootChainWTONFlag,
					utils.DeveloperKeyFlag,
					utils.RootChainGasPriceFlag,
				},
				Description: `
    geth manage-staking deployManagers <withdrawalDelay> <seigPerBlock>

Deploy new manager contracts.

NOTE:
set manager contracts or use --rootchain.ton, --rootchain.wton flags to use already deployed token contracts
`,
			},
			{
				Name:      "deploy-powerton",
				Usage:     "Deploy Power TON contract",
				ArgsUsage: "<roundDuration>",
				Action:    utils.MigrateFlags(deployPowerTON),
				Category:  "TON STAKING MANAGE COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.UnlockedAccountFlag,
					utils.PasswordFileFlag,
					utils.RootChainSenderFlag,
					utils.RootChainGasPriceFlag,
					utils.RootChainWTONFlag,
					utils.RootChainSeigManagerFlag,
				},
				Description: `
    geth manage-staking deployPowerTON <roundDuration>

Deploy new PowerTON contract.

NOTE:
set manager contracts or use --rootchain.wton, --rootchain.seigManager flags to use already deployed token contracts
`,
			},
			{
				Name:     "start-powerton",
				Usage:    "Start Power TON first round",
				Action:   utils.MigrateFlags(startPowerTON),
				Category: "TON STAKING MANAGE COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.UnlockedAccountFlag,
					utils.PasswordFileFlag,
					utils.RootChainSenderFlag,
					utils.RootChainPowerTONFlag,
					utils.DeveloperKeyFlag,
					utils.RootChainGasPriceFlag,
				},
				Description: `
    geth manage-staking startPowerTON

Start first round of PowerTON

NOTE:
set manager contracts or use --rootchain.powerton flag to use already deployed token contracts
`,
			},
			{
				Name:     "register",
				Usage:    "Register RootChain contract",
				Action:   utils.MigrateFlags(registerRootChain),
				Category: "TON STAKING MANAGE COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.UnlockedAccountFlag,
					utils.PasswordFileFlag,
					utils.RootChainSenderFlag,
					utils.DeveloperKeyFlag,
					utils.RootChainGasPriceFlag,
				},
				Description: `
				geth manage-staking register

Register RootChain contract to RootChainRegistry
`,
			},
			{
				Name:      "set-commission-rate",
				Usage:     "Set commission rate",
				Action:    utils.MigrateFlags(setCommissionRate),
				ArgsUsage: "<rate> <isCommissionRateNegative>",
				Category:  "TON STAKING MANAGE COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.UnlockedAccountFlag,
					utils.PasswordFileFlag,
					utils.RootChainSenderFlag,
					utils.DeveloperKeyFlag,
					utils.RootChainGasPriceFlag,
				},
				Description: `
				geth manage-staking set-commission-rate <rate> <isCommissionRateNegative>

Set commission rate of the root chain (operator only)

Default value of <isCommissionRateNegative> is false

Example:
	geth manage-staking set-commission-rate 0
	geth manage-staking set-commission-rate 0.1
	geth manage-staking set-commission-rate 0.1 false
	geth manage-staking set-commission-rate 0.1 true

NOTE:
rate should be 0 or between 0.01 and 1.00
`,
			},
			{
				Name:      "get-managers",
				Usage:     "Get staking managers addresses in database",
				Action:    utils.MigrateFlags(getManagers),
				ArgsUsage: "<path>",
				Category:  "TON STAKING MANAGE COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
				},
				Description: `
    geth manage-staking getManagers <path>

Get staking contracts addresses. If path is given, contracts are stored in the path as JSON.
`,
			},
			{
				Name:      "set-managers",
				Usage:     "Set staking managers addresses in database",
				ArgsUsage: "<uri>",
				Action:    utils.MigrateFlags(setManagers),
				Category:  "TON STAKING MANAGE COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.RootChainTONFlag,
					utils.RootChainWTONFlag,
					utils.RootChainDepositManagerFlag,
					utils.RootChainRegistryFlag,
					utils.RootChainSeigManagerFlag,
					utils.RootChainPowerTONFlag,
					utils.RootChainGasPriceFlag,
				},
				Description: `
    geth manage-staking setManagers <uri>

Set staking contracts addresses

NOTE:
use --rootchain.ton, --rootchain.wton, --rootchain.depositmanager, --rootchain.registry, --rootchain.seigmanager, --rootchain.powerton flags to use already deployed token contracts
`,
			},
			{
				Name:      "mint-ton",
				Usage:     "Mint TON to account (for dev)",
				ArgsUsage: "<to> <amount>",
				Action:    utils.MigrateFlags(mintTON),
				Category:  "TON STAKING MANAGE COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.UnlockedAccountFlag,
					utils.PasswordFileFlag,
					utils.RootChainSenderFlag,
					utils.RootChainTONFlag,
					utils.RootChainGasPriceFlag,
				},
				Description: `
				geth manage-staking mintTON <to> <amount>

Mint TON to account

NOTE:
use --rootchain.ton flags to use already deployed token contracts
`,
			},
		},
	}

	stakingCmd = cli.Command{
		Name:     "staking",
		Usage:    "Stake TON",
		Category: "TON STAKING COMMANDS",
		Description: `

The staking command
`,
		Subcommands: []cli.Command{

			{
				Name:      "balances",
				Usage:     "Print balances of token and stake",
				ArgsUsage: "<address>",
				Action:    utils.MigrateFlags(getBalances),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.RootChainSenderFlag,
					utils.RootChainTONFlag,
					utils.RootChainWTONFlag,
					utils.RootChainDepositManagerFlag,
					utils.RootChainSeigManagerFlag,
				},
				Description: `
				geth staking balances <address>

Mint TON to account

NOTE:
use --rootchain.ton, --rootchain.wton, --rootchain.depositmanager, --rootchain.seigmanager flags to use already deployed token contracts
`,
			},
			{
				Name:      "swap-from-ton",
				Usage:     "Swap TON to WTON",
				ArgsUsage: "<tonAmount>",
				Action:    utils.MigrateFlags(swapFromTON),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.UnlockedAccountFlag,
					utils.PasswordFileFlag,
					utils.RootChainSenderFlag,
					utils.RootChainTONFlag,
					utils.RootChainWTONFlag,
					utils.RootChainDepositManagerFlag,
					utils.RootChainSeigManagerFlag,
					utils.RootChainGasPriceFlag,
				},
				Description: `
				geth staking swapFromTON <tonAmount>

Change TON to WTON

CAVEAT: <tonAmount> must be a float

NOTE:
use --rootchain.ton, --rootchain.wton, --rootchain.depositmanager, --rootchain.seigmanager flags to use already deployed token contracts
`,
			},
			{
				Name:      "swap-to-ton",
				Usage:     "Swap WTON to TON",
				ArgsUsage: "<wtonAmount>",
				Action:    utils.MigrateFlags(swapToTON),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.UnlockedAccountFlag,
					utils.PasswordFileFlag,
					utils.RootChainSenderFlag,
					utils.RootChainTONFlag,
					utils.RootChainWTONFlag,
					utils.RootChainDepositManagerFlag,
					utils.RootChainSeigManagerFlag,
					utils.RootChainGasPriceFlag,
				},
				Description: `
				geth staking swapToTON <wtonAmount>

Change WTON to TON

CAVEAT: <wtonAmount> must be a float

NOTE:
use --rootchain.ton, --rootchain.wton, --rootchain.depositmanager, --rootchain.seigmanager flags to use already deployed token contracts
`,
			},
			{
				Name:      "stake-ton",
				Usage:     "Stake TON",
				ArgsUsage: "<amount>",
				Action:    utils.MigrateFlags(stakeTON),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.UnlockedAccountFlag,
					utils.PasswordFileFlag,
					utils.RootChainSenderFlag,
					utils.RootChainTONFlag,
					utils.RootChainWTONFlag,
					utils.RootChainDepositManagerFlag,
					utils.RootChainSeigManagerFlag,
					utils.RootChainGasPriceFlag,
				},
				Description: `
				geth staking stakeTON <amount>

Stake WTON

CAVEAT: <amount> must be a float

NOTE:
use --rootchain.ton, --rootchain.wton, --rootchain.depositmanager, --rootchain.seigmanager flags to use already deployed token contracts
`,
			},
			{
				Name:      "stake-wton",
				Usage:     "Stake WTON",
				ArgsUsage: "<amount>",
				Action:    utils.MigrateFlags(stakeWTON),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.UnlockedAccountFlag,
					utils.PasswordFileFlag,
					utils.RootChainSenderFlag,
					utils.RootChainTONFlag,
					utils.RootChainWTONFlag,
					utils.RootChainDepositManagerFlag,
					utils.RootChainSeigManagerFlag,
					utils.RootChainGasPriceFlag,
				},
				Description: `
				geth staking stakeWTON <amount>

Stake WTON

CAVEAT: <amount> must be a float

NOTE:
use --rootchain.ton, --rootchain.wton, --rootchain.depositmanager, --rootchain.seigmanager flags to use already deployed token contracts
`,
			},
			{
				Name:      "restake",
				Usage:     "Restake pending withdrawal request",
				ArgsUsage: "<numRequests>",
				Action:    utils.MigrateFlags(restake),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.UnlockedAccountFlag,
					utils.PasswordFileFlag,
					utils.RootChainSenderFlag,
					utils.RootChainTONFlag,
					utils.RootChainWTONFlag,
					utils.RootChainDepositManagerFlag,
					utils.RootChainSeigManagerFlag,
					utils.RootChainGasPriceFlag,
				},
				Description: `
				geth staking restake <numRequests>

Restake pending request withdrawal

If <numRequests> is not provided or zero, restake all pending requests

NOTE:
use --rootchain.ton, --rootchain.wton, --rootchain.depositmanager, --rootchain.seigmanager flags to use already deployed token contracts
`,
			},
			{
				Name:      "request-withdrawal",
				Usage:     "Make a withdrawal request",
				ArgsUsage: "<amount>",
				Action:    utils.MigrateFlags(requestWithdrawal),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.UnlockedAccountFlag,
					utils.PasswordFileFlag,
					utils.RootChainSenderFlag,
					utils.RootChainDepositManagerFlag,
					utils.RootChainGasPriceFlag,
				},
				Description: `
				geth staking requestWithdrawal <amount>

Make an unstaking request. If amount is not given, make a request with all staked amount.

CAVEAT: <amount> must be a float

NOTE:
use --rootchain.depositmanager flags to use already deployed token contracts
`,
			},
			{
				Name:      "process-withdrawal",
				Usage:     "Process pending withdrawals",
				ArgsUsage: "<numRequests>",
				Action:    utils.MigrateFlags(processWithdrawal),
				Category:  "TON STAKING COMMANDS",
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.RootChainUrlFlag,
					utils.UnlockedAccountFlag,
					utils.PasswordFileFlag,
					utils.RootChainSenderFlag,
					utils.RootChainDepositManagerFlag,
					utils.RootChainGasPriceFlag,
				},
				Description: `
				geth staking processWithdrawal <numRequests>

Process unstaking requests

CAVEAT: <amount> must be a float

NOTE:
use --rootchain.depositmanager flags to use already deployed token contracts
`,
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
	PowerTON          common.Address `json:PowerTON`
}

func getManagerConfig(reader ethdb.Reader, ctx *cli.Context, override bool) *ManagerConfig {
	managers := &ManagerConfig{
		TON:               rawdb.ReadTON(reader),
		WTON:              rawdb.ReadWTON(reader),
		DepositManager:    rawdb.ReadDepositManager(reader),
		RootChainRegistry: rawdb.ReadRegistry(reader),
		SeigManager:       rawdb.ReadSeigManager(reader),
		PowerTON:          rawdb.ReadPowerTON(reader),
	}

	if override {
		tonAddr := common.HexToAddress(ctx.GlobalString(utils.RootChainTONFlag.Name))
		wtonAddr := common.HexToAddress(ctx.GlobalString(utils.RootChainWTONFlag.Name))
		depositManagerAddr := common.HexToAddress(ctx.GlobalString(utils.RootChainDepositManagerFlag.Name))
		registryAddr := common.HexToAddress(ctx.GlobalString(utils.RootChainRegistryFlag.Name))
		seigManagerAddr := common.HexToAddress(ctx.GlobalString(utils.RootChainSeigManagerFlag.Name))
		powertonAddr := common.HexToAddress(ctx.GlobalString(utils.RootChainPowerTONFlag.Name))

		if (managers.TON != common.Address{} && tonAddr != common.Address{}) && managers.TON != tonAddr {
			log.Warn("Override TON address", "previous", managers.TON, "new", tonAddr)
			managers.TON = tonAddr
		}

		if (managers.WTON != common.Address{} && wtonAddr != common.Address{}) && managers.WTON != wtonAddr {
			log.Warn("Override WTON address", "previous", managers.WTON, "new", wtonAddr)
			managers.WTON = wtonAddr
		}

		if (managers.DepositManager != common.Address{} && depositManagerAddr != common.Address{}) && managers.DepositManager != depositManagerAddr {
			log.Warn("Override DepositManager address", "previous", managers.DepositManager, "new", depositManagerAddr)
			managers.DepositManager = depositManagerAddr
		}

		if (managers.RootChainRegistry != common.Address{} && registryAddr != common.Address{}) && managers.RootChainRegistry != registryAddr {
			log.Warn("Override RootChainRegistry address", "previous", managers.RootChainRegistry, "new", registryAddr)
			managers.RootChainRegistry = registryAddr
		}

		if (managers.SeigManager != common.Address{} && seigManagerAddr != common.Address{}) && managers.SeigManager != seigManagerAddr {
			log.Warn("Override SeigManager address", "previous", managers.SeigManager, "new", seigManagerAddr)
			managers.SeigManager = seigManagerAddr
		}

		if (managers.PowerTON != common.Address{} && powertonAddr != common.Address{}) && managers.PowerTON != powertonAddr {
			log.Warn("Override WTON address", "previous", managers.PowerTON, "new", powertonAddr)
			managers.PowerTON = powertonAddr
		}
	}

	return managers
}

func getRootChainAddr(datadir string) (rootchainAddr common.Address) {
	path := filepath.Join(datadir, "geth", "genesis.json")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		utils.Fatalf("Failed to read genesis file: %v", err)
	}

	genesis := new(core.Genesis)
	if err = json.Unmarshal(data, genesis); err != nil {
		utils.Fatalf("Failed to parse genesis file: %v", err)
	}

	rootchainAddr = common.BytesToAddress(genesis.ExtraData)
	if (rootchainAddr == common.Address{}) {
		utils.Fatalf("RootChain is not set in genesis")
	}
	return
}

func parseFloatString(str string, decimals int) *big.Int {
	if decimals != 18 && decimals != 27 {
		utils.Fatalf("decimals should be 18 or 27, not %d", decimals)
	}

	i := strings.Index(str, ".")

	v := str
	n := decimals

	// split string with "."
	if i >= 0 {
		a := str[:i]
		b := str[i+1:]
		n = n - len(b)

		if n < 0 {
			utils.Fatalf("out of decimals precision: %d", decimals)
		}

		v = a + b
	}

	v = v + strings.Repeat("0", n)

	bi, ok := big.NewInt(0).SetString(v, 10)
	if !ok {
		utils.Fatalf(fmt.Sprintf("Failed to parse integer: %s", str))
	}

	return bi
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

func logManagers(managers *ManagerConfig) {
	log.Info("Using manager contracts", "TON", managers.TON, "WTON", managers.WTON, "DepositManager", managers.DepositManager, "RootChainRegistry", managers.RootChainRegistry, "SeigManager", managers.SeigManager, "PowerTON", managers.PowerTON)
}

func initOpts(ctx *cli.Context, stack *node.Node, cfg *pls.Config) (*bind.TransactOpts, *ethclient.Client) {
	unlockAccounts(ctx, stack)

	ks := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)

	sender := common.HexToAddress(ctx.GlobalString(utils.RootChainSenderFlag.Name))

	var opt *bind.TransactOpts
	if (sender != common.Address{}) {
		if !ks.HasAddress(sender) {
			utils.Fatalf("Unknown sender account: %s", sender)
		}

		log.Info("Root chain transaction sender found", "address", sender)

		senderAccount := accounts.Account{Address: sender}

		opt = bind.NewAccountTransactor(ks, senderAccount)
		if ctx.IsSet(utils.RootChainDeployGasPriceFlag.Name) {
			opt.GasPrice = utils.GlobalGasPrice(ctx, utils.RootChainDeployGasPriceFlag.Name)
		} else {
			opt.GasPrice = utils.GlobalGasPrice(ctx, utils.RootChainGasPriceFlag.Name)
		}
	}

	backend, err := ethclient.Dial(cfg.RootChainURL)
	if err != nil {
		utils.Fatalf("Failed to connect root chain: %v", err)
	}

	return opt, backend
}

func deployManagers(ctx *cli.Context) error {
	if len(ctx.Args()) != 2 {
		utils.Fatalf("Expected 2 parameters, not %d", len(ctx.Args()))
	}

	stack, cfg := makeConfigNode(ctx)
	opt, backend := initOpts(ctx, stack, &cfg.Pls)

	stakedb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	defer stakedb.Close()
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}

	decimals := params.WTONDecimals

	var (
		withdrawalDelay *big.Int
		seigPerBlock    *big.Int
	)

	withdrawalDelayStr := ctx.Args().Get(0)
	seigPerBlockStr := ctx.Args().Get(1)

	// parse int string
	withdrawalDelay, ok := big.NewInt(0).SetString(withdrawalDelayStr, 10)
	if !ok {
		return errors.New(fmt.Sprintf("Failed to parse integer: %s", withdrawalDelayStr))
	}

	// parse float string e.g., 12.4 to RAY value
	seigPerBlock = parseFloatString(seigPerBlockStr, decimals)

	_tonAddr := common.HexToAddress(ctx.String(utils.RootChainTONFlag.Name))
	_wtonAddr := common.HexToAddress(ctx.String(utils.RootChainWTONFlag.Name))

	tonAddr, wtonAddr, registryAddr, depositManagerAddr, seigManagerAddr, err := plasma.DeployManagers(opt, backend, withdrawalDelay, seigPerBlock, _tonAddr, _wtonAddr)

	if err != nil {
		return err
	}

	log.Info("Staking manager contract deployed", "TON", tonAddr, "WTON", wtonAddr, "RootChainRegistry", registryAddr, "DepositManager", depositManagerAddr, "SeigManager", seigManagerAddr)

	rawdb.WriteTON(stakedb, tonAddr)
	rawdb.WriteWTON(stakedb, wtonAddr)
	rawdb.WriteRegistry(stakedb, registryAddr)
	rawdb.WriteDepositManager(stakedb, depositManagerAddr)
	rawdb.WriteSeigManager(stakedb, seigManagerAddr)

	return nil
}

func deployPowerTON(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		utils.Fatalf("Expected 1 parameters, not %d", len(ctx.Args()))
	}

	stack, cfg := makeConfigNode(ctx)
	opt, backend := initOpts(ctx, stack, &cfg.Pls)

	stakedb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	defer stakedb.Close()
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}
	managers := getManagerConfig(stakedb, ctx, true)

	// parse duration
	roundDurationTime, err := time.ParseDuration(ctx.Args().Get(0))
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to parse duration: %v", err))
	}

	roundDuration := big.NewInt(int64(roundDurationTime.Seconds()))

	wtonAddr := managers.WTON
	seigManagerAddr := managers.SeigManager

	if ctx.IsSet(utils.RootChainWTONFlag.Name) {
		addr := common.HexToAddress(ctx.String(utils.RootChainWTONFlag.Name))
		if (wtonAddr != common.Address{}) && wtonAddr != addr {
			log.Warn("Override WTON address", "previous", wtonAddr, "current", addr)
		}
		wtonAddr = addr
	}

	if ctx.IsSet(utils.RootChainSeigManagerFlag.Name) {
		addr := common.HexToAddress(ctx.String(utils.RootChainSeigManagerFlag.Name))
		if (seigManagerAddr != common.Address{}) && seigManagerAddr != addr {
			log.Warn("Override SeigManager address", "previous", seigManagerAddr, "current", addr)
		}
		seigManagerAddr = addr
	}

	log.Info("Deploy PowerTON", "TON", wtonAddr, "SeigManager", seigManagerAddr, "roundDuration", roundDuration.String())

	powertonAddr, err := plasma.DeployPowerTON(opt, backend, wtonAddr, seigManagerAddr, roundDuration)

	if err != nil {
		return err
	}

	log.Info("PowerTON deployed", "PowerTON", powertonAddr, "WTON", wtonAddr, "SeigManager", seigManagerAddr)

	rawdb.WritePowerTON(stakedb, powertonAddr)

	return nil
}

func startPowerTON(ctx *cli.Context) error {
	stack, cfg := makeConfigNode(ctx)
	opt, backend := initOpts(ctx, stack, &cfg.Pls)

	stakedb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	defer stakedb.Close()
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}
	managers := getManagerConfig(stakedb, ctx, true)

	powertonAddr := managers.PowerTON

	if ctx.GlobalIsSet(utils.RootChainPowerTONFlag.Name) {
		addr := common.HexToAddress(ctx.GlobalString(utils.RootChainPowerTONFlag.Name))
		if (powertonAddr != common.Address{}) && powertonAddr != addr {
			log.Warn("Override WTON address", "previous", powertonAddr, "current", addr)
		}
		powertonAddr = addr
	}

	pton, err := powerton.NewPowerTON(powertonAddr, backend)

	if err != nil {
		utils.Fatalf("Failed to instantiate PowerTON contract: %v", err)
	}

	tx, err := pton.Start(opt)
	if err != nil {
		utils.Fatalf("Failed to start PowerTON: %v", err)
	}

	if err := plasma.WaitTx(backend, tx.Hash()); err != nil {
		utils.Fatalf("Failed to start PowerTON: %v", err)
	}

	log.Info("PowerTON started", "PowerTON", powertonAddr)

	return nil
}

func getManagers(ctx *cli.Context) error {
	configPath := ctx.Args().First()

	stack, _ := makeConfigNode(ctx)

	stakedb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	defer stakedb.Close()
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}
	managers := getManagerConfig(stakedb, ctx, false)

	b, err := json.MarshalIndent(managers, "", "  ")
	data := string(b)

	if err != nil {
		utils.Fatalf("Failed to parse manager data: %v", err)
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

	stack, _ := makeConfigNode(ctx)

	stakedb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	defer stakedb.Close()
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}

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
	if ctx.GlobalIsSet(utils.RootChainPowerTONFlag.Name) {
		powertonAddr := common.HexToAddress(ctx.GlobalString(utils.RootChainPowerTONFlag.Name))
		if (managers.PowerTON != common.Address{}) && managers.PowerTON != powertonAddr {
			log.Warn("Override PowerTON address", "previous", managers.PowerTON, "current", powertonAddr)
		}
		managers.PowerTON = powertonAddr
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
		{
			name:  "PowerTON",
			addr:  managers.PowerTON,
			read:  rawdb.ReadPowerTON,
			write: rawdb.WritePowerTON,
		},
	}

	for _, target := range targets {
		// short circuit if no target address is provided
		if (target.addr == common.Address{}) {
			continue
		}

		addr := target.read(stakedb)

		switch addr {
		case common.Address{}:
			log.Info("Set address", "name", target.name, "addr", target.addr)
			target.write(stakedb, target.addr)
		case target.addr:
			log.Info("Address is already set", "name", target.name, "addr", target.addr)
		default:
			log.Error("Known address is different", "name", target.name, "known", addr, "target", target.addr)
		}
	}

	return nil
}

func registerRootChain(ctx *cli.Context) error {
	stack, cfg := makeConfigNode(ctx)
	opt, backend := initOpts(ctx, stack, &cfg.Pls)

	stakedb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	defer stakedb.Close()
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}
	managers := getManagerConfig(stakedb, ctx, true)

	if (managers.RootChainRegistry == common.Address{}) || (managers.SeigManager == common.Address{}) {
		return errors.New("manager contract addresses is empty. please write contracts before register using `geth staking setManagers`")
	}

	rootchainAddr := getRootChainAddr(cfg.Node.DataDir)

	logManagers(managers)

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

func setCommissionRate(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 && len(ctx.Args()) != 2 {
		utils.Fatalf("Expected 1 or 2 parameters, not %d", len(ctx.Args()))
	}

	stack, cfg := makeConfigNode(ctx)
	opt, backend := initOpts(ctx, stack, &cfg.Pls)

	stakedb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	defer stakedb.Close()
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}
	managers := getManagerConfig(stakedb, ctx, true)

	if (managers.RootChainRegistry == common.Address{}) || (managers.SeigManager == common.Address{}) {
		return errors.New("manager contract addresses is empty. please write contracts before register using `geth staking setManagers`")
	}

	rootchainAddr := getRootChainAddr(cfg.Node.DataDir)

	decimals := params.WTONDecimals

	rate := parseFloatString(ctx.Args().Get(0), decimals)
	isCommissionRateNegative := false

	if rate.Cmp(big.NewInt(0)) < 0 {
		utils.Fatalf("Commission rate cannot be negative")
	}

	switch ctx.Args().Get(1) {
	case "":
	case "false":
	case "true":
		isCommissionRateNegative = true
	default:
		utils.Fatalf("Cannot parse isCommissionRateNegative: %s", ctx.Args().Get(1))
	}

	logManagers(managers)

	// load contract instances
	rootchainCtr, err := rootchain.NewRootChain(rootchainAddr, backend)
	if err != nil {
		utils.Fatalf("Failed to load RootChain contract: %v", err)
	}
	seigManager, err := seigmanager.NewSeigManager(managers.SeigManager, backend)
	if err != nil {
		utils.Fatalf("Failed to load SeigManager contract: %v", err)
	}

	operator, err := rootchainCtr.Operator(&bind.CallOpts{Pending: false})
	if err != nil {
		utils.Fatalf("Failed to read operator: %v", err)
	}

	if operator != opt.From {
		utils.Fatalf("Transaction sender is not the operator: %s", opt.From)
	}

	minRate, err := seigManager.MINVALIDCOMMISSION(&bind.CallOpts{Pending: false})
	if err != nil {
		utils.Fatalf("Failed to MIN_VALID_COMMISSION_RATE: %v", err)
	}
	maxRate, err := seigManager.MAXVALIDCOMMISSION(&bind.CallOpts{Pending: false})
	if err != nil {
		utils.Fatalf("Failed to MAX_VALID_COMMISSION_RATE: %v", err)
	}

	if rate.Cmp(big.NewInt(0)) != 0 && (minRate.Cmp(rate) > 0 || maxRate.Cmp(rate) < 0) {
		utils.Fatalf("Commission rate should be 0 or between %.2f and %.2f", params.ToRayFloat64(minRate), params.ToRayFloat64(maxRate))
	}

	log.Info("Set commission rate", "rootchain", rootchainAddr, "commissionRate", params.ToRayFloat64(rate), "isCommissionRateNegative", isCommissionRateNegative)

	// send transaction
	tx, err := seigManager.SetCommissionRate(opt, rootchainAddr, rate, isCommissionRateNegative)
	if err != nil {
		utils.Fatalf("Failed to send transaction: %v", err)
	}

	if err = plasma.WaitTx(backend, tx.Hash()); err != nil {
		utils.Fatalf("Failed to send transaction: %v", err)
	}

	return nil
}

// TODO: pending withdrawal amount
func getBalances(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		utils.Fatalf("Expected 1 parameters, not %d", len(ctx.Args()))
	}

	stack, cfg := makeConfigNode(ctx)
	_, backend := initOpts(ctx, stack, &cfg.Pls)

	stakedb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	defer stakedb.Close()
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}
	managers := getManagerConfig(stakedb, ctx, true)
	opt := &bind.CallOpts{Pending: false}
	rootchainAddr := getRootChainAddr(cfg.Node.DataDir)

	var (
		depositor common.Address
	)

	depositor = common.HexToAddress(ctx.Args().Get(0))

	if (managers.TON == common.Address{}) ||
		(managers.WTON == common.Address{}) ||
		(managers.DepositManager == common.Address{}) ||
		(managers.RootChainRegistry == common.Address{}) ||
		(managers.SeigManager == common.Address{}) {
		return errors.New("manager contract addresses is empty. please write contracts before register using `geth staking setManagers`")
	}

	logManagers(managers)

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

		commissionRate           *big.Int
		isCommissionRateNegative bool
	)

	// load contract instances
	if TON, err = ton.NewTON(managers.TON, backend); err != nil {
		utils.Fatalf("Failed to load TON contract: %v", err)
	}
	if WTON, err = wton.NewWTON(managers.WTON, backend); err != nil {
		utils.Fatalf("Failed to load WTON contract: %v", err)
	}
	if depositManager, err = depositmanager.NewDepositManager(managers.DepositManager, backend); err != nil {
		utils.Fatalf("Failed to load DepositManager contract: %v", err)
	}
	if seigManager, err = seigmanager.NewSeigManager(managers.SeigManager, backend); err != nil {
		utils.Fatalf("Failed to load SeigManager contract: %v", err)
	}

	totAddr, err := seigManager.Tot(opt)
	if err != nil {
		utils.Fatalf("Failed to load tot address: %v", err)
	}
	coinageAddr, err := seigManager.Coinages(opt, rootchainAddr)
	if err != nil {
		utils.Fatalf("Failed to load coinage address: %v", err)
	}

	if tot, err = seigmanager.NewERC20(totAddr, backend); err != nil {
		utils.Fatalf("Failed to load tot contract: %v", err)
	}
	if coinage, err = seigmanager.NewERC20(coinageAddr, backend); err != nil {
		utils.Fatalf("Failed to load tot contract: %v", err)
	}

	// read balances
	if tonBalance, err = TON.BalanceOf(opt, depositor); err != nil {
		utils.Fatalf("Failed to read TON balance: %v", err)
	}
	if wtonBalance, err = WTON.BalanceOf(opt, depositor); err != nil {
		utils.Fatalf("Failed to read WTON balance: %v", err)
	}
	if accStaked, err = depositManager.AccStaked(opt, rootchainAddr, depositor); err != nil {
		utils.Fatalf("Failed to read accumulated stake: %v", err)
	}
	if accUnstaked, err = depositManager.AccUnstaked(opt, rootchainAddr, depositor); err != nil {
		utils.Fatalf("Failed to read accumulated unstake: %v", err)
	}
	if numPendingRequests, err = depositManager.NumPendingRequests(opt, rootchainAddr, depositor); err != nil {
		utils.Fatalf("Failed to read num pending requests: %v", err)
	}
	if pendingUnstaked, err = depositManager.PendingUnstaked(opt, rootchainAddr, depositor); err != nil {
		utils.Fatalf("Failed to read pending withdrawal amount: %v", err)
	}

	if totalStake, err = tot.TotalSupply(opt); err != nil {
		log.Warn("Failed to read total stake", "err", err)
		totalStake = big.NewInt(0)
	}
	if totalStakeRootChain, err = coinage.TotalSupply(opt); err != nil {
		log.Warn("Failed to read total stake of root chain", "err", err)
		totalStakeRootChain = big.NewInt(0)
	}

	if uncomittedStakeOf, err = seigManager.UncomittedStakeOf(opt, rootchainAddr, depositor); err != nil {
		log.Warn("Failed to read uncomitted stake", "err", err)
		uncomittedStakeOf = big.NewInt(0)
	}
	if stakeOf, err = seigManager.StakeOf(opt, rootchainAddr, depositor); err != nil {
		log.Warn("Failed to read stake", "err", err)
		stakeOf = big.NewInt(0)
	}
	if commissionRate, err = seigManager.CommissionRates(opt, rootchainAddr); err != nil {
		log.Warn("Failed to read commission rate stake", "err", err)
		commissionRate = big.NewInt(0)
	}
	if isCommissionRateNegative, err = seigManager.IsCommissionRateNegative(opt, rootchainAddr); err != nil {
		log.Warn("Failed to read commission rate stake", "err", err)
	}
	if isCommissionRateNegative {
		commissionRate = new(big.Int).Neg(commissionRate)
	}

	deposit = new(big.Int).Sub(accStaked, accUnstaked)

	// print balances
	log.Info("TON Balance", "amount", bigIntToString(tonBalance, params.TONDecimals)+" TON", "depositor", depositor)
	log.Info("WTON Balance", "amount", bigIntToString(wtonBalance, params.WTONDecimals)+" WTON", "depositor", depositor)
	log.Info("Deposit", "amount", bigIntToString(deposit, params.WTONDecimals)+" WTON", "rootchain", rootchainAddr, "depositor", depositor)

	log.Info("Pending withdrawal requests", "num", numPendingRequests)
	log.Info("Pending withdrawal WTON", "amount", bigIntToString(pendingUnstaked, params.WTONDecimals)+" WTON", "rootchain", rootchainAddr, "depositor", depositor)

	log.Info("Total Stake", "amount", bigIntToString(totalStake, params.WTONDecimals)+" WTON")
	log.Info("Total Stake of Root Chain", "amount", bigIntToString(totalStakeRootChain, params.WTONDecimals)+" WTON", "rootchain", rootchainAddr)

	log.Info("Uncomitted Stake", "amount", bigIntToString(uncomittedStakeOf, params.WTONDecimals)+" WTON", "rootchain", rootchainAddr, "depositor", depositor)
	log.Info("Comitted Stake", "amount", bigIntToString(stakeOf, params.WTONDecimals)+" WTON", "rootchain", rootchainAddr, "depositor", depositor)

	log.Info("Commission Rate", "rate", params.ToRayFloat64(commissionRate))

	return nil
}

func mintTON(ctx *cli.Context) error {
	if len(ctx.Args()) != 2 {
		utils.Fatalf("Expected 2 parameters, not %d", len(ctx.Args()))
	}

	stack, cfg := makeConfigNode(ctx)
	opt, backend := initOpts(ctx, stack, &cfg.Pls)

	stakedb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	defer stakedb.Close()
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}
	managers := getManagerConfig(stakedb, ctx, true)

	decimals := params.TONDecimals

	var (
		to     common.Address
		amount *big.Int
	)

	to = common.HexToAddress(ctx.Args().Get(0))
	amount = parseFloatString(ctx.Args().Get(1), decimals)

	if (managers.TON == common.Address{}) {
		return errors.New("manager contract addresses is empty. please write contracts before register using `geth staking setManagers`")
	}

	logManagers(managers)

	TON, err := ton.NewTON(managers.TON, backend)
	if err != nil {
		utils.Fatalf("Failed to instantiate TON: %v", err)
	}

	var tx *types.Transaction
	if tx, err = TON.Mint(opt, to, amount); err != nil {
		utils.Fatalf("Failed to mint TON: %v", err)
	}
	log.Info("Minting TON", "to", to, "amount", bigIntToString(amount, decimals)+" TON", "tx", tx.Hash())

	if err = plasma.WaitTx(backend, tx.Hash()); err != nil {
		utils.Fatalf("Failed to wait transaction: %v", err)
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

	stack, cfg := makeConfigNode(ctx)
	opt, backend := initOpts(ctx, stack, &cfg.Pls)

	stakedb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	defer stakedb.Close()
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}
	managers := getManagerConfig(stakedb, ctx, true)

	decimals := params.TONDecimals

	var (
		amount *big.Int
	)

	amount = parseFloatString(ctx.Args().Get(0), decimals)

	if (managers.WTON == common.Address{}) || (managers.TON == common.Address{}) {
		return errors.New("manager contract addresses is empty. please write contracts before register using `geth staking setManagers`")
	}

	logManagers(managers)

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

	stack, cfg := makeConfigNode(ctx)
	opt, backend := initOpts(ctx, stack, &cfg.Pls)

	stakedb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	defer stakedb.Close()
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}
	managers := getManagerConfig(stakedb, ctx, true)

	decimals := params.WTONDecimals

	var (
		amount *big.Int
	)

	amount = parseFloatString(ctx.Args().Get(0), decimals)

	if (managers.WTON == common.Address{}) || (managers.TON == common.Address{}) {
		return errors.New("manager contract addresses is empty. please write contracts before register using `geth staking setManagers`")
	}

	logManagers(managers)

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

func stakeTON(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		utils.Fatalf("Expected 1 parameters, not %d", len(ctx.Args()))
	}

	stack, cfg := makeConfigNode(ctx)
	opt, backend := initOpts(ctx, stack, &cfg.Pls)

	stakedb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	defer stakedb.Close()
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}
	managers := getManagerConfig(stakedb, ctx, true)
	rootchainAddr := getRootChainAddr(cfg.Node.DataDir)

	decimals := params.TONDecimals

	var (
		amount *big.Int

		tx *types.Transaction
	)

	amount = parseFloatString(ctx.Args().Get(0), decimals)

	if (managers.TON == common.Address{}) ||
		(managers.WTON == common.Address{}) ||
		(managers.DepositManager == common.Address{}) ||
		(managers.RootChainRegistry == common.Address{}) ||
		(managers.SeigManager == common.Address{}) {
		return errors.New("manager contract addresses is empty. please write contracts before register using `geth staking setManagers`")
	}

	logManagers(managers)

	var (
		TON *ton.TON
	)

	// load contract instances
	if TON, err = ton.NewTON(managers.TON, backend); err != nil {
		utils.Fatalf("Failed to load TON contract: %v", err)
	}

	// check TON balance
	tonBalance, err := TON.BalanceOf(&bind.CallOpts{Pending: false}, opt.From)
	if err != nil {
		utils.Fatalf("Failed to read TON balance: %v", err)
	}

	if tonBalance.Cmp(amount) < 0 {
		utils.Fatalf("Insufficient TON Balance (%s)", bigIntToString(tonBalance, decimals))
	}

	// send transaction
	pad := make([]byte, 12)
	approveData := append(append(pad, managers.DepositManager.Bytes()...), append(pad, rootchainAddr.Bytes()...)...)
	tx, err = TON.ApproveAndCall(opt, managers.WTON, amount, approveData)

	if err != nil {
		utils.Fatalf("Failed to send transaction: %v", err)
	}

	if err = plasma.WaitTx(backend, tx.Hash()); err != nil {
		return err
	}

	log.Info("Deposit TON to RootChain", "rootchain", rootchainAddr, "amount", bigIntToString(amount, decimals)+" TON", "tx", tx.Hash())

	return nil
}

func stakeWTON(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		utils.Fatalf("Expected 1 parameters, not %d", len(ctx.Args()))
	}

	stack, cfg := makeConfigNode(ctx)
	opt, backend := initOpts(ctx, stack, &cfg.Pls)

	stakedb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	defer stakedb.Close()
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}
	managers := getManagerConfig(stakedb, ctx, true)
	rootchainAddr := getRootChainAddr(cfg.Node.DataDir)

	decimals := params.WTONDecimals

	var (
		amount *big.Int

		tx *types.Transaction
	)

	amount = parseFloatString(ctx.Args().Get(0), decimals)

	if (managers.TON == common.Address{}) ||
		(managers.WTON == common.Address{}) ||
		(managers.DepositManager == common.Address{}) ||
		(managers.RootChainRegistry == common.Address{}) ||
		(managers.SeigManager == common.Address{}) {
		return errors.New("manager contract addresses is empty. please write contracts before register using `geth staking setManagers`")
	}

	logManagers(managers)

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
		utils.Fatalf("Insufficient WTON Balance (%s)", bigIntToString(wtonBalance, decimals))
	}

	// send transaction(s)
	approveToken("WTON", WTON, backend, opt, managers.DepositManager, amount, decimals)

	if tx, err = depositManager.Deposit(opt, rootchainAddr, amount); err != nil {
		return err
	}

	if err = plasma.WaitTx(backend, tx.Hash()); err != nil {
		return err
	}

	log.Info("Deposit WTON to RootChain", "rootchain", rootchainAddr, "amount", bigIntToString(amount, decimals)+" WTON", "tx", tx.Hash())

	return nil
}

func restake(ctx *cli.Context) error {
	if len(ctx.Args()) > 1 {
		utils.Fatalf("Expected 1 or 0 parameters, not %d", len(ctx.Args()))
	}

	stack, cfg := makeConfigNode(ctx)
	opt, backend := initOpts(ctx, stack, &cfg.Pls)

	stakedb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	defer stakedb.Close()
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}
	managers := getManagerConfig(stakedb, ctx, true)
	rootchainAddr := getRootChainAddr(cfg.Node.DataDir)

	var (
		n                  int
		numPendingRequests *big.Int
		numRequests        *big.Int

		tx *types.Transaction
	)

	if len(ctx.Args()) == 1 {
		if n, err = strconv.Atoi(ctx.Args().Get(0)); err != nil {
			return err
		}
	}

	if (managers.TON == common.Address{}) ||
		(managers.WTON == common.Address{}) ||
		(managers.DepositManager == common.Address{}) ||
		(managers.RootChainRegistry == common.Address{}) ||
		(managers.SeigManager == common.Address{}) {
		return errors.New("manager contract addresses is empty. please write contracts before register using `geth staking setManagers`")
	}

	logManagers(managers)

	var (
		depositManager *depositmanager.DepositManager
	)

	if depositManager, err = depositmanager.NewDepositManager(managers.DepositManager, backend); err != nil {
		utils.Fatalf("Failed to load DepositManager contract: %v", err)
	}

	if numPendingRequests, err = depositManager.NumPendingRequests(&bind.CallOpts{Pending: false}, rootchainAddr, opt.From); err != nil {
		utils.Fatalf("Failed to read the number of pending requests: %v", err)
	}

	if n == 0 {
		numRequests = new(big.Int).Set(numPendingRequests)
	} else {
		numRequests = big.NewInt(int64(n))
	}

	if numRequests.Cmp(numPendingRequests) > 0 {
		utils.Fatalf("the number of request to restake exceeds the number of pending request (pendingRequest=%d, numRequest=%d)", numPendingRequests.Uint64(), numRequests.Uint64())
	}

	if tx, err = depositManager.RedepositMulti(opt, rootchainAddr, numRequests); err != nil {
		return err
	}

	if err = plasma.WaitTx(backend, tx.Hash()); err != nil {
		return err
	}

	log.Info("Redeposit pending requests", "rootchain", rootchainAddr, "numRequests", numRequests, "tx", tx.Hash())

	return nil
}

func requestWithdrawal(ctx *cli.Context) error {
	if len(ctx.Args()) > 1 {
		utils.Fatalf("Expected 1 or 0 parameters, not %d", len(ctx.Args()))
	}

	stack, cfg := makeConfigNode(ctx)
	opt, backend := initOpts(ctx, stack, &cfg.Pls)

	stakedb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	defer stakedb.Close()
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}
	managers := getManagerConfig(stakedb, ctx, true)
	rootchainAddr := getRootChainAddr(cfg.Node.DataDir)

	decimals := params.WTONDecimals

	var (
		amount *big.Int

		tx *types.Transaction
	)

	if len(ctx.Args()) == 1 {
		amount = parseFloatString(ctx.Args().Get(0), decimals)
	}

	if (managers.TON == common.Address{}) ||
		(managers.WTON == common.Address{}) ||
		(managers.DepositManager == common.Address{}) ||
		(managers.RootChainRegistry == common.Address{}) ||
		(managers.SeigManager == common.Address{}) {
		return errors.New("manager contract addresses is empty. please write contracts before register using `geth staking setManagers`")
	}

	logManagers(managers)

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
	staked, err := seigManager.StakeOf(&bind.CallOpts{Pending: false}, rootchainAddr, opt.From)
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

	stack, cfg := makeConfigNode(ctx)
	opt, backend := initOpts(ctx, stack, &cfg.Pls)

	stakedb, err := stack.OpenDatabase("stakingdata", 0, 0, "")
	defer stakedb.Close()
	if err != nil {
		utils.Fatalf("Failed to open database: %v", err)
	}
	managers := getManagerConfig(stakedb, ctx, true)
	rootchainAddr := getRootChainAddr(cfg.Node.DataDir)

	var (
		n int

		tx *types.Transaction
	)

	if len(ctx.Args()) == 1 {
		if n, err = strconv.Atoi(ctx.Args().Get(0)); err != nil {
			return err
		}
	}

	if (managers.TON == common.Address{}) ||
		(managers.WTON == common.Address{}) ||
		(managers.DepositManager == common.Address{}) ||
		(managers.RootChainRegistry == common.Address{}) ||
		(managers.SeigManager == common.Address{}) {
		return errors.New("manager contract addresses is empty. please write contracts before register using `geth staking setManagers`")
	}

	logManagers(managers)

	var (
		numPendingRequests *big.Int

		depositManager *depositmanager.DepositManager
	)

	// load contract instances
	if depositManager, err = depositmanager.NewDepositManager(managers.DepositManager, backend); err != nil {
		utils.Fatalf("Failed to load DepositManager contract: %v", err)
	}

	if numPendingRequests, err = depositManager.NumPendingRequests(&bind.CallOpts{Pending: false}, rootchainAddr, opt.From); err != nil {
		utils.Fatalf("Failed to read num pending requests: %v", err)
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
