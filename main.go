package main

import (
	"BitTorrentTracker/bittorrent"
	"BitTorrentTracker/dbwrapper"
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

// ProcessPeerRequest A function to process the peer request
func ProcessPeerRequest(peerRequest bittorrent.PeerRequest) {
	var peerID = peerRequest.PeerID
	var ip = "124.123.125.12"
	var port = peerRequest.Port

	peer := dbwrapper.CreatePeer(peerID, port, ip)

	fmt.Print(peer)
}

// main function to boot up everything
func main() {
	dbwrapper.Migrate()
	fmt.Print("Finished")
	return

	router := mux.NewRouter()
	router.HandleFunc("/", HandleClientRequest).Methods("POST")

	fmt.Println("Server Listening at 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
