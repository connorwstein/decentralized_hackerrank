package main

import (
	"testing"
	"time"
	"encoding/hex"
	"fmt"
	"log"
)

func TestAdder(t *testing.T) {
	// Test that a successful implementation of adder 
	// passes the tests
    sim, auth := createSimulator()	
	logs, contractAbi, testerAddress, session := deployTester(sim, auth)
	fmt.Println(testerAddress)
	// Start listening for logs
	result := make(chan bool)
	go func(result chan bool) {
		timeout := time.After(5 * time.Second)
		for {
			select {
// 			case err := <-sub.Err():
// 				log.Fatal(err)
// 				break
			case <- timeout:
				result <- false
				break
			case vLog := <-logs:
				var event bool 
				fmt.Println(contractAbi, vLog.Data)
				  err := contractAbi.Unpack(&event, "TestPass", vLog.Data)
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
// 	fmt.Println("adder bytecode err", err)
	tx, err := session.Test(adderCode)
	fmt.Println("Result:", tx.Hash(), err)
	sim.Commit() // like mining, should trigger the log
	// Wait until either the log comes in or we timeout
	testCaseResult := <-result
	if !testCaseResult {
		t.Fail()
	}
}
