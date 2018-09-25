package main

import (
	"fmt"
	"encoding/hex"

    "log"
    "math/big"
//     "time"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
    "github.com/ethereum/go-ethereum/core"
    "github.com/ethereum/go-ethereum/crypto"
)

// params: [{
//   "from": "0xb60e8dd61c5d32be8058bb8eb970870f07233155",
//   "to": "0xd46e8dd67c5d32be8058bb8eb970870f07244567",
//   "gas": "0x76c0", // 30400
//   "gasPrice": "0x9184e72a000", // 10000000000000
//   "value": "0x9184e72a", // 2441406250
//   "data": "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"
// }]

func main() {
    key, _ := crypto.GenerateKey()
    auth := bind.NewKeyedTransactor(key)
    alloc := make(core.GenesisAlloc)
    alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)} // Genesis account
    sim := backends.NewSimulatedBackend(alloc, 100000000) // 100000000 is gas limit
	// This factory contract uses the create opcode to deploy an aribitrary contract
	tester, _, contract, err := DeployTester(auth, sim)

	// Goal here is to pass in the byte code of adder as an argument, and have that contract deployed
	// Adder bytecode
	// 608060405234801561001057600080fd5b5060c58061001f6000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063771602f7146044575b600080fd5b348015604f57600080fd5b5060766004803603810190808035906020019092919080359060200190929190505050608c565b6040518082815260200191505060405180910390f35b60008183019050929150505600a165627a7a72305820e80809f762456af961d7357d0db84ff34c4cf49807222c472768ee39b60ac7a80029
	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
	}
	fmt.Println("Contract is at address:", tester)
	instance, err := NewTester(tester, sim)
	if err != nil {
	  log.Fatal(err)
	}
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
	adderCode, err := hex.DecodeString("608060405234801561001057600080fd5b5060c58061001f6000396000f300608060405260043610603f576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063771602f7146044575b600080fd5b348015604f57600080fd5b5060766004803603810190808035906020019092919080359060200190929190505050608c565b6040518082815260200191505060405180910390f35b60008183019050929150505600a165627a7a72305820e80809f762456af961d7357d0db84ff34c4cf49807222c472768ee39b60ac7a80029")
	fmt.Println("adder bytecode err", err)
	tx, err := session.Test(adderCode)
	fmt.Println("Result:", tx.Hash(), err)
	sim.Commit() // like mining
	val, err := instance.Result(nil)
	fmt.Println("Add result", val, err)
}

