package main

//
//func main() {
//	if err := godotenv.Load(); err != nil {
//		panic(fmt.Sprintf("Error loading .env file: %v", err))
//	}
//	//client, err := ethclient.Dial(os.Getenv("NODE_RPC_URL"))
//	client_bsc, err := ethclient.Dial(os.Getenv("NODE_RPC_URL_BSC"))
//	if err != nil {
//		panic(fmt.Sprintf("Error connecting to eth client: %v", err))
//	}
//	auth := GetAccountAuth(client_bsc, os.Getenv("METAMASK_PRIVATE_KEY"))
//
//	//eth_cAddress, tx, instance, err := tokenomics.DeployTokenEth(auth, client, "GToken", "GTP")
//	//if err != nil {
//	//	panic(fmt.Sprintf("Error deploying contract to eth sepolia testnet: %v", err))
//	//}
//	//fmt.Printf("Contract ETH deployed to %v\n %v\n %v", eth_cAddress.Hex(), tx, instance)
//	bsc_cAddress, tx_bsc, instance_bsc, err := tokenomics.De(auth, client_bsc, "GToken", "GTP")
//	if err != nil {
//		panic(fmt.Sprintf("Error deploying contract to bsc testnet: %v", err))
//	}
//	fmt.Printf("Contract BSC deployed to %v\n %v\n %v", bsc_cAddress.Hex(), tx_bsc, instance_bsc)
//}
//
//// GetAccountAuth connect to account with Private Key
//func GetAccountAuth(client *ethclient.Client, privKey string) *bind.TransactOpts {
//	privateKey, err := crypto.HexToECDSA(privKey)
//	if err != nil {
//		panic(fmt.Sprintf("Error while transforming private key: %v", err))
//	}
//	pubKey := privateKey.Public()
//	pubkeyEcdsa, ok := pubKey.(*ecdsa.PublicKey)
//	if !ok {
//		panic("Error casting public key to ECDSA: Invalid key")
//	}
//	fromAddress := crypto.PubkeyToAddress(*pubkeyEcdsa)
//
//	// fetch the last nonce of account
//	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
//	if err != nil {
//		panic(fmt.Sprintf("Error while fetching nonce: %v", err))
//	}
//	chainId, err := client.ChainID(context.Background())
//	if err != nil {
//		panic(fmt.Sprintf("Error getting chainId: %v", err))
//	}
//	user, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
//	user.Nonce = big.NewInt(int64(nonce))
//	//user.Value = big.NewInt(0) // in wei
//	//user.GasLimit = uint64(3000000) // in units
//	//user.GasPrice = big.NewInt(1000000)
//	return user
//}
