package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"testing"
	"time"
)

// Used for low level testing
func TestAdder(t *testing.T) {
	// Test that a successful implementation of adder
	// passes the tests
	var s SubmissionRunner
	s.create()
	s.deployRunner()
	// Start listening for logs
	result := make(chan bool)
	go func(result chan bool) {
		timeout := time.After(5 * time.Second)
		for {
			select {
			// 			case err := <-sub.Err():
			// 				log.Fatal(err)
			// 				break
			case <-timeout:
				result <- false
				break
			case vLog := <-s.logs:
				var event bool
				fmt.Println(s.contractAbi, vLog.Data)
				err := s.contractAbi.Unpack(&event, "TestPass", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Got a log", event) // pointer to event log
				result <- event
				break
			}
		}
	}(result)
	// If pass in correct adder testcode and never
	// get the positive log then we fail the test
	adderCode, _ := hex.DecodeString("608060405234801561001057600080fd5b5060c58061001f6000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063771602f7146044575b600080fd5b348015604f57600080fd5b5060766004803603810190808035906020019092919080359060200190929190505050608c565b6040518082815260200191505060405180910390f35b60008183019050929150505600a165627a7a72305820e80809f762456af961d7357d0db84ff34c4cf49807222c472768ee39b60ac7a80029")
	// Note optimized code does not work, not sure why
// 	optimizedCode := "6080604052348015600f57600080fd5b50609c8061001e6000396000f300608060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663c6888fa181146043575b600080fd5b348015604e57600080fd5b506058600435606a565b60408051918252519081900360200190f35b600702905600a165627a7a72305820e1d273f96d734b39265bc6859c83427dbaedd9f692e888b87f79c6ee5b318c940029"
	adderCode, _ := hex.DecodeString(optimizedCode)
	tx, err := s.session.Test(adderCode)
	fmt.Println("Result:", tx.Hash(), err)
	s.sim.Commit() // like mining, should trigger the log
	// Wait until either the log comes in or we timeout
	testCaseResult := <-result
	if !testCaseResult {
		t.Fail()
	}
}
