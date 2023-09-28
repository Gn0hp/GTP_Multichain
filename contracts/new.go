package tokenomics

import (
	logger2 "eth_bsc_multichain/pkg/logger"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"logur.dev/logur"
)

type Tokenomics interface {
	TokenEth() *TokenomicsEth
	TokenBsc() *TokenomicsBsc
}

type impl struct {
	logger   logur.LoggerFacade
	tokenEth *TokenomicsEth
	tokenBsc *TokenomicsBsc
}

func (i impl) TokenEth() *TokenomicsEth {
	return i.tokenEth
}

func (i impl) TokenBsc() *TokenomicsBsc {
	return i.tokenBsc
}

func New(addressEth, addressBsc common.Address, backendEth, backendBsc bind.ContractBackend) Tokenomics {
	logger := logger2.NewLogger(logger2.DefaultConfig())
	tE, err := NewTokenomicsEth(addressEth, backendEth)
	if err != nil {
		logger.Error(fmt.Sprintf("[NewContractInstance] Error getting TokenomicsEth: %v", err))
		panic(err)
	}
	tB, err := NewTokenomicsBsc(addressBsc, backendBsc)
	if err != nil {
		logger.Error(fmt.Sprintf("[NewContractInstance] Error getting TokenomicsBsc: %v", err))
		panic(err)
	}
	return &impl{
		logger:   logger,
		tokenEth: tE,
		tokenBsc: tB,
	}

}
