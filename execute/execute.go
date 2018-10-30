// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package execute

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// AdderABI is the input ABI used to generate the binding from.
const AdderABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"int256\"},{\"name\":\"b\",\"type\":\"int256\"}],\"name\":\"add\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AdderBin is the compiled bytecode used for deploying new contracts.
const AdderBin = `0x`

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

// Add is a paid mutator transaction binding the contract method 0xa5f3c23b.
//
// Solidity: function add(a int256, b int256) returns(int256)
func (_Adder *AdderTransactor) Add(opts *bind.TransactOpts, a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _Adder.contract.Transact(opts, "add", a, b)
}

// Add is a paid mutator transaction binding the contract method 0xa5f3c23b.
//
// Solidity: function add(a int256, b int256) returns(int256)
func (_Adder *AdderSession) Add(a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _Adder.Contract.Add(&_Adder.TransactOpts, a, b)
}

// Add is a paid mutator transaction binding the contract method 0xa5f3c23b.
//
// Solidity: function add(a int256, b int256) returns(int256)
func (_Adder *AdderTransactorSession) Add(a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _Adder.Contract.Add(&_Adder.TransactOpts, a, b)
}

// StringReverseABI is the input ABI used to generate the binding from.
const StringReverseABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"input\",\"type\":\"string\"}],\"name\":\"stringReverse\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StringReverseBin is the compiled bytecode used for deploying new contracts.
const StringReverseBin = `0x`

