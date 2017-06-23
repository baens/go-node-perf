package main

import (
	"net/http"
	"fmt"
	"bytes"
	"time"
)

func main() {
	url := "http://localhost:3000"

	startingTime := time.Now().UTC()
	// Contact the server and print out its response.
	for i := 0; i < 10000; i++ {
		var jsonStr = []byte(`{"name":"heloworld"}`)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		resp.Body.Close()
	}
	endingTime := time.Now().UTC()
	var duration time.Duration = endingTime.Sub(startingTime)
	fmt.Printf("Duration [%v]\n", duration)
}
