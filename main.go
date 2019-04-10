package main

import (
	"BitTorrentTracker/bittorrent"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HandleClientRequest :
func HandleClientRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	var peerRequest bittorrent.PeerRequest
	_ = json.NewDecoder(r.Body).Decode(&peerRequest)
	fmt.Println(peerRequest)
	var trackerResponse bittorrent.TrackerResponse
	json.NewEncoder(w).Encode(trackerResponse)
}

// main function to boot up everything
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HandleClientRequest).Methods("POST")

	fmt.Println("Server Listening at 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
