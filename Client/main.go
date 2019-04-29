package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func announce(infoHash string, peerID string, port string, uploaded int, downloaded int, left int, event string) {
	url := "http://localhost:3000"

	params := fmt.Sprintf("{\n\t\"infoHash\": \"%s\",\n\t\"peerID\": \"%s\",\n\t\"port\": \"%s\",\n\t\"uploaded\": %d,\n\t\"downloaded\": %d,\n\t\"left\": %d,\n\t\"event\": \"%s\"\n}",
		infoHash, peerID, port, uploaded, downloaded, left, event)
	payload := strings.NewReader(params)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "9d56bd58-3781-4713-92a0-7fbfcce4558c")

	http.DefaultClient.Do(req)
}

func updateStatus(infoHash string, peerID string, port string, uploaded int, downloaded int, left int, event string) {
	url := "http://localhost:3000/stat"

	params := fmt.Sprintf("{\n\t\"infoHash\": \"%s\",\n\t\"peerID\": \"%s\",\n\t\"port\": \"%s\",\n\t\"uploaded\": %d,\n\t\"downloaded\": %d,\n\t\"left\": %d,\n\t\"event\": \"%s\"\n}",
		infoHash, peerID, port, uploaded, downloaded, left, event)
	payload := strings.NewReader(params)

	req, _ := http.NewRequest("PUT", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "f1f550cd-bbd5-46a4-8f24-93c75f790981")

	http.DefaultClient.Do(req)
}

func work(infoHash string, peerID string, port string, uploaded int, downloaded int, left int, event string) {
	announce(infoHash, peerID, port, uploaded, downloaded, left, event)

	for range time.Tick(1000) {
		if left > 0 {
			downloaded++
			left--
		} else if left == 0 {
			event = "completed"
			uploaded++
		}
		updateStatus(infoHash, peerID, port, uploaded, downloaded, left, event)
		log.Println("  Uploaded:", uploaded)
		log.Println("  Downloaded:", downloaded)
		log.Println("  Left:", left)
		log.Println("  Status:", event)
	}
}

func main() {
	args := os.Args
	infoHash := args[1]
	peerID := args[2]
	port := args[3]
	uploaded, _ := strconv.Atoi(args[4])
	downloaded, _ := strconv.Atoi(args[5])
	left, _ := strconv.Atoi(args[6])
	event := "started"

	work(infoHash, peerID, port, uploaded, downloaded, left, event)
}
