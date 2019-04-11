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
	var peerRequest bittorrent.PeerRequest
	_ = json.NewDecoder(r.Body).Decode(&peerRequest)
	fmt.Println("[SERVER] Client IP Address ", r.RemoteAddr)
	peers := ProcessPeerRequest(peerRequest)
	fmt.Println(peers)
	var trackerResponse bittorrent.TrackerResponse

	json.NewEncoder(w).Encode(trackerResponse)

}

// ProcessPeerRequest A function to process the peer request
func ProcessPeerRequest(peerRequest bittorrent.PeerRequest) []bittorrent.Peer {
	var peerID = peerRequest.PeerID
	var ip = "124.123.125.12"
	var port = peerRequest.Port
	var infoHash = peerRequest.InfoHash

	var uploaded = peerRequest.Uploaded
	var downloaded = peerRequest.Downloaded
	var left = peerRequest.Left
	var event = peerRequest.Event

	fmt.Println("[SERVER] Creating Peer")
	peer := dbwrapper.CreatePeer(peerID, port, ip)
	fmt.Print("[SERVER] ", peer)

	fmt.Println("[SERVER] Creating Download")
	download := dbwrapper.CreateDownload(infoHash)
	fmt.Print("[SERVER] ", download)

	fmt.Println("[SERVER] Creating DownloadPeer")
	peerDownload := dbwrapper.CreatePeerDownload(uploaded, downloaded, left, event, peerID, infoHash)
	fmt.Println(peerDownload)
	fmt.Print("[SERVER] ", peerDownload)

	fmt.Println("[SERVER] Getting Peers for download " + infoHash)
	peers := dbwrapper.GetPeers(infoHash)
	return peers
}

// main function to boot up everything
func main() {
	dbwrapper.Migrate()
	fmt.Println("Migrated Database")

	router := mux.NewRouter()
	router.HandleFunc("/", HandleClientRequest).Methods("POST")

	fmt.Println("Server Listening at 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
