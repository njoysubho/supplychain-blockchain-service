package contracts

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type SupplyChainEthClient struct {
	ContractAddress common.Address
	FromAddress     common.Address
	Client          bind.ContractBackend
	TOps            *bind.TransactOpts
	COps            *bind.CallOpts
}

func (s *SupplyChainEthClient) RegisterBeneficiary(seller *KisanSupplyChainBeneficiary) (*types.Transaction, error) {
	s.refreshTransactionOpts(s.TOps)
	instance, _ := NewContracts(s.ContractAddress, s.Client)
	tx, err := instance.RegisterBeneficiary(s.TOps, seller.BeneficiaryId, seller.Name, seller.Uid, seller.Pan, seller.BankAccount)
	return tx, err
}

func (s *SupplyChainEthClient) RegisterBuyer(buyer *KisanSupplyChainBuyer) (*types.Transaction, error) {
	s.refreshTransactionOpts(s.TOps)
	instance, _ := NewContracts(s.ContractAddress, s.Client)
	tx, err := instance.RegisterBuyer(s.TOps, buyer.BuyerId, buyer.Name, buyer.Uid, buyer.Pan, buyer.Tan, buyer.BankAccount)
	return tx, err
}

func (s *SupplyChainEthClient) refreshTransactionOpts(ops *bind.TransactOpts) {
	nonce, _ := s.Client.PendingNonceAt(context.Background(), s.FromAddress)
	gasPrice, _ := s.Client.SuggestGasPrice(context.Background())
	ops.Nonce = big.NewInt(int64(nonce))
	ops.GasPrice = gasPrice
}

func (s *SupplyChainEthClient) GetBeneficiary(sellerId string) (KisanSupplyChainBeneficiary, error) {

	instance, _ := NewContracts(s.ContractAddress, s.Client)
	ksb, err := instance.GetBenificiary(s.COps, sellerId)
	return ksb, err
}

func (s *SupplyChainEthClient) GetBuyer(buyerId string) (KisanSupplyChainBuyer, error) {

	instance, _ := NewContracts(s.ContractAddress, s.Client)
	ksb, err := instance.GetBuyer(s.COps, buyerId)
	return ksb, err
}
