package main

//
// import (
// 	"encoding/hex"
// 	"fmt"
// 	"log"
// 	"testing"
// 	"time"
// )
//
// var runner SubmissionRunner
//
// type SubmissionRunner struct {
// 	sim *backends.SimulatedBackend
// 	auth *bind.TransactOpts
// 	logs chan types.Log
// 	contractAbi abi.ABI
// 	address common.Address
// 	session *TesterSession
// 	results []Result
// }
//
// func (s *SubmissionRunner) create() {
// 	key, _ := crypto.GenerateKey()
// 	auth := bind.NewKeyedTransactor(key)
// 	alloc := make(core.GenesisAlloc)
// 	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)} // Genesis account
// 	sim := backends.NewSimulatedBackend(alloc, 100000000)                  // 100000000 is gas limit
// 	s.sim = sim
// 	s.auth = auth
// 	s.results = make([]Result, 0)
// }
//
// func (s *SubmissionRunner) deployRunner() {
// 	// Returns a channel to listen for test case results
// 	// and the address at which the tester is deployed
// 	tester, _, contract, err := DeployTester(s.auth, s.sim)
// 	if err != nil {
// 		log.Fatalf("Could not deploy contract: %v", err)
// 	}
// 	fmt.Println("Contract is at address:", tester)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	s.sim.Commit() // like mining
// 	// Subscribe to events logs from tester
// 	query := ethereum.FilterQuery{
// 		Addresses: []common.Address{tester},
// 	}
// 	contractAbi, err := abi.JSON(strings.NewReader(string(TesterABI)))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	logs := make(chan types.Log)
// 	_, err = s.sim.SubscribeFilterLogs(context.Background(), query, logs)
// 	session := &TesterSession{
// 		Contract: contract,
// 		CallOpts: bind.CallOpts{
// 			Pending: true,
// 		},
// 		TransactOpts: bind.TransactOpts{
// 			From:     s.auth.From,
// 			Signer:   s.auth.Signer,
// 			GasLimit: 3141592,
// 		},
// 	}
// 	s.logs = logs
// 	s.contractAbi = contractAbi
// 	s.address = tester
// 	s.session = session
// }
//
//
//
// // Used for low level testing
// func TestAdder(t *testing.T) {
// 	// Test that a successful implementation of adder
// 	// passes the tests
// 	var s SubmissionRunner
// 	s.create()
// 	s.deployRunner()
// 	// Start listening for logs
// 	result := make(chan bool)
// 	go func(result chan bool) {
// 		timeout := time.After(5 * time.Second)
// 		for {
// 			select {
// 			// 			case err := <-sub.Err():
// 			// 				log.Fatal(err)
// 			// 				break
// 			case <-timeout:
// 				result <- false
// 				break
// 			case vLog := <-s.logs:
// 				var event bool
// 				fmt.Println(s.contractAbi, vLog.Data)
// 				err := s.contractAbi.Unpack(&event, "TestPass", vLog.Data)
// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 				fmt.Println("Got a log", event) // pointer to event log
// 				result <- event
// 				break
// 			}
// 		}
// 	}(result)
// 	// If pass in correct adder testcode and never
// 	// get the positive log then we fail the test
// 	adderCode, _ := hex.DecodeString("608060405234801561001057600080fd5b5060c58061001f6000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063771602f7146044575b600080fd5b348015604f57600080fd5b5060766004803603810190808035906020019092919080359060200190929190505050608c565b6040518082815260200191505060405180910390f35b60008183019050929150505600a165627a7a72305820e80809f762456af961d7357d0db84ff34c4cf49807222c472768ee39b60ac7a80029")
// 	// Note optimized code does not work, not sure why
// // 	optimizedCode := "6080604052348015600f57600080fd5b50609c8061001e6000396000f300608060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663c6888fa181146043575b600080fd5b348015604e57600080fd5b506058600435606a565b60408051918252519081900360200190f35b600702905600a165627a7a72305820e1d273f96d734b39265bc6859c83427dbaedd9f692e888b87f79c6ee5b318c940029"
// // 	adderCode, _ := hex.DecodeString(optimizedCode)
// 	tx, err := s.session.Test(adderCode)
// 	fmt.Println("Result:", tx.Hash(), err)
// 	s.sim.Commit() // like mining, should trigger the log
// 	// Wait until either the log comes in or we timeout
// 	testCaseResult := <-result
// 	if !testCaseResult {
// 		t.Fail()
// 	}
// }
