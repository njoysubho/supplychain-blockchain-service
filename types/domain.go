package types

type VerifiedActor struct {
	Name string
	UID  string
	PAN  string
}
type Seller struct {
	VerifiedActor
	SellerId    string
	BankAccount string
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
