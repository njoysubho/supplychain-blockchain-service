package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/njoysubho/supplychain-blockchain-service/contracts"
	"github.com/njoysubho/supplychain-blockchain-service/domain"
	"github.com/pace/bricks/maintenance/log"
	"github.com/pborman/uuid"
	"net/http"
)

type SupplyChainService struct {
	SCMEthClient *contracts.SupplyChainEthClient
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
