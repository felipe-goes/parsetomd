package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// Make the query to stack exchange's api
func Query(site, query, tag string) {
	// Set up the request parameters
	baseURL := "https://api.stackexchange.com/2.3" + Endpoint
	params := url.Values{}
	params.Add("site", site)
	params.Add("q", query)
	if tag != "" {
		params.Add("tagged", tag)
	}
	params.Add("accepted", "true")
	params.Add("key", "rPwuqWr4Kp4R)yT6b1w*ZQ((")

	// Build the request URL with parameters
	requestURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Make the HTTP GET request
	resp, err := http.Get(requestURL)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read and parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Unmarshal the JSON response into the ApiResponse struct
	var apiResponse ApiResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		log.Fatalf("Failed to unmarshal JSON response: %v", err)
	}

	// Print the remaining quota and the titles and links of the questions
	fmt.Printf("# %s\n\n", query)
	fmt.Printf("Quota Remaining: %d\n", apiResponse.QuotaRemaining)
	fmt.Printf("Matches: %d\n\n", len(apiResponse.Items))
	fmt.Printf("## Results:\n\n")

	for _, item := range apiResponse.Items {
		fmt.Printf("- [%s](%s)\n", item.Title, item.Link)
	}
}
