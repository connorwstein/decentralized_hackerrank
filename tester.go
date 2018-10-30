// This module should be agnostic to what the actual backend
// chain is whether that is geth itself or ganache-cli
// and maybe even the simulated backend as well
package main

import (
	execute "./execute"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

// Represents the ethereum node we are talking to
// Accounts will be the accounts that we control
type EthBackend struct {
	client  *ethclient.Client
	address string
	// Could have multiple keys	for now just support one
	privateKey string
	publicKey  string          // address is the keccak 256 hash of this
	instance   *execute.Tester // instance of our tester contract
}

// Only thing we need to know is the address to reach the eth backend
// Connect to it
func (n *EthBackend) initializeEthBackend(address string) {
	n.address = address
	client, err := ethclient.Dial(address)
	if err != nil {
		log.Fatal(err)
	}
	n.client = client
}

// Set the keys to be used
func (n *EthBackend) setKeys(public, private string) {
	n.privateKey = private
	n.publicKey = public
}

// Returns the address of the contract
// Use the client and clients private key to do this deploy
// XXX: make sense to make this more generic so it can deploy any contract
func (n *EthBackend) deployTester() (string, error) {
	privateKey, err := crypto.HexToECDSA(n.privateKey)
	if err != nil {
		return "", err
	}
	nonce, err := n.client.PendingNonceAt(context.Background(), common.HexToAddress(n.publicKey))
	if err != nil {
		return "", err
	}

	gasPrice, err := n.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	fmt.Println(gasPrice)
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	// XXX: way to estimate this?
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	// 2nd return value is type.Transaction
	address, _, instance, err := execute.DeployTester(auth, n.client)
	// Instance is *Tester which we can use
	n.instance = instance
	if err != nil {
		return "", nil
	}
	return address.Hex(), nil
}

func (n *EthBackend) loadTester(addressHex string) error {
	// Load a pre-deployed Tester contract
	address := common.HexToAddress(addressHex)
	instance, err := execute.NewTester(address, n.client)
	if err != nil {
		return err
	}
	n.instance = instance
	return nil
}

func (n EthBackend) submit(challenge, codeSubmission string) error {
	if n.instance == nil {
		return fmt.Errorf("Tester contract not deployed")
	}
	privateKey, err := crypto.HexToECDSA(n.privateKey)
	signerFn := func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		// 			signer := types.NewEIP155Signer(nil)
		data := signer.Hash(tx).Bytes()
		signature, err := crypto.Sign(data, privateKey)
		if err != nil {
			return nil, err
		}
		return tx.WithSignature(signer, signature)
	}
	// If you omit the gas price and gas limit
	// it will try and make an estimate, but sometimes
	// the estimate isn't enough 
	transOpts := bind.TransactOpts{
					From: common.HexToAddress(n.publicKey), 
					Signer: signerFn, 
					GasPrice: big.NewInt(20000000000), // price in wei
					GasLimit: 5000000} 
	decoded, err := hex.DecodeString(codeSubmission)
	if err != nil {
		return err
	}
	switch challenge {
		case "Adder":
			tx, err := n.instance.TestAdder(&transOpts, decoded) 
			if err != nil {
				return err
			}
			fmt.Println(tx, err)
		case "StringReverse": 
			tx, err := n.instance.TestStringReverse(&transOpts, decoded) 
			if err != nil {
				return err
			}
			fmt.Println(tx, err)
	}
	return nil
}

func (n EthBackend) getSubmissions() ([]Result, error) {
	count, err := n.instance.SubmissionCount(nil) // submission count working, now we need to find out how to
	if err != nil {
		return nil, err
	}
	var i int64
	results := make([]Result, 0)
	for i = 0; i < count.Int64(); i++ {
		s, err := n.instance.Submissions(nil, big.NewInt(i)) // submission count working, now we need to find out how to
		if err != nil {
			return nil, err
		}
		fmt.Println("Submission:", s, err)
		results = append(results, Result{Submitter: s.Submitter.Hex(), Pass: s.Pass})
	}
	return results, nil
}
