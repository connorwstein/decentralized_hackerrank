package main

import (
	"bytes"
	"flag"
	"fmt"
	"sort"
	"encoding/json"
	"html/template"
	"io/ioutil"
// 	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"path"
)

type DHR struct {
	client      EthBackend // client
	results     []Result   // actual results
	testAddress string     // address of the tester contract
	challenges  map[string]Challenge
}

type Challenge struct {
	Name       string
	Difficulty string
	Interface  string
}

type Result struct {
	Submitter string // public key
	Challenge string
	Pass      bool
}

var dhr DHR

func submit(w http.ResponseWriter, r *http.Request) {
	// This POST takes in the user's contract code
	// Compile the smart contract code and
	// submit it to the blockchain for tests
	body := r.FormValue("body")
	challenge := path.Base(r.URL.Path)
	buf := bytes.NewBufferString(body)
	// Body will be a smart contract which we need to compile and then submit
	submissionSource := fmt.Sprintf("%v.sol", challenge)
	err := ioutil.WriteFile(submissionSource, buf.Bytes(), 0644)
	submissionByteCode := compileContract(challenge, submissionSource)
	err = dhr.client.submit(challenge, submissionByteCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func edit(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/edit.tmpl")
	if err != nil {
		fmt.Println("Error when parsing template ", err)
	}
	challenge := r.URL.Path[len("/edit/"):]
	fmt.Println("Editing challenge: ", challenge, dhr.challenges[challenge])
	t.Execute(w, dhr.challenges[challenge])
}

type LeaderItem struct {
	Leader string
	Score int
}

type M map[string]interface{}
// type M struct


func home(w http.ResponseWriter, r *http.Request) {
	// Walk through all the available challenges
	t, err := template.ParseFiles("template/index.tmpl")
	if err != nil {
		fmt.Println("Error when parsing template ", err)
	}
	dhr.results, err = dhr.client.getSubmissions()
	// Compute leaders
	// Show challenges
	// Show recent submissions
	competitors := make(map[string]int)
	for _, res := range dhr.results {
		if res.Pass {
			competitors[res.Submitter] += 10
		} else {
			competitors[res.Submitter] -= 1
		}
	}
	leaderboard := make([]LeaderItem, 0)
	for k, v := range competitors {
		leaderboard = append(leaderboard, LeaderItem{Leader: k, Score: v})
	}
	sort.Slice(leaderboard, func(i, j int) bool {
		return leaderboard[i].Score > leaderboard[j].Score
	})
	challenges := make([]Challenge, 0)
	for _, v := range dhr.challenges {
		challenges = append(challenges, v)
	}
// 	toShow := struct{leaderboard []LeaderItem
// 					 challenges []Challenge
// 					 results []Result}{leaderboard, challenges, dhr.results}
	toShow := M{"challenges":challenges, "leaderboard": leaderboard, 
				"submissions": dhr.results}
	t.Execute(w, toShow)
}

func clearCompiledContracts() {
	files, err := filepath.Glob("*.bin")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			panic(err)
		}
	}
}

// Given path to tester contract, compile it and
// and return a string of bytecode
func compileContract(name, path string) string {
	// Assumes solc is install locally
	// cant use the compiler in geth codebase because of
	// the optimization bug
	// solc --bin execute.sol -o .
	clearCompiledContracts()
	cmd := exec.Command("solc", "--bin", path, "-o", ".")
	fmt.Println(cmd)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("err running cmd", err)
		return ""
	}
	contents, err := ioutil.ReadFile(fmt.Sprintf("%v.bin", name))
	if err != nil {
		fmt.Println("err reading file", err)
		return ""
	}
	buf := bytes.NewBuffer(contents)
	return buf.String()
}

type Config struct {
	Address string 
	PrivateKey string
	Deploy bool
	TesterAddress string
	Port int
	EthBackend string
}

func main() {
	// Take in a index into the 10 accounts created by ganache as a command line argument
	// Could be replaced with a pub key if it was actually geth
	configFile := flag.String("config", "", "config file")

	flag.Parse()

	if *configFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	config, err := os.Open(*configFile)
	if err != nil {
		fmt.Println("Unable to read file")
		os.Exit(1)
	}
	configBytes, err := ioutil.ReadAll(config)
	var configuration Config	
	json.Unmarshal(configBytes, &configuration)
	fmt.Println(configuration)	
	// XXX: Need to validate configuration
	if configuration.Deploy {
		fmt.Println("Deploying new tester")
		// Deploy the tester and print the address
		var backend EthBackend
		backend.initializeEthBackend(configuration.EthBackend)
		backend.setKeys(configuration.Address, configuration.PrivateKey)
		address, err := backend.deployTester()
		fmt.Println(address, err)
		os.Exit(0)
	} 
	challenges := make(map[string]Challenge)
	challenges["Adder"] = Challenge{"Adder", "intro", "function add(int a, int b) returns (int)"}
	challenges["StringReverse"] = Challenge{"StringReverse", "easy", "function stringReverse(string input) public returns(string)"}
	var backend EthBackend
	backend.initializeEthBackend(configuration.EthBackend)
	backend.setKeys(configuration.Address, configuration.PrivateKey)
	err = backend.loadTester(configuration.TesterAddress)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dhr = DHR{backend, make([]Result, 0), configuration.TesterAddress, challenges}
	http.HandleFunc("/", home)
	http.HandleFunc("/edit/", edit)
	http.HandleFunc("/submit/", submit)
	http.ListenAndServe(fmt.Sprintf(":%v", configuration.Port), nil)
}
