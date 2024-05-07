package main

import (
	"fmt"
	"os"
	"querystack/api"
)

// Package querystack request a query to stack exchange's api and return a list of answered question.
func main() {
	// Check if there are any command-line arguments (excluding the program name itself)
	// expected: (site query tag)
	if len(os.Args) != 4 {
		fmt.Println(api.Help)
	} else {
		site  := os.Args[1]
		query := os.Args[2]
		tag   := os.Args[3]
		api.Query(site, query, tag)
	}
}
