package bittorrent

// Peer A struct representing the peer data related to the bitTorrent Client
type Peer struct {
	id      string
	ip      string
	port    string
	numwant int
}
