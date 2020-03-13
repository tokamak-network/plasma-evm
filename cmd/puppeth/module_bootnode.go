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
	"fmt"
	"github.com/Onther-Tech/plasma-evm/log"
	"html/template"
	"math/rand"
	"path/filepath"
	"strconv"
)

// faucetDockerfile is the Dockerfile required to build a faucet container to
// grant crypto tokens based on GitHub authentications.
var bootnodeDockerfile = `
FROM onthertech/plasma-evm:alltools-latest

EXPOSE 30301 30301/udp

WORKDIR /usr/local/bin
RUN echo $'bootnode -nodekeyhex {{.NodeKey}} -nat extip:{{.Host}} -addr :{{.Port}} -verbosity 6' > run.bootnode.sh

ENTRYPOINT ["/bin/sh", "run.bootnode.sh"]
`

// bootnodeComposefile is the docker-compose.yaml file
var bootnodeComposefile = `
version: '2'
services:
  bootnode:
    build: .
    image: {{.Network}}/bootnode
    container_name: {{.Network}}_bootnode_1
    network_mode: "host"
    environment:
      - BOOTNODE_ADDR={{.Enode}}
      - BOOTNODE_PORT={{.Port}}
      - HOST={{.Host}}
    logging:
      driver: "json-file"
      options:
        max-size: "1m"
        max-file: "10"
    restart: always
`

// bootnodeInfos is returned from a bootnode status check to allow reporting various
// configuration parameters.
type bootnodeInfos struct {
	host    string
	nodekey string
	enode   string
	port    int
}

// Report converts the typed struct into a plain string->string map, containing
// most - but not all - fields for reporting to the user.
func (info *bootnodeInfos) Report() map[string]string {
	report := map[string]string{
		"IP address":    info.host,
		"listener port": strconv.Itoa(info.port),
		"enode address": info.enode[:14] + " ... ",
	}
	return report
}

// deploybootnode deploys a new bootnode container to a remote machine via SSH,
// docker and docker-compose. If an instance with the specified network name
// already exists there, it will be overwritten!
func deployBootnode(client *sshClient, network string, config *bootnodeInfos, nocache bool) ([]byte, error) {
	// Generate the content to upload to the server
	workdir := fmt.Sprintf("%d", rand.Int63())
	files := make(map[string][]byte)

	dockerfile := new(bytes.Buffer)
	template.Must(template.New("").Parse(bootnodeDockerfile)).Execute(dockerfile, map[string]interface{}{
		"NodeKey": config.nodekey,
		"Host":    config.host,
		"Port":    config.port,
	})
	files[filepath.Join(workdir, "Dockerfile")] = dockerfile.Bytes()

	composefile := new(bytes.Buffer)
	template.Must(template.New("").Parse(bootnodeComposefile)).Execute(composefile, map[string]interface{}{
		"Network": network,
		"Enode":   config.enode,
		"Host":    config.host,
		"Port":    config.port,
	})
	files[filepath.Join(workdir, "docker-compose.yaml")] = composefile.Bytes()

	// Upload the deployment files to the remote server (and clean up afterwards)
	if out, err := client.Upload(files); err != nil {
		return out, err
	}
	defer client.Run("rm -rf " + workdir)

	// Build and deploy the bootnode service
	if nocache {
		return nil, client.Stream(fmt.Sprintf("cd %s && docker-compose -p %s build --pull --no-cache && docker-compose -p %s up -d --force-recreate --timeout 60", workdir, network, network))
	}
	return nil, client.Stream(fmt.Sprintf("cd %s && docker-compose -p %s up -d --build --force-recreate --timeout 60", workdir, network))
}

// checkbootnode does a health-check against a bootnode server to verify whether
// it's running, and if yes, gathering a collection of useful infos about it.
// TODO : Make that not deploy same instance with operator or usernode
func checkBootnode(client *sshClient, network string) (*bootnodeInfos, error) {
	// Inspect a possible bootnode container on the host
	infos, err := inspectContainer(client, fmt.Sprintf("%s_bootnode_1", network))
	if err != nil {
		return nil, err
	}
	if !infos.running {
		return nil, ErrServiceOffline
	}

	// Resolve the host from container environments
	host := infos.envvars["HOST"]
	port, err := strconv.Atoi(infos.envvars["BOOTNODE_PORT"])
	if err != nil {
		log.Error("Failed convert port number", "err", err)
	}
	enode := infos.envvars["BOOTNODE_ADDR"]

	// Run a sanity check to see if the port is reachable
	if err = checkUDPPort(host, port); err != nil {
		log.Warn("bootnode service seems unreachable", "server", host, "port", port, "err", err)
	}
	// Container available, assemble and return the useful infos
	return &bootnodeInfos{
		host: host,
		port: port,
		enode: enode,
	}, nil
}
