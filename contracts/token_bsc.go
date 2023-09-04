// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenomics_bsc

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// TokenomicsBscMetaData contains all meta data concerning the TokenomicsBsc contract.
var TokenomicsBscMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Transfer\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"indexed\":true},{\"name\":\"_to\",\"type\":\"address\",\"indexed\":true},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"_spender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ReceiveEther\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"indexed\":true},{\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"_current_bal\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"_gasLeft\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"_oldOwner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"_newOwner\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_totalSupply\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"sendEther\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"sendAllEther\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"checkContractBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"OWNER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"TOKEN_NAME\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"TOKEN_SYMBOL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"TOTAL_SUPPLY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]}]",
}

// TokenomicsBscABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenomicsBscMetaData.ABI instead.
var TokenomicsBscABI = TokenomicsBscMetaData.ABI

// TokenomicsBsc is an auto generated Go binding around an Ethereum contract.
type TokenomicsBsc struct {
	TokenomicsBscCaller     // Read-only binding to the contract
	TokenomicsBscTransactor // Write-only binding to the contract
	TokenomicsBscFilterer   // Log filterer for contract events
}

// TokenomicsBscCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenomicsBscCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenomicsBscTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenomicsBscTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenomicsBscFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenomicsBscFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenomicsBscSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenomicsBscSession struct {
	Contract     *TokenomicsBsc    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenomicsBscCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenomicsBscCallerSession struct {
	Contract *TokenomicsBscCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TokenomicsBscTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenomicsBscTransactorSession struct {
	Contract     *TokenomicsBscTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TokenomicsBscRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenomicsBscRaw struct {
	Contract *TokenomicsBsc // Generic contract binding to access the raw methods on
}

// TokenomicsBscCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenomicsBscCallerRaw struct {
	Contract *TokenomicsBscCaller // Generic read-only contract binding to access the raw methods on
}

// TokenomicsBscTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenomicsBscTransactorRaw struct {
	Contract *TokenomicsBscTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenomicsBsc creates a new instance of TokenomicsBsc, bound to a specific deployed contract.
func NewTokenomicsBsc(address common.Address, backend bind.ContractBackend) (*TokenomicsBsc, error) {
	contract, err := bindTokenomicsBsc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenomicsBsc{TokenomicsBscCaller: TokenomicsBscCaller{contract: contract}, TokenomicsBscTransactor: TokenomicsBscTransactor{contract: contract}, TokenomicsBscFilterer: TokenomicsBscFilterer{contract: contract}}, nil
}

// NewTokenomicsBscCaller creates a new read-only instance of TokenomicsBsc, bound to a specific deployed contract.
func NewTokenomicsBscCaller(address common.Address, caller bind.ContractCaller) (*TokenomicsBscCaller, error) {
	contract, err := bindTokenomicsBsc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenomicsBscCaller{contract: contract}, nil
}

// NewTokenomicsBscTransactor creates a new write-only instance of TokenomicsBsc, bound to a specific deployed contract.
func NewTokenomicsBscTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenomicsBscTransactor, error) {
	contract, err := bindTokenomicsBsc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenomicsBscTransactor{contract: contract}, nil
}

// NewTokenomicsBscFilterer creates a new log filterer instance of TokenomicsBsc, bound to a specific deployed contract.
func NewTokenomicsBscFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenomicsBscFilterer, error) {
	contract, err := bindTokenomicsBsc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenomicsBscFilterer{contract: contract}, nil
}

// bindTokenomicsBsc binds a generic wrapper to an already deployed contract.
func bindTokenomicsBsc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenomicsBscMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenomicsBsc *TokenomicsBscRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenomicsBsc.Contract.TokenomicsBscCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenomicsBsc *TokenomicsBscRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.TokenomicsBscTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenomicsBsc *TokenomicsBscRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.TokenomicsBscTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenomicsBsc *TokenomicsBscCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenomicsBsc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenomicsBsc *TokenomicsBscTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenomicsBsc *TokenomicsBscTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.contract.Transact(opts, method, params...)
}

