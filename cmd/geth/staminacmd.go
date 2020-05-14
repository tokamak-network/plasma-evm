package main

import (
	"context"
	"fmt"
	"math/big"
	"path/filepath"

	"github.com/Onther-Tech/plasma-evm/accounts"
	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind"
	"github.com/Onther-Tech/plasma-evm/accounts/keystore"
	"github.com/Onther-Tech/plasma-evm/common"
	stamina "github.com/Onther-Tech/plasma-evm/contracts/stamina/contract"
	"github.com/Onther-Tech/plasma-evm/log"
	"github.com/Onther-Tech/plasma-evm/node"
	"github.com/Onther-Tech/plasma-evm/params"
	"github.com/Onther-Tech/plasma-evm/plsclient"
	"gopkg.in/urfave/cli.v1"

	"github.com/Onther-Tech/plasma-evm/cmd/utils"
)

var (
	staminaCmd = cli.Command{
		Name:     "stamina",
		Usage:    "Manage Stamina Contracts",
		Category: "STAMINA COMMANDS",
		Description: `
The stamina command send transaction to interact with stamina contract.

NOTE: uint256 type parameters must be a float
`,
		Subcommands: []cli.Command{
			{
				Name:      "get-delegatee",
				Usage:     "Get delegatee of account",
				ArgsUsage: "<address>",
				Action:    utils.MigrateFlags(getDelegatee),
				Flags: []cli.Flag{
					utils.ChildChainUrlFlag,
				},
			},
			{
				Name:      "get-stamina",
				Usage:     "Get stamina of account",
				ArgsUsage: "<address>",
				Action:    utils.MigrateFlags(getStamina),
				Flags: []cli.Flag{
					utils.ChildChainUrlFlag,
				},
			},
			{
				Name:      "get-totaldeposit",
				Usage:     "Get total deposit of account",
				ArgsUsage: "<address>",
				Action:    utils.MigrateFlags(getTotalDeposit),
				Flags: []cli.Flag{
					utils.ChildChainUrlFlag,
				},
			},
			{
				Name:      "get-deposit",
				Usage:     "Get deposit of account from the depositor",
				ArgsUsage: "<depositor> <delegatee>",
				Action:    utils.MigrateFlags(getDeposit),
				Flags: []cli.Flag{
					utils.ChildChainUrlFlag,
				},
			},
			{
				Name:      "get-lastrecoveryblock",
				Usage:     "Get last recovery block of the delegatee",
				ArgsUsage: "<delegatee>",
				Action:    utils.MigrateFlags(getLastRecoveryBlock),
				Flags: []cli.Flag{
					utils.ChildChainUrlFlag,
				},
			},
			{
				Name:      "get-withdrawal",
				Usage:     "Get withdrawal requests",
				ArgsUsage: "<depositor>",
				Action:    utils.MigrateFlags(getWithdrawal),
				Flags: []cli.Flag{
					utils.ChildChainUrlFlag,
				},
			},
			{
				Name:      "set-delegator",
				Usage:     "Set delegator",
				ArgsUsage: "<delegator>",
				Action:    utils.MigrateFlags(setDelegator),
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.UnlockedAccountFlag,
					utils.PasswordFileFlag,
					utils.ChildChainUrlFlag,
					utils.ChildChainSenderFlag,
					utils.ChildChainGasPriceFlag,
				},
			},
			{
				Name:      "deposit",
				Usage:     "Deposit PETH to gain stamina",
				ArgsUsage: "<delegatee> <value>",
				Action:    utils.MigrateFlags(deposit),
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.UnlockedAccountFlag,
					utils.PasswordFileFlag,
					utils.ChildChainUrlFlag,
					utils.ChildChainSenderFlag,
					utils.ChildChainGasPriceFlag,
				},
			},
			{
				Name:      "request-withdrawal",
				Usage:     "Request withdraw",
				ArgsUsage: "<delegatee> <value>",
				Action:    utils.MigrateFlags(requestStaminaWithdrawal),
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.UnlockedAccountFlag,
					utils.PasswordFileFlag,
					utils.ChildChainUrlFlag,
					utils.ChildChainSenderFlag,
					utils.ChildChainGasPriceFlag,
				},
			},
			{
				Name:      "withdraw",
				Usage:     "Process withdraw",
				ArgsUsage: "",
				Action:    utils.MigrateFlags(withdraw),
				Flags: []cli.Flag{
					utils.DataDirFlag,
					utils.UnlockedAccountFlag,
					utils.PasswordFileFlag,
					utils.ChildChainUrlFlag,
					utils.ChildChainSenderFlag,
					utils.ChildChainGasPriceFlag,
				},
			},
		},
	}
)

