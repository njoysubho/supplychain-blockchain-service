package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/magiconair/properties"
	"github.com/njoysubho/supplychain-blockchain-service/scm-solidity"
	"github.com/njoysubho/supplychain-blockchain-service/service"
	"log"
	"math/big"
)

func main(){
	infuraKey:=service.GetSecretByKey("infura-key")
	pk:=service.GetSecretByKey("private-key")
	client,err:=ethclient.Dial("https://rinkeby.infura.io/v3/"+infuraKey)
	if err != nil {
		log.Fatal(err)
	}
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
	auth.GasLimit = uint64(3500000) // in units
	auth.GasPrice = gasPrice

	props:=properties.MustLoadFile("application.properties",properties.UTF8)
	contract:=props.MustGet("migrate.contractAddress")
	proxy:=props.MustGet("migrate.proxy")
	instance, err:= contracts.NewContracts(common.HexToAddress(contract),client)
	refreshTransactionOpts(auth,client)
	log.Println(proxy)
	log.Println(common.HexToAddress(proxy))
	instance.AddAdmin(auth,common.HexToAddress(proxy))
	refreshTransactionOpts(auth,client)
	tx,_:=instance.RegisterBuyer(auth,"a","","","","","")
	log.Println(tx.Hash().String())
	if err != nil {
		log.Fatal(err)
	}
	_ = instance
}

func refreshTransactionOpts(ops *bind.TransactOpts, backend bind.ContractBackend) {
	nonce, _ := backend.PendingNonceAt(context.Background(), ops.From)
	gasPrice, _ := backend.SuggestGasPrice(context.Background())
	ops.Nonce = big.NewInt(int64(nonce))
	ops.GasPrice = gasPrice
}
