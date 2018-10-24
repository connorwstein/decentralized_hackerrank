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
	"log"
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

type DHR struct {
	client Client // client 
	filterID string // ID to query for results
	results []Result // actual results
	testAddress string // address of the tester contract
	challenges map[string]Challenge
}

type Challenge struct {
	Name string
	Difficulty string	
	Interface string
}

type Result struct {
	Submitter string // public key
	Challenge string	
	Pass bool
}

var dhr DHR

func submit(w http.ResponseWriter, r *http.Request) {
	// This POST takes in the user's contract code
	// Compile the smart contract code and 
	// submit it to the blockchain for tests
    body := r.FormValue("body")
	log.Println(body)
	buf := bytes.NewBufferString(body)
	// Body will be a smart contract which we need to compile and then submit
	err := ioutil.WriteFile("adder.sol", buf.Bytes(), 0644)
	adder := compileContract("Adder", "adder.sol")
// adder := "608060405234801561001057600080fd5b5060c58061001f6000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063771602f7146044575b600080fd5b348015604f57600080fd5b5060766004803603810190808035906020019092919080359060200190929190505050608c565b6040518082815260200191505060405180910390f35b60008183019050929150505600a165627a7a72305820f364b5f1249cb19c2df18521b0ffe9bbf8849deb1584f4df2f9580e0a27768f00029"
	fmt.Println("from:", dhr.client.Accounts[0])
	fmt.Println("to:", dhr.testAddress)
	res, err := dhr.client.ethSendTransaction(dhr.client.Accounts[0], dhr.testAddress, buildInput(adder))
	fmt.Println(res, err)

	// Query filter for results
	// given filterID
	dhr.client.ethGetFilterChanges(dhr.filterID)
	// Update results
	http.Redirect(w, r, "/view/", http.StatusFound)
}

func edit(w http.ResponseWriter, r *http.Request) {
 	t, err := template.ParseFiles("edit.html")
	if err != nil {
		fmt.Println("Error when parsing template ", err)
	}
	challenge := r.URL.Path[len("/edit/"):]
	fmt.Println("Editing challenge: ", challenge, dhr.challenges[challenge])
    t.Execute(w, dhr.challenges[challenge])
}

func viewResults(w http.ResponseWriter, r *http.Request) {
	// View open challenges, reputations etc.
	// We should be able to query the chain to get this info
	// Clicking on a challenge should take you to the edit page
 	t, err := template.ParseFiles("view.html")
	if err != nil {
		fmt.Println("Error when parsing template ", err)
	}
    t.Execute(w, dhr.results)
}

func viewChallenges(w http.ResponseWriter, r *http.Request) {
	// Walk through all the available challenges
 	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("Error when parsing template ", err)
	}
    t.Execute(w, dhr.challenges)
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
	challenges := make(map[string]Challenge) 
	challenges["Adder"] = Challenge{"Adder", "easy", "function add(uint a, uint b) returns (uint)"}
	tester := compileContract("Tester", "execute.sol")
	var client Client
	client.initializeClient("http://localhost:8545")
	address, _ := client.deployContract(tester)
	fmt.Println(address)
	// Get filter for reading results (submissions to that address)
	filterID := client.ethNewFilter(address, "da82b96f")
	dhr = DHR{client, filterID, make([]Result, 0), address, challenges}	
	// Call test at the tester address with argument adder 
	// https://ethereum.stackexchange.com/questions/3780/how-can-i-create-a-listener-for-new-transaction-with-ethereum-rpc-calls
	http.HandleFunc("/", viewChallenges)
	http.HandleFunc("/edit/", edit)
	http.HandleFunc("/submit/", submit)
	http.HandleFunc("/view/", viewResults)
	http.ListenAndServe(":8080", nil)
}
