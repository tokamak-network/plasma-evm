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

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/log"
)

// nodeDockerfile is the Dockerfile required to run an Ethereum node.
var nodeDockerfile = `
FROM {{.Image}}

ADD genesis.json /genesis.json
{{if .Operator}}
  ADD operator.json /operator.json
	ADD operator.pass /operator.pass
{{end}}
{{if .Challenger}}
  ADD challenger.json /challenger.json
	ADD challenger.pass /challenger.pass
{{end}}
{{if .Unlock}}
	ADD signer.pass /signer.pass
{{end}}
RUN \
  echo $'geth --cache 512 init --rootchain.url {{.RootChainURL}} /genesis.json' > geth.sh && \{{if .Unlock}}
	echo 'mkdir -p /root/.ethereum/keystore/' >> geth.sh && \{{end}}{{if .Operator}}
  echo 'cp /operator.json /root/.ethereum/keystore/' >> geth.sh && \{{end}}{{if .Challenger}}
	echo 'cp /challenger.json /root/.ethereum/keystore/' >> geth.sh && \{{end}}
	echo $'exec geth --syncmode="full" --networkid {{.NetworkID}} --rootchain.url {{.RootChainURL}} {{if .Operator}}--operator {{.Operator}} --operator.password /operator.pass {{end}} {{if .Challenger}}--rootchain.challenger {{.Challenger}} --challenger.password /challenger.pass{{end}} {{if .RPCPort}}--rpc --rpcaddr \'0.0.0.0\' --rpcport {{.RPCPort}} --rpcapi eth,net,debug --rpccorsdomain "*" {{if .VHOST}}--rpcvhosts={{.VHOST}}{{end}}{{end}} {{if .WSPort}}--ws --wsorigins \'*\' --wsaddr \'0.0.0.0\' --wsport {{.WSPort}}{{end}} --cache 512 --port {{.Port}} --nat extip:{{.IP}} --maxpeers {{.Peers}} {{.LightFlag}} --ethstats \'{{.Ethstats}}\' {{if .Bootnodes}}--bootnodes {{.Bootnodes}}{{end}} {{if .Etherbase}}--miner.etherbase {{.Etherbase}} --mine --miner.threads 1 --tx.interval "10s" --tx.maxgasprice 10000000000 --miner.recommit "15s" --{{end}} {{if .Unlock}}--unlock {{.Unlock}} --password /signer.pass --mine{{end}} --miner.gastarget {{.GasTarget}} --miner.gaslimit {{.GasLimit}} --miner.gasprice {{.GasPrice}}' >> geth.sh

ENTRYPOINT ["/bin/sh", "geth.sh"]
`

// nodeComposefile is the docker-compose.yml file required to deploy and maintain
// an Ethereum node (bootnode or miner for now).
var nodeComposefile = `
version: '2'
services:
  {{.Type}}:
    build: .
    image: {{.Network}}/{{.Type}}
    container_name: {{.Network}}_{{.Type}}_1
    ports:
      - "{{.Port}}:{{.Port}}"
      - "{{.Port}}:{{.Port}}/udp"{{if .RPCPort}}
      - "{{.RPCPort}}:{{.RPCPort}}"{{end}}{{if .WSPort}}
      - "{{.WSPort}}:{{.WSPort}}"{{end}}
    volumes:
      - {{.Datadir}}:/root/.ethereum{{if .Ethashdir}}
      - {{.Ethashdir}}:/root/.ethash{{end}}
    environment:
      - ROOTCHAIN_URL={{.RootChainURL}}{{if .Operator}}
      - Operator={{.Operator}}{{end}}{{if .Challenger}}
      - Challenger={{.Challenger}}{{end}}
      - PORT={{.Port}}/tcp{{if .RPCPort}}
      - RPC_PORT={{.RPCPort}}{{end}}{{if .WSPort}}
      - WS_PORT={{.WSPort}}{{end}}{{if .VHOST}}
      - VIRTUAL_HOST={{.VHOST}}{{end}}
      - TOTAL_PEERS={{.TotalPeers}}
      - LIGHT_PEERS={{.LightPeers}}
      - STATS_NAME={{.Ethstats}}
      - MINER_NAME={{.Etherbase}}
      - GAS_TARGET={{.GasTarget}}
      - GAS_LIMIT={{.GasLimit}}
      - GAS_PRICE={{.GasPrice}}
    logging:
      driver: "json-file"
      options:
        max-size: "1m"
        max-file: "10"
    restart: always
`