// OWNER is a free data retrieval call binding the contract method 0x117803e3.
//
// Solidity: function OWNER() view returns(address)
func (_TokenomicsBsc *TokenomicsBscCaller) OWNER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenomicsBsc.contract.Call(opts, &out, "OWNER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OWNER is a free data retrieval call binding the contract method 0x117803e3.
//
// Solidity: function OWNER() view returns(address)
func (_TokenomicsBsc *TokenomicsBscSession) OWNER() (common.Address, error) {
	return _TokenomicsBsc.Contract.OWNER(&_TokenomicsBsc.CallOpts)
}

// OWNER is a free data retrieval call binding the contract method 0x117803e3.
//
// Solidity: function OWNER() view returns(address)
func (_TokenomicsBsc *TokenomicsBscCallerSession) OWNER() (common.Address, error) {
	return _TokenomicsBsc.Contract.OWNER(&_TokenomicsBsc.CallOpts)
}

// TOKENNAME is a free data retrieval call binding the contract method 0x18821400.
//
// Solidity: function TOKEN_NAME() view returns(string)
func (_TokenomicsBsc *TokenomicsBscCaller) TOKENNAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenomicsBsc.contract.Call(opts, &out, "TOKEN_NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TOKENNAME is a free data retrieval call binding the contract method 0x18821400.
//
// Solidity: function TOKEN_NAME() view returns(string)
func (_TokenomicsBsc *TokenomicsBscSession) TOKENNAME() (string, error) {
	return _TokenomicsBsc.Contract.TOKENNAME(&_TokenomicsBsc.CallOpts)
}

// TOKENNAME is a free data retrieval call binding the contract method 0x18821400.
//
// Solidity: function TOKEN_NAME() view returns(string)
func (_TokenomicsBsc *TokenomicsBscCallerSession) TOKENNAME() (string, error) {
	return _TokenomicsBsc.Contract.TOKENNAME(&_TokenomicsBsc.CallOpts)
}

// TOKENSYMBOL is a free data retrieval call binding the contract method 0x2a905318.
//
// Solidity: function TOKEN_SYMBOL() view returns(string)
func (_TokenomicsBsc *TokenomicsBscCaller) TOKENSYMBOL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenomicsBsc.contract.Call(opts, &out, "TOKEN_SYMBOL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TOKENSYMBOL is a free data retrieval call binding the contract method 0x2a905318.
//
// Solidity: function TOKEN_SYMBOL() view returns(string)
func (_TokenomicsBsc *TokenomicsBscSession) TOKENSYMBOL() (string, error) {
	return _TokenomicsBsc.Contract.TOKENSYMBOL(&_TokenomicsBsc.CallOpts)
}

// TOKENSYMBOL is a free data retrieval call binding the contract method 0x2a905318.
//
// Solidity: function TOKEN_SYMBOL() view returns(string)
func (_TokenomicsBsc *TokenomicsBscCallerSession) TOKENSYMBOL() (string, error) {
	return _TokenomicsBsc.Contract.TOKENSYMBOL(&_TokenomicsBsc.CallOpts)
}

// TOTALSUPPLY is a free data retrieval call binding the contract method 0x902d55a5.
//
// Solidity: function TOTAL_SUPPLY() view returns(uint256)
func (_TokenomicsBsc *TokenomicsBscCaller) TOTALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenomicsBsc.contract.Call(opts, &out, "TOTAL_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALSUPPLY is a free data retrieval call binding the contract method 0x902d55a5.
//
// Solidity: function TOTAL_SUPPLY() view returns(uint256)
func (_TokenomicsBsc *TokenomicsBscSession) TOTALSUPPLY() (*big.Int, error) {
	return _TokenomicsBsc.Contract.TOTALSUPPLY(&_TokenomicsBsc.CallOpts)
}

// TOTALSUPPLY is a free data retrieval call binding the contract method 0x902d55a5.
//
// Solidity: function TOTAL_SUPPLY() view returns(uint256)
func (_TokenomicsBsc *TokenomicsBscCallerSession) TOTALSUPPLY() (*big.Int, error) {
	return _TokenomicsBsc.Contract.TOTALSUPPLY(&_TokenomicsBsc.CallOpts)
}

// Allowance is a paid mutator transaction binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) returns(uint256)
func (_TokenomicsBsc *TokenomicsBscTransactor) Allowance(opts *bind.TransactOpts, owner common.Address, spender common.Address) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "allowance", owner, spender)
}

// Allowance is a paid mutator transaction binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) returns(uint256)
func (_TokenomicsBsc *TokenomicsBscSession) Allowance(owner common.Address, spender common.Address) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.Allowance(&_TokenomicsBsc.TransactOpts, owner, spender)
}

// Allowance is a paid mutator transaction binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) returns(uint256)
func (_TokenomicsBsc *TokenomicsBscTransactorSession) Allowance(owner common.Address, spender common.Address) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.Allowance(&_TokenomicsBsc.TransactOpts, owner, spender)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenomicsBsc *TokenomicsBscTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenomicsBsc *TokenomicsBscSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.Approve(&_TokenomicsBsc.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenomicsBsc *TokenomicsBscTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.Approve(&_TokenomicsBsc.TransactOpts, spender, amount)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) returns(uint256)
func (_TokenomicsBsc *TokenomicsBscTransactor) BalanceOf(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "balanceOf", account)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) returns(uint256)
func (_TokenomicsBsc *TokenomicsBscSession) BalanceOf(account common.Address) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.BalanceOf(&_TokenomicsBsc.TransactOpts, account)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) returns(uint256)
func (_TokenomicsBsc *TokenomicsBscTransactorSession) BalanceOf(account common.Address) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.BalanceOf(&_TokenomicsBsc.TransactOpts, account)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_TokenomicsBsc *TokenomicsBscTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_TokenomicsBsc *TokenomicsBscSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.Burn(&_TokenomicsBsc.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_TokenomicsBsc *TokenomicsBscTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.Burn(&_TokenomicsBsc.TransactOpts, account, amount)
}

