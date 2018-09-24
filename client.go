package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"bytes"
	"io/ioutil"

    "log"
    "math/big"
//     "time"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
    "github.com/ethereum/go-ethereum/core"
    "github.com/ethereum/go-ethereum/crypto"
)

type Request struct {
	Jsonrpc string  `json:"jsonrpc"` 
	Method string   `json:"method"` 
	Params []string `json:"params"` 
	Id int	  		`json:"id"` 
}

type Response struct {
	Id int `json:"id"` 
	Jsonrpc string `json:"jsonrpc"`
	Result []string `json:"result"`
}
type Response2 struct {
	Id int `json:"id"` 
	Jsonrpc string `json:"jsonrpc"`
	Result string `json:"result"`
}

// params: [{
//   "from": "0xb60e8dd61c5d32be8058bb8eb970870f07233155",
//   "to": "0xd46e8dd67c5d32be8058bb8eb970870f07244567",
//   "gas": "0x76c0", // 30400
//   "gasPrice": "0x9184e72a000", // 10000000000000
//   "value": "0x9184e72a", // 2441406250
//   "data": "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"
// }]

// type EthSendTransactionRequest struct {
// 	From string `json:"from"`
// 	To string `json:"to"`
// 	Gas string `json:"gas"`
// 	GasPrice string `json:"gasPrice"`
// 	Value string `json:"value"`
// 	Data string `json:"data"`
// } 
// 
// type EthSendTransactionResponse struct {
// 	Result string 
// }
// 

func sendRequest(rpcName string, params []string) {
	fmt.Println("sending request")
	t := Request{"2.0", rpcName, params, 1}
	b, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(err)
	resp, err := http.Post("http://localhost:8545", "application/json", bytes.NewReader(b))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	if err != nil {
		fmt.Println(err)
		return	
	}
	if rpcName == "eth_accounts" {
		var r Response 
		err = json.Unmarshal(body, &r)
		if err != nil {
			fmt.Println(err)
			return	
		}
		fmt.Println("response", r)
	} else {
		var r Response2
		err = json.Unmarshal(body, &r)
		if err != nil {
			fmt.Println(err)
			return	
		}
		fmt.Println("response", r)
	}
}

func get(contract *SimpleStorage) {
	opts := &bind.CallOpts{}
	res, err := contract.Get(opts) // Calling a contract function
	if err != nil {
		fmt.Println("Failed to get")
		return
	}
	fmt.Println("GET:", res, err)
}

func set(auth *bind.TransactOpts, contract *SimpleStorage, value int64) {
	opts2 := &bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: 2381623,
		Value:    big.NewInt(10000),
	}
	res2, err := contract.Set(opts2, big.NewInt(value)) // Calling a contract function
	if err != nil {
		fmt.Println("Failed to set")
		return
	}
	fmt.Println("SET:", res2, err)
}
// 
// Interact with ganach via RPC 
func main() {
// 	sendRequest("eth_accounts", []string{})
// 	sendRequest("eth_getBalance", []string{"0x8c19973ef264e80f930c9a095b20fac7e9607174", "latest"})
// 	sendRequest("eth_getBalance", []string{"0x8c19973ef264e80f930c9a095b20fac7e9607174", "latest"})
// params: [{
//   "from": "0xb60e8dd61c5d32be8058bb8eb970870f07233155",
//   "to": "0xd46e8dd67c5d32be8058bb8eb970870f07244567",
//   "gas": "0x76c0", // 30400
//   "gasPrice": "0x9184e72a000", // 10000000000000
//   "value": "0x9184e72a", // 2441406250
//   "data": "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"
// }]
// 	sendRequest("eth_sendTransaction", []string{"0x8c19973ef264e80f930c9a095b20fac7e9607174", "0x00", "0x76c0", "0x9184e72a000", "0x9184e72a",   "latest"})
    key, _ := crypto.GenerateKey()
    auth := bind.NewKeyedTransactor(key)
    alloc := make(core.GenesisAlloc)
    alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
    sim := backends.NewSimulatedBackend(alloc, 100000000) // 100000000 is gas limit
	// Deploying the contract, auth represents our identity
	// This contract is an instance of 
// type SimpleStorage struct {
// 	SimpleStorageCaller     // Read-only binding to the contract
// 	SimpleStorageTransactor // Write-only binding to the contract
// 	SimpleStorageFilterer   // Log filterer for contract events
// }
// 	addr, _, contract, err := DeploySimpleStorage(auth, sim)
	factAddr, _, _, err := DeployFactory(auth, sim)
	addr, _, contract, err := DeployTester(auth, sim, factAddr)
	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
	}
	fmt.Println("Contract is at address:", addr)
	sim.Commit() // like mining
	session := &TesterSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: 3141592,
		},
	}
	value, err := session.Test(big.NewInt(100), big.NewInt(101))
	fmt.Println(value, err)
// 	trans, err := session.Set(big.NewInt(100))
// 	fmt.Println(trans, err)
// 	sim.Commit() // like mining
// 	v2, err := session.Get()
// 	fmt.Println(v2, err)
// 	fmt.Println(set(auth, contract, 25)
// 	fmt.Println(sim)
// 	get(contract)
}

