package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gosoln/dto"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			block := getLatestBlock()
			// send metrics to OTLP protocol
			sendMetrics(block)

		case <-quit:
			ticker.Stop()
			return
		}
	}

}

func getLatestBlock() int {
	url := "https://ethereum.publicnode.com"
	method := "POST"

	payload := strings.NewReader(`{
    "jsonrpc": "2.0",
    "method": "eth_getBlockByNumber",
    "params": ["latest",false],
    "id": 1
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return 0
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	// fmt.Println(string(body))

	var resp dto.EthLatestBlockResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	decimal, err := strconv.ParseInt(resp.Result.Number[2:], 16, 32)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return int(decimal)
}

func sendMetrics(block int) {
	// send metrics
	log.Println(block)
}
