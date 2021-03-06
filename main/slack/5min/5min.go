package slack5min

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Run slackに通知する
func Run(url string) (err error){

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	type Payload struct {
		Text string `json:"text"`
	}

	time.Local = time.FixedZone("Asia/Tokyo", 9*60*60)
	data := Payload{
		// fill struct
		Text: time.Now().String(),
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error slack request")
		return
	}
	defer resp.Body.Close()

	return
}