// CheckContractBalance is a paid mutator transaction binding the contract method 0x50312c9e.
//
// Solidity: function checkContractBalance() returns(uint256)
func (_TokenomicsBsc *TokenomicsBscTransactor) CheckContractBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "checkContractBalance")
}

// CheckContractBalance is a paid mutator transaction binding the contract method 0x50312c9e.
//
// Solidity: function checkContractBalance() returns(uint256)
func (_TokenomicsBsc *TokenomicsBscSession) CheckContractBalance() (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.CheckContractBalance(&_TokenomicsBsc.TransactOpts)
}

// CheckContractBalance is a paid mutator transaction binding the contract method 0x50312c9e.
//
// Solidity: function checkContractBalance() returns(uint256)
func (_TokenomicsBsc *TokenomicsBscTransactorSession) CheckContractBalance() (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.CheckContractBalance(&_TokenomicsBsc.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TokenomicsBsc *TokenomicsBscTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TokenomicsBsc *TokenomicsBscSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.DecreaseAllowance(&_TokenomicsBsc.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TokenomicsBsc *TokenomicsBscTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.DecreaseAllowance(&_TokenomicsBsc.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TokenomicsBsc *TokenomicsBscTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TokenomicsBsc *TokenomicsBscSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.IncreaseAllowance(&_TokenomicsBsc.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TokenomicsBsc *TokenomicsBscTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.IncreaseAllowance(&_TokenomicsBsc.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_TokenomicsBsc *TokenomicsBscTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_TokenomicsBsc *TokenomicsBscSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.Mint(&_TokenomicsBsc.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_TokenomicsBsc *TokenomicsBscTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.Mint(&_TokenomicsBsc.TransactOpts, to, amount)
}

// SendAllEther is a paid mutator transaction binding the contract method 0x085b3b00.
//
// Solidity: function sendAllEther(address to) returns()
func (_TokenomicsBsc *TokenomicsBscTransactor) SendAllEther(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "sendAllEther", to)
}

// SendAllEther is a paid mutator transaction binding the contract method 0x085b3b00.
//
// Solidity: function sendAllEther(address to) returns()
func (_TokenomicsBsc *TokenomicsBscSession) SendAllEther(to common.Address) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.SendAllEther(&_TokenomicsBsc.TransactOpts, to)
}

// SendAllEther is a paid mutator transaction binding the contract method 0x085b3b00.
//
// Solidity: function sendAllEther(address to) returns()
func (_TokenomicsBsc *TokenomicsBscTransactorSession) SendAllEther(to common.Address) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.SendAllEther(&_TokenomicsBsc.TransactOpts, to)
}

// SendEther is a paid mutator transaction binding the contract method 0xc1756a2c.
//
// Solidity: function sendEther(address to, uint256 amount) returns()
func (_TokenomicsBsc *TokenomicsBscTransactor) SendEther(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "sendEther", to, amount)
}

// SendEther is a paid mutator transaction binding the contract method 0xc1756a2c.
//
// Solidity: function sendEther(address to, uint256 amount) returns()
func (_TokenomicsBsc *TokenomicsBscSession) SendEther(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.SendEther(&_TokenomicsBsc.TransactOpts, to, amount)
}

// SendEther is a paid mutator transaction binding the contract method 0xc1756a2c.
//
// Solidity: function sendEther(address to, uint256 amount) returns()
func (_TokenomicsBsc *TokenomicsBscTransactorSession) SendEther(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.SendEther(&_TokenomicsBsc.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TokenomicsBsc *TokenomicsBscTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TokenomicsBsc *TokenomicsBscSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.Transfer(&_TokenomicsBsc.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TokenomicsBsc *TokenomicsBscTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.Transfer(&_TokenomicsBsc.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TokenomicsBsc *TokenomicsBscTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TokenomicsBsc *TokenomicsBscSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.TransferFrom(&_TokenomicsBsc.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TokenomicsBsc *TokenomicsBscTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.TransferFrom(&_TokenomicsBsc.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenomicsBsc *TokenomicsBscTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenomicsBsc *TokenomicsBscSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.TransferOwnership(&_TokenomicsBsc.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenomicsBsc *TokenomicsBscTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.TransferOwnership(&_TokenomicsBsc.TransactOpts, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TokenomicsBsc *TokenomicsBscTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TokenomicsBsc *TokenomicsBscSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.Fallback(&_TokenomicsBsc.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TokenomicsBsc *TokenomicsBscTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.Fallback(&_TokenomicsBsc.TransactOpts, calldata)
}

// TokenomicsBscApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TokenomicsBsc contract.
type TokenomicsBscApprovalIterator struct {
	Event *TokenomicsBscApproval // Event containing the contract specifics and raw log

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
func (it *TokenomicsBscApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenomicsBscApproval)
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
		it.Event = new(TokenomicsBscApproval)
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
func (it *TokenomicsBscApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenomicsBscApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenomicsBscApproval represents a Approval event raised by the TokenomicsBsc contract.
type TokenomicsBscApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_TokenomicsBsc *TokenomicsBscFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*TokenomicsBscApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _TokenomicsBsc.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenomicsBscApprovalIterator{contract: _TokenomicsBsc.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_TokenomicsBsc *TokenomicsBscFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenomicsBscApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _TokenomicsBsc.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenomicsBscApproval)
				if err := _TokenomicsBsc.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_TokenomicsBsc *TokenomicsBscFilterer) ParseApproval(log types.Log) (*TokenomicsBscApproval, error) {
	event := new(TokenomicsBscApproval)
	if err := _TokenomicsBsc.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenomicsBscOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenomicsBsc contract.
type TokenomicsBscOwnershipTransferredIterator struct {
	Event *TokenomicsBscOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenomicsBscOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenomicsBscOwnershipTransferred)
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
		it.Event = new(TokenomicsBscOwnershipTransferred)
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
func (it *TokenomicsBscOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenomicsBscOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenomicsBscOwnershipTransferred represents a OwnershipTransferred event raised by the TokenomicsBsc contract.
type TokenomicsBscOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _oldOwner, address indexed _newOwner)
func (_TokenomicsBsc *TokenomicsBscFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _oldOwner []common.Address, _newOwner []common.Address) (*TokenomicsBscOwnershipTransferredIterator, error) {

	var _oldOwnerRule []interface{}
	for _, _oldOwnerItem := range _oldOwner {
		_oldOwnerRule = append(_oldOwnerRule, _oldOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _TokenomicsBsc.contract.FilterLogs(opts, "OwnershipTransferred", _oldOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenomicsBscOwnershipTransferredIterator{contract: _TokenomicsBsc.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _oldOwner, address indexed _newOwner)
func (_TokenomicsBsc *TokenomicsBscFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenomicsBscOwnershipTransferred, _oldOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _oldOwnerRule []interface{}
	for _, _oldOwnerItem := range _oldOwner {
		_oldOwnerRule = append(_oldOwnerRule, _oldOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _TokenomicsBsc.contract.WatchLogs(opts, "OwnershipTransferred", _oldOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenomicsBscOwnershipTransferred)
				if err := _TokenomicsBsc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _oldOwner, address indexed _newOwner)
func (_TokenomicsBsc *TokenomicsBscFilterer) ParseOwnershipTransferred(log types.Log) (*TokenomicsBscOwnershipTransferred, error) {
	event := new(TokenomicsBscOwnershipTransferred)
	if err := _TokenomicsBsc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenomicsBscReceiveEtherIterator is returned from FilterReceiveEther and is used to iterate over the raw logs and unpacked data for ReceiveEther events raised by the TokenomicsBsc contract.
type TokenomicsBscReceiveEtherIterator struct {
	Event *TokenomicsBscReceiveEther // Event containing the contract specifics and raw log

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
func (it *TokenomicsBscReceiveEtherIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenomicsBscReceiveEther)
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
		it.Event = new(TokenomicsBscReceiveEther)
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
func (it *TokenomicsBscReceiveEtherIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenomicsBscReceiveEtherIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenomicsBscReceiveEther represents a ReceiveEther event raised by the TokenomicsBsc contract.
type TokenomicsBscReceiveEther struct {
	From       common.Address
	Amount     *big.Int
	CurrentBal *big.Int
	GasLeft    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReceiveEther is a free log retrieval operation binding the contract event 0x16831e3995a0d3f34fe61238d5df9e0e2c2b51c29df181683822627edd94e200.
//
// Solidity: event ReceiveEther(address indexed _from, uint256 _amount, uint256 _current_bal, uint256 _gasLeft)
func (_TokenomicsBsc *TokenomicsBscFilterer) FilterReceiveEther(opts *bind.FilterOpts, _from []common.Address) (*TokenomicsBscReceiveEtherIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _TokenomicsBsc.contract.FilterLogs(opts, "ReceiveEther", _fromRule)
	if err != nil {
		return nil, err
	}
	return &TokenomicsBscReceiveEtherIterator{contract: _TokenomicsBsc.contract, event: "ReceiveEther", logs: logs, sub: sub}, nil
}

// WatchReceiveEther is a free log subscription operation binding the contract event 0x16831e3995a0d3f34fe61238d5df9e0e2c2b51c29df181683822627edd94e200.
//
// Solidity: event ReceiveEther(address indexed _from, uint256 _amount, uint256 _current_bal, uint256 _gasLeft)
func (_TokenomicsBsc *TokenomicsBscFilterer) WatchReceiveEther(opts *bind.WatchOpts, sink chan<- *TokenomicsBscReceiveEther, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _TokenomicsBsc.contract.WatchLogs(opts, "ReceiveEther", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenomicsBscReceiveEther)
				if err := _TokenomicsBsc.contract.UnpackLog(event, "ReceiveEther", log); err != nil {
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

// ParseReceiveEther is a log parse operation binding the contract event 0x16831e3995a0d3f34fe61238d5df9e0e2c2b51c29df181683822627edd94e200.
//
// Solidity: event ReceiveEther(address indexed _from, uint256 _amount, uint256 _current_bal, uint256 _gasLeft)
func (_TokenomicsBsc *TokenomicsBscFilterer) ParseReceiveEther(log types.Log) (*TokenomicsBscReceiveEther, error) {
	event := new(TokenomicsBscReceiveEther)
	if err := _TokenomicsBsc.contract.UnpackLog(event, "ReceiveEther", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenomicsBscTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TokenomicsBsc contract.
type TokenomicsBscTransferIterator struct {
	Event *TokenomicsBscTransfer // Event containing the contract specifics and raw log

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
func (it *TokenomicsBscTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenomicsBscTransfer)
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
		it.Event = new(TokenomicsBscTransfer)
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
func (it *TokenomicsBscTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenomicsBscTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenomicsBscTransfer represents a Transfer event raised by the TokenomicsBsc contract.
type TokenomicsBscTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_TokenomicsBsc *TokenomicsBscFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*TokenomicsBscTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _TokenomicsBsc.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &TokenomicsBscTransferIterator{contract: _TokenomicsBsc.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_TokenomicsBsc *TokenomicsBscFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenomicsBscTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _TokenomicsBsc.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenomicsBscTransfer)
				if err := _TokenomicsBsc.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_TokenomicsBsc *TokenomicsBscFilterer) ParseTransfer(log types.Log) (*TokenomicsBscTransfer, error) {
	event := new(TokenomicsBscTransfer)
	if err := _TokenomicsBsc.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
