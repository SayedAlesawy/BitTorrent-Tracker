package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	url := "http://localhost:3000"
	infoHash := args[1]
	peerID := args[2]
	port := args[3]
	uploaded, _ := strconv.Atoi(args[4])
	downloaded, _ := strconv.Atoi(args[5])
	left, _ := strconv.Atoi(args[6])

	params := fmt.Sprintf("{\n\t\"infoHash\": \"%s\",\n\t\"peerID\": \"%s\",\n\t\"port\": \"%s\",\n\t\"uploaded\": %d,\n\t\"downloaded\": %d,\n\t\"left\": %d,\n\t\"event\": \"started\"\n}",
		infoHash, peerID, port, uploaded, downloaded, left)
	payload := strings.NewReader(params)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "9d56bd58-3781-4713-92a0-7fbfcce4558c")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
