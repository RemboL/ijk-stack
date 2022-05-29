package main

import (
	"log"
	"net/http"
	"math/rand"
    	"time"

	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("ok"))
}

func postFunction(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(rand.Intn(100) + 200) * time.Millisecond)

	json := simplejson.New()

	
	w.Header().Set("Content-Type", "application/json")
	payload, _ := json.MarshalJSON()
	w.Write(payload)
}



func setupRouter(router *mux.Router) {
  	router.
  		Path("/health").
  		HandlerFunc(healthCheck)
  		
	router.
		Path("/useGiftCard").
		HandlerFunc(postFunction)
}

func main() {
	router := mux.NewRouter()

	setupRouter(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
