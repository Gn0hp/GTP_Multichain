// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenomics

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

// TokenomicsMetaData contains all meta data concerning the Tokenomics contract.
var TokenomicsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentAllowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestedDecrease\",\"type\":\"uint256\"}],\"name\":\"ERC20FailedDecreaseAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526012600660146101000a81548161ffff021916908361ffff1602179055503480156200002e575f80fd5b506040516200238738038062002387833981810160405281019062000054919062000521565b620000646200012060201b60201c565b8383838260039081620000789190620007e6565b5081600290816200008a9190620007e6565b5080600481905550620000a26200012060201b60201c565b60055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050620000f5816200012760201b60201c565b50620001176200010a6200012060201b60201c565b82620001ea60201b60201c565b505050620009db565b5f33905090565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036200025b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002529062000928565b60405180910390fd5b6200026e5f83836200035a60201b60201c565b805f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254620002bb919062000975565b925050819055508060045f828254620002d5919062000975565b925050819055508173ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516200033b9190620009c0565b60405180910390a3620003565f83836200035f60201b60201c565b5050565b505050565b505050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b620003c5826200037d565b810181811067ffffffffffffffff82111715620003e757620003e66200038d565b5b80604052505050565b5f620003fb62000364565b9050620004098282620003ba565b919050565b5f67ffffffffffffffff8211156200042b576200042a6200038d565b5b62000436826200037d565b9050602081019050919050565b5f5b838110156200046257808201518184015260208101905062000445565b5f8484015250505050565b5f620004836200047d846200040e565b620003f0565b905082815260208101848484011115620004a257620004a162000379565b5b620004af84828562000443565b509392505050565b5f82601f830112620004ce57620004cd62000375565b5b8151620004e08482602086016200046d565b91505092915050565b5f819050919050565b620004fd81620004e9565b811462000508575f80fd5b50565b5f815190506200051b81620004f2565b92915050565b5f805f606084860312156200053b576200053a6200036d565b5b5f84015167ffffffffffffffff8111156200055b576200055a62000371565b5b6200056986828701620004b7565b935050602084015167ffffffffffffffff8111156200058d576200058c62000371565b5b6200059b86828701620004b7565b9250506040620005ae868287016200050b565b9150509250925092565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806200060757607f821691505b6020821081036200061d576200061c620005c2565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620006817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000644565b6200068d868362000644565b95508019841693508086168417925050509392505050565b5f819050919050565b5f620006ce620006c8620006c284620004e9565b620006a5565b620004e9565b9050919050565b5f819050919050565b620006e983620006ae565b62000701620006f882620006d5565b84845462000650565b825550505050565b5f90565b6200071762000709565b62000724818484620006de565b505050565b5b818110156200074b576200073f5f826200070d565b6001810190506200072a565b5050565b601f8211156200079a57620007648162000623565b6200076f8462000635565b810160208510156200077f578190505b620007976200078e8562000635565b83018262000729565b50505b505050565b5f82821c905092915050565b5f620007bc5f19846008026200079f565b1980831691505092915050565b5f620007d68383620007ab565b9150826002028217905092915050565b620007f182620005b8565b67ffffffffffffffff8111156200080d576200080c6200038d565b5b620008198254620005ef565b620008268282856200074f565b5f60209050601f8311600181146200085c575f841562000847578287015190505b620008538582620007c9565b865550620008c2565b601f1984166200086c8662000623565b5f5b8281101562000895578489015182556001820191506020850194506020810190506200086e565b86831015620008b55784890151620008b1601f891682620007ab565b8355505b6001600288020188555050505b505050505050565b5f82825260208201905092915050565b7f45524332303a206d696e7420746f20746865207a65726f2061646472657373005f82015250565b5f62000910601f83620008ca565b91506200091d82620008da565b602082019050919050565b5f6020820190508181035f830152620009418162000902565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6200098182620004e9565b91506200098e83620004e9565b9250828201905080821115620009a957620009a862000948565b5b92915050565b620009ba81620004e9565b82525050565b5f602082019050620009d55f830184620009af565b92915050565b61199e80620009e95f395ff3fe608060405234801561000f575f80fd5b50600436106100fe575f3560e01c8063715018a611610095578063a457c2d711610064578063a457c2d71461029a578063a9059cbb146102ca578063dd62ed3e146102fa578063f2fde38b1461032a576100fe565b8063715018a6146102385780638da5cb5b1461024257806395d89b41146102605780639dc29fac1461027e576100fe565b8063313ce567116100d1578063313ce5671461019e57806339509351146101bc57806340c10f19146101ec57806370a0823114610208576100fe565b806306fdde0314610102578063095ea7b31461012057806318160ddd1461015057806323b872dd1461016e575b5f80fd5b61010a610346565b6040516101179190611227565b60405180910390f35b61013a600480360381019061013591906112d8565b6103d6565b6040516101479190611330565b60405180910390f35b6101586103f8565b6040516101659190611358565b60405180910390f35b61018860048036038101906101839190611371565b610401565b6040516101959190611330565b60405180910390f35b6101a661042f565b6040516101b391906113dd565b60405180910390f35b6101d660048036038101906101d191906112d8565b610446565b6040516101e39190611330565b60405180910390f35b610206600480360381019061020191906112d8565b6104ed565b005b610222600480360381019061021d91906113f6565b610503565b60405161022f9190611358565b60405180910390f35b610240610548565b005b61024a61055b565b6040516102579190611430565b60405180910390f35b610268610583565b6040516102759190611227565b60405180910390f35b610298600480360381019061029391906112d8565b610613565b005b6102b460048036038101906102af91906112d8565b610629565b6040516102c19190611330565b60405180910390f35b6102e460048036038101906102df91906112d8565b61070f565b6040516102f19190611330565b60405180910390f35b610314600480360381019061030f9190611449565b610731565b6040516103219190611358565b60405180910390f35b610344600480360381019061033f91906113f6565b6107b3565b005b606060038054610355906114b4565b80601f0160208091040260200160405190810160405280929190818152602001828054610381906114b4565b80156103cc5780601f106103a3576101008083540402835291602001916103cc565b820191905f5260205f20905b8154815290600101906020018083116103af57829003601f168201915b5050505050905090565b5f806103e0610837565b90506103ed81858561083e565b600191505092915050565b5f600454905090565b5f8061040b610837565b9050610418858285610850565b6104238585856108e2565b60019150509392505050565b5f600660149054906101000a900461ffff16905090565b5f6104e3610452610837565b848460015f61045f610837565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546104de9190611511565b61083e565b6001905092915050565b6104f5610b57565b6104ff8282610bde565b5050565b5f805f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b610550610b57565b6105595f610d35565b565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606060028054610592906114b4565b80601f01602080910402602001604051908101604052809291908181526020018280546105be906114b4565b80156106095780601f106105e057610100808354040283529160200191610609565b820191905f5260205f20905b8154815290600101906020018083116105ec57829003601f168201915b5050505050905090565b61061b610b57565b6106258282610df8565b5050565b5f8060015f610636610837565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050828110156106f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e7906115b4565b60405180910390fd5b6107046106fb610837565b8585840361083e565b600191505092915050565b5f80610719610837565b90506107268185856108e2565b600191505092915050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b6107bb610b57565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361082b575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016108229190611430565b60405180910390fd5b61083481610d35565b50565b5f33905090565b61084b8383836001610fc4565b505050565b5f61085b8484610731565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146108dc57818110156108cd578281836040517ffb8f41b20000000000000000000000000000000000000000000000000000000081526004016108c4939291906115d2565b60405180910390fd5b6108db84848484035f610fc4565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610950576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161094790611677565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036109be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109b590611705565b60405180910390fd5b6109c9838383611193565b5f805f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610a4c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4390611793565b60405180910390fd5b8181035f808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550815f808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610ada9190611511565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610b3e9190611358565b60405180910390a3610b51848484611198565b50505050565b610b5f610837565b73ffffffffffffffffffffffffffffffffffffffff16610b7d61055b565b73ffffffffffffffffffffffffffffffffffffffff1614610bdc57610ba0610837565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401610bd39190611430565b60405180910390fd5b565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610c4c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c43906117fb565b60405180910390fd5b610c575f8383611193565b805f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610ca29190611511565b925050819055508060045f828254610cba9190611511565b925050819055508173ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610d1e9190611358565b60405180910390a3610d315f8383611198565b5050565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610e66576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e5d90611889565b60405180910390fd5b610e71825f83611193565b5f805f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610ef4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eeb90611917565b60405180910390fd5b8181035f808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508160045f828254610f489190611935565b925050819055505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610fac9190611358565b60405180910390a3610fbf835f84611198565b505050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603611034575f6040517fe602df0500000000000000000000000000000000000000000000000000000000815260040161102b9190611430565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036110a4575f6040517f94280d6200000000000000000000000000000000000000000000000000000000815260040161109b9190611430565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550801561118d578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516111849190611358565b60405180910390a35b50505050565b505050565b505050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156111d45780820151818401526020810190506111b9565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6111f98261119d565b61120381856111a7565b93506112138185602086016111b7565b61121c816111df565b840191505092915050565b5f6020820190508181035f83015261123f81846111ef565b905092915050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6112748261124b565b9050919050565b6112848161126a565b811461128e575f80fd5b50565b5f8135905061129f8161127b565b92915050565b5f819050919050565b6112b7816112a5565b81146112c1575f80fd5b50565b5f813590506112d2816112ae565b92915050565b5f80604083850312156112ee576112ed611247565b5b5f6112fb85828601611291565b925050602061130c858286016112c4565b9150509250929050565b5f8115159050919050565b61132a81611316565b82525050565b5f6020820190506113435f830184611321565b92915050565b611352816112a5565b82525050565b5f60208201905061136b5f830184611349565b92915050565b5f805f6060848603121561138857611387611247565b5b5f61139586828701611291565b93505060206113a686828701611291565b92505060406113b7868287016112c4565b9150509250925092565b5f61ffff82169050919050565b6113d7816113c1565b82525050565b5f6020820190506113f05f8301846113ce565b92915050565b5f6020828403121561140b5761140a611247565b5b5f61141884828501611291565b91505092915050565b61142a8161126a565b82525050565b5f6020820190506114435f830184611421565b92915050565b5f806040838503121561145f5761145e611247565b5b5f61146c85828601611291565b925050602061147d85828601611291565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806114cb57607f821691505b6020821081036114de576114dd611487565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61151b826112a5565b9150611526836112a5565b925082820190508082111561153e5761153d6114e4565b5b92915050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f775f8201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b5f61159e6025836111a7565b91506115a982611544565b604082019050919050565b5f6020820190508181035f8301526115cb81611592565b9050919050565b5f6060820190506115e55f830186611421565b6115f26020830185611349565b6115ff6040830184611349565b949350505050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f2061645f8201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b5f6116616025836111a7565b915061166c82611607565b604082019050919050565b5f6020820190508181035f83015261168e81611655565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f20616464725f8201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b5f6116ef6023836111a7565b91506116fa82611695565b604082019050919050565b5f6020820190508181035f83015261171c816116e3565b9050919050565b7f45524332303a207472616e7366657220616d6f756e74206578636565647320625f8201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b5f61177d6026836111a7565b915061178882611723565b604082019050919050565b5f6020820190508181035f8301526117aa81611771565b9050919050565b7f45524332303a206d696e7420746f20746865207a65726f2061646472657373005f82015250565b5f6117e5601f836111a7565b91506117f0826117b1565b602082019050919050565b5f6020820190508181035f830152611812816117d9565b9050919050565b7f45524332303a206275726e2066726f6d20746865207a65726f206164647265735f8201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b5f6118736021836111a7565b915061187e82611819565b604082019050919050565b5f6020820190508181035f8301526118a081611867565b9050919050565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e5f8201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b5f6119016022836111a7565b915061190c826118a7565b604082019050919050565b5f6020820190508181035f83015261192e816118f5565b9050919050565b5f61193f826112a5565b915061194a836112a5565b9250828203905081811115611962576119616114e4565b5b9291505056fea2646970667358221220ee0e93f8c9aa92451a9eb2ed5fafc0a494a884573496eb55f734ee3a0ce6552064736f6c63430008150033",
}

// TokenomicsABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenomicsMetaData.ABI instead.
var TokenomicsABI = TokenomicsMetaData.ABI

// TokenomicsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenomicsMetaData.Bin instead.
var TokenomicsBin = TokenomicsMetaData.Bin

// DeployTokenomics deploys a new Ethereum contract, binding an instance of Tokenomics to it.
func DeployTokenomics(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, totalSupply_ *big.Int) (common.Address, *types.Transaction, *Tokenomics, error) {
	parsed, err := TokenomicsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenomicsBin), backend, _name, _symbol, totalSupply_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tokenomics{TokenomicsCaller: TokenomicsCaller{contract: contract}, TokenomicsTransactor: TokenomicsTransactor{contract: contract}, TokenomicsFilterer: TokenomicsFilterer{contract: contract}}, nil
}

// Tokenomics is an auto generated Go binding around an Ethereum contract.
type Tokenomics struct {
	TokenomicsCaller     // Read-only binding to the contract
	TokenomicsTransactor // Write-only binding to the contract
	TokenomicsFilterer   // Log filterer for contract events
}

// TokenomicsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenomicsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenomicsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenomicsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenomicsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenomicsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenomicsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenomicsSession struct {
	Contract     *Tokenomics       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenomicsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenomicsCallerSession struct {
	Contract *TokenomicsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TokenomicsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenomicsTransactorSession struct {
	Contract     *TokenomicsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TokenomicsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenomicsRaw struct {
	Contract *Tokenomics // Generic contract binding to access the raw methods on
}

// TokenomicsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenomicsCallerRaw struct {
	Contract *TokenomicsCaller // Generic read-only contract binding to access the raw methods on
}

// TokenomicsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenomicsTransactorRaw struct {
	Contract *TokenomicsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenomics creates a new instance of Tokenomics, bound to a specific deployed contract.
func NewTokenomics(address common.Address, backend bind.ContractBackend) (*Tokenomics, error) {
	contract, err := bindTokenomics(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokenomics{TokenomicsCaller: TokenomicsCaller{contract: contract}, TokenomicsTransactor: TokenomicsTransactor{contract: contract}, TokenomicsFilterer: TokenomicsFilterer{contract: contract}}, nil
}

// NewTokenomicsCaller creates a new read-only instance of Tokenomics, bound to a specific deployed contract.
func NewTokenomicsCaller(address common.Address, caller bind.ContractCaller) (*TokenomicsCaller, error) {
	contract, err := bindTokenomics(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenomicsCaller{contract: contract}, nil
}

// NewTokenomicsTransactor creates a new write-only instance of Tokenomics, bound to a specific deployed contract.
func NewTokenomicsTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenomicsTransactor, error) {
	contract, err := bindTokenomics(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenomicsTransactor{contract: contract}, nil
}

// NewTokenomicsFilterer creates a new log filterer instance of Tokenomics, bound to a specific deployed contract.
func NewTokenomicsFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenomicsFilterer, error) {
	contract, err := bindTokenomics(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenomicsFilterer{contract: contract}, nil
}

// bindTokenomics binds a generic wrapper to an already deployed contract.
func bindTokenomics(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenomicsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenomics *TokenomicsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenomics.Contract.TokenomicsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenomics *TokenomicsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenomics.Contract.TokenomicsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenomics *TokenomicsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenomics.Contract.TokenomicsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenomics *TokenomicsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenomics.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenomics *TokenomicsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenomics.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenomics *TokenomicsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenomics.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Tokenomics *TokenomicsCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tokenomics.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Tokenomics *TokenomicsSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Tokenomics.Contract.Allowance(&_Tokenomics.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Tokenomics *TokenomicsCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Tokenomics.Contract.Allowance(&_Tokenomics.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Tokenomics *TokenomicsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tokenomics.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Tokenomics *TokenomicsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Tokenomics.Contract.BalanceOf(&_Tokenomics.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Tokenomics *TokenomicsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Tokenomics.Contract.BalanceOf(&_Tokenomics.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint16)
func (_Tokenomics *TokenomicsCaller) Decimals(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Tokenomics.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint16)
func (_Tokenomics *TokenomicsSession) Decimals() (uint16, error) {
	return _Tokenomics.Contract.Decimals(&_Tokenomics.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint16)
func (_Tokenomics *TokenomicsCallerSession) Decimals() (uint16, error) {
	return _Tokenomics.Contract.Decimals(&_Tokenomics.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokenomics *TokenomicsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tokenomics.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokenomics *TokenomicsSession) Name() (string, error) {
	return _Tokenomics.Contract.Name(&_Tokenomics.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokenomics *TokenomicsCallerSession) Name() (string, error) {
	return _Tokenomics.Contract.Name(&_Tokenomics.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokenomics *TokenomicsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenomics.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokenomics *TokenomicsSession) Owner() (common.Address, error) {
	return _Tokenomics.Contract.Owner(&_Tokenomics.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokenomics *TokenomicsCallerSession) Owner() (common.Address, error) {
	return _Tokenomics.Contract.Owner(&_Tokenomics.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokenomics *TokenomicsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tokenomics.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokenomics *TokenomicsSession) Symbol() (string, error) {
	return _Tokenomics.Contract.Symbol(&_Tokenomics.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokenomics *TokenomicsCallerSession) Symbol() (string, error) {
	return _Tokenomics.Contract.Symbol(&_Tokenomics.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokenomics *TokenomicsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokenomics.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokenomics *TokenomicsSession) TotalSupply() (*big.Int, error) {
	return _Tokenomics.Contract.TotalSupply(&_Tokenomics.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokenomics *TokenomicsCallerSession) TotalSupply() (*big.Int, error) {
	return _Tokenomics.Contract.TotalSupply(&_Tokenomics.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Tokenomics *TokenomicsTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokenomics.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Tokenomics *TokenomicsSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokenomics.Contract.Approve(&_Tokenomics.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Tokenomics *TokenomicsTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokenomics.Contract.Approve(&_Tokenomics.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_Tokenomics *TokenomicsTransactor) Burn(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tokenomics.contract.Transact(opts, "burn", _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_Tokenomics *TokenomicsSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tokenomics.Contract.Burn(&_Tokenomics.TransactOpts, _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_Tokenomics *TokenomicsTransactorSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tokenomics.Contract.Burn(&_Tokenomics.TransactOpts, _from, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Tokenomics *TokenomicsTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Tokenomics.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Tokenomics *TokenomicsSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Tokenomics.Contract.DecreaseAllowance(&_Tokenomics.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Tokenomics *TokenomicsTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Tokenomics.Contract.DecreaseAllowance(&_Tokenomics.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Tokenomics *TokenomicsTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Tokenomics.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Tokenomics *TokenomicsSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Tokenomics.Contract.IncreaseAllowance(&_Tokenomics.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Tokenomics *TokenomicsTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Tokenomics.Contract.IncreaseAllowance(&_Tokenomics.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Tokenomics *TokenomicsTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tokenomics.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Tokenomics *TokenomicsSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tokenomics.Contract.Mint(&_Tokenomics.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Tokenomics *TokenomicsTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tokenomics.Contract.Mint(&_Tokenomics.TransactOpts, _to, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenomics *TokenomicsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenomics.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenomics *TokenomicsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tokenomics.Contract.RenounceOwnership(&_Tokenomics.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenomics *TokenomicsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tokenomics.Contract.RenounceOwnership(&_Tokenomics.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Tokenomics *TokenomicsTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokenomics.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Tokenomics *TokenomicsSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokenomics.Contract.Transfer(&_Tokenomics.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Tokenomics *TokenomicsTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokenomics.Contract.Transfer(&_Tokenomics.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Tokenomics *TokenomicsTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokenomics.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Tokenomics *TokenomicsSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokenomics.Contract.TransferFrom(&_Tokenomics.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Tokenomics *TokenomicsTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokenomics.Contract.TransferFrom(&_Tokenomics.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokenomics *TokenomicsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Tokenomics.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokenomics *TokenomicsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokenomics.Contract.TransferOwnership(&_Tokenomics.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokenomics *TokenomicsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokenomics.Contract.TransferOwnership(&_Tokenomics.TransactOpts, newOwner)
}

// TokenomicsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Tokenomics contract.
type TokenomicsApprovalIterator struct {
	Event *TokenomicsApproval // Event containing the contract specifics and raw log

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
func (it *TokenomicsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenomicsApproval)
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
		it.Event = new(TokenomicsApproval)
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
func (it *TokenomicsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenomicsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenomicsApproval represents a Approval event raised by the Tokenomics contract.
type TokenomicsApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tokenomics *TokenomicsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokenomicsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Tokenomics.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenomicsApprovalIterator{contract: _Tokenomics.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tokenomics *TokenomicsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenomicsApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Tokenomics.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenomicsApproval)
				if err := _Tokenomics.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tokenomics *TokenomicsFilterer) ParseApproval(log types.Log) (*TokenomicsApproval, error) {
	event := new(TokenomicsApproval)
	if err := _Tokenomics.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenomicsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Tokenomics contract.
type TokenomicsOwnershipTransferredIterator struct {
	Event *TokenomicsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenomicsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenomicsOwnershipTransferred)
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
		it.Event = new(TokenomicsOwnershipTransferred)
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
func (it *TokenomicsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenomicsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenomicsOwnershipTransferred represents a OwnershipTransferred event raised by the Tokenomics contract.
type TokenomicsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokenomics *TokenomicsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenomicsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tokenomics.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenomicsOwnershipTransferredIterator{contract: _Tokenomics.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokenomics *TokenomicsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenomicsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tokenomics.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenomicsOwnershipTransferred)
				if err := _Tokenomics.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokenomics *TokenomicsFilterer) ParseOwnershipTransferred(log types.Log) (*TokenomicsOwnershipTransferred, error) {
	event := new(TokenomicsOwnershipTransferred)
	if err := _Tokenomics.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenomicsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Tokenomics contract.
type TokenomicsTransferIterator struct {
	Event *TokenomicsTransfer // Event containing the contract specifics and raw log

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
func (it *TokenomicsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenomicsTransfer)
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
		it.Event = new(TokenomicsTransfer)
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
func (it *TokenomicsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenomicsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenomicsTransfer represents a Transfer event raised by the Tokenomics contract.
type TokenomicsTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tokenomics *TokenomicsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenomicsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Tokenomics.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenomicsTransferIterator{contract: _Tokenomics.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tokenomics *TokenomicsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenomicsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Tokenomics.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenomicsTransfer)
				if err := _Tokenomics.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tokenomics *TokenomicsFilterer) ParseTransfer(log types.Log) (*TokenomicsTransfer, error) {
	event := new(TokenomicsTransfer)
	if err := _Tokenomics.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