// deployNode deploys a new Ethereum node container to a remote machine via SSH,
// docker and docker-compose. If an instance with the specified network name
// already exists there, it will be overwritten!
func deployNode(client *sshClient, network string, image string, bootnodes []string, config *nodeInfos, nocache bool) ([]byte, error) {
	kind := "sealnode"
	if config.operatorKeyJSON == "" && config.etherbase == "" {
		kind = "usernode"
	}
	// Generate the content to upload to the server
	workdir := fmt.Sprintf("%d", rand.Int63())
	files := make(map[string][]byte)

	lightFlag := ""
	if config.peersLight > 0 {
		lightFlag = fmt.Sprintf("--lightpeers=%d --lightserv=50", config.peersLight)
	}
	dockerfile := new(bytes.Buffer)

	var key struct {
		Address string `json:"address"`
	}
	var (
		operator   string
		challenger string
		accounts   []string
		passwords  []string
	)

	if config.operatorKeyJSON != "" {
		if err := json.Unmarshal([]byte(config.operatorKeyJSON), &key); err == nil {
			operator = common.HexToAddress(key.Address).Hex()
		} else {
			log.Error("Failed to retrieve operator address", "err", err)
		}
	}

	if config.challengerKeyJSON != "" {
		if err := json.Unmarshal([]byte(config.challengerKeyJSON), &key); err == nil {
			challenger = common.HexToAddress(key.Address).Hex()
		} else {
			log.Error("Failed to retrieve operator address", "err", err)
		}
	}

	if operator != "" {
		accounts = append(accounts, operator)
		passwords = append(passwords, config.operatorKeyPass)
		files[filepath.Join(workdir, "operator.json")] = []byte(config.operatorKeyJSON)
		files[filepath.Join(workdir, "operator.pass")] = []byte(config.operatorKeyPass)
	}
	if challenger != "" {
		accounts = append(accounts, challenger)
		passwords = append(passwords, config.challengerKeyPass)
		files[filepath.Join(workdir, "challenger.json")] = []byte(config.challengerKeyJSON)
		files[filepath.Join(workdir, "challenger.pass")] = []byte(config.challengerKeyPass)
	}
	unlock := strings.Join(accounts, ",")

	template.Must(template.New("").Parse(nodeDockerfile)).Execute(dockerfile, map[string]interface{}{
		"Image":        image,
		"NetworkID":    config.network,
		"RootChainURL": config.rootchainURL,
		"Operator":     operator,
		"Challenger":   challenger,
		"Port":         config.port,
		"VHOST":        config.vhost,
		"RPCPort":      config.rpcPort,
		"WSPort":       config.wsPort,
		"IP":           client.address,
		"Peers":        config.peersTotal,
		"LightFlag":    lightFlag,
		"Bootnodes":    strings.Join(bootnodes, ","),
		"Ethstats":     config.ethstats,
		"Etherbase":    config.etherbase,
		"GasTarget":    uint64(1000000 * config.gasTarget),
		"GasLimit":     uint64(1000000 * config.gasLimit),
		"GasPrice":     uint64(1000000000 * config.gasPrice),
		"Unlock":       unlock,
	})
	files[filepath.Join(workdir, "Dockerfile")] = dockerfile.Bytes()

	composefile := new(bytes.Buffer)
	template.Must(template.New("").Parse(nodeComposefile)).Execute(composefile, map[string]interface{}{
		"Type":         kind,
		"Operator":     operator,
		"Challenger":   challenger,
		"Datadir":      config.datadir,
		"Ethashdir":    config.ethashdir,
		"Network":      network,
		"VHOST":        config.vhost,
		"RootChainURL": config.rootchainURL,
		"Port":         config.port,
		"RPCPort":      config.rpcPort,
		"WSPort":       config.wsPort,
		"TotalPeers":   config.peersTotal,
		"Light":        config.peersLight > 0,
		"LightPeers":   config.peersLight,
		"Ethstats":     config.ethstats[:strings.Index(config.ethstats, ":")],
		"Etherbase":    config.etherbase,
		"GasTarget":    config.gasTarget,
		"GasLimit":     config.gasLimit,
		"GasPrice":     config.gasPrice,
	})
	files[filepath.Join(workdir, "docker-compose.yaml")] = composefile.Bytes()

	files[filepath.Join(workdir, "genesis.json")] = config.genesis
	if unlock != "" {
		files[filepath.Join(workdir, "signer.pass")] = []byte(strings.Join(passwords, "\n"))
	}
	// Upload the deployment files to the remote server (and clean up afterwards)
	if out, err := client.Upload(files); err != nil {
		return out, err
	}
	defer client.Run("rm -rf " + workdir)

	// Build and deploy the boot or seal node service
	if nocache {
		return nil, client.Stream(fmt.Sprintf("cd %s && docker-compose -p %s build --pull --no-cache && docker-compose -p %s up -d --force-recreate --timeout 60", workdir, network, network))
	}
	return nil, client.Stream(fmt.Sprintf("cd %s && docker-compose -p %s up -d --build --force-recreate --timeout 60", workdir, network))
}

