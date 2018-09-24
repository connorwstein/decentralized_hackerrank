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
const AdderBin = `0x6080604052348015600f57600080fd5b50609d8061001e6000396000f300608060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663771602f781146043575b600080fd5b348015604e57600080fd5b50605b600435602435606d565b60408051918252519081900360200190f35b01905600a165627a7a723058206d993819755025d0891a9219f94a2838ea17ee03802979e005097e7d50e8404b0029`

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

// FactoryABI is the input ABI used to generate the binding from.
const FactoryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"code\",\"type\":\"bytes\"}],\"name\":\"create\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FactoryBin is the compiled bytecode used for deploying new contracts.
const FactoryBin = `0x608060405234801561001057600080fd5b5060ff8061001f6000396000f300608060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663cf5ba53f81146043575b600080fd5b348015604e57600080fd5b506040805160206004803580820135601f8101849004840285018401909552848452609994369492936024939284019190819084018382808284375094975060c29650505050505050565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60008151602083016000f0929150505600a165627a7a72305820db76f4c173bf8970aa5992fb81ee7ba619d67482e8b2e7615a86dced4a7179ba0029`

// DeployFactory deploys a new Ethereum contract, binding an instance of Factory to it.
func DeployFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Factory, error) {
	parsed, err := abi.JSON(strings.NewReader(FactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// Factory is an auto generated Go binding around an Ethereum contract.
type Factory struct {
	FactoryCaller     // Read-only binding to the contract
	FactoryTransactor // Write-only binding to the contract
	FactoryFilterer   // Log filterer for contract events
}

// FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactorySession struct {
	Contract     *Factory          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactoryCallerSession struct {
	Contract *FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactoryTransactorSession struct {
	Contract     *FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactoryRaw struct {
	Contract *Factory // Generic contract binding to access the raw methods on
}

// FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactoryCallerRaw struct {
	Contract *FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactoryTransactorRaw struct {
	Contract *FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactory creates a new instance of Factory, bound to a specific deployed contract.
func NewFactory(address common.Address, backend bind.ContractBackend) (*Factory, error) {
	contract, err := bindFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// NewFactoryCaller creates a new read-only instance of Factory, bound to a specific deployed contract.
func NewFactoryCaller(address common.Address, caller bind.ContractCaller) (*FactoryCaller, error) {
	contract, err := bindFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryCaller{contract: contract}, nil
}

// NewFactoryTransactor creates a new write-only instance of Factory, bound to a specific deployed contract.
func NewFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FactoryTransactor, error) {
	contract, err := bindFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryTransactor{contract: contract}, nil
}

// NewFactoryFilterer creates a new log filterer instance of Factory, bound to a specific deployed contract.
func NewFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FactoryFilterer, error) {
	contract, err := bindFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryFilterer{contract: contract}, nil
}

// bindFactory binds a generic wrapper to an already deployed contract.
func bindFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transact(opts, method, params...)
}

// Create is a paid mutator transaction binding the contract method 0xcf5ba53f.
//
// Solidity: function create(code bytes) returns(addr address)
func (_Factory *FactoryTransactor) Create(opts *bind.TransactOpts, code []byte) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "create", code)
}

// Create is a paid mutator transaction binding the contract method 0xcf5ba53f.
//
// Solidity: function create(code bytes) returns(addr address)
func (_Factory *FactorySession) Create(code []byte) (*types.Transaction, error) {
	return _Factory.Contract.Create(&_Factory.TransactOpts, code)
}

// Create is a paid mutator transaction binding the contract method 0xcf5ba53f.
//
// Solidity: function create(code bytes) returns(addr address)
func (_Factory *FactoryTransactorSession) Create(code []byte) (*types.Transaction, error) {
	return _Factory.Contract.Create(&_Factory.TransactOpts, code)
}

// TesterABI is the input ABI used to generate the binding from.
const TesterABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"test\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"factory\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TesterBin is the compiled bytecode used for deploying new contracts.
const TesterBin = `0x608060405234801561001057600080fd5b5060405160208061031283398101604081905290517fcf5ba53f00000000000000000000000000000000000000000000000000000000825260206004830190815260c5602484018190529192600160a060020a0384169263cf5ba53f9291829160449091019061024d823960e001915050602060405180830381600087803b15801561009b57600080fd5b505af11580156100af573d6000803e3d6000fd5b505050506040513d60208110156100c557600080fd5b505160008054600160a060020a031916600160a060020a0392831617908190551615156100f157600080fd5b5061014c806101016000396000f3006080604052600436106100405763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663eb8ac9218114610045575b600080fd5b34801561005157600080fd5b50610060600435602435610072565b60408051918252519081900360200190f35b60008054604080517f771602f70000000000000000000000000000000000000000000000000000000081526004810186905260248101859052905173ffffffffffffffffffffffffffffffffffffffff9092169163771602f79160448082019260209290919082900301818787803b1580156100ed57600080fd5b505af1158015610101573d6000803e3d6000fd5b505050506040513d602081101561011757600080fd5b505193925050505600a165627a7a72305820360cbd057678a2ca84de0a60aed8cb693c172223f3d2f61e45b088d3b2dc2b260029606060405234610000575b60ad806100186000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063771602f714603c575b6000565b34600057605d60048080359060200190919080359060200190919050506073565b6040518082815260200191505060405180910390f35b600081830190505b929150505600a165627a7a723058205d7bec00c6d410f7ea2a3b03112b597bb3ef544439889ecc1294a77b85eab15e0029`

// DeployTester deploys a new Ethereum contract, binding an instance of Tester to it.
func DeployTester(auth *bind.TransactOpts, backend bind.ContractBackend, factory common.Address) (common.Address, *types.Transaction, *Tester, error) {
	parsed, err := abi.JSON(strings.NewReader(TesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TesterBin), backend, factory)
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

// Test is a free data retrieval call binding the contract method 0xeb8ac921.
//
// Solidity: function test(x uint256, y uint256) constant returns(uint256)
func (_Tester *TesterCaller) Test(opts *bind.CallOpts, x *big.Int, y *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Tester.contract.Call(opts, out, "test", x, y)
	return *ret0, err
}

// Test is a free data retrieval call binding the contract method 0xeb8ac921.
//
// Solidity: function test(x uint256, y uint256) constant returns(uint256)
func (_Tester *TesterSession) Test(x *big.Int, y *big.Int) (*big.Int, error) {
	return _Tester.Contract.Test(&_Tester.CallOpts, x, y)
}

// Test is a free data retrieval call binding the contract method 0xeb8ac921.
//
// Solidity: function test(x uint256, y uint256) constant returns(uint256)
func (_Tester *TesterCallerSession) Test(x *big.Int, y *big.Int) (*big.Int, error) {
	return _Tester.Contract.Test(&_Tester.CallOpts, x, y)
}
