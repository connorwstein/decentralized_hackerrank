package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum"
	"strings"
	"net/http"
	"html/template"
	"context"
	"encoding/hex"
	"reflect"
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
}

func (s *SubmissionRunner) create() {
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)} // Genesis account
	sim := backends.NewSimulatedBackend(alloc, 100000000)                  // 100000000 is gas limit
	s.sim = sim
	s.auth = auth
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
	workingCode := "608060405234801561001057600080fd5b5060c58061001f6000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063771602f7146044575b600080fd5b348015604f57600080fd5b5060766004803603810190808035906020019092919080359060200190929190505050608c565b6040518082815260200191505060405180910390f35b60008183019050929150505600a165627a7a72305820e80809f762456af961d7357d0db84ff34c4cf49807222c472768ee39b60ac7a80029"
	adderCode, _ := hex.DecodeString(contractSubmission.Code)
	log.Printf("Contract info %v", contractSubmission.Info)
// 	adderCode, _ := hex.DecodeString(workingCode)
	// 	fmt.Println("adder bytecode err", err)
	tx, err := s.session.Test(adderCode)
	log.Println("Result:", tx.Hash(), err)
	log.Printf("Submitted code %v working %v", contractSubmission.Code, workingCode)
	s.sim.Commit() // like mining, should trigger the log
	return <- resultChan
}

func (s SubmissionRunner) getResult() chan bool {
	// Start listening for logs
	result := make(chan bool)
	go func(result chan bool) {
		timeout := time.After(10 * time.Second)
		for {
			select {
			// 			case err := <-sub.Err():
			// 				log.Fatal(err)
			// 				break
			case <-timeout:
				log.Println("Timeout waiting for response")
				result <- false
				break
			case vLog := <-s.logs:
				var event bool
				log.Println(s.contractAbi, vLog.Data)
				err := s.contractAbi.Unpack(&event, "TestPass", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				log.Println("Got a log", event) // pointer to event log
				result <- event
				break
			}
		}
	}(result)
	return result 
}

type Submission struct {
	Title string
}

func submit(w http.ResponseWriter, r *http.Request) {
	// This POST takes in the user's contract code
	// Compile the smart contract code and 
	// submit it to the blockchain for tests
    body := r.FormValue("body")
	log.Println(body)
	// Compile this code returns 
	// map[string]*Contract
	contracts, err := compiler.CompileSolidityString("", body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}
	log.Println(contracts, err)
	// If it compiles submit it to the tester and get a result
	for k,_ := range contracts {
		log.Println(reflect.TypeOf(k))
		log.Println("KEY:", k)
	}
	if adder, ok := contracts["<stdin>:Adder"]; ok {
// 		log.Println("Adder code", adder.Code)
		// Submit the code
		res := runner.submitCode(adder)
		log.Printf("Submission result was %v", res)
	} else {
		http.Error(w, "Does not adhere to interface", http.StatusBadRequest)
		return	
	}
// 	contracts["
	http.Redirect(w, r, "/view/", http.StatusFound)
}

func edit(w http.ResponseWriter, r *http.Request) {
 	t, _ := template.ParseFiles("index.html")
    t.Execute(w, Submission{Title:"adder"})
}

func view(w http.ResponseWriter, r *http.Request) {
	// View open challenges, reputations etc.
	// We should be able to query the chain to get this info
	// Clicking on a challenge should take you to the edit page
 	t, _ := template.ParseFiles("view.html")
    t.Execute(w, struct{}{})
}

func main() {
	// initialize the simulated backend and deploy the tester 
	// contract
	runner.create()
	runner.deployRunner()
	http.HandleFunc("/", edit)
	http.HandleFunc("/submit/", submit)
	http.HandleFunc("/view/", view)
	http.ListenAndServe(":8080", nil)
}
