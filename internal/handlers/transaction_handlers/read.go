package transaction_handlers

import (
	"context"
	"encoding/json"
	tokenomics "eth_bsc_multichain/contracts"
	"eth_bsc_multichain/pkg/middlewares"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"math/big"
	"net/http"
	"os"
)

var (
	SystemWalletPrv string
	ClientEth       *ethclient.Client
	ClientBsc       *ethclient.Client
	Token           tokenomics.Tokenomics
	SystemWallet    common.Address
)

func init() {
	_ = godotenv.Load()

	SystemWalletPrv = os.Getenv("METAMASK_PRIVATE_KEY")
	SystemWallet = common.HexToAddress(os.Getenv("METAMASK_ADDRESS"))
	fmt.Println("getenv: ", os.Getenv("NODE_RPC_URL"))
	ClientEth, err := ethclient.Dial(os.Getenv("NODE_RPC_URL"))
	if err != nil {
		panic(err)
	}
	addressEthToken := common.HexToAddress(os.Getenv("ETH_VYPER_CONTRACT_ADDRESS"))
	addressBscToken := common.HexToAddress(os.Getenv("BSC_SOL_CONTRACT_ADDRESS"))
	if addressEthToken == common.HexToAddress("") || addressBscToken == common.HexToAddress("") {
		panic(err)
	}

	ClientBsc, err = ethclient.Dial(os.Getenv("NODE_RPC_URL_BSC"))
	Token = tokenomics.New(addressEthToken, addressBscToken,
		ClientEth,
		ClientBsc)
	if err != nil {
		panic(err)
	}
}

// Query  {string} address ex: ?address=0x123
func (t TransactionHandler) getBalanceBsc(address string, ctx context.Context) (*types.Transaction, error) {
	chainId, err := ClientBsc.ChainID(ctx)
	if err != nil {
		t.logger.Error(fmt.Sprintf("[GetBalanceBsc] Error getting chainId: %v", err))
		return nil, err
	}
	accountBound, _ := TransactOptsAccountBinding(chainId)
	balance, err := Token.TokenBsc().BalanceOf(accountBound, common.HexToAddress(address))
	if err != nil {
		t.logger.Error(fmt.Sprintf("[GetBalanceBsc] Error getting balance: %v", err))
		return nil, err
	}
	return balance, nil
}
func TransactOptsAccountBinding(chainId *big.Int) (*bind.TransactOpts, error) {
	toEcdsa, err := crypto.HexToECDSA(SystemWalletPrv)
	if err != nil {
		return nil, err
	}
	user, err := bind.NewKeyedTransactorWithChainID(toEcdsa, chainId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (t TransactionHandler) getBalanceEth(address string) (*big.Int, error) {
	callOpts := &bind.CallOpts{
		From: SystemWallet,
	}
	balance, err := Token.TokenEth().BalanceOf(callOpts, common.HexToAddress(address))
	if err != nil {
		t.logger.Error(fmt.Sprintf("[GetBalanceBsc] Error getting balance: %v", err))
		return nil, err
	}
	return balance, nil
}

func (t TransactionHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	userAddress := r.URL.Query().Get("address")
	if userAddress == "" {
		middlewares.WriteResponse(&w, http.StatusBadRequest, "address is required")
		return
	}
	ctx := context.TODO()
	balanceBsc, err := t.getBalanceBsc(userAddress, ctx)
	if err != nil {
		middlewares.WriteResponse(&w, http.StatusBadRequest, "error getting balance")
		return
	}
	balanceEth, err := t.getBalanceEth(userAddress)
	if err != nil {
		middlewares.WriteResponse(&w, http.StatusBadRequest, "error getting balance")
		return
	}
	body, _ := json.Marshal(map[string]interface{}{
		"balanceBsc": balanceBsc,
		"balanceEth": balanceEth,
	})
	middlewares.WriteResponse(&w, http.StatusOK, body)
}