// DeployStringReverse deploys a new Ethereum contract, binding an instance of StringReverse to it.
func DeployStringReverse(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StringReverse, error) {
	parsed, err := abi.JSON(strings.NewReader(StringReverseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StringReverseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StringReverse{StringReverseCaller: StringReverseCaller{contract: contract}, StringReverseTransactor: StringReverseTransactor{contract: contract}, StringReverseFilterer: StringReverseFilterer{contract: contract}}, nil
}

// StringReverse is an auto generated Go binding around an Ethereum contract.
type StringReverse struct {
	StringReverseCaller     // Read-only binding to the contract
	StringReverseTransactor // Write-only binding to the contract
	StringReverseFilterer   // Log filterer for contract events
}

// StringReverseCaller is an auto generated read-only Go binding around an Ethereum contract.
type StringReverseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringReverseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StringReverseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringReverseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StringReverseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringReverseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StringReverseSession struct {
	Contract     *StringReverse    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StringReverseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StringReverseCallerSession struct {
	Contract *StringReverseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StringReverseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StringReverseTransactorSession struct {
	Contract     *StringReverseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StringReverseRaw is an auto generated low-level Go binding around an Ethereum contract.
type StringReverseRaw struct {
	Contract *StringReverse // Generic contract binding to access the raw methods on
}

// StringReverseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StringReverseCallerRaw struct {
	Contract *StringReverseCaller // Generic read-only contract binding to access the raw methods on
}

// StringReverseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StringReverseTransactorRaw struct {
	Contract *StringReverseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStringReverse creates a new instance of StringReverse, bound to a specific deployed contract.
func NewStringReverse(address common.Address, backend bind.ContractBackend) (*StringReverse, error) {
	contract, err := bindStringReverse(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StringReverse{StringReverseCaller: StringReverseCaller{contract: contract}, StringReverseTransactor: StringReverseTransactor{contract: contract}, StringReverseFilterer: StringReverseFilterer{contract: contract}}, nil
}

// NewStringReverseCaller creates a new read-only instance of StringReverse, bound to a specific deployed contract.
func NewStringReverseCaller(address common.Address, caller bind.ContractCaller) (*StringReverseCaller, error) {
	contract, err := bindStringReverse(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StringReverseCaller{contract: contract}, nil
}

// NewStringReverseTransactor creates a new write-only instance of StringReverse, bound to a specific deployed contract.
func NewStringReverseTransactor(address common.Address, transactor bind.ContractTransactor) (*StringReverseTransactor, error) {
	contract, err := bindStringReverse(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StringReverseTransactor{contract: contract}, nil
}

// NewStringReverseFilterer creates a new log filterer instance of StringReverse, bound to a specific deployed contract.
func NewStringReverseFilterer(address common.Address, filterer bind.ContractFilterer) (*StringReverseFilterer, error) {
	contract, err := bindStringReverse(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StringReverseFilterer{contract: contract}, nil
}

// bindStringReverse binds a generic wrapper to an already deployed contract.
func bindStringReverse(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StringReverseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StringReverse *StringReverseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StringReverse.Contract.StringReverseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StringReverse *StringReverseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StringReverse.Contract.StringReverseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StringReverse *StringReverseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StringReverse.Contract.StringReverseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StringReverse *StringReverseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StringReverse.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StringReverse *StringReverseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StringReverse.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StringReverse *StringReverseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StringReverse.Contract.contract.Transact(opts, method, params...)
}

// StringReverse is a paid mutator transaction binding the contract method 0xf3582c69.
//
// Solidity: function stringReverse(input string) returns(string)
func (_StringReverse *StringReverseTransactor) StringReverse(opts *bind.TransactOpts, input string) (*types.Transaction, error) {
	return _StringReverse.contract.Transact(opts, "stringReverse", input)
}

// StringReverse is a paid mutator transaction binding the contract method 0xf3582c69.
//
// Solidity: function stringReverse(input string) returns(string)
func (_StringReverse *StringReverseSession) StringReverse(input string) (*types.Transaction, error) {
	return _StringReverse.Contract.StringReverse(&_StringReverse.TransactOpts, input)
}

// StringReverse is a paid mutator transaction binding the contract method 0xf3582c69.
//
// Solidity: function stringReverse(input string) returns(string)
func (_StringReverse *StringReverseTransactorSession) StringReverse(input string) (*types.Transaction, error) {
	return _StringReverse.Contract.StringReverse(&_StringReverse.TransactOpts, input)
}

// TesterABI is the input ABI used to generate the binding from.
const TesterABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"code\",\"type\":\"bytes\"}],\"name\":\"testStringReverse\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"code\",\"type\":\"bytes\"}],\"name\":\"testAdder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submissionCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"submissions\",\"outputs\":[{\"name\":\"pass\",\"type\":\"bool\"},{\"name\":\"challenge\",\"type\":\"string\"},{\"name\":\"submitter\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"res\",\"type\":\"bool\"}],\"name\":\"TestPass\",\"type\":\"event\"}]"

// TesterBin is the compiled bytecode used for deploying new contracts.
const TesterBin = `0x608060405234801561001057600080fd5b50610b9a806100206000396000f3006080604052600436106100615763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663267a4728811461006657806366dfce07146100c15780636ec02be91461011a578063ad73349e14610141575b600080fd5b34801561007257600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100bf9436949293602493928401919081908401838280828437509497506101eb9650505050505050565b005b3480156100cd57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100bf9436949293602493928401919081908401838280828437509497506105e09650505050505050565b34801561012657600080fd5b5061012f61098e565b60408051918252519081900360200190f35b34801561014d57600080fd5b50610159600435610994565b604080518415158152600160a060020a03831691810191909152606060208083018281528551928401929092528451608084019186019080838360005b838110156101ae578181015183820152602001610196565b50505050905090810190601f1680156101db5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b6000806101f783610a62565b9150600160a060020a038216151561020e57600080fd5b50600080546001018155604080517f66656463626100000000000000000000000000000000000000000000000000008152815190819003600690810182207ff3582c690000000000000000000000000000000000000000000000000000000083526020600484015260248301919091527f6162636465660000000000000000000000000000000000000000000000000000604483015291518493600160a060020a0385169263f3582c69926064808301939282900301818387803b1580156102d557600080fd5b505af11580156102e9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561031257600080fd5b81019080805164010000000081111561032a57600080fd5b8201602081018481111561033d57600080fd5b815164010000000081118282018710171561035757600080fd5b50509291905050506040518082805190602001908083835b6020831061038e5780518252601f19909201916020918201910161036f565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000191614156104d45760408051600181529051600080516020610b0f8339815191529181900360200190a160408051606081018252600180825282518084018452600d81527f537472696e6752657665727365000000000000000000000000000000000000006020828101919091528084019182523394840194909452815480830180845560009390935283516003909102600080516020610b2f8339815191528101805492151560ff19909316929092178255915180519395919361049793600080516020610b4f83398151915201929190910190610a73565b50604091909101516002909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909216919091179055506105db565b60408051600081529051600080516020610b0f8339815191529181900360200190a160408051606081018252600080825282518084018452600d81527f537472696e6752657665727365000000000000000000000000000000000000006020828101919091528084019182523394840194909452600180548082018083559190935283516003909302600080516020610b2f8339815191528101805494151560ff199095169490941784559151805191956105a293600080516020610b4f8339815191520192910190610a73565b50604091909101516002909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909216919091179055505b505050565b6000806105ec83610a62565b9150600160a060020a038216151561060357600080fd5b50600080546001018155604080517fa5f3c23b000000000000000000000000000000000000000000000000000000008152600760048201526003602482015290518392600160a060020a0384169263a5f3c23b9260448083019360209383900390910190829087803b15801561067857600080fd5b505af115801561068c573d6000803e3d6000fd5b505050506040513d60208110156106a257600080fd5b5051600a1480156107465750604080517fa5f3c23b0000000000000000000000000000000000000000000000000000000081526002196004820152600360248201529051600160a060020a0383169163a5f3c23b9160448083019260209291908290030181600087803b15801561071857600080fd5b505af115801561072c573d6000803e3d6000fd5b505050506040513d602081101561074257600080fd5b5051155b80156107e95750604080517fa5f3c23b000000000000000000000000000000000000000000000000000000008152600219600482015260061960248201529051600160a060020a0383169163a5f3c23b9160448083019260209291908290030181600087803b1580156107b857600080fd5b505af11580156107cc573d6000803e3d6000fd5b505050506040513d60208110156107e257600080fd5b5051600919145b156108c05760408051600181529051600080516020610b0f8339815191529181900360200190a160408051606081018252600180825282518084018452600581527f41646465720000000000000000000000000000000000000000000000000000006020828101919091528084019182523394840194909452815480830180845560009390935283516003909102600080516020610b2f8339815191528101805492151560ff19909316929092178255915180519395919361049793600080516020610b4f83398151915201929190910190610a73565b60408051600081529051600080516020610b0f8339815191529181900360200190a160408051606081018252600080825282518084018452600581527f41646465720000000000000000000000000000000000000000000000000000006020828101919091528084019182523394840194909452600180548082018083559190935283516003909302600080516020610b2f8339815191528101805494151560ff199095169490941784559151805191956105a293600080516020610b4f8339815191520192910190610a73565b60005481565b60018054829081106109a257fe5b600091825260209182902060039091020180546001808301805460408051601f600260001996851615610100029690960190931694909404918201879004870284018701905280835260ff9093169550929390929190830182828015610a495780601f10610a1e57610100808354040283529160200191610a49565b820191906000526020600020905b815481529060010190602001808311610a2c57829003601f168201915b50505060029093015491925050600160a060020a031683565b60008151602083016000f092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610ab457805160ff1916838001178555610ae1565b82800160010185558215610ae1579182015b82811115610ae1578251825591602001919060010190610ac6565b50610aed929150610af1565b5090565b610b0b91905b80821115610aed5760008155600101610af7565b905600e7230e41682a55f671747a59595e23f58680339803ff4a8a09cbc4d1a433f84eb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6b10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf7a165627a7a723058203d7145e9d0c0a7ec71382975e53a9fe43324afbc6ee801615eb38661079cd2c10029`

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

// SubmissionCount is a free data retrieval call binding the contract method 0x6ec02be9.
//
// Solidity: function submissionCount() constant returns(uint256)
func (_Tester *TesterCaller) SubmissionCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Tester.contract.Call(opts, out, "submissionCount")
	return *ret0, err
}

// SubmissionCount is a free data retrieval call binding the contract method 0x6ec02be9.
//
// Solidity: function submissionCount() constant returns(uint256)
func (_Tester *TesterSession) SubmissionCount() (*big.Int, error) {
	return _Tester.Contract.SubmissionCount(&_Tester.CallOpts)
}

// SubmissionCount is a free data retrieval call binding the contract method 0x6ec02be9.
//
// Solidity: function submissionCount() constant returns(uint256)
func (_Tester *TesterCallerSession) SubmissionCount() (*big.Int, error) {
	return _Tester.Contract.SubmissionCount(&_Tester.CallOpts)
}

// Submissions is a free data retrieval call binding the contract method 0xad73349e.
//
// Solidity: function submissions( uint256) constant returns(pass bool, challenge string, submitter address)
func (_Tester *TesterCaller) Submissions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Pass      bool
	Challenge string
	Submitter common.Address
}, error) {
	ret := new(struct {
		Pass      bool
		Challenge string
		Submitter common.Address
	})
	out := ret
	err := _Tester.contract.Call(opts, out, "submissions", arg0)
	return *ret, err
}

// Submissions is a free data retrieval call binding the contract method 0xad73349e.
//
// Solidity: function submissions( uint256) constant returns(pass bool, challenge string, submitter address)
func (_Tester *TesterSession) Submissions(arg0 *big.Int) (struct {
	Pass      bool
	Challenge string
	Submitter common.Address
}, error) {
	return _Tester.Contract.Submissions(&_Tester.CallOpts, arg0)
}

// Submissions is a free data retrieval call binding the contract method 0xad73349e.
//
// Solidity: function submissions( uint256) constant returns(pass bool, challenge string, submitter address)
func (_Tester *TesterCallerSession) Submissions(arg0 *big.Int) (struct {
	Pass      bool
	Challenge string
	Submitter common.Address
}, error) {
	return _Tester.Contract.Submissions(&_Tester.CallOpts, arg0)
}

// TestAdder is a paid mutator transaction binding the contract method 0x66dfce07.
//
// Solidity: function testAdder(code bytes) returns()
func (_Tester *TesterTransactor) TestAdder(opts *bind.TransactOpts, code []byte) (*types.Transaction, error) {
	return _Tester.contract.Transact(opts, "testAdder", code)
}

// TestAdder is a paid mutator transaction binding the contract method 0x66dfce07.
//
// Solidity: function testAdder(code bytes) returns()
func (_Tester *TesterSession) TestAdder(code []byte) (*types.Transaction, error) {
	return _Tester.Contract.TestAdder(&_Tester.TransactOpts, code)
}

// TestAdder is a paid mutator transaction binding the contract method 0x66dfce07.
//
// Solidity: function testAdder(code bytes) returns()
func (_Tester *TesterTransactorSession) TestAdder(code []byte) (*types.Transaction, error) {
	return _Tester.Contract.TestAdder(&_Tester.TransactOpts, code)
}

// TestStringReverse is a paid mutator transaction binding the contract method 0x267a4728.
//
// Solidity: function testStringReverse(code bytes) returns()
func (_Tester *TesterTransactor) TestStringReverse(opts *bind.TransactOpts, code []byte) (*types.Transaction, error) {
	return _Tester.contract.Transact(opts, "testStringReverse", code)
}

// TestStringReverse is a paid mutator transaction binding the contract method 0x267a4728.
//
// Solidity: function testStringReverse(code bytes) returns()
func (_Tester *TesterSession) TestStringReverse(code []byte) (*types.Transaction, error) {
	return _Tester.Contract.TestStringReverse(&_Tester.TransactOpts, code)
}

// TestStringReverse is a paid mutator transaction binding the contract method 0x267a4728.
//
// Solidity: function testStringReverse(code bytes) returns()
func (_Tester *TesterTransactorSession) TestStringReverse(code []byte) (*types.Transaction, error) {
	return _Tester.Contract.TestStringReverse(&_Tester.TransactOpts, code)
}

// TesterTestPassIterator is returned from FilterTestPass and is used to iterate over the raw logs and unpacked data for TestPass events raised by the Tester contract.
type TesterTestPassIterator struct {
	Event *TesterTestPass // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TesterTestPassIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TesterTestPass)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TesterTestPass)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TesterTestPassIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TesterTestPassIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TesterTestPass represents a TestPass event raised by the Tester contract.
type TesterTestPass struct {
	Res bool
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTestPass is a free log retrieval operation binding the contract event 0xe7230e41682a55f671747a59595e23f58680339803ff4a8a09cbc4d1a433f84e.
//
// Solidity: e TestPass(res bool)
func (_Tester *TesterFilterer) FilterTestPass(opts *bind.FilterOpts) (*TesterTestPassIterator, error) {

	logs, sub, err := _Tester.contract.FilterLogs(opts, "TestPass")
	if err != nil {
		return nil, err
	}
	return &TesterTestPassIterator{contract: _Tester.contract, event: "TestPass", logs: logs, sub: sub}, nil
}

// WatchTestPass is a free log subscription operation binding the contract event 0xe7230e41682a55f671747a59595e23f58680339803ff4a8a09cbc4d1a433f84e.
//
// Solidity: e TestPass(res bool)
func (_Tester *TesterFilterer) WatchTestPass(opts *bind.WatchOpts, sink chan<- *TesterTestPass) (event.Subscription, error) {

	logs, sub, err := _Tester.contract.WatchLogs(opts, "TestPass")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TesterTestPass)
				if err := _Tester.contract.UnpackLog(event, "TestPass", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
