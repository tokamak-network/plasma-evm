package registry

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Onther-Tech/plasma-evm/common"
)

// TODO: check connection status.
type Manager struct {
	TON               common.Address `json:TON`
	WTON              common.Address `json:WTON`
	RootChainRegistry common.Address `json:RootChainRegistry`
	DepositManager    common.Address `json:DepositManager`
	SeigManager       common.Address `json:SeigManager`
	// PowerTON          common.Address `json:PowerTON`
}

func SetManagers(tonAddr, wtonAddr, registryAddr, depositManagerAddr, seigManagerAddr common.Address) error {
	manager := Manager{tonAddr, wtonAddr, registryAddr, depositManagerAddr, seigManagerAddr}
	pbytes, _ := json.Marshal(manager)
	buffer := bytes.NewBuffer(pbytes)

	// TODO: use environment variable
	resp, err := http.Post("http://localhost:9000/managers", "application/json", buffer)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func GetManagers() (*Manager, error) {
	resp, err := http.Get("http://localhost:9000/managers")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	managers, err := parse(body)
	if err != nil {
		return nil, err
	}

	return managers, nil
}

func SetRootChain(addr common.Address) error {
	requestBody, err := json.Marshal(map[string]common.Address{
		"address": addr,
	})
	if err != nil {
		return err
	}

	resp, err := http.Post("http://localhost:9000/rootchains", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func parse(body []byte) (*Manager, error) {
	var m = new(Manager)
	err := json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}
	return m, err
}
