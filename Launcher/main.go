package main

import (
	dbwrapper "BitTorrentTracker/Database"
	tracker "BitTorrentTracker/Tracker"
	logger "BitTorrentTracker/Utils/Log"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// LogSign Used for logging tracker messages
const LogSign string = "[Tracker]"

// HandleAnnounceRequest A function to handle the announce request
func HandleAnnounceRequest(writer http.ResponseWriter, req *http.Request) {
	var peerRequest tracker.PeerRequest

	json.NewDecoder(req.Body).Decode(&peerRequest)
	logger.LogMsg(LogSign, fmt.Sprintf("Received request from IP = %s", req.RemoteAddr))

	peerList := RegisterPeer(peerRequest)
	logger.LogMsg(LogSign, "Peer List:")
	for _, peer := range peerList {
		tracker.PrintPeer(peer)
	}
	//TODO bencode and send json
	var trackerResponse tracker.TrackerResponse
	json.NewEncoder(writer).Encode(trackerResponse)
}

// RegisterPeer A function to register a new peer
func RegisterPeer(peerRequest tracker.PeerRequest) []tracker.Peer {
	peerID := peerRequest.PeerID
	ip := "124.123.125.12"
	port := peerRequest.Port
	infoHash := peerRequest.InfoHash
	uploaded := peerRequest.Uploaded
	downloaded := peerRequest.Downloaded
	left := peerRequest.Left
	event := peerRequest.Event

	logger.LogMsg(LogSign, "Registering Peer")
	peer := dbwrapper.CreatePeer(peerID, port, ip)
	tracker.PrintPeer(peer)

	logger.LogMsg(LogSign, "Creating Download")
	download := dbwrapper.CreateDownload(infoHash)
	tracker.PrintDownload(download)

	logger.LogMsg(LogSign, "Creating Peer-Download")
	peerDownload := dbwrapper.CreatePeerDownload(uploaded, downloaded, left, event, peerID, infoHash)
	tracker.PrintPeerDownload(peerDownload)

	logger.LogMsg(LogSign, "Acquiring Peer list")
	peerList := dbwrapper.GetPeerList(infoHash)

	return peerList
}

// HandleSwarmsRequest A function to handle the swarms request
func HandleSwarmsRequest(writer http.ResponseWriter, req *http.Request) {
	swarms := dbwrapper.GetSwarms()
	json.NewEncoder(writer).Encode(swarms)
}

// HandleStatUpdate A function to handle the status update request
func HandleStatUpdate(writer http.ResponseWriter, req *http.Request) {
	var peerRequest tracker.PeerRequest

	json.NewDecoder(req.Body).Decode(&peerRequest)
	logger.LogMsg(LogSign, fmt.Sprintf("Received update request from IP = %s", req.RemoteAddr))

	ok := dbwrapper.UpdatePeerDownload(peerRequest.Downloaded, peerRequest.Uploaded, peerRequest.Left,
		peerRequest.Event, peerRequest.PeerID, peerRequest.InfoHash)

	if ok == true {
		logger.LogMsg(LogSign, "Update succeeded")
	}
}

func main() {
	dbwrapper.CleanUP()
	dbwrapper.Migrate()

	router := mux.NewRouter()
	router.HandleFunc("/", HandleAnnounceRequest).Methods("POST")
	router.HandleFunc("/swarms", HandleSwarmsRequest).Methods("GET")
	router.HandleFunc("/stat", HandleStatUpdate).Methods("PUT")

	logger.LogMsg(LogSign, "Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