// nodeInfos is returned from a boot or seal node status check to allow reporting
// various configuration parameters.
type nodeInfos struct {
	genesis           []byte
	network           int64
	datadir           string
	ethashdir         string
	ethstats          string
	rootchainURL      string
	vhost             string
	rpcPort           int
	wsPort            int
	port              int
	enode             string
	peersTotal        int
	peersLight        int
	etherbase         string
	operatorKeyJSON   string
	operatorKeyPass   string
	challengerKeyJSON string
	challengerKeyPass string
	gasTarget         float64
	gasLimit          float64
	gasPrice          float64
}

// Report converts the typed struct into a plain string->string map, containing
// most - but not all - fields for reporting to the user.
func (info *nodeInfos) Report() map[string]string {
	report := map[string]string{
		"Data directory":           info.datadir,
		"Listener port":            strconv.Itoa(info.port),
		"Peer count (all total)":   strconv.Itoa(info.peersTotal),
		"Peer count (light nodes)": strconv.Itoa(info.peersLight),
		"Ethstats username":        info.ethstats,
		"Root chain JSONRPC URL":   info.rootchainURL,
	}
	if info.gasTarget > 0 {
		// Miner or signer node
		report["Gas price (minimum accepted)"] = fmt.Sprintf("%0.3f GWei", info.gasPrice)
		report["Gas floor (baseline target)"] = fmt.Sprintf("%0.3f MGas", info.gasTarget)
		report["Gas ceil  (target maximum)"] = fmt.Sprintf("%0.3f MGas", info.gasLimit)

		if info.etherbase != "" {
			// Ethash proof-of-work miner
			report["Ethash directory"] = info.ethashdir
			report["Miner account"] = info.etherbase
		}
		if info.operatorKeyJSON != "" {
			// Clique proof-of-authority signer
			var key struct {
				Address string `json:"address"`
			}
			if err := json.Unmarshal([]byte(info.operatorKeyJSON), &key); err == nil {
				report["Operator account"] = common.HexToAddress(key.Address).Hex()
			} else {
				log.Error("Failed to retrieve operator address", "err", err)
			}
			if err := json.Unmarshal([]byte(info.challengerKeyJSON), &key); err == nil {
				report["CHallenger account"] = common.HexToAddress(key.Address).Hex()
			} else {
				log.Error("Failed to retrieve challenger address", "err", err)
			}
		}
	}
	if info.rpcPort != 0 {
		report["JSONRPC HTTP port"] = strconv.Itoa(info.rpcPort)
	}
	if info.wsPort != 0 {
		report["JSONRPC WebSocket port"] = strconv.Itoa(info.wsPort)
	}
	if info.vhost != "" {
		report["JSONRPC VHOST"] = info.vhost
	}
	return report
}

