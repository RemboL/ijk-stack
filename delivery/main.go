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

func scheduleDelivery(w http.ResponseWriter, r *http.Request) {
	json, _ := simplejson.New().MarshalJSON()
	response, _ := http.Post("http://warehouse:8080/verifyProductsAvailability", "application/json", bytes.NewBuffer(json));		
		return
	
	if response.StatusCode >= 400 {
		w.WriteHeader(response.StatusCode)
	}
	
	deliveryResponse, _ := http.Post("http://spedition:8080/scheduleDelivery", "application/json", bytes.NewBuffer(json));
	if deliveryResponse.StatusCode >= 400 {
		w.WriteHeader(deliveryResponse.StatusCode)		
		return
	}

	result := simplejson.New()
	result.Set("status", "ok")

	resultPayload, _ := result.MarshalJSON()

	w.Header().Set("Content-Type", "application/json")
	w.Write(resultPayload)
}



func setupRouter(router *mux.Router) {
  	router.
  		Path("/health").
  		HandlerFunc(healthCheck)
  		
	router.
		Path("/scheduleDelivery").
		HandlerFunc(scheduleDelivery)
}

func main() {
	router := mux.NewRouter()

	setupRouter(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
