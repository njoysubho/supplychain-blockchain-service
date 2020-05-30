package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/njoysubho/supplychain-blockchain-service/contracts"
	"github.com/njoysubho/supplychain-blockchain-service/service"
	"github.com/pace/bricks/http"
	"log"
	"math/big"
)

func main() {
	scmEthClient := setupEthClient()
	scmService := service.SupplyChainService{SCMEthClient: scmEthClient}
	r := http.Router()
	r.HandleFunc("/v1/sellers", scmService.RegisterSeller).Methods("POST")
	r.HandleFunc("/v1/sellers/{id}", scmService.GetSellerById)
	r.HandleFunc("/v1/buyers", scmService.RegisterBuyer).Methods("POST")
	r.HandleFunc("/v1/buyers/{id}", scmService.GetBuyerById)
	_ = http.Server(r).ListenAndServe()
}

func setupEthClient() *contracts.SupplyChainEthClient {
	infuraKey:=service.GetSecretByKey("infura-key")
	pk:=service.GetSecretByKey("private-key")
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/"+infuraKey)
	address := common.HexToAddress("0x21818de2feF10F2D44130ae1D5DD29c228E12914")
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	cOps := bind.CallOpts{From: address, Pending: false}
	scmEthClient := contracts.SupplyChainEthClient{ContractAddress: address, FromAddress: fromAddress, Client: client, TOps: auth, COps: &cOps}
	return &scmEthClient
}
