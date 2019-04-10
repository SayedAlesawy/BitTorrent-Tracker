package bittorrent

// EventType A struct to represent the different status types of a download
type EventType string

// Enumerating the different types of an event
const (
	started   EventType = "started"
	stopped   EventType = "stopped"
	completed EventType = "completed"
)

// PeerDownload ;
type PeerDownload struct {
	Uploaded   int
	Downloaded int
	Left       int
	Event      EventType
}
