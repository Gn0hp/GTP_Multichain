package main

import (
	"testing"
)

const (
	SepoliaChainId    = 11155111
	BscTestnetChainId = 97
	GenesisBalance    = 1337000000000000000
)

// TestDeployEthToken test the contract gets deployed correctly
func TestDeployEthToken(t *testing.T) {
	//key, _ := crypto.GenerateKey()
	//auth, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(SepoliaChainId))
	//alloc := make(core.GenesisAlloc)
	//
	//alloc[auth.From] = core.GenesisAccount{
	//	Balance: big.NewInt(GenesisBalance)}
	//blockchain := backends.NewSimulatedBackend(alloc, 8000000)
	//address, _, _, err := tokenomics.(auth, blockchain, "GToken", "GTP")
	//blockchain.Commit() // commit all pending transactions
	//
	//if err != nil {
	//	t.Fatalf("Failed to deploy new token contract: %v", err)
	//}
	//if len(address.Bytes()) == 0 {
	//	t.Error("Received empty address .Expected a valid deployment address")
	//}
}

//func TestDeployBscToken(t *testing.T) {
//	key, _ := crypto.GenerateKey()
//	auth, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(BscTestnetChainId))
//}
