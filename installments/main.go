package main

import (
	"log"
	"net/http"
	"bytes"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("ok"))
}

func setUpInstallmentPayments(w http.ResponseWriter, r *http.Request) {
	json, _ := simplejson.New().MarshalJSON()
	response, _ := http.Post("http://debtors:8080/checkDebtorStatus", "application/json", bytes.NewBuffer(json));
	
	if response.StatusCode >= 400 {
		w.WriteHeader(response.StatusCode)		
		return
	}
	
	creditCardResponse, _ := http.Post("http://credit-cards:8080/verifyCreditCard", "application/json", bytes.NewBuffer(json));
	if creditCardResponse.StatusCode >= 400 {
		w.WriteHeader(creditCardResponse.StatusCode)		
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}



func setupRouter(router *mux.Router) {
  	router.
  		Path("/health").
  		HandlerFunc(healthCheck)
  		
	router.
		Path("/setUpInstallmentPayments").
		HandlerFunc(setUpInstallmentPayments)
}

func main() {
	router := mux.NewRouter()

	setupRouter(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
