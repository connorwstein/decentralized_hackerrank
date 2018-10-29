package main

import (
	"bytes"
	"flag"
	"fmt"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
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
	log.Println(body)
	buf := bytes.NewBufferString(body)
	// Body will be a smart contract which we need to compile and then submit
	err := ioutil.WriteFile("adder.sol", buf.Bytes(), 0644)
	adder := compileContract("Adder", "adder.sol")
	err = dhr.client.submit(adder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// Update results
	dhr.results, err = dhr.client.getSubmissions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Update results
	http.Redirect(w, r, "/view/", http.StatusFound)
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

func viewResults(w http.ResponseWriter, r *http.Request) {
	// View open challenges, reputations etc.
	// We should be able to query the chain to get this info
	// Clicking on a challenge should take you to the edit page
	t, err := template.ParseFiles("template/view.tmpl")
	if err != nil {
		fmt.Println("Error when parsing template ", err)
	}
	t.Execute(w, dhr.results)
}

func viewChallenges(w http.ResponseWriter, r *http.Request) {
	// Walk through all the available challenges
	t, err := template.ParseFiles("template/index.tmpl")
	if err != nil {
		fmt.Println("Error when parsing template ", err)
	}
	t.Execute(w, dhr.challenges)
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
	// XXX: Easier to just use a config file here
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
		// Deploy the tester and print the address
		var backend EthBackend
		backend.initializeEthBackend(configuration.EthBackend)
		backend.setKeys(configuration.Address, configuration.PrivateKey)
		address, err := backend.deployTester()
		fmt.Println(address, err)
		os.Exit(0)
	} 
	challenges := make(map[string]Challenge)
	challenges["Adder"] = Challenge{"Adder", "easy", "function add(uint a, uint b) returns (uint)"}
	var backend EthBackend
	backend.initializeEthBackend(configuration.EthBackend)
	backend.setKeys(configuration.Address, configuration.PrivateKey)
	err = backend.loadTester(configuration.TesterAddress)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dhr = DHR{backend, make([]Result, 0), configuration.TesterAddress, challenges}
	http.HandleFunc("/", viewChallenges)
	http.HandleFunc("/edit/", edit)
	http.HandleFunc("/submit/", submit)
	http.HandleFunc("/view/", viewResults)
	http.ListenAndServe(fmt.Sprintf(":%v", configuration.Port), nil)
}
