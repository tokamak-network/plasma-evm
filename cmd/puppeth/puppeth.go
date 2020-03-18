// Copyright 2017 The go-ethereum Authors
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

// puppeth is a command to assemble and maintain private networks.
package main

import (
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/Onther-Tech/plasma-evm/log"
	"gopkg.in/urfave/cli.v1"
)

// main is just a boring entry point to set up the CLI app.
func main() {
	app := cli.NewApp()
	app.Name = "puppeth"
	app.Usage = "assemble and maintain private Ethereum networks"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "network",
			Usage: "name of the network to administer (no spaces or hyphens, please)",
		},
		cli.IntFlag{
			Name:  "loglevel",
			Value: 3,
			Usage: "log level to emit to the screen",
		},
		cli.StringFlag{
			Name:  "images.ethstats",
			Usage: "name of ethstats docker image",
			Value: "puppeth/ethstats:latest",
		},
		cli.StringFlag{
			Name:  "images.bootnode",
			Usage: "name of bootnode docker image",
			Value: "onthertech/plasma-evm:alltools-latest",
		},
		cli.StringFlag{
			Name:  "images.node",
			Usage: "name of node docker image",
			Value: "onthertech/plasma-evm:latest",
		},
		cli.StringFlag{
			Name:  "images.explorer",
			Usage: "name of explorer docker image",
			Value: "onthertech/blockscout:latest",
		},
		cli.StringFlag{
			Name:  "images.nginx",
			Usage: "name of nginx docker image",
			Value: "jwilder/nginx-proxy",
		},
		cli.StringFlag{
			Name:  "images.wallet",
			Usage: "name of wallet docker image",
			Value: "puppeth/wallet:latest",
		},
		cli.StringFlag{
			Name:  "images.faucet",
			Usage: "name of faucet docker image",
			Value: "onthertech/plasma-evm:alltools-latest",
		},
		cli.StringFlag{
			Name:  "images.dashboard",
			Usage: "name of dashboard docker image",
			Value: "mhart/alpine-node:latest",
		},
		cli.StringFlag{
			Name:  "bootnodes",
			Usage: "Comma separated enode URLs for P2P discovery bootstrap for Sealer or usernode",
			Value: "",
		},
	}
	app.Before = func(c *cli.Context) error {
		// Set up the logger to print everything and the random generator
		log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(c.Int("loglevel")), log.StreamHandler(os.Stdout, log.TerminalFormat(true))))
		rand.Seed(time.Now().UnixNano())

		return nil
	}
	app.Action = runWizard
	app.Run(os.Args)
}

// runWizard start the wizard and relinquish control to it.
func runWizard(c *cli.Context) error {
	network := c.String("network")
	images := map[string]string{
		"ethstats":  c.String("images.ethstats"),
		"bootnode":  c.String("images.bootnode"),
		"node":      c.String("images.node"),
		"explorer":  c.String("images.explorer"),
		"nginx":     c.String("images.nginx"),
		"wallet":    c.String("images.wallet"),
		"faucet":    c.String("images.faucet"),
		"dashboard": c.String("images.dashboard"),
	}
	bootnodes := c.String("bootnodes")

	urls := []string{}
	urls = strings.Split(bootnodes, ",")

	if strings.Contains(network, " ") || strings.Contains(network, "-") || strings.ToLower(network) != network {
		log.Crit("No spaces, hyphens or capital letters allowed in network name")
	}

	// Checking enode address
	for _, url := range urls {
		if url != "" {
			// prefix check
			if !strings.HasPrefix(url, "enode://") {
				log.Crit("Bootstrap URL invalid, has not `enode://`")
			}
			url = url[8:]
			addrs := strings.Split(url, "@")
			if len(addrs[0]) != 128 {
				log.Crit("Bootstrap URL invalid, wrong length, want 128 hex chars")
			}
		}
	}

	makeWizard(c.String("network"), images, urls).run()
	return nil
}
