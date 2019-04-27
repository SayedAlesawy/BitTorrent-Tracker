package tracker

import "time"

// EventType Represents the different status types of a download
type EventType string

// Enumerating the different types events
const (
	//Started The first requst to the tracker should have this value
	Started EventType = "started"

	//Stopped Should be sent to the tracker if the client is shutting down
	Stopped EventType = "stopped"

	//Completed Must be sent to the tracker when the download is complete
	//But not send if the download was already at 100% when it first started
	Completed EventType = "completed"
)

// Peer Represents a peer associated with the Bitorrent Client
type Peer struct {
	ID   string //The Peer ID
	IP   string //The Peer IP
	Port string //The Peer Port
}

// Download Represents the file being downloaded
type Download struct {
	InfoHash string //A unique download identifer
}

// PeerRequest Represents the peer request that's sent to the tracker
type PeerRequest struct {
	InfoHash   string    `json:"infoHash,omitempty"`   //Unique identifier of the download
	PeerID     string    `json:"peerID,omitempty"`     //The Peer ID
	Port       string    `json:"port,omitempty"`       //The Peer Port
	Uploaded   int       `json:"uploaded,omitempty"`   //The amount of bytes uploaded
	Downloaded int       `json:"downloaded,omitempty"` //The amount of bytes downloaded
	Left       int       `json:"left,omitempty"`       //The amount of bytes left
	Event      EventType `json:"event,omitempty"`      //The current status of the download
}

// PeerDownload Represents the relationship between Peers and Downloads
type PeerDownload struct {
	Uploaded   int       //Total number of bytes uploaded since the start event was sent
	Downloaded int       //Total number of bytes downloaded since the start event was sent
	Left       int       //Total number of bytes the peer has to download to have a complete file
	Event      EventType //The current status of the download
}

//TrackerResponse Represents a response sent by the tracker to the peer
type TrackerResponse struct {
	FailureReason string        `json:"infoHash,omitempty"`   //Reason why the failed request. If present, no other key should
	Interval      time.Duration `json:"interval,omitempty"`   //Interval in secs that the client should wait for re-requesting
	TrackerID     string        `json:"trackerID,omitempty"`  //The tracker ID
	Complete      int           `json:"complete,omitempty"`   //Number of seeders (peers with complete files)
	Incomplete    int           `json:"incomplete,omitempty"` //Number of leechers (peers with non-complete files)
	Peers         string        `json:"peers,omitempty"`      //Bencoded dictionary (list of peers, each has id, ip, port)
}
