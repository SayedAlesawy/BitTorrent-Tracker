package bittorrent

// PeerRequest A struct to represent the fields of a PeerRequest
type PeerRequest struct {
	infoHash   string
	peerID     string
	port       string
	uploaded   int
	downloaded int
	left       int
	event      EventType
}
