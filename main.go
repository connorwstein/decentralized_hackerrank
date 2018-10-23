package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum"
	"strings"
	"net/http"
	"html/template"
	"context"
	"encoding/hex"
// 	"reflect"
	"log"
	"math/big"
    "time"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var runner SubmissionRunner 

type SubmissionRunner struct {
	sim *backends.SimulatedBackend
	auth *bind.TransactOpts
	logs chan types.Log
	contractAbi abi.ABI
	address common.Address
	session *TesterSession
	results []Result
}

type Challenge struct {
	Name string
	Difficulty string	
	Interface string
}

var challenges map[string]Challenge

func (s *SubmissionRunner) create() {
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)} // Genesis account
	sim := backends.NewSimulatedBackend(alloc, 100000000)                  // 100000000 is gas limit
	s.sim = sim
	s.auth = auth
	s.results = make([]Result, 0)
}

func (s *SubmissionRunner) deployRunner() {
	// Returns a channel to listen for test case results
	// and the address at which the tester is deployed
	tester, _, contract, err := DeployTester(s.auth, s.sim)
	if err != nil {
		log.Fatalf("Could not deploy contract: %v", err)
	}
	fmt.Println("Contract is at address:", tester)
	if err != nil {
		log.Fatal(err)
	}
	s.sim.Commit() // like mining
	// Subscribe to events logs from tester
	query := ethereum.FilterQuery{
		Addresses: []common.Address{tester},
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(TesterABI)))
	if err != nil {
		log.Fatal(err)
	}
	logs := make(chan types.Log)
	_, err = s.sim.SubscribeFilterLogs(context.Background(), query, logs)
	session := &TesterSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     s.auth.From,
			Signer:   s.auth.Signer,
			GasLimit: 3141592,
		},
	}
	s.logs = logs
	s.contractAbi = contractAbi
	s.address = tester  
	s.session = session
}

func (s SubmissionRunner) submitCode(contractSubmission *compiler.Contract) bool {
	resultChan := s.getResult()
	adderCode, err := hex.DecodeString(contractSubmission.Code[2:])
	tx, err := s.session.Test(adderCode)
	receipt, rerr := s.sim.TransactionReceipt(context.Background(), tx.Hash())
	log.Println("Receipt: ", receipt, rerr)
	log.Println("Result:", tx.Hash(), err)
	log.Printf("Submitted code %v", adderCode)
	s.sim.Commit() // like mining, should trigger the log
	return <- resultChan
}

func (s SubmissionRunner) getResult() chan bool {
	// Start listening for logs
	result := make(chan bool)
	go func(result chan bool) {
		timeout := time.After(10 * time.Second)
		select {
			case <-timeout:
				log.Println("Timeout waiting for response")
				result <- false
			case vLog := <-s.logs:
				var event bool
				log.Println(s.contractAbi, vLog.Data)
				err := s.contractAbi.Unpack(&event, "TestPass", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				log.Println("Got a log", event) // pointer to event log
				result <- event
		}
	}(result)
	return result 
}

func submit(w http.ResponseWriter, r *http.Request) {
	// This POST takes in the user's contract code
	// Compile the smart contract code and 
	// submit it to the blockchain for tests
    body := r.FormValue("body")
	log.Println(body)
	// Compile this code returns 
	// NOTE this compiler wrapper from the geth codebase
	// automatically enables optimization, which for unknown reasons
	// does not work with my tester contract to make this work you
	// need to comment out --optimize in makeArgs in common/compiler/solidity.go 
	fmt.Println(compiler.SolidityVersion(""))
	contracts, err := compiler.CompileSolidityString("", body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}
	log.Println(contracts, err)
	// If it compiles submit it to the tester and get a result
	if adder, ok := contracts["<stdin>:Adder"]; ok {
		log.Println("Adder code", adder.Code)
		// Submit the code
		res := runner.submitCode(adder)
		log.Printf("Submission result was %v", res)
		// Based on the result update the screen
		runner.results = append(runner.results, Result{runner.auth.From.Hex(), "Adder", res})
	} else {
		http.Error(w, "Does not adhere to interface", http.StatusBadRequest)
		return	
	}
	http.Redirect(w, r, "/view/", http.StatusFound)
}

func edit(w http.ResponseWriter, r *http.Request) {
 	t, err := template.ParseFiles("edit.html")
	if err != nil {
		fmt.Println("Error when parsing template ", err)
	}
	challenge := r.URL.Path[len("/edit/"):]
	fmt.Println("Editing challenge: ", challenge, challenges[challenge])
    t.Execute(w, challenges[challenge])
}

type Result struct {
	Submitter string // public key
	Challenge string	
	Pass bool
}

func viewResults(w http.ResponseWriter, r *http.Request) {
	// View open challenges, reputations etc.
	// We should be able to query the chain to get this info
	// Clicking on a challenge should take you to the edit page
 	t, err := template.ParseFiles("view.html")
	if err != nil {
		fmt.Println("Error when parsing template ", err)
	}
    t.Execute(w, runner.results)
}

func viewChallenges(w http.ResponseWriter, r *http.Request) {
	// Walk through all the available challenges
 	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("Error when parsing template ", err)
	}
    t.Execute(w, challenges)
}

func main() {
	// initialize the simulated backend and deploy the tester 
	// contract
	challenges = make(map[string]Challenge) 
	challenges["Adder"] = Challenge{"Adder", "easy", "function add(uint a, uint b) returns (uint)"}
	runner.create()
	runner.deployRunner()
	http.HandleFunc("/", viewChallenges)
	http.HandleFunc("/edit/", edit)
	http.HandleFunc("/submit/", submit)
	http.HandleFunc("/view/", viewResults)
	http.ListenAndServe(":8080", nil)
}
