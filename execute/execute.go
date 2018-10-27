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
const AdderABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AdderBin is the compiled bytecode used for deploying new contracts.
const AdderBin = `0x6080604052348015600f57600080fd5b5060a18061001e6000396000f300608060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663771602f781146043575b600080fd5b348015604e57600080fd5b50605b600435602435606d565b60408051918252519081900360200190f35b6000929150505600a165627a7a72305820d2de68dacb72250f7466fce9d3f863932e831567882cf4e355035bdcee7319f80029`

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
const TesterABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"code\",\"type\":\"bytes\"}],\"name\":\"test\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"submissionCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"submissions\",\"outputs\":[{\"name\":\"pass\",\"type\":\"bool\"},{\"name\":\"submitter\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getSubmitters\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"res\",\"type\":\"bool\"}],\"name\":\"TestPass\",\"type\":\"event\"}]"

// TesterBin is the compiled bytecode used for deploying new contracts.
const TesterBin = `0x608060405234801561001057600080fd5b5061049e806100206000396000f3006080604052600436106100615763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632f570a2381146100665780636ec02be9146100c1578063ad73349e146100e8578063ae4dac4114610130575b600080fd5b34801561007257600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100bf9436949293602493928401919081908401838280828437509497506101959650505050505050565b005b3480156100cd57600080fd5b506100d6610418565b60408051918252519081900360200190f35b3480156100f457600080fd5b5061010060043561041e565b60408051921515835273ffffffffffffffffffffffffffffffffffffffff90911660208301528051918290030190f35b34801561013c57600080fd5b5061014561045c565b60408051602080825283518183015283519192839290830191858101910280838360005b83811015610181578181015183820152602001610169565b505050509050019250505060405180910390f35b6000806101a183610461565b915073ffffffffffffffffffffffffffffffffffffffff821615156101c557600080fd5b50600080546001018155604080517f771602f7000000000000000000000000000000000000000000000000000000008152600a6004820181905260248201529051839273ffffffffffffffffffffffffffffffffffffffff84169263771602f79260448083019360209383900390910190829087803b15801561024757600080fd5b505af115801561025b573d6000803e3d6000fd5b505050506040513d602081101561027157600080fd5b50516014141561034a57604080516001815290517fe7230e41682a55f671747a59595e23f58680339803ff4a8a09cbc4d1a433f84e9181900360200190a160408051808201909152600180825233602083019081528154808301835560009290925291517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf69091018054925173ffffffffffffffffffffffffffffffffffffffff166101000274ffffffffffffffffffffffffffffffffffffffff001992151560ff199094169390931791909116919091179055610413565b604080516000815290517fe7230e41682a55f671747a59595e23f58680339803ff4a8a09cbc4d1a433f84e9181900360200190a16040805180820190915260008082523360208301908152600180548082018255925291517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf69091018054925173ffffffffffffffffffffffffffffffffffffffff166101000274ffffffffffffffffffffffffffffffffffffffff001992151560ff1990941693909317919091169190911790555b505050565b60005481565b600180548290811061042c57fe5b60009182526020909120015460ff81169150610100900473ffffffffffffffffffffffffffffffffffffffff1682565b606090565b60008151602083016000f0929150505600a165627a7a7230582005dea32886e23ceb8e48b063007010320dc592f3f3b31247769a9a38ed2781430029`

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
// Solidity: function submissions( uint256) constant returns(pass bool, submitter address)
func (_Tester *TesterCaller) Submissions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Pass      bool
	Submitter common.Address
}, error) {
	ret := new(struct {
		Pass      bool
		Submitter common.Address
	})
	out := ret
	err := _Tester.contract.Call(opts, out, "submissions", arg0)
	return *ret, err
}

// Submissions is a free data retrieval call binding the contract method 0xad73349e.
//
// Solidity: function submissions( uint256) constant returns(pass bool, submitter address)
func (_Tester *TesterSession) Submissions(arg0 *big.Int) (struct {
	Pass      bool
	Submitter common.Address
}, error) {
	return _Tester.Contract.Submissions(&_Tester.CallOpts, arg0)
}

// Submissions is a free data retrieval call binding the contract method 0xad73349e.
//
// Solidity: function submissions( uint256) constant returns(pass bool, submitter address)
func (_Tester *TesterCallerSession) Submissions(arg0 *big.Int) (struct {
	Pass      bool
	Submitter common.Address
}, error) {
	return _Tester.Contract.Submissions(&_Tester.CallOpts, arg0)
}

// GetSubmitters is a paid mutator transaction binding the contract method 0xae4dac41.
//
// Solidity: function getSubmitters() returns(address[])
func (_Tester *TesterTransactor) GetSubmitters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tester.contract.Transact(opts, "getSubmitters")
}

// GetSubmitters is a paid mutator transaction binding the contract method 0xae4dac41.
//
// Solidity: function getSubmitters() returns(address[])
func (_Tester *TesterSession) GetSubmitters() (*types.Transaction, error) {
	return _Tester.Contract.GetSubmitters(&_Tester.TransactOpts)
}

// GetSubmitters is a paid mutator transaction binding the contract method 0xae4dac41.
//
// Solidity: function getSubmitters() returns(address[])
func (_Tester *TesterTransactorSession) GetSubmitters() (*types.Transaction, error) {
	return _Tester.Contract.GetSubmitters(&_Tester.TransactOpts)
}

// Test is a paid mutator transaction binding the contract method 0x2f570a23.
//
// Solidity: function test(code bytes) returns()
func (_Tester *TesterTransactor) Test(opts *bind.TransactOpts, code []byte) (*types.Transaction, error) {
	return _Tester.contract.Transact(opts, "test", code)
}

// Test is a paid mutator transaction binding the contract method 0x2f570a23.
//
// Solidity: function test(code bytes) returns()
func (_Tester *TesterSession) Test(code []byte) (*types.Transaction, error) {
	return _Tester.Contract.Test(&_Tester.TransactOpts, code)
}

// Test is a paid mutator transaction binding the contract method 0x2f570a23.
//
// Solidity: function test(code bytes) returns()
func (_Tester *TesterTransactorSession) Test(code []byte) (*types.Transaction, error) {
	return _Tester.Contract.Test(&_Tester.TransactOpts, code)
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
