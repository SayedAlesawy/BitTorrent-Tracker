package bittorrent

// PeerRequest A struct to represent the fields of a PeerRequest
type PeerRequest struct {
	InfoHash   string    `json:"infoHash,omitempty"`
	PeerID     string    `json:"peerID,omitempty"`
	Port       string    `json:"port,omitempty"`
	Uploaded   int       `json:"uploaded,omitempty"`
	Downloaded int       `json:"downloaded,omitempty"`
	Left       int       `json:"left,omitempty"`
	Event      EventType `json:"event,omitempty"`
}
