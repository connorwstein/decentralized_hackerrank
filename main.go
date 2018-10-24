package main

import (
	"fmt"
// 	"github.com/ethereum/go-ethereum"
	"strings"
	"bytes"
	"os/exec"
	"io/ioutil"
	"net/http"
	"html/template"
// 	"context"
// 	"encoding/hex"
// 	"reflect"
// 	"log"
// 	"math/big"
//     "time"
// 	"github.com/ethereum/go-ethereum/accounts/abi"
// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/common/compiler"
// 	"github.com/ethereum/go-ethereum/core"
// 	"github.com/ethereum/go-ethereum/core/types"
// 	"github.com/ethereum/go-ethereum/crypto"
)
type Challenge struct {
	Name string
	Difficulty string	
	Interface string
}

var challenges map[string]Challenge

// func (s SubmissionRunner) submitCode(contractSubmission *compiler.Contract) bool {
// 	resultChan := s.getResult()
// 	adderCode, err := hex.DecodeString(contractSubmission.Code[2:])
// 	tx, err := s.session.Test(adderCode)
// 	receipt, rerr := s.sim.TransactionReceipt(context.Background(), tx.Hash())
// 	log.Println("Receipt: ", receipt, rerr)
// 	log.Println("Result:", tx.Hash(), err)
// 	log.Printf("Submitted code %v", adderCode)
// 	s.sim.Commit() // like mining, should trigger the log
// 	return <- resultChan
// }
// 
// func (s SubmissionRunner) getResult() chan bool {
// 	// Start listening for logs
// 	result := make(chan bool)
// 	go func(result chan bool) {
// 		timeout := time.After(10 * time.Second)
// 		select {
// 			case <-timeout:
// 				log.Println("Timeout waiting for response")
// 				result <- false
// 			case vLog := <-s.logs:
// 				var event bool
// 				log.Println(s.contractAbi, vLog.Data)
// 				err := s.contractAbi.Unpack(&event, "TestPass", vLog.Data)
// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 				log.Println("Got a log", event) // pointer to event log
// 				result <- event
// 		}
// 	}(result)
// 	return result 
// }
// 

// func submit(w http.ResponseWriter, r *http.Request) {
// 	// This POST takes in the user's contract code
// 	// Compile the smart contract code and 
// 	// submit it to the blockchain for tests
//     body := r.FormValue("body")
// 	log.Println(body)
// 	// Compile this code returns 
// 	// NOTE this compiler wrapper from the geth codebase
// 	// automatically enables optimization, which for unknown reasons
// 	// does not work with my tester contract to make this work you
// 	// need to comment out --optimize in makeArgs in common/compiler/solidity.go 
// 	fmt.Println(compiler.SolidityVersion(""))
// 	contracts, err := compiler.CompileSolidityString("", body)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return 
// 	}
// 	log.Println(contracts, err)
// 	// If it compiles submit it to the tester and get a result
// 	if adder, ok := contracts["<stdin>:Adder"]; ok {
// 		log.Println("Adder code", adder.Code)
// 		// Submit the code
// 		res := runner.submitCode(adder)
// 		log.Printf("Submission result was %v", res)
// 		// Based on the result update the screen
// 		runner.results = append(runner.results, Result{runner.auth.From.Hex(), "Adder", res})
// 	} else {
// 		http.Error(w, "Does not adhere to interface", http.StatusBadRequest)
// 		return	
// 	}
// 	http.Redirect(w, r, "/view/", http.StatusFound)
// }
// 
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

// func viewResults(w http.ResponseWriter, r *http.Request) {
// 	// View open challenges, reputations etc.
// 	// We should be able to query the chain to get this info
// 	// Clicking on a challenge should take you to the edit page
//  	t, err := template.ParseFiles("view.html")
// 	if err != nil {
// 		fmt.Println("Error when parsing template ", err)
// 	}
//     t.Execute(w, runner.results)
// }

func viewChallenges(w http.ResponseWriter, r *http.Request) {
	// Walk through all the available challenges
 	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("Error when parsing template ", err)
	}
    t.Execute(w, challenges)
}

// Given path to tester contract, compile it and 
// and return a string of bytecode 
func compileContract(name, path string) string {
	// Assumes solc is install locally
	// cant use the compiler in geth codebase because of
	// the optimization bug
	// solc --bin execute.sol -o .
	cmd := exec.Command("solc", "--bin",  path, "-o .", "--overwrite")
	fmt.Println(cmd)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("err running cmd", err)
		return "" 
	}
	contents, err := ioutil.ReadFile(fmt.Sprintf("%v.bin", name))
	buf := bytes.NewBuffer(contents)
	return buf.String()
}

func padRightSide(str string, item string, count int) string {
	return str + strings.Repeat(item, count)
}

func buildInput(testContract string) string {
	// Given a testContract string of bytes
	// Assemble that into input data for a call to 
	// test in the Tester contract
	functionSelector := "2f570a23" // 4 bytes of keccak(test(bytes))
	// Dynamic arguments have their offset stored first 
	// In our case test only has one argument (bytes)
	// So it starts 32 bytes (0x20) after the start of the arguments
	// section
	offset := fmt.Sprintf("%064x", 32) 
	length := fmt.Sprintf("%064x", len(testContract))
	// Now the data section
	// Will need to pad the testContract with zeroes 
	// to make it a multiple of 32
	testContract = padRightSide(testContract, "0", (len(testContract) / 32 + 1) * 32 - len(testContract))
	return strings.Join([]string{functionSelector, offset, length, testContract}, "")
}

func main() {
	challenges = make(map[string]Challenge) 
	challenges["Adder"] = Challenge{"Adder", "easy", "function add(uint a, uint b) returns (uint)"}
	tester := compileContract("Tester", "execute.sol")
	var n Node
	n.initializeNode("http://localhost:8545")
	fmt.Println(n.Accounts)
	adder := "608060405234801561001057600080fd5b5060c58061001f6000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063771602f7146044575b600080fd5b348015604f57600080fd5b5060766004803603810190808035906020019092919080359060200190929190505050608c565b6040518082815260200191505060405180910390f35b60008183019050929150505600a165627a7a72305820f364b5f1249cb19c2df18521b0ffe9bbf8849deb1584f4df2f9580e0a27768f00029"
	address, _ := n.deployContract(tester)
	fmt.Println(address)
	// Call test at the tester address with argument adder 
	// https://ethereum.stackexchange.com/questions/3780/how-can-i-create-a-listener-for-new-transaction-with-ethereum-rpc-calls
	filterID := n.ethNewFilter(address, "da82b96f")
	fmt.Println("input", buildInput(adder))	
	res, err := n.ethSendTransaction(n.Accounts[0], address, buildInput(adder))
	fmt.Println(res, err)

	// Query filter for results
	// given filterID
	n.ethGetFilterChanges(filterID)

// 	runner.create()
// 	runner.deployRunner()
// 	http.HandleFunc("/", viewChallenges)
// 	http.HandleFunc("/edit/", edit)
// 	http.HandleFunc("/submit/", submit)
// 	http.HandleFunc("/view/", viewResults)
// 	http.ListenAndServe(":8080", nil)
}
