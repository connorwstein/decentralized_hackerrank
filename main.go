package main

import (
	"fmt"
// 	"github.com/ethereum/go-ethereum"
// 	"strings"
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

func main() {
	// initialize the simulated backend and deploy the tester 
	// contract
	challenges = make(map[string]Challenge) 
	challenges["Adder"] = Challenge{"Adder", "easy", "function add(uint a, uint b) returns (uint)"}
	var n Node
	n.initializeNode("http://localhost:8545")
	fmt.Println(n.Accounts)
// 	adder := "608060405234801561001057600080fd5b5060c58061001f6000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063771602f7146044575b600080fd5b348015604f57600080fd5b5060766004803603810190808035906020019092919080359060200190929190505050608c565b6040518082815260200191505060405180910390f35b60008183019050929150505600a165627a7a72305820f364b5f1249cb19c2df18521b0ffe9bbf8849deb1584f4df2f9580e0a27768f00029"
	tester := "608060405234801561001057600080fd5b50610260806100206000396000f300608060405260043610610041576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632f570a2314610046575b600080fd5b34801561005257600080fd5b506100ad600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506100af565b005b6000806100bb83610222565b915060008273ffffffffffffffffffffffffffffffffffffffff1614156100e157600080fd5b81905060148173ffffffffffffffffffffffffffffffffffffffff1663771602f7600a806040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083815260200182815260200192505050602060405180830381600087803b15801561015e57600080fd5b505af1158015610172573d6000803e3d6000fd5b505050506040513d602081101561018857600080fd5b810190808051906020019092919050505014156101e0577fe7230e41682a55f671747a59595e23f58680339803ff4a8a09cbc4d1a433f84e6001604051808215151515815260200191505060405180910390a161021d565b7fe7230e41682a55f671747a59595e23f58680339803ff4a8a09cbc4d1a433f84e6000604051808215151515815260200191505060405180910390a15b505050565b60008151602083016000f090509190505600a165627a7a72305820f83f43cb3cdf51339142eeb9bbbcaf3f82d6ba75faa507cc3f39e433c932e25e0029" 
	address, _ := n.deployContract(tester)
	fmt.Println(address)
	// Call test at the tester address with argument adder 
	// https://ethereum.stackexchange.com/questions/3780/how-can-i-create-a-listener-for-new-transaction-with-ethereum-rpc-calls
	filterID := n.ethNewFilter(address, "da82b96f")
	// Run the tests for the adder code
// keccak("test(bytes)") 2f570a234f56174a0be5cf2fff788ff394b02e8140a68e91a993c49f6c1e0219
	// Since bytes is a dynamic type 
	// We will need
	// Offset of 32 as the first 32 
// "0000000000000000000000000000000000000000000000000000000000000020"
// Then the data part is the length in bytes in one 32 byte word
// Followed by all the bytes padded
// 	funcHash := "2f570a23"
// 	offset := "0000000000000000000000000000000000000000000000000000000000000020"
// 	length := len(adder)   	 // 456 --> 1c8
	data := "2f570a23000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000001c8608060405234801561001057600080fd5b5060c58061001f6000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063771602f7146044575b600080fd5b348015604f57600080fd5b5060766004803603810190808035906020019092919080359060200190929190505050608c565b6040518082815260200191505060405180910390f35b60008183019050929150505600a165627a7a72305820f364b5f1249cb19c2df18521b0ffe9bbf8849deb1584f4df2f9580e0a27768f0002900000000000000000000000000000000000000000000000000000000"
	res, err := n.ethSendTransaction(n.Accounts[0], address, data)
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
