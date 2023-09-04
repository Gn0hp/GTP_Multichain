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
)

// TokenomicsMetaData contains all meta data concerning the Tokenomics contract.
var TokenomicsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentAllowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestedDecrease\",\"type\":\"uint256\"}],\"name\":\"ERC20FailedDecreaseAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DECIMALS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b506040516200205d3803806200205d833981810160405281019062000036919062000326565b62000046620000d760201b60201c565b82828160039081620000599190620005e0565b5080600290816200006b9190620005e0565b506200007c620000d760201b60201c565b60055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050620000ce81620000de60201b60201c565b505050620006c4565b5f33905090565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6200020282620001ba565b810181811067ffffffffffffffff82111715620002245762000223620001ca565b5b80604052505050565b5f62000238620001a1565b9050620002468282620001f7565b919050565b5f67ffffffffffffffff821115620002685762000267620001ca565b5b6200027382620001ba565b9050602081019050919050565b5f5b838110156200029f57808201518184015260208101905062000282565b5f8484015250505050565b5f620002c0620002ba846200024b565b6200022d565b905082815260208101848484011115620002df57620002de620001b6565b5b620002ec84828562000280565b509392505050565b5f82601f8301126200030b576200030a620001b2565b5b81516200031d848260208601620002aa565b91505092915050565b5f80604083850312156200033f576200033e620001aa565b5b5f83015167ffffffffffffffff8111156200035f576200035e620001ae565b5b6200036d85828601620002f4565b925050602083015167ffffffffffffffff811115620003915762000390620001ae565b5b6200039f85828601620002f4565b9150509250929050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680620003f857607f821691505b6020821081036200040e576200040d620003b3565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620004727fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000435565b6200047e868362000435565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f620004c8620004c2620004bc8462000496565b6200049f565b62000496565b9050919050565b5f819050919050565b620004e383620004a8565b620004fb620004f282620004cf565b84845462000441565b825550505050565b5f90565b6200051162000503565b6200051e818484620004d8565b505050565b5b818110156200054557620005395f8262000507565b60018101905062000524565b5050565b601f82111562000594576200055e8162000414565b620005698462000426565b8101602085101562000579578190505b62000591620005888562000426565b83018262000523565b50505b505050565b5f82821c905092915050565b5f620005b65f198460080262000599565b1980831691505092915050565b5f620005d08383620005a5565b9150826002028217905092915050565b620005eb82620003a9565b67ffffffffffffffff811115620006075762000606620001ca565b5b620006138254620003e0565b6200062082828562000549565b5f60209050601f83116001811462000656575f841562000641578287015190505b6200064d8582620005c3565b865550620006bc565b601f198416620006668662000414565b5f5b828110156200068f5784890151825560018201915060208501945060208101905062000668565b86831015620006af5784890151620006ab601f891682620005a5565b8355505b6001600288020188555050505b505050505050565b61198b80620006d25f395ff3fe608060405234801561000f575f80fd5b50600436106100fe575f3560e01c8063715018a611610095578063a457c2d711610064578063a457c2d71461029a578063a9059cbb146102ca578063dd62ed3e146102fa578063f2fde38b1461032a576100fe565b8063715018a6146102385780638da5cb5b1461024257806395d89b41146102605780639dc29fac1461027e576100fe565b80632e0f2625116100d15780632e0f26251461019e57806339509351146101bc57806340c10f19146101ec57806370a0823114610208576100fe565b806306fdde0314610102578063095ea7b31461012057806318160ddd1461015057806323b872dd1461016e575b5f80fd5b61010a610346565b6040516101179190611215565b60405180910390f35b61013a600480360381019061013591906112c6565b6103d6565b604051610147919061131e565b60405180910390f35b6101586103f8565b6040516101659190611346565b60405180910390f35b6101886004803603810190610183919061135f565b610401565b604051610195919061131e565b60405180910390f35b6101a661042f565b6040516101b391906113ca565b60405180910390f35b6101d660048036038101906101d191906112c6565b610434565b6040516101e3919061131e565b60405180910390f35b610206600480360381019061020191906112c6565b6104db565b005b610222600480360381019061021d91906113e3565b6104f1565b60405161022f9190611346565b60405180910390f35b610240610536565b005b61024a610549565b604051610257919061141d565b60405180910390f35b610268610571565b6040516102759190611215565b60405180910390f35b610298600480360381019061029391906112c6565b610601565b005b6102b460048036038101906102af91906112c6565b610617565b6040516102c1919061131e565b60405180910390f35b6102e460048036038101906102df91906112c6565b6106fd565b6040516102f1919061131e565b60405180910390f35b610314600480360381019061030f9190611436565b61071f565b6040516103219190611346565b60405180910390f35b610344600480360381019061033f91906113e3565b6107a1565b005b606060038054610355906114a1565b80601f0160208091040260200160405190810160405280929190818152602001828054610381906114a1565b80156103cc5780601f106103a3576101008083540402835291602001916103cc565b820191905f5260205f20905b8154815290600101906020018083116103af57829003601f168201915b5050505050905090565b5f806103e0610825565b90506103ed81858561082c565b600191505092915050565b5f600454905090565b5f8061040b610825565b905061041885828561083e565b6104238585856108d0565b60019150509392505050565b601281565b5f6104d1610440610825565b848460015f61044d610825565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546104cc91906114fe565b61082c565b6001905092915050565b6104e3610b45565b6104ed8282610bcc565b5050565b5f805f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b61053e610b45565b6105475f610d23565b565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606060028054610580906114a1565b80601f01602080910402602001604051908101604052809291908181526020018280546105ac906114a1565b80156105f75780601f106105ce576101008083540402835291602001916105f7565b820191905f5260205f20905b8154815290600101906020018083116105da57829003601f168201915b5050505050905090565b610609610b45565b6106138282610de6565b5050565b5f8060015f610624610825565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050828110156106de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d5906115a1565b60405180910390fd5b6106f26106e9610825565b8585840361082c565b600191505092915050565b5f80610707610825565b90506107148185856108d0565b600191505092915050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b6107a9610b45565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610819575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610810919061141d565b60405180910390fd5b61082281610d23565b50565b5f33905090565b6108398383836001610fb2565b505050565b5f610849848461071f565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146108ca57818110156108bb578281836040517ffb8f41b20000000000000000000000000000000000000000000000000000000081526004016108b2939291906115bf565b60405180910390fd5b6108c984848484035f610fb2565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361093e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093590611664565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036109ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109a3906116f2565b60405180910390fd5b6109b7838383611181565b5f805f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610a3a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a3190611780565b60405180910390fd5b8181035f808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550815f808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610ac891906114fe565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610b2c9190611346565b60405180910390a3610b3f848484611186565b50505050565b610b4d610825565b73ffffffffffffffffffffffffffffffffffffffff16610b6b610549565b73ffffffffffffffffffffffffffffffffffffffff1614610bca57610b8e610825565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401610bc1919061141d565b60405180910390fd5b565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610c3a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c31906117e8565b60405180910390fd5b610c455f8383611181565b805f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610c9091906114fe565b925050819055508060045f828254610ca891906114fe565b925050819055508173ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610d0c9190611346565b60405180910390a3610d1f5f8383611186565b5050565b5f60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610e54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4b90611876565b60405180910390fd5b610e5f825f83611181565b5f805f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610ee2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ed990611904565b60405180910390fd5b8181035f808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508160045f828254610f369190611922565b925050819055505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610f9a9190611346565b60405180910390a3610fad835f84611186565b505050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603611022575f6040517fe602df05000000000000000000000000000000000000000000000000000000008152600401611019919061141d565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603611092575f6040517f94280d62000000000000000000000000000000000000000000000000000000008152600401611089919061141d565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550801561117b578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516111729190611346565b60405180910390a35b50505050565b505050565b505050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156111c25780820151818401526020810190506111a7565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6111e78261118b565b6111f18185611195565b93506112018185602086016111a5565b61120a816111cd565b840191505092915050565b5f6020820190508181035f83015261122d81846111dd565b905092915050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61126282611239565b9050919050565b61127281611258565b811461127c575f80fd5b50565b5f8135905061128d81611269565b92915050565b5f819050919050565b6112a581611293565b81146112af575f80fd5b50565b5f813590506112c08161129c565b92915050565b5f80604083850312156112dc576112db611235565b5b5f6112e98582860161127f565b92505060206112fa858286016112b2565b9150509250929050565b5f8115159050919050565b61131881611304565b82525050565b5f6020820190506113315f83018461130f565b92915050565b61134081611293565b82525050565b5f6020820190506113595f830184611337565b92915050565b5f805f6060848603121561137657611375611235565b5b5f6113838682870161127f565b93505060206113948682870161127f565b92505060406113a5868287016112b2565b9150509250925092565b5f60ff82169050919050565b6113c4816113af565b82525050565b5f6020820190506113dd5f8301846113bb565b92915050565b5f602082840312156113f8576113f7611235565b5b5f6114058482850161127f565b91505092915050565b61141781611258565b82525050565b5f6020820190506114305f83018461140e565b92915050565b5f806040838503121561144c5761144b611235565b5b5f6114598582860161127f565b925050602061146a8582860161127f565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806114b857607f821691505b6020821081036114cb576114ca611474565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61150882611293565b915061151383611293565b925082820190508082111561152b5761152a6114d1565b5b92915050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f775f8201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b5f61158b602583611195565b915061159682611531565b604082019050919050565b5f6020820190508181035f8301526115b88161157f565b9050919050565b5f6060820190506115d25f83018661140e565b6115df6020830185611337565b6115ec6040830184611337565b949350505050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f2061645f8201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b5f61164e602583611195565b9150611659826115f4565b604082019050919050565b5f6020820190508181035f83015261167b81611642565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f20616464725f8201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b5f6116dc602383611195565b91506116e782611682565b604082019050919050565b5f6020820190508181035f830152611709816116d0565b9050919050565b7f45524332303a207472616e7366657220616d6f756e74206578636565647320625f8201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b5f61176a602683611195565b915061177582611710565b604082019050919050565b5f6020820190508181035f8301526117978161175e565b9050919050565b7f45524332303a206d696e7420746f20746865207a65726f2061646472657373005f82015250565b5f6117d2601f83611195565b91506117dd8261179e565b602082019050919050565b5f6020820190508181035f8301526117ff816117c6565b9050919050565b7f45524332303a206275726e2066726f6d20746865207a65726f206164647265735f8201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b5f611860602183611195565b915061186b82611806565b604082019050919050565b5f6020820190508181035f83015261188d81611854565b9050919050565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e5f8201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b5f6118ee602283611195565b91506118f982611894565b604082019050919050565b5f6020820190508181035f83015261191b816118e2565b9050919050565b5f61192c82611293565b915061193783611293565b925082820390508181111561194f5761194e6114d1565b5b9291505056fea264697066735822122014f350909c18a90de4981ec8edac912219e4a1b6b2cbea696406e90fe9c2c4c064736f6c63430008150033",
}

// TokenomicsABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenomicsMetaData.ABI instead.
var TokenomicsABI = TokenomicsMetaData.ABI

// TokenomicsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenomicsMetaData.Bin instead.
var TokenomicsBin = TokenomicsMetaData.Bin

// DeployTokenomics deploys a new Ethereum contract, binding an instance of Tokenomics to it.
func DeployTokenomics(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string) (common.Address, *types.Transaction, *Tokenomics, error) {
	parsed, err := TokenomicsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenomicsBin), backend, _name, _symbol)
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
	parsed, err := abi.JSON(strings.NewReader(TokenomicsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_Tokenomics *TokenomicsCaller) DECIMALS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenomics.contract.Call(opts, &out, "DECIMALS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_Tokenomics *TokenomicsSession) DECIMALS() (uint8, error) {
	return _Tokenomics.Contract.DECIMALS(&_Tokenomics.CallOpts)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_Tokenomics *TokenomicsCallerSession) DECIMALS() (uint8, error) {
	return _Tokenomics.Contract.DECIMALS(&_Tokenomics.CallOpts)
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
