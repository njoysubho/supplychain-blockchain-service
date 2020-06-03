package contracts

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pborman/uuid"
	"math/big"
	"os"
	"testing"
)

type ContractWrapper struct {
	Auth            *bind.TransactOpts
	ContractBackend *backends.SimulatedBackend
	Instance *Contracts
}

var client ContractWrapper

func Setup() *ContractWrapper {
	key, _ := crypto.GenerateKey()
	client.Auth = bind.NewKeyedTransactor(key)

	address := client.Auth.From
	gAlloc := map[common.Address]core.GenesisAccount{
		address: {Balance: big.NewInt(10000000000)},
	}

	client.ContractBackend = backends.NewSimulatedBackend(gAlloc, 3000000)
	_, _, client.Instance, _ = DeployContracts(client.Auth, client.ContractBackend)
	return &client
}
func TestMain(m *testing.M) {
	Setup()
	os.Exit(m.Run())

}


func TestCreateBeneficiary(t *testing.T) {
    
	tx,err:=client.Instance.RegisterBeneficiary(client.Auth,uuid.New(),"TestBeneficiary","TestUID",
		"TestPAN","UBI1234567")
	if err !=nil{
		t.Fail()
	}
	if tx ==nil{
		t.Fail()
	}
}

func TestCreateBuyer(t *testing.T) {

	tx,err:=client.Instance.RegisterBuyer(client.Auth,uuid.New(),"TestBuyer","TestUID",
		"TestPAN","TestTan","ICIC123456")
	if err !=nil{
		t.Fail()
	}
	if tx ==nil{
		t.Fail()
	}
}


