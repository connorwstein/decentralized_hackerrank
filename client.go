package main

import (
	"net/http"
	"reflect"
// 	"encoding/hex"
	"math/big"
	"fmt"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)


type EthRequest struct {
	Jsonrpc string  `json:"jsonrpc"` 
	Method string   `json:"method"` 
	Params []interface{} `json:"params"` 
	Id int	  		`json:"id"` 
}

type EthResponseList struct {
	Id int `json:"id"` 
	Jsonrpc string `json:"jsonrpc"`
	Result []string `json:"result"` 
}

type EthResponse struct {
	Id int `json:"id"` 
	Jsonrpc string `json:"jsonrpc"`
	Result string `json:"result"` 
}


// Represents the ethereum node we are talking to
// Accounts will be the accounts that we control
type Client struct {
	Accounts []string	
	Address string // http://host:port
	AccountIndex int // which account is ours
}

func (n *Client) initializeClient(address string) (error) {
	// Query the ganache RPC simulator to get all of the accounts and store them
	n.Address = address
	accounts, err := n.ethAccounts()
	if err != nil {
		return err
	}
	n.Accounts = accounts
	return nil
}

// Returns the address of the contract
func (n Client) deployContract(contractCode string) (string, error) {
	params := make(map[string]string)
	params["from"] = n.Accounts[n.AccountIndex] // Just pick the first account
	// TODO: look into using the suggested gas price
	params["gas"] = "0xfffff" // fffff seems to be sufficient for contract deployment (default of 90k is not)
	params["data"] = contractCode 
	tx, err := n.makeRequest("eth_sendTransaction", []interface{}{params})
	// Result is a single string which is the contract address 
	receipt, _ := n.ethGetTransactionReceipt(tx.(string))
	return receipt["contractAddress"].(string), err
}

func intToHex(i int) string {
	return fmt.Sprintf("0x%x", i)
}


func (n Client) getSubmissions(testerAddress string) []Result {
// curl -X POST --data '{"jsonrpc":"2.0", "method": "eth_getStorageAt", "params": ["0x295a70b2de5e3953354a6a8344e616ed314d7251", "0x0", "latest"], "id": 1}' localhost:8545
// 	resp, err := n.makeRequest("eth_getStorageAt", []interface{}{testerAddress, intToHex(0) , "latest"})
	address := common.HexToAddress(testerAddress)
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		fmt.Println("error: ", err)
	}
	instance, err := NewTester(address, client)
	count, err := instance.SubmissionCount(nil) // submission count working, now we need to find out how to 
	var i int64
	results := make([]Result, 0)
	for i = 0; i < count.Int64(); i++ {
		s, err := instance.Submissions(nil, big.NewInt(i)) // submission count working, now we need to find out how to 
		fmt.Println("Submission:", s, err)
		results = append(results, Result{Submitter: s.Submitter.Hex(), Pass: s.Pass})
	}
	return results
}

func (n Client) ethSendTransaction(from, to, data string) (interface{}, error) {
	params := make(map[string]string)
	params["from"] = from
	params["to"] = to
	params["gas"] = "0xfffff" // fffff seems to be sufficient for most transactions (default of 90k is not),
	params["data"] = data
	resp, err := n.makeRequest("eth_sendTransaction", []interface{}{params})
	fmt.Println(resp, err)
	// TODO: handle generic transactions
	return resp, err
}

func (n Client) ethGetTransactionReceipt(tx string) (map[string]interface{}, error) {
	resp, err := n.makeRequest("eth_getTransactionReceipt", []interface{}{tx})
	// Return contract address
	fmt.Println(resp, err)
	return resp.(map[string]interface{}), err
}

// Returns the filter id
func (n Client) ethNewFilter(address string, topic string) string {
// Hash of event TestPass(bool res) da82b96fcb1b4eeeec40c3fa1a155f1490b43b3b2a3eed181a3bcc7f61a0dccc
// curl -X POST --data '{"jsonrpc":"2.0","method":"eth_newFilter","params":[{"address": "0xaddress", topics:["KECCAK_HASH", "0x6"], "fromBlock":"latest","toBlock":"latest"}],"id":2}' localhost:8545 {"id":2,"jsonrpc":"2.0","result":"0x0"}
	params := make(map[string]string)
	params["address"] = address
	params["topic"] = topic
	resp, err := n.makeRequest("eth_newFilter", []interface{}{params})
	fmt.Println(err)
	return resp.(string)
}

func (n Client) ethGetFilterChanges(filterID string) string {
	// eth_getFilterChanges
	resp, err := n.makeRequest("eth_getFilterChanges", []interface{}{[]string{filterID}})
	fmt.Println(resp, err)
	// Consume the first log
	resp = resp.([]interface{})[0]
	change := resp.(map[string]interface{})
	fmt.Println(reflect.TypeOf(change["data"].(string)))
	// Extract just the data from this change
	return change["data"].(string)
}

// Returns an interface which is a string or []string depending on the call
func (n Client) makeRequest(api string, params []interface{}) (interface{}, error) {
	t := EthRequest{"2.0", api, params, 1}
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	rawResp, err := http.Post(n.Address, "application/json", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	defer rawResp.Body.Close()
	body, err := ioutil.ReadAll(rawResp.Body)
	if err != nil {
		return nil, err
	}
	var r interface{}
	respErr := json.Unmarshal(body, &r)
	if respErr != nil {
		return nil, respErr
	}
	fmt.Println("Result from unmarshaling: ", r)
	// TODO: handle error if result not present
	return r.(map[string]interface{})["result"], nil
}

func (n Client) ethAccounts() ([]string, error) {
	resp, err := n.makeRequest("eth_accounts", []interface{}{})
	results := make([]string, 0)
	for i := range resp.([]interface{}) {
		results = append(results, resp.([]interface{})[i].(string))
	}
	return results, err
}