func initPlsOpts(ctx *cli.Context) (*bind.TransactOpts, *plsclient.Client) {
	var opt *bind.TransactOpts

	if ctx.GlobalIsSet(utils.ChildChainSenderFlag.Name) {
		stack, _ := makeConfigNode(ctx)
		unlockAccounts(ctx, stack)

		ks := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)

		sender := common.HexToAddress(ctx.GlobalString(utils.ChildChainSenderFlag.Name))

		if (sender != common.Address{}) {
			if !ks.HasAddress(sender) {
				utils.Fatalf("Unknown sender account: %s", sender)
			}

			log.Info("Child chain transaction sender found", "address", sender)

			senderAccount := accounts.Account{Address: sender}

			opt = bind.NewAccountTransactor(ks, senderAccount)
			opt.GasPrice = utils.GlobalGasPrice(ctx, utils.ChildChainGasPriceFlag.Name)
		}
	}

	endpoint := ctx.GlobalString(utils.ChildChainUrlFlag.Name)
	if endpoint == "" {
		path := node.DefaultDataDir()
		if ctx.GlobalIsSet(utils.DataDirFlag.Name) {
			path = ctx.GlobalString(utils.DataDirFlag.Name)
		}
		if path != "" {
			if ctx.GlobalBool(utils.TestnetFlag.Name) {
				path = filepath.Join(path, "testnet")
			} else if ctx.GlobalBool(utils.RinkebyFlag.Name) {
				path = filepath.Join(path, "rinkeby")
			}
		}
		endpoint = fmt.Sprintf("%s/geth.ipc", path)
	}

	client, err := dialRPC(endpoint)
	if err != nil {
		utils.Fatalf("Failed to connect endpoint: %v", err)
	}

	return opt, plsclient.NewClient(client)
}

func loadStaminaContract(backend *plsclient.Client) *stamina.Stamina {
	staminaCtr, err := stamina.NewStamina(params.StaminaAddress, backend)
	if err != nil {
		utils.Fatalf("Failed to load stamina contract: %v", err)
	}
	log.Info("Load stamina contract", "addr", params.StaminaAddress)
	return staminaCtr
}

func getDelegatee(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		utils.Fatalf("Expected 1 parameters...")
	}

	delegator := common.HexToAddress(ctx.Args().Get(0))

	_, backend := initPlsOpts(ctx)
	staminaCtr := loadStaminaContract(backend)

	delegatee, err := staminaCtr.GetDelegatee(&bind.CallOpts{Pending: false}, delegator)
	if err != nil {
		utils.Fatalf("Failed to get delegatee: %v", err)
	}

	log.Info("Stamina.getDelegatee", "delegator", delegator, "delegatee", delegatee)

	return nil
}

func getStamina(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		utils.Fatalf("Expected 1 parameters...")
	}

	addr := common.HexToAddress(ctx.Args().Get(0))

	_, backend := initPlsOpts(ctx)
	staminaCtr := loadStaminaContract(backend)

	staminaAmount, err := staminaCtr.GetStamina(&bind.CallOpts{Pending: false}, addr)
	if err != nil {
		utils.Fatalf("Failed to get delegatee: %v", err)
	}

	log.Info("Stamina.getStamina", "addr", addr, "stamina", params.ToEtherFloat64(staminaAmount))

	return nil
}

