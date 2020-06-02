package domain

import "github.com/njoysubho/supplychain-blockchain-service/contracts"

type VerifiedActor struct {
	ID   string
	Name string
	UID  string
	PAN  string
}
type Seller struct {
	VerifiedActor
	SellerId    string
	BankAccount string
}

func (s *Seller) FromKSB(beneficiary contracts.KisanSupplyChainBeneficiary) *Seller {
	s.ID = beneficiary.BeneficiaryId
	s.BankAccount = beneficiary.BankAccount
	s.Name = beneficiary.Name
	s.UID = beneficiary.Uid
	s.PAN = beneficiary.Pan
	return s
}



type Buyer struct {
	VerifiedActor
	BuyerId     string
	TAN         string
	BankAccount string
}

func (b *Buyer) FromKSBuyer(buyer contracts.KisanSupplyChainBuyer) *Buyer {
	b.ID = buyer.BuyerId
	b.BankAccount = buyer.BankAccount
	b.Name = buyer.Name
	b.UID = buyer.Uid
	b.PAN = buyer.Pan
	b.TAN = buyer.Tan
	return b
}

type Sales struct {
	InvoiceId     string
	Item          string
	Unit          string
	Amount        uint
	AmountPerUnit uint
	SalesDate     uint
	BeneficiaryId string
	BuyerId       string
	Status        string
}

