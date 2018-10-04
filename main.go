package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum"
	"strings"
	"net/http"
	"html/template"
	"context"
	"log"
	"math/big"
	//     "time"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func createSimulator() (*backends.SimulatedBackend, *bind.TransactOpts) {
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)} // Genesis account
	sim := backends.NewSimulatedBackend(alloc, 100000000)                  // 100000000 is gas limit
	return sim, auth
}

func deployTester(sim *backends.SimulatedBackend, auth *bind.TransactOpts) (chan types.Log, abi.ABI, common.Address, *TesterSession) {
	// Returns a channel to listen for test case results
	// and the address at which the tester is deployed
	tester, _, contract, err := DeployTester(auth, sim)
	if err != nil {
		log.Fatalf("Could not deploy contract: %v", err)
	}
	fmt.Println("Contract is at address:", tester)
	if err != nil {
		log.Fatal(err)
	}
	sim.Commit() // like mining
	// Subscribe to events logs from tester
	query := ethereum.FilterQuery{
		Addresses: []common.Address{tester},
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(TesterABI)))
	if err != nil {
		log.Fatal(err)
	}
	logs := make(chan types.Log)
	_, err = sim.SubscribeFilterLogs(context.Background(), query, logs)
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
	return logs, contractAbi, tester, session
}

type Page struct {
	Title string
	Body string
}

func submit(w http.ResponseWriter, r *http.Request) {
	// This POST takes in the user's contract code
	// Compile the smart contract code and 
	// submit it to the blockchain for tests
    body := r.FormValue("body")
	log.Println(body)
	// Compile this code
	contracts, err := compiler.CompileSolidityString("", body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}
	log.Println(contracts, err)
	http.Redirect(w, r, "/view/", http.StatusFound)
}

func edit(w http.ResponseWriter, r *http.Request) {
 	t, _ := template.ParseFiles("index.html")
    t.Execute(w, Page{Title:"test", Body:"test"})
}

func view(w http.ResponseWriter, r *http.Request) {
	// View open challenges, reputations etc.
 	t, _ := template.ParseFiles("view.html")
    t.Execute(w, struct{}{})
}

func main() {
	http.HandleFunc("/", edit)
	http.HandleFunc("/submit/", submit)
	http.HandleFunc("/view/", view)
	http.ListenAndServe(":8080", nil)
}