func getTotalDeposit(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		utils.Fatalf("Expected 1 parameters...")
	}

	delegatee := common.HexToAddress(ctx.Args().Get(0))

	_, backend := initPlsOpts(ctx)
	staminaCtr := loadStaminaContract(backend)

	deposits, err := staminaCtr.GetTotalDeposit(&bind.CallOpts{Pending: false}, delegatee)
	if err != nil {
		utils.Fatalf("Failed to get delegatee: %v", err)
	}

	log.Info("Stamina.getTotalDeposit", "delegatee", delegatee, "deposits", params.ToEtherFloat64(deposits))

	return nil
}

func getDeposit(ctx *cli.Context) error {
	if len(ctx.Args()) != 2 {
		utils.Fatalf("Expected 2 parameters...")
	}

	depositor := common.HexToAddress(ctx.Args().Get(0))
	delegatee := common.HexToAddress(ctx.Args().Get(1))

	_, backend := initPlsOpts(ctx)
	staminaCtr := loadStaminaContract(backend)

	deposits, err := staminaCtr.GetDeposit(&bind.CallOpts{Pending: false}, depositor, delegatee)
	if err != nil {
		utils.Fatalf("Failed to get delegatee: %v", err)
	}

	log.Info("Stamina.getStamina", "depositor", depositor, "delegatee", delegatee, "stamina", params.ToEtherFloat64(deposits))

	return nil
}

func getLastRecoveryBlock(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		utils.Fatalf("Expected 1 parameters...")
	}

	delegatee := common.HexToAddress(ctx.Args().Get(0))

	_, backend := initPlsOpts(ctx)
	staminaCtr := loadStaminaContract(backend)

	block, err := staminaCtr.GetLastRecoveryBlock(&bind.CallOpts{Pending: false}, delegatee)
	if err != nil {
		utils.Fatalf("Failed to get delegatee: %v", err)
	}

	log.Info("Stamina.getLastRecoveryBlock", "delegatee", delegatee, "block", block)

	return nil
}

func getWithdrawal(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		utils.Fatalf("Expected 1 parameters...")
	}

	depositor := common.HexToAddress(ctx.Args().Get(0))

	_, backend := initPlsOpts(ctx)
	staminaCtr := loadStaminaContract(backend)

	header, err := backend.HeaderByNumber(context.Background(), nil)
	if err != nil {
		utils.Fatalf("Failed to get latest block number: %v", err)
	}
	blockNumber := header.Number

	index, err := staminaCtr.GetLastProcessedWithdrawalIndex(&bind.CallOpts{Pending: false}, depositor)
	if err != nil {
		utils.Fatalf("Failed to get last processed withdrawal index: %v", err)
	}

	num, err := staminaCtr.GetNumWithdrawals(&bind.CallOpts{Pending: false}, depositor)
	if err != nil {
		utils.Fatalf("Failed to get the number of withdrawal requests: %v", err)
	}

	delay, err := staminaCtr.WITHDRAWALDELAY(&bind.CallOpts{Pending: false})
	if err != nil {
		utils.Fatalf("Failed to get withdrawal delay: %v", err)
	}

	i := index.Uint64()
	n := num.Uint64()

	var numWithdrawable int64 = 0

	totalAmount := big.NewInt(0)
	withdrawableAmount := big.NewInt(0)

	for ; i < n; i++ {
		w, err := staminaCtr.GetWithdrawal(&bind.CallOpts{Pending: false}, depositor, big.NewInt(int64(i)))
		if err != nil {
			utils.Fatalf("Failed to get withdrawal request: %v", err)
		}

		totalAmount = new(big.Int).Add(totalAmount, w.Amount)
		if blockNumber.Cmp(new(big.Int).Add(w.RequestBlockNumber, delay)) >= 0 {
			numWithdrawable++
			withdrawableAmount = new(big.Int).Add(withdrawableAmount, w.Amount)
		}
	}

	log.Info("Total Request", "n", num, "amount", params.ToEtherFloat64(totalAmount))
	log.Info("Pending Request", "n", new(big.Int).Sub(num, big.NewInt(numWithdrawable)), "amount", params.ToEtherFloat64(new(big.Int).Sub(totalAmount, withdrawableAmount)))
	log.Info("Withdrawable Request", "n", numWithdrawable, "amount", params.ToEtherFloat64(withdrawableAmount))

	return nil
}

