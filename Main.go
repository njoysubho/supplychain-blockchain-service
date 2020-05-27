package main

import (
	"github.com/njoysubho/supplychain-blockchain-service/service"
	"github.com/pace/bricks/http"
)

func main(){
	r:=http.Router()
	r.HandleFunc("/v1/sellers",service.RegisterSeller).Methods("POST")
	r.HandleFunc("/v1/Buyers",service.RegisterBuyer).Methods("POST")
	_ = http.Server(r).ListenAndServe()
}
