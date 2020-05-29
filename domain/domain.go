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

type Sales struct {
	InvoiceId     string
	Item          string
	Unit          string
	amount        uint
	amountPerUnit uint
	SalesDate     uint
	BeneficiaryId uint
	BuyerId       uint
	Status        string
}
