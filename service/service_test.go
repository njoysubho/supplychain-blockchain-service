package service

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/njoysubho/supplychain-blockchain-service/scm-solidity"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)


var service *SupplyChainService
func Setup() {
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	address := auth.From
	gAlloc := map[common.Address]core.GenesisAccount{
		address: {Balance: big.NewInt(10000000000)},
	}

	contractBackend := backends.NewSimulatedBackend(gAlloc, 3000000)
	contractAddress, _,_, _ := contracts.DeployContracts(auth,contractBackend)
	ethClient:= SupplyChainEthClient{
		ContractAddress: contractAddress,
		FromAddress:     auth.From,
		Client:          contractBackend,
		TOps:            &bind.TransactOpts{
			                     From:     address,
			                     GasLimit: 300000,
		                  },
		COps:            &bind.CallOpts{From: address,Pending: false},
	}
	service= &SupplyChainService{
		&ethClient}
}
func TestMain(m *testing.M) {
	Setup()
	os.Exit(m.Run())

}


func TestCreateBeneficiary(t *testing.T) {
    body:=`{
    "Name": "MSP",
    "UID": "a",
    "PAN": "a",
    "BankAccount": "a"
}
          `
    bodyReader:=strings.NewReader(body)
	req, err := http.NewRequest("POST", "/v1/sellers", bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.RegisterSeller)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusAccepted)
	}
}

func TestCreateBuyer(t *testing.T) {
	body:=`{
    "Name": "Buyer",
    "UID": "a",
    "PAN": "a",
    "TAN": "a",
    "BankAccount": "a"
}
          `
	bodyReader:=strings.NewReader(body)
	req, err := http.NewRequest("POST", "/v1/buyers", bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.RegisterBuyer)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusAccepted)
	}
}





