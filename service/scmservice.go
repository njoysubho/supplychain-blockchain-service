package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/njoysubho/supplychain-blockchain-service/scm-solidity"
	"github.com/njoysubho/supplychain-blockchain-service/domain"
	"github.com/pace/bricks/maintenance/log"
	"github.com/pborman/uuid"
	"math/big"
	"net/http"
	"time"
)

type SupplyChainService struct {
	SCMEthClient *SupplyChainEthClient
}

func (s *SupplyChainService) RegisterSeller(w http.ResponseWriter, r *http.Request) {
	seller := domain.Seller{}
	json.NewDecoder(r.Body).Decode(&seller)
	ksb := contracts.KisanSupplyChainBeneficiary{
		BeneficiaryId: uuid.New(),
		Name:          seller.Name,
		Uid:           seller.UID,
		Pan:           seller.PAN,
		BankAccount:   seller.BankAccount,
	}
	go s.SCMEthClient.RegisterBeneficiary(&ksb)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(202)
	createdSeller := domain.Seller{}
	createdSeller.ID = ksb.BeneficiaryId
	json.NewEncoder(w).Encode(createdSeller)
}

func (s *SupplyChainService) RegisterBuyer(w http.ResponseWriter, r *http.Request) {
	buyer := domain.Buyer{}
	json.NewDecoder(r.Body).Decode(&buyer)
	ksb := contracts.KisanSupplyChainBuyer{
		BuyerId:     uuid.New(),
		Name:        buyer.Name,
		Uid:         buyer.UID,
		Pan:         buyer.PAN,
		Tan:         buyer.TAN,
		BankAccount: buyer.BankAccount,
	}
	go s.SCMEthClient.RegisterBuyer(&ksb)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(202)
	createdBuyer := domain.Buyer{}
	createdBuyer.ID = ksb.BuyerId
	json.NewEncoder(w).Encode(createdBuyer)
}

func (s *SupplyChainService) GetSellerById(w http.ResponseWriter, r *http.Request) {
	pathvars := mux.Vars(r)
	sellerId := pathvars["id"]
	ksb, err := s.SCMEthClient.GetBeneficiary(sellerId)
	if err != nil {
		log.Print(err)
		w.WriteHeader(400)
	}
	seller := domain.Seller{}
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(seller.FromKSB(ksb))
}

func (s *SupplyChainService) GetBuyerById(w http.ResponseWriter, r *http.Request) {
	pathvars := mux.Vars(r)
	buyerId := pathvars["id"]
	ksb, err := s.SCMEthClient.GetBuyer(buyerId)
	if err != nil {
		log.Print(err)
		w.WriteHeader(400)
	}
	buyer := domain.Buyer{}
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(buyer.FromKSBuyer(ksb))
}

func (s *SupplyChainService) CreateSales(w http.ResponseWriter, r *http.Request) {
	sales := domain.Sales{}
	json.NewDecoder(r.Body).Decode(&sales)
	kisanSupplySales := contracts.KisanSupplyChainSales{
		InvoiceId:     uuid.New(),
		Item:          sales.Item,
		Unit:          sales.Unit,
		Amount:        big.NewInt(int64(sales.Amount)),
		AmountPerUnit: big.NewInt(int64(sales.AmountPerUnit)),
		SalesDate:     big.NewInt(time.Now().Unix()),
		BeneficiaryId: sales.BeneficiaryId,
		BuyerId:       sales.BuyerId,
	}
	go s.SCMEthClient.CreateInvoice(&kisanSupplySales)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(202)
	createdSales := domain.Sales{}
	createdSales.InvoiceId = kisanSupplySales.InvoiceId
	json.NewEncoder(w).Encode(createdSales)
}

func (s *SupplyChainService) ApproveByBuyer(w http.ResponseWriter, r *http.Request) {
	approval := domain.Approval{}
	json.NewDecoder(r.Body).Decode(&approval)
	go s.SCMEthClient.ApproveBuyer(approval.ApproverId, approval.InvoiceId)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(202)
}

func (s *SupplyChainService) ApproveBySeller(w http.ResponseWriter, r *http.Request) {
	approval := domain.Approval{}
	json.NewDecoder(r.Body).Decode(&approval)
	go s.SCMEthClient.ApproveSeller(approval.ApproverId, approval.InvoiceId)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(202)
}
