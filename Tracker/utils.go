package tracker

import "fmt"

// PrintPeer A function print a peer
func PrintPeer(peer Peer) {
	fmt.Println("Peer:")
	fmt.Printf("   ID:%s\n", peer.ID)
	fmt.Printf("   IP:%s\n", peer.IP)
	fmt.Printf("   Port:%s\n", peer.Port)
}

// PrintDownload A function to print a download
func PrintDownload(download Download) {
	fmt.Println("Download:")
	fmt.Printf("   InfoHash:%s\n", download.InfoHash)
}

// PrintPeerRequest A function to print a peer request
func PrintPeerRequest(peerRequest PeerRequest) {
	fmt.Println("Peer Request:")
	fmt.Printf("   InfoHash:%s\n", peerRequest.InfoHash)
	fmt.Printf("   PeerID:%s\n", peerRequest.PeerID)
	fmt.Printf("   Port:%s\n", peerRequest.Port)
	fmt.Printf("   Uploaded:%d\n", peerRequest.Uploaded)
	fmt.Printf("   Downloaded:%d\n", peerRequest.Downloaded)
	fmt.Printf("   Left:%d\n", peerRequest.Left)
	fmt.Printf("   Event:%s\n", peerRequest.Event)
}

// PrintPeerDownload A function to print a peer-download
func PrintPeerDownload(peerDownload PeerDownload) {
	fmt.Println("Peer Download:")
	fmt.Printf("   Uploaded:%d\n", peerDownload.Uploaded)
	fmt.Printf("   Downloaded:%d\n", peerDownload.Downloaded)
	fmt.Printf("   Left:%d\n", peerDownload.Left)
	fmt.Printf("   Event:%s\n", peerDownload.Event)
}