func setDelegator(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		utils.Fatalf("Expected 1 parameters...")
	}

	delegator := common.HexToAddress(ctx.Args().Get(0))

	opt, backend := initPlsOpts(ctx)
	staminaCtr := loadStaminaContract(backend)

	tx, err := staminaCtr.SetDelegator(opt, delegator)
	if err != nil {
		utils.Fatalf("Failed to set delegator: %v", err)
	}

	receipt, err := bind.WaitMined(context.Background(), backend, tx)
	if err != nil {
		utils.Fatalf("Transaction not mined: %v", err)
	}

	if receipt.Status != 1 {
		utils.Fatalf("Transaction reverted: %s", tx.Hash())
	}

	log.Info("SetDelegator", "hash", tx.Hash(), "delegatee", opt.From, "deleator", delegator)

	return nil
}

func deposit(ctx *cli.Context) error {
	if len(ctx.Args()) != 2 {
		utils.Fatalf("Expected 2 parameters...")
	}

	delegatee := common.HexToAddress(ctx.Args().Get(0))
	amount := parseFloatString(ctx.Args().Get(1), 18)

	opt, backend := initPlsOpts(ctx)
	staminaCtr := loadStaminaContract(backend)

	opt.Value = new(big.Int).Set(amount)

	tx, err := staminaCtr.Deposit(opt, delegatee)
	if err != nil {
		utils.Fatalf("Failed to deposit: %v", err)
	}

	receipt, err := bind.WaitMined(context.Background(), backend, tx)
	if err != nil {
		utils.Fatalf("Transaction not mined: %v", err)
	}

	if receipt.Status != 1 {
		utils.Fatalf("Transaction reverted: %s", tx.Hash())
	}

	log.Info("Deposit", "hash", tx.Hash(), "depositor", opt.From, "deleatee", delegatee)

	return nil
}

func requestStaminaWithdrawal(ctx *cli.Context) error {
	if len(ctx.Args()) != 2 {
		utils.Fatalf("Expected 2 parameters...")
	}

	opt, backend := initPlsOpts(ctx)
	staminaCtr := loadStaminaContract(backend)

	delegatee := common.HexToAddress(ctx.Args().Get(0))
	amount := parseFloatString(ctx.Args().Get(1), 18)

	tx, err := staminaCtr.RequestWithdrawal(opt, delegatee, amount)
	if err != nil {
		utils.Fatalf("Failed to withdraw: %v", err)
	}

	receipt, err := bind.WaitMined(context.Background(), backend, tx)
	if err != nil {
		utils.Fatalf("Transaction not mined: %v", err)
	}

	if receipt.Status != 1 {
		utils.Fatalf("Transaction reverted: %s", tx.Hash())
	}

	log.Info("RequestWithdrawal", "hash", tx.Hash(), "depositor", opt.From, "deleatee", delegatee, "amount", params.ToEtherFloat64(amount))

	return nil
}

func withdraw(ctx *cli.Context) error {
	if len(ctx.Args()) != 0 {
		utils.Fatalf("Expected 0 parameters...")
	}

	opt, backend := initPlsOpts(ctx)
	staminaCtr := loadStaminaContract(backend)

	tx, err := staminaCtr.Withdraw(opt)
	if err != nil {
		utils.Fatalf("Failed to withdraw: %v", err)
	}

	receipt, err := bind.WaitMined(context.Background(), backend, tx)
	if err != nil {
		utils.Fatalf("Transaction not mined: %v", err)
	}

	if receipt.Status != 1 {
		utils.Fatalf("Transaction reverted: %s", tx.Hash())
	}

	// TODO: parse receipt logs and log delegatee and amount
	log.Info("Withdraw", "hash", tx.Hash(), "depositor", opt.From)

	return nil
}
