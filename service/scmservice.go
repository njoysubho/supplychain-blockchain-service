package service

import (
	"encoding/json"
	"github.com/njoysubho/supplychain-blockchain-service/types"
	"net/http"
)

func RegisterSeller(w http.ResponseWriter, r *http.Request) {
	seller := types.Seller{}
	err := json.NewDecoder(r.Body).Decode(&seller)
	if err!=nil{

	}
   
}

func RegisterBuyer(w http.ResponseWriter, r *http.Request) {

}

func CreateInvoice(w http.ResponseWriter, r *http.Request) {

}
