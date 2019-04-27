package dbwrapper

import (
	tracker "BitTorrentTracker/Tracker"
	logger "BitTorrentTracker/Utils/Log"
	"database/sql"
	"fmt"
)

// CreatePeer A function to insert a peer in the Peers table
func CreatePeer(peerID string, port string, ip string) tracker.Peer {
	sqlStatement := InsertPeer

	logMsgs := logger.LogInfo{
		Success: "Peer created Successfully",
		Error:   "Peer creation failed",
	}

	ok := ExecuteQuery(sqlStatement, logMsgs, false, peerID, port, ip)

	var peer tracker.Peer
	if ok == true {
		peer.ID = peerID
		peer.IP = ip
		peer.Port = port
	}

	return peer
}

// CreateDownload A function to insert a peer in the Downloads table
func CreateDownload(infoHash string) tracker.Download {
	sqlStatement := InsertDownload

	logMsgs := logger.LogInfo{
		Success: "Download created Successfully",
		Error:   "Download creation failed",
	}

	ok := ExecuteQuery(sqlStatement, logMsgs, false, infoHash)

	var download tracker.Download
	if ok == true {
		download.InfoHash = infoHash
	}

	return download
}

// CreatePeerDownload A function to insert an entry in the Peer_Download relationship
func CreatePeerDownload(uploaded int, downloaded int, left int,
	event tracker.EventType, peerID string, infoHash string) tracker.PeerDownload {

	sqlStatement := InsertPeerDownload

	logMsgs := logger.LogInfo{
		Success: "PeerDownload created Successfully",
		Error:   "PeerDownload creation failed",
	}

	ok := ExecuteQuery(sqlStatement, logMsgs, false, uploaded, downloaded, left, event, peerID, infoHash)

	var peerDownload tracker.PeerDownload
	if ok == true {
		peerDownload.Uploaded = uploaded
		peerDownload.Downloaded = downloaded
		peerDownload.Left = left
		peerDownload.Event = event
	}

	return peerDownload
}

// GetPeer A function to select a peer from the Peers table identified with its id
func GetPeer(peerID string) tracker.Peer {
	sqlStatement := SelectPeer

	row := ExecuteRowQuery(sqlStatement, peerID)

	var dummyPeer tracker.Peer
	var peer tracker.Peer
	var serialID int

	switch err := row.Scan(&serialID, &dummyPeer.ID, &dummyPeer.Port, &dummyPeer.IP); err {
	case sql.ErrNoRows:
		logger.LogMsg(LogSign, fmt.Sprintf("No such peer with id = %s", peerID))
	case nil:
		peer = dummyPeer
	default:
		panic(err)
	}

	return peer
}

// GetPeerList A function to select a peer list based on a common file they are downloading
func GetPeerList(infoHash string) []tracker.Peer {
	sqlStatement := SelectPeerList

	logMsgs := logger.LogInfo{
		Success: "PeerList selected Successfully",
		Error:   "PeerList selection failed",
	}

	rows, ok := ExecuteRowsQuery(sqlStatement, logMsgs, false, infoHash)
	defer rows.Close()

	var peerList []tracker.Peer
	for rows.Next() {
		//Ma3lsh ya linter, need to change col name in DB
		var peer_id string

		err := rows.Scan(&peer_id)
		logger.LogErr(err, LogSign, "GetPeerList(): Error while extracting results", false)
		peerList = append(peerList, GetPeer(peer_id))
	}

	err := rows.Err()
	logger.LogErr(err, LogSign, "GetPeerList(): Error while extracting results", false)
	logger.LogSuccess(err, LogSign, "Peer list extracted successfully")

	if ok == false {
		peerList = []tracker.Peer{}
	}

	return peerList
}