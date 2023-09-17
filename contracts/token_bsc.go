// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenomics

import (
	"math/big"
	"strings"
	"errors"

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
	Bin: "0x��0x60206111275f395f516032602082611127015f395f511161112357602081611127015f395f516020820181816111270161014039508061012052505060206111475f395f516032602082611127015f395f511161112357602081611127015f395f51602082018181611127016101a03950806101805250503461112357335f55610120515f81601f0160051c600281116111235780156100b457905b8060051b6101400151816002015560010181811861009b575b50508060015550610180515f81601f0160051c600281116111235780156100f057905b8060051b6101a0015181600501556001018181186100d7575b505080600455505f6007553360a05260206111675f395f5160c052610113611021565b610ef861012561000039610ef8610000f36003361161000c57610788565b5f3560e01c63117803e3811861002c5734610ee7575f5460405260206040f35b631882140081186100ac5734610ee75760208060405280604001600154602082015f82601f0160051c60028111610ee757801561007c57905b80600201548160051b840152600101818118610065575b505050808252508051806020830101601f825f03163682375050601f19601f825160200101169050810190506040f35b632a905318811861012c5734610ee75760208060405280604001600454602082015f82601f0160051c60028111610ee75780156100fc57905b80600501548160051b8401526001018181186100e5575b505050808252508051806020830101601f825f03163682375050601f19601f825160200101169050810190506040f35b63902d55a581186101485734610ee75760075460405260206040f35b63f2fde38b81186102055760243610610ee7576004358060a01c610ee75760805234610ee7576080516101f557602960a0527f4f776e61626c653a205472616e73666572206f776e65727368697020746f206160c0527f646472657373283029000000000000000000000000000000000000000000000060e05260a05060a0518060c001601f825f031636823750506308c379a06060526020608052601f19601f60a0510116604401607cfd5b6080516040526102036108c3565b005b6370a0823181186102435760243610610ee7576004358060a01c610ee75760405234610ee75760086040516020525f5260405f205460605260206060f35b63a9059cbb81186102985760443610610ee7576004358060a01c610ee7576101805234610ee757336101a0526101a05160a0526101805160c05260243560e05261028b6108fa565b60016101c05260206101c0f35b63dd62ed3e81186102f35760443610610ee7576004358060a01c610ee7576040526024358060a01c610ee75760605234610ee75760096040516020525f5260405f20806060516020525f5260405f2090505460805260206080f35b63095ea7b381186103485760443610610ee7576004358060a01c610ee7576101005234610ee7573361012052610120516040526101005160605260243560805261033b610b26565b6001610140526020610140f35b6323b872dd81186103cb5760643610610ee7576004358060a01c610ee7576101c0526024358060a01c610ee7576101e05234610ee75733610200526101c051610100526102005161012052604435610140526103a2610c89565b6101c05160a0526101e05160c05260443560e0526103be6108fa565b6001610220526020610220f35b633950935181186104435760443610610ee7576004358060a01c610ee7576101005234610ee75733604052610100516060526009336020525f5260405f2080610100516020525f5260405f20905054602435808201828110610ee75790509050608052610436610b26565b6001610120526020610120f35b63a457c2d7811861055a5760443610610ee7576004358060a01c610ee7576101005234610ee7576009336020525f5260405f2080610100516020525f5260405f2090505461012052602435610120511015610522576022610140527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77610160527f20300000000000000000000000000000000000000000000000000000000000006101805261014050610140518061016001601f825f031636823750506308c379a061010052602061012052601f19601f61014051011660440161011cfd5b336040526101005160605261012051602435808203828111610ee7579050905060805261054d610b26565b6001610140526020610140f35b6340c10f19811861062b5760443610610ee7576004358060a01c610ee7576101205234610ee757335f541815610614576033610140527f47546f6b656e3a20596f7520617265206e6f7420616c6c6f7765642061636365610160527f7373696e672074686973207265736f75726365000000000000000000000000006101805261014050610140518061016001601f825f031636823750506308c379a061010052602061012052601f19601f61014051011660440161011cfd5b6101205160a05260243560c0526106296107c1565b005b639dc29fac81186106fc5760443610610ee7576004358060a01c610ee7576101205234610ee757335f5418156106e5576033610140527f47546f6b656e3a20596f7520617265206e6f7420616c6c6f7765642061636365610160527f7373696e672074686973207265736f75726365000000000000000000000000006101805261014050610140518061016001601f825f031636823750506308c379a061010052602061012052601f19601f61014051011660440161011cfd5b6101205160a05260243560c0526106fa610d7b565b005b63c1756a2c81186107355760443610610ee7576004358060a01c610ee75760405234610ee7575f5f5f5f6024356040515ff115610ee757005b63085b3b00811861076c5760243610610ee7576004358060a01c610ee75760405234610ee7575f5f5f5f476040515ff115610ee757005b6350312c9e81186107865734610ee7574760405260206040f35b505b337f16831e3995a0d3f34fe61238d5df9e0e2c2b51c29df181683822627edd94e20034604052476060525a60805260606040a2005b565b565b60a05161082657601960e0527f45524332303a206d696e7420746f2061646472657373283029000000000000006101005260e05060e0518061010001601f825f031636823750506308c379a060a052602060c052601f19601f60e051011660440160bcfd5b5f60405260a05160605260c05160805261083e6107bd565b600860a0516020525f5260405f20805460c051808201828110610ee7579050905081555060075460c051808201828110610ee7579050905060075560a0515f7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60c05160e052602060e0a35f60405260a05160605260c0516080526108c16107bf565b565b5f546060526040515f556040516060517f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f6080a3565b60a05161096357601f610100527f45524332303a207472616e736665722066726f6d2061646472657373283029006101205261010050610100518061012001601f825f031636823750506308c379a060c052602060e052601f19601f61010051011660440160dcfd5b60c0516109cc57601d610100527f45524332303b207472616e7366657220746f20616464726573732830290000006101205261010050610100518061012001601f825f031636823750506308c379a060c052602060e052601f19601f61010051011660440160dcfd5b60a05160405260c05160605260e0516080526109e66107bd565b600860a0516020525f5260405f20546101005260e051610100511015610a8e576026610120527f45524332303a207472616e7366657220616d6f756e7420657863656564732062610140527f616c616e636500000000000000000000000000000000000000000000000000006101605261012050610120518061014001601f825f031636823750506308c379a060e052602061010052601f19601f61012051011660440160fcfd5b600860a0516020525f5260405f20805460e051808203828111610ee75790509050815550600860c0516020525f5260405f20805460e051808201828110610ee7579050905081555060c05160a0517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60e051610120526020610120a360a05160405260c05160605260e051608052610b246107bf565b565b604051610bad57602560a0527f45524332303a20496e76616c696420617070726f766572206f6620616464726560c0527f737328302900000000000000000000000000000000000000000000000000000060e05260a05060a0518060c001601f825f031636823750506308c379a06060526020608052601f19601f60a0510116604401607cfd5b606051610c3457602460a0527f45524332303a20496e76616c6964207370656e646572206f662061646472657360c0527f732830290000000000000000000000000000000000000000000000000000000060e05260a05060a0518060c001601f825f031636823750506308c379a06060526020608052601f19601f60a0510116604401607cfd5b60805160096040516020525f5260405f20806060516020525f5260405f209050556060516040517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560805160a052602060a0a3565b6009610100516020525f5260405f2080610120516020525f5260405f20905054610160527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6101605114610d795761014051610160511015610d4a57601d610180527f45524332303a20496e73756666696369656e7420616c6c6f77616e63650000006101a0526101805061018051806101a001601f825f031636823750506308c379a061014052602061016052601f19601f61018051011660440161015cfd5b61010051604052610120516060526101605161014051808203828111610ee75790509050608052610d79610b26565b565b60a051610de057601b60e0527f45524332303a206275726e2066726f6d206164647265737328302900000000006101005260e05060e0518061010001601f825f031636823750506308c379a060a052602060c052601f19601f60e051011660440160bcfd5b60c0516007541015610e4a57601960e0527f45524332303a206e6f7420656e6f7567682062616c616e6365000000000000006101005260e05060e0518061010001601f825f031636823750506308c379a060a052602060c052601f19601f60e051011660440160bcfd5b60a0516040525f60605260c051608052610e626107bd565b600860a0516020525f5260405f20805460c051808203828111610ee7579050905081555060075460c051808203828111610ee757905090506007555f60a0517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60c05160e052602060e0a360a0516040525f60605260c051608052610ee56107bf565b565b5f80fda165767970657283000309000b5b565b565b60a05161108657601960e0527f45524332303a206d696e7420746f2061646472657373283029000000000000006101005260e05060e0518061010001601f825f031636823750506308c379a060a052602060c052601f19601f60e051011660440160bcfd5b5f60405260a05160605260c05160805261109e61101d565b600860a0516020525f5260405f20805460c051808201828110611123579050905081555060075460c051808201828110611123579050905060075560a0515f7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60c05160e052602060e0a35f60405260a05160605260c05160805261112161101f565b565b5f80fd",

}
// TokenomicsBscABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenomicsBscMetaData.ABI instead.
var TokenomicsBscABI = TokenomicsBscMetaData.ABI




// TokenomicsBscBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenomicsBscMetaData.Bin instead.
var TokenomicsBscBin = TokenomicsBscMetaData.Bin

// DeployTokenomicsBsc deploys a new Ethereum contract, binding an instance of TokenomicsBsc to it.
func DeployTokenomicsBsc(auth *bind.TransactOpts, backend bind.ContractBackend , _name string, _symbol string, _totalSupply *big.Int) (common.Address, *types.Transaction, *TokenomicsBsc, error) {
	parsed, err := TokenomicsBscMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenomicsBscBin), backend , _name, _symbol, _totalSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenomicsBsc{ TokenomicsBscCaller: TokenomicsBscCaller{contract: contract}, TokenomicsBscTransactor: TokenomicsBscTransactor{contract: contract}, TokenomicsBscFilterer: TokenomicsBscFilterer{contract: contract} }, nil
}


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
	Contract     *TokenomicsBsc        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenomicsBscCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenomicsBscCallerSession struct {
	Contract *TokenomicsBscCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TokenomicsBscTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenomicsBscTransactorSession struct {
	Contract     *TokenomicsBscTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
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
	return &TokenomicsBsc{ TokenomicsBscCaller: TokenomicsBscCaller{contract: contract}, TokenomicsBscTransactor: TokenomicsBscTransactor{contract: contract}, TokenomicsBscFilterer: TokenomicsBscFilterer{contract: contract} }, nil
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
func (_TokenomicsBsc *TokenomicsBscCaller) OWNER(opts *bind.CallOpts ) (common.Address, error) {
	var out []interface{}
	err := _TokenomicsBsc.contract.Call(opts, &out, "OWNER" )

	if err != nil {
		return *new(common.Address),  err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0,  err

}

// OWNER is a free data retrieval call binding the contract method 0x117803e3.
//
// Solidity: function OWNER() view returns(address)
func (_TokenomicsBsc *TokenomicsBscSession) OWNER() ( common.Address,  error) {
	return _TokenomicsBsc.Contract.OWNER(&_TokenomicsBsc.CallOpts )
}

// OWNER is a free data retrieval call binding the contract method 0x117803e3.
//
// Solidity: function OWNER() view returns(address)
func (_TokenomicsBsc *TokenomicsBscCallerSession) OWNER() ( common.Address,  error) {
	return _TokenomicsBsc.Contract.OWNER(&_TokenomicsBsc.CallOpts )
}

// TOKENNAME is a free data retrieval call binding the contract method 0x18821400.
//
// Solidity: function TOKEN_NAME() view returns(string)
func (_TokenomicsBsc *TokenomicsBscCaller) TOKENNAME(opts *bind.CallOpts ) (string, error) {
	var out []interface{}
	err := _TokenomicsBsc.contract.Call(opts, &out, "TOKEN_NAME" )

	if err != nil {
		return *new(string),  err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0,  err

}

// TOKENNAME is a free data retrieval call binding the contract method 0x18821400.
//
// Solidity: function TOKEN_NAME() view returns(string)
func (_TokenomicsBsc *TokenomicsBscSession) TOKENNAME() ( string,  error) {
	return _TokenomicsBsc.Contract.TOKENNAME(&_TokenomicsBsc.CallOpts )
}

// TOKENNAME is a free data retrieval call binding the contract method 0x18821400.
//
// Solidity: function TOKEN_NAME() view returns(string)
func (_TokenomicsBsc *TokenomicsBscCallerSession) TOKENNAME() ( string,  error) {
	return _TokenomicsBsc.Contract.TOKENNAME(&_TokenomicsBsc.CallOpts )
}

// TOKENSYMBOL is a free data retrieval call binding the contract method 0x2a905318.
//
// Solidity: function TOKEN_SYMBOL() view returns(string)
func (_TokenomicsBsc *TokenomicsBscCaller) TOKENSYMBOL(opts *bind.CallOpts ) (string, error) {
	var out []interface{}
	err := _TokenomicsBsc.contract.Call(opts, &out, "TOKEN_SYMBOL" )

	if err != nil {
		return *new(string),  err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0,  err

}

// TOKENSYMBOL is a free data retrieval call binding the contract method 0x2a905318.
//
// Solidity: function TOKEN_SYMBOL() view returns(string)
func (_TokenomicsBsc *TokenomicsBscSession) TOKENSYMBOL() ( string,  error) {
	return _TokenomicsBsc.Contract.TOKENSYMBOL(&_TokenomicsBsc.CallOpts )
}

// TOKENSYMBOL is a free data retrieval call binding the contract method 0x2a905318.
//
// Solidity: function TOKEN_SYMBOL() view returns(string)
func (_TokenomicsBsc *TokenomicsBscCallerSession) TOKENSYMBOL() ( string,  error) {
	return _TokenomicsBsc.Contract.TOKENSYMBOL(&_TokenomicsBsc.CallOpts )
}

// TOTALSUPPLY is a free data retrieval call binding the contract method 0x902d55a5.
//
// Solidity: function TOTAL_SUPPLY() view returns(uint256)
func (_TokenomicsBsc *TokenomicsBscCaller) TOTALSUPPLY(opts *bind.CallOpts ) (*big.Int, error) {
	var out []interface{}
	err := _TokenomicsBsc.contract.Call(opts, &out, "TOTAL_SUPPLY" )

	if err != nil {
		return *new(*big.Int),  err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0,  err

}

// TOTALSUPPLY is a free data retrieval call binding the contract method 0x902d55a5.
//
// Solidity: function TOTAL_SUPPLY() view returns(uint256)
func (_TokenomicsBsc *TokenomicsBscSession) TOTALSUPPLY() ( *big.Int,  error) {
	return _TokenomicsBsc.Contract.TOTALSUPPLY(&_TokenomicsBsc.CallOpts )
}

// TOTALSUPPLY is a free data retrieval call binding the contract method 0x902d55a5.
//
// Solidity: function TOTAL_SUPPLY() view returns(uint256)
func (_TokenomicsBsc *TokenomicsBscCallerSession) TOTALSUPPLY() ( *big.Int,  error) {
	return _TokenomicsBsc.Contract.TOTALSUPPLY(&_TokenomicsBsc.CallOpts )
}



// Allowance is a paid mutator transaction binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) returns(uint256)
func (_TokenomicsBsc *TokenomicsBscTransactor) Allowance(opts *bind.TransactOpts , owner common.Address , spender common.Address ) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "allowance" , owner, spender)
}

// Allowance is a paid mutator transaction binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) returns(uint256)
func (_TokenomicsBsc *TokenomicsBscSession) Allowance( owner common.Address , spender common.Address ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.Allowance(&_TokenomicsBsc.TransactOpts , owner, spender)
}

// Allowance is a paid mutator transaction binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) returns(uint256)
func (_TokenomicsBsc *TokenomicsBscTransactorSession) Allowance( owner common.Address , spender common.Address ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.Allowance(&_TokenomicsBsc.TransactOpts , owner, spender)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenomicsBsc *TokenomicsBscTransactor) Approve(opts *bind.TransactOpts , spender common.Address , amount *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "approve" , spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenomicsBsc *TokenomicsBscSession) Approve( spender common.Address , amount *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.Approve(&_TokenomicsBsc.TransactOpts , spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenomicsBsc *TokenomicsBscTransactorSession) Approve( spender common.Address , amount *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.Approve(&_TokenomicsBsc.TransactOpts , spender, amount)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) returns(uint256)
func (_TokenomicsBsc *TokenomicsBscTransactor) BalanceOf(opts *bind.TransactOpts , account common.Address ) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "balanceOf" , account)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) returns(uint256)
func (_TokenomicsBsc *TokenomicsBscSession) BalanceOf( account common.Address ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.BalanceOf(&_TokenomicsBsc.TransactOpts , account)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) returns(uint256)
func (_TokenomicsBsc *TokenomicsBscTransactorSession) BalanceOf( account common.Address ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.BalanceOf(&_TokenomicsBsc.TransactOpts , account)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_TokenomicsBsc *TokenomicsBscTransactor) Burn(opts *bind.TransactOpts , account common.Address , amount *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "burn" , account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_TokenomicsBsc *TokenomicsBscSession) Burn( account common.Address , amount *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.Burn(&_TokenomicsBsc.TransactOpts , account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_TokenomicsBsc *TokenomicsBscTransactorSession) Burn( account common.Address , amount *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.Burn(&_TokenomicsBsc.TransactOpts , account, amount)
}

// CheckContractBalance is a paid mutator transaction binding the contract method 0x50312c9e.
//
// Solidity: function checkContractBalance() returns(uint256)
func (_TokenomicsBsc *TokenomicsBscTransactor) CheckContractBalance(opts *bind.TransactOpts ) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "checkContractBalance" )
}

// CheckContractBalance is a paid mutator transaction binding the contract method 0x50312c9e.
//
// Solidity: function checkContractBalance() returns(uint256)
func (_TokenomicsBsc *TokenomicsBscSession) CheckContractBalance() (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.CheckContractBalance(&_TokenomicsBsc.TransactOpts )
}

// CheckContractBalance is a paid mutator transaction binding the contract method 0x50312c9e.
//
// Solidity: function checkContractBalance() returns(uint256)
func (_TokenomicsBsc *TokenomicsBscTransactorSession) CheckContractBalance() (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.CheckContractBalance(&_TokenomicsBsc.TransactOpts )
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TokenomicsBsc *TokenomicsBscTransactor) DecreaseAllowance(opts *bind.TransactOpts , spender common.Address , subtractedValue *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "decreaseAllowance" , spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TokenomicsBsc *TokenomicsBscSession) DecreaseAllowance( spender common.Address , subtractedValue *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.DecreaseAllowance(&_TokenomicsBsc.TransactOpts , spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TokenomicsBsc *TokenomicsBscTransactorSession) DecreaseAllowance( spender common.Address , subtractedValue *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.DecreaseAllowance(&_TokenomicsBsc.TransactOpts , spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TokenomicsBsc *TokenomicsBscTransactor) IncreaseAllowance(opts *bind.TransactOpts , spender common.Address , addedValue *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "increaseAllowance" , spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TokenomicsBsc *TokenomicsBscSession) IncreaseAllowance( spender common.Address , addedValue *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.IncreaseAllowance(&_TokenomicsBsc.TransactOpts , spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TokenomicsBsc *TokenomicsBscTransactorSession) IncreaseAllowance( spender common.Address , addedValue *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.IncreaseAllowance(&_TokenomicsBsc.TransactOpts , spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_TokenomicsBsc *TokenomicsBscTransactor) Mint(opts *bind.TransactOpts , to common.Address , amount *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "mint" , to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_TokenomicsBsc *TokenomicsBscSession) Mint( to common.Address , amount *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.Mint(&_TokenomicsBsc.TransactOpts , to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_TokenomicsBsc *TokenomicsBscTransactorSession) Mint( to common.Address , amount *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.Mint(&_TokenomicsBsc.TransactOpts , to, amount)
}

// SendAllEther is a paid mutator transaction binding the contract method 0x085b3b00.
//
// Solidity: function sendAllEther(address to) returns()
func (_TokenomicsBsc *TokenomicsBscTransactor) SendAllEther(opts *bind.TransactOpts , to common.Address ) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "sendAllEther" , to)
}

// SendAllEther is a paid mutator transaction binding the contract method 0x085b3b00.
//
// Solidity: function sendAllEther(address to) returns()
func (_TokenomicsBsc *TokenomicsBscSession) SendAllEther( to common.Address ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.SendAllEther(&_TokenomicsBsc.TransactOpts , to)
}

// SendAllEther is a paid mutator transaction binding the contract method 0x085b3b00.
//
// Solidity: function sendAllEther(address to) returns()
func (_TokenomicsBsc *TokenomicsBscTransactorSession) SendAllEther( to common.Address ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.SendAllEther(&_TokenomicsBsc.TransactOpts , to)
}

// SendEther is a paid mutator transaction binding the contract method 0xc1756a2c.
//
// Solidity: function sendEther(address to, uint256 amount) returns()
func (_TokenomicsBsc *TokenomicsBscTransactor) SendEther(opts *bind.TransactOpts , to common.Address , amount *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "sendEther" , to, amount)
}

// SendEther is a paid mutator transaction binding the contract method 0xc1756a2c.
//
// Solidity: function sendEther(address to, uint256 amount) returns()
func (_TokenomicsBsc *TokenomicsBscSession) SendEther( to common.Address , amount *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.SendEther(&_TokenomicsBsc.TransactOpts , to, amount)
}

// SendEther is a paid mutator transaction binding the contract method 0xc1756a2c.
//
// Solidity: function sendEther(address to, uint256 amount) returns()
func (_TokenomicsBsc *TokenomicsBscTransactorSession) SendEther( to common.Address , amount *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.SendEther(&_TokenomicsBsc.TransactOpts , to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TokenomicsBsc *TokenomicsBscTransactor) Transfer(opts *bind.TransactOpts , recipient common.Address , amount *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "transfer" , recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TokenomicsBsc *TokenomicsBscSession) Transfer( recipient common.Address , amount *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.Transfer(&_TokenomicsBsc.TransactOpts , recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TokenomicsBsc *TokenomicsBscTransactorSession) Transfer( recipient common.Address , amount *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.Transfer(&_TokenomicsBsc.TransactOpts , recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TokenomicsBsc *TokenomicsBscTransactor) TransferFrom(opts *bind.TransactOpts , sender common.Address , recipient common.Address , amount *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "transferFrom" , sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TokenomicsBsc *TokenomicsBscSession) TransferFrom( sender common.Address , recipient common.Address , amount *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.TransferFrom(&_TokenomicsBsc.TransactOpts , sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TokenomicsBsc *TokenomicsBscTransactorSession) TransferFrom( sender common.Address , recipient common.Address , amount *big.Int ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.TransferFrom(&_TokenomicsBsc.TransactOpts , sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenomicsBsc *TokenomicsBscTransactor) TransferOwnership(opts *bind.TransactOpts , newOwner common.Address ) (*types.Transaction, error) {
	return _TokenomicsBsc.contract.Transact(opts, "transferOwnership" , newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenomicsBsc *TokenomicsBscSession) TransferOwnership( newOwner common.Address ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.TransferOwnership(&_TokenomicsBsc.TransactOpts , newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenomicsBsc *TokenomicsBscTransactorSession) TransferOwnership( newOwner common.Address ) (*types.Transaction, error) {
	return _TokenomicsBsc.Contract.TransferOwnership(&_TokenomicsBsc.TransactOpts , newOwner)
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
	if (it.fail != nil) {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if (it.done) {
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
	Owner common.Address;
	Spender common.Address;
	Value *big.Int;
	Raw types.Log // Blockchain specific contextual infos
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
	if (it.fail != nil) {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if (it.done) {
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
	OldOwner common.Address;
	NewOwner common.Address;
	Raw types.Log // Blockchain specific contextual infos
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
	if (it.fail != nil) {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if (it.done) {
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
	From common.Address;
	Amount *big.Int;
	CurrentBal *big.Int;
	GasLeft *big.Int;
	Raw types.Log // Blockchain specific contextual infos
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
	if (it.fail != nil) {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if (it.done) {
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
	From common.Address;
	To common.Address;
	Value *big.Int;
	Raw types.Log // Blockchain specific contextual infos
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



