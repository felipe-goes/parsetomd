package api

// Define a struct to parse the JSON response. This struct will only extract
// part of the response.
// Adjust it according to the data you need.
type ApiResponse struct {
	Items []struct {
		Title string `json:"title"`
		Link  string `json:"link" `
	} `json:"items"`
	QuotaRemaining int `json:"quota_remaining"`
}

// Endpoint to make the request
const Endpoint string = "/search/advanced"

// Help text to show when any of the expected parameters is missing
const Help string = `querystack is a tool for quering to stack exchange's api.

Usage: querystack [arguments]

Arguments:

  site:  choose one of the availables sites of the stack exchange api
  query: make a query to the chosen site
  tag:   choose up to one tag to help on the search`
