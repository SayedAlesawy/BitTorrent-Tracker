package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "http://localhost:3000"

	payload := strings.NewReader("{\n\t\"infoHash\": \"movie1\",\n\t\"peerID\": \"1\",\n\t\"port\": \"3001\",\n\t\"uploaded\": 20,\n\t\"downloaded\": 3,\n\t\"left\": 6,\n\t\"event\": \"started\"\n}")

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
