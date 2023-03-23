package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/batch", handleBatchRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleBatchRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/http" {
		http.Error(w, "Unsupported media type", http.StatusUnsupportedMediaType)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// split the request body by blank lines
	requests := splitRequests(string(body))

	// process each individual request
	for _, req := range requests {
		// send the request to the target server
		client := &http.Client{}
		newReq, err := http.ReadRequest(bufio.NewReader(strings.NewReader(req)))
		if err != nil {
			fmt.Println("Error reading request:", err)
			continue
		}
		resp, err := client.Do(newReq)
		if err != nil {
			fmt.Println("Error sending request:", err)
			continue
		}

		// read the response body
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			continue
		}

		// write the response to the batch response
		fmt.Fprintf(w, "HTTP/1.1 %s\n", resp.Status)
		for k, v := range resp.Header {
			for _, h := range v {
				fmt.Fprintf(w, "%s: %s\n", k, h)
			}
		}
		fmt.Fprintf(w, "Content-Length: %d\n\n", len(respBody))
		fmt.Fprintf(w, "%s", string(respBody))
		fmt.Fprintf(w, "\n")
	}
}

func splitRequests(requests string) []string {
	// split the requests by double newline
	splitRequests := strings.Split(requests, "\n\n")
	var result []string
	for _, req := range splitRequests {
		// ignore any empty lines
		if len(req) > 0 {
			result = append(result, req)
		}
	}
	return result
}
