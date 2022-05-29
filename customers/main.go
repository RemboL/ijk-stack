package main

import (
	"bytes"
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("ok"))
}

func registerOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customerId"]
	
	http.Get("http://customer-orders:8080/customers/" + customerId + "/orders");
	
	consents := simplejson.New()
	consents.Set("terms&conditions", "yes")
	consents.Set("digital marketing", "no")

	consentsPayload, _ := consents.MarshalJSON()
	
	http.Post("http://customer-consents:8080/customers/" + customerId + "/consents", "application/json", bytes.NewBuffer(consentsPayload))

	response := simplejson.New()
	response.Set("status", "ok")
	response.Set("id", customerId)

	responsePayload, _ := response.MarshalJSON()

	w.Header().Set("Content-Type", "application/json")
	w.Write(responsePayload)
}



func setupRouter(router *mux.Router) {
  	router.
  		Path("/health").
  		HandlerFunc(healthCheck)
  		
	router.
		Path("/customers/{customerId}/registerOrder").
		HandlerFunc(registerOrder)
}

func main() {
	router := mux.NewRouter()

	setupRouter(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