// checkNode does a health-check against a boot or seal node server to verify
// whether it's running, and if yes, whether it's responsive.
func checkNode(client *sshClient, network string, user bool) (*nodeInfos, error) {
	kind := "usernode"
	if !user {
		kind = "sealnode"
	}
	// Inspect a possible usernode container on the host
	infos, err := inspectContainer(client, fmt.Sprintf("%s_%s_1", network, kind))
	if err != nil {
		return nil, err
	}
	if !infos.running {
		return nil, ErrServiceOffline
	}
	// Resolve a few types from the environmental variables
	totalPeers, _ := strconv.Atoi(infos.envvars["TOTAL_PEERS"])
	lightPeers, _ := strconv.Atoi(infos.envvars["LIGHT_PEERS"])
	gasTarget, _ := strconv.ParseFloat(infos.envvars["GAS_TARGET"], 64)
	gasLimit, _ := strconv.ParseFloat(infos.envvars["GAS_LIMIT"], 64)
	gasPrice, _ := strconv.ParseFloat(infos.envvars["GAS_PRICE"], 64)
	rpcPort, _ := strconv.Atoi(infos.envvars["RPC_PORT"])
	wsPort, _ := strconv.Atoi(infos.envvars["WS_PORT"])

	// Container available, retrieve its node ID and its genesis json
	var out []byte
	if out, err = client.Run(fmt.Sprintf("docker exec %s_%s_1 geth --exec admin.nodeInfo.enode --cache=16 attach", network, kind)); err != nil {
		return nil, ErrServiceUnreachable
	}
	enode := bytes.Trim(bytes.TrimSpace(out), "\"")

	if out, err = client.Run(fmt.Sprintf("docker exec %s_%s_1 cat /genesis.json", network, kind)); err != nil {
		return nil, ErrServiceUnreachable
	}
	genesis := bytes.TrimSpace(out)

	operatorKeyJSON, challengerKeyJSON := "", ""
	operatorKeyPass, challengerKeyPass := "", ""

	if out, err = client.Run(fmt.Sprintf("docker exec %s_%s_1 cat /operator.json", network, kind)); err == nil {
		operatorKeyJSON = string(bytes.TrimSpace(out))
	}
	if out, err = client.Run(fmt.Sprintf("docker exec %s_%s_1 cat /challenger.json", network, kind)); err == nil {
		challengerKeyJSON = string(bytes.TrimSpace(out))
	}
	if out, err = client.Run(fmt.Sprintf("docker exec %s_%s_1 cat /signer.pass", network, kind)); err == nil {
		keyPass := string(bytes.TrimSpace(out))
		passwords := strings.Split(keyPass, "\n")

		if len(passwords) == 2 {
			operatorKeyPass = passwords[0]
			challengerKeyPass = passwords[1]
		}
	}
	// Run a sanity check to see if the devp2p is reachable
	port := infos.portmap[infos.envvars["PORT"]]
	if err = checkPort(client.server, port); err != nil {
		log.Warn(fmt.Sprintf("%s devp2p port seems unreachable", strings.Title(kind)), "server", client.server, "port", port, "err", err)
	}
	// Assemble and return the useful infos
	stats := &nodeInfos{
		genesis:           genesis,
		datadir:           infos.volumes["/root/.ethereum"],
		ethashdir:         infos.volumes["/root/.ethash"],
		rootchainURL:      infos.envvars["ROOTCHAIN_URL"],
		port:              port,
		vhost:             infos.envvars["VIRTUAL_HOST"],
		peersTotal:        totalPeers,
		peersLight:        lightPeers,
		ethstats:          infos.envvars["STATS_NAME"],
		etherbase:         infos.envvars["MINER_NAME"],
		operatorKeyJSON:   operatorKeyJSON,
		operatorKeyPass:   operatorKeyPass,
		challengerKeyJSON: challengerKeyJSON,
		challengerKeyPass: challengerKeyPass,
		gasTarget:         gasTarget,
		gasLimit:          gasLimit,
		gasPrice:          gasPrice,
	}
	stats.enode = string(enode)

	if rpcPort != 0 {
		if err = checkPort(client.server, rpcPort); err != nil {
			log.Warn(fmt.Sprintf("%s HTTP JSONRPC endpoint seems unreachable", strings.Title(kind)), "server", client.server, "port", rpcPort, "err", err)
		} else {
			stats.rpcPort = rpcPort
		}
	}
	if wsPort != 0 {
		if err = checkPort(client.server, wsPort); err != nil {
			log.Warn(fmt.Sprintf("%s WebSocket JSONRPC endpoint seems unreachable", strings.Title(kind)), "server", client.server, "port", wsPort, "err", err)
		} else {
			stats.wsPort = wsPort
		}
	}

	return stats, nil
}
