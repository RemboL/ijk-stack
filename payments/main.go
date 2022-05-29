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

func processPayment(w http.ResponseWriter, r *http.Request) {
	json, _ := simplejson.New().MarshalJSON()
	response, _ := http.Post("http://gift-cards:8080/useGiftCard", "application/json", bytes.NewBuffer(json));
	
	if response.StatusCode >= 400 {
		w.WriteHeader(response.StatusCode)		
		return
	}
	
	installmentResponse, _ := http.Post("http://installments:8080/setUpInstallmentPayments", "application/json", bytes.NewBuffer(json));
	if installmentResponse.StatusCode >= 400 {
		w.WriteHeader(installmentResponse.StatusCode)		
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
		Path("/processPayment").
		HandlerFunc(processPayment)
}

func main() {
	router := mux.NewRouter()

	setupRouter(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
