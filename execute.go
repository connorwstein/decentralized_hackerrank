// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// AdderABI is the input ABI used to generate the binding from.
const AdderABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AdderBin is the compiled bytecode used for deploying new contracts.
const AdderBin = `0x6080604052348015600f57600080fd5b5060a18061001e6000396000f300608060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663771602f781146043575b600080fd5b348015604e57600080fd5b50605b600435602435606d565b60408051918252519081900360200190f35b6000929150505600a165627a7a72305820fa31dfa6ebdd2b3964dc10209349ac76b16eb156c553b258e8c317c00461238c0029`

// DeployAdder deploys a new Ethereum contract, binding an instance of Adder to it.
func DeployAdder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Adder, error) {
	parsed, err := abi.JSON(strings.NewReader(AdderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AdderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Adder{AdderCaller: AdderCaller{contract: contract}, AdderTransactor: AdderTransactor{contract: contract}, AdderFilterer: AdderFilterer{contract: contract}}, nil
}

// Adder is an auto generated Go binding around an Ethereum contract.
type Adder struct {
	AdderCaller     // Read-only binding to the contract
	AdderTransactor // Write-only binding to the contract
	AdderFilterer   // Log filterer for contract events
}

// AdderCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdderSession struct {
	Contract     *Adder            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdderCallerSession struct {
	Contract *AdderCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AdderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdderTransactorSession struct {
	Contract     *AdderTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdderRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdderRaw struct {
	Contract *Adder // Generic contract binding to access the raw methods on
}

// AdderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdderCallerRaw struct {
	Contract *AdderCaller // Generic read-only contract binding to access the raw methods on
}

// AdderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdderTransactorRaw struct {
	Contract *AdderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdder creates a new instance of Adder, bound to a specific deployed contract.
func NewAdder(address common.Address, backend bind.ContractBackend) (*Adder, error) {
	contract, err := bindAdder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Adder{AdderCaller: AdderCaller{contract: contract}, AdderTransactor: AdderTransactor{contract: contract}, AdderFilterer: AdderFilterer{contract: contract}}, nil
}

// NewAdderCaller creates a new read-only instance of Adder, bound to a specific deployed contract.
func NewAdderCaller(address common.Address, caller bind.ContractCaller) (*AdderCaller, error) {
	contract, err := bindAdder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdderCaller{contract: contract}, nil
}

// NewAdderTransactor creates a new write-only instance of Adder, bound to a specific deployed contract.
func NewAdderTransactor(address common.Address, transactor bind.ContractTransactor) (*AdderTransactor, error) {
	contract, err := bindAdder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdderTransactor{contract: contract}, nil
}

// NewAdderFilterer creates a new log filterer instance of Adder, bound to a specific deployed contract.
func NewAdderFilterer(address common.Address, filterer bind.ContractFilterer) (*AdderFilterer, error) {
	contract, err := bindAdder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdderFilterer{contract: contract}, nil
}

// bindAdder binds a generic wrapper to an already deployed contract.
func bindAdder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Adder *AdderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Adder.Contract.AdderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Adder *AdderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Adder.Contract.AdderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Adder *AdderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Adder.Contract.AdderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Adder *AdderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Adder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Adder *AdderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Adder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Adder *AdderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Adder.Contract.contract.Transact(opts, method, params...)
}

// Add is a paid mutator transaction binding the contract method 0x771602f7.
//
// Solidity: function add(a uint256, b uint256) returns(uint256)
func (_Adder *AdderTransactor) Add(opts *bind.TransactOpts, a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _Adder.contract.Transact(opts, "add", a, b)
}

// Add is a paid mutator transaction binding the contract method 0x771602f7.
//
// Solidity: function add(a uint256, b uint256) returns(uint256)
func (_Adder *AdderSession) Add(a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _Adder.Contract.Add(&_Adder.TransactOpts, a, b)
}

// Add is a paid mutator transaction binding the contract method 0x771602f7.
//
// Solidity: function add(a uint256, b uint256) returns(uint256)
func (_Adder *AdderTransactorSession) Add(a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _Adder.Contract.Add(&_Adder.TransactOpts, a, b)
}

// TesterABI is the input ABI used to generate the binding from.
const TesterABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"code\",\"type\":\"bytes\"}],\"name\":\"test\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"result\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TesterBin is the compiled bytecode used for deploying new contracts.
const TesterBin = `0x6080604052600360005534801561001557600080fd5b506101ff806100256000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632f570a23811461005057806365372147146100bb575b600080fd5b34801561005c57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100a99436949293602493928401919081908401838280828437509497506100d09650505050505050565b60408051918252519081900360200190f35b3480156100c757600080fd5b506100a96101bc565b60008060006100de846101c2565b915073ffffffffffffffffffffffffffffffffffffffff8216151561010257600080fd5b819050600a6000819055508073ffffffffffffffffffffffffffffffffffffffff1663771602f7600a806040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083815260200182815260200192505050602060405180830381600087803b15801561018557600080fd5b505af1158015610199573d6000803e3d6000fd5b505050506040513d60208110156101af57600080fd5b5051600055509092915050565b60005481565b60008151602083016000f0929150505600a165627a7a72305820a953bbf1a8ed7003040411c669caad885ca69abf5864fe98a5a9f1fc8682b25c0029`

// DeployTester deploys a new Ethereum contract, binding an instance of Tester to it.
func DeployTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tester, error) {
	parsed, err := abi.JSON(strings.NewReader(TesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tester{TesterCaller: TesterCaller{contract: contract}, TesterTransactor: TesterTransactor{contract: contract}, TesterFilterer: TesterFilterer{contract: contract}}, nil
}

// Tester is an auto generated Go binding around an Ethereum contract.
type Tester struct {
	TesterCaller     // Read-only binding to the contract
	TesterTransactor // Write-only binding to the contract
	TesterFilterer   // Log filterer for contract events
}

// TesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TesterSession struct {
	Contract     *Tester           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TesterCallerSession struct {
	Contract *TesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TesterTransactorSession struct {
	Contract     *TesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TesterRaw struct {
	Contract *Tester // Generic contract binding to access the raw methods on
}

// TesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TesterCallerRaw struct {
	Contract *TesterCaller // Generic read-only contract binding to access the raw methods on
}

// TesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TesterTransactorRaw struct {
	Contract *TesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTester creates a new instance of Tester, bound to a specific deployed contract.
func NewTester(address common.Address, backend bind.ContractBackend) (*Tester, error) {
	contract, err := bindTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tester{TesterCaller: TesterCaller{contract: contract}, TesterTransactor: TesterTransactor{contract: contract}, TesterFilterer: TesterFilterer{contract: contract}}, nil
}

// NewTesterCaller creates a new read-only instance of Tester, bound to a specific deployed contract.
func NewTesterCaller(address common.Address, caller bind.ContractCaller) (*TesterCaller, error) {
	contract, err := bindTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TesterCaller{contract: contract}, nil
}

// NewTesterTransactor creates a new write-only instance of Tester, bound to a specific deployed contract.
func NewTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*TesterTransactor, error) {
	contract, err := bindTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TesterTransactor{contract: contract}, nil
}

// NewTesterFilterer creates a new log filterer instance of Tester, bound to a specific deployed contract.
func NewTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*TesterFilterer, error) {
	contract, err := bindTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TesterFilterer{contract: contract}, nil
}

// bindTester binds a generic wrapper to an already deployed contract.
func bindTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tester *TesterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tester.Contract.TesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tester *TesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tester.Contract.TesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tester *TesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tester.Contract.TesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tester *TesterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tester *TesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tester *TesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tester.Contract.contract.Transact(opts, method, params...)
}

// Result is a free data retrieval call binding the contract method 0x65372147.
//
// Solidity: function result() constant returns(uint256)
func (_Tester *TesterCaller) Result(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Tester.contract.Call(opts, out, "result")
	return *ret0, err
}

// Result is a free data retrieval call binding the contract method 0x65372147.
//
// Solidity: function result() constant returns(uint256)
func (_Tester *TesterSession) Result() (*big.Int, error) {
	return _Tester.Contract.Result(&_Tester.CallOpts)
}

// Result is a free data retrieval call binding the contract method 0x65372147.
//
// Solidity: function result() constant returns(uint256)
func (_Tester *TesterCallerSession) Result() (*big.Int, error) {
	return _Tester.Contract.Result(&_Tester.CallOpts)
}

// Test is a paid mutator transaction binding the contract method 0x2f570a23.
//
// Solidity: function test(code bytes) returns(uint256)
func (_Tester *TesterTransactor) Test(opts *bind.TransactOpts, code []byte) (*types.Transaction, error) {
	return _Tester.contract.Transact(opts, "test", code)
}

// Test is a paid mutator transaction binding the contract method 0x2f570a23.
//
// Solidity: function test(code bytes) returns(uint256)
func (_Tester *TesterSession) Test(code []byte) (*types.Transaction, error) {
	return _Tester.Contract.Test(&_Tester.TransactOpts, code)
}

// Test is a paid mutator transaction binding the contract method 0x2f570a23.
//
// Solidity: function test(code bytes) returns(uint256)
func (_Tester *TesterTransactorSession) Test(code []byte) (*types.Transaction, error) {
	return _Tester.Contract.Test(&_Tester.TransactOpts, code)
}
