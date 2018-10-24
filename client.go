package main

import (
	"net/http"
// 	"reflect"
	"fmt"
	"encoding/json"
	"bytes"
	"io/ioutil"
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
	params["from"] = n.Accounts[0] // Just pick the first account
	params["gas"] = "0xfffff" // fffff seems to be sufficient for contract deployment (default of 90k is not)
	params["data"] = contractCode 
	tx, err := n.makeRequest("eth_sendTransaction", []interface{}{params})
	// Result is a single string which is the contract address 
	receipt, _ := n.ethGetTransactionReceipt(tx.(string))
	return receipt["contractAddress"].(string), err
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
	fmt.Println(resp, err)
	return resp.(string)
}

func (n Client) ethGetFilterChanges(filterID string) {
	// eth_getFilterChanges
	resp, err := n.makeRequest("eth_getFilterChanges", []interface{}{[]string{filterID}})
	fmt.Println(resp, err)
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
	fmt.Println(r)
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
