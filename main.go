package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// create a new item
func HandleClientRequest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var peerRequest PeerRequest = json.NewDecoder(r.Body).Decode(&peerRequest)

	//json.NewEncoder(w).Encode(people)
}

// main function to boot up everything
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HandleClientRequest).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
