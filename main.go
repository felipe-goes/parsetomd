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
	switch len(os.Args) {
	case 3:
		// without tag
		site := os.Args[1]
		query := os.Args[2]
		api.Query(site, query, "")
	case 4:
		// with tag
		site := os.Args[1]
		query := os.Args[2]
		tag := os.Args[3]
		api.Query(site, query, tag)
	default:
		fmt.Println(api.Help)
	}
}
