package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	trunc := 80
	data := make(map[string]string)
	var result string

	// Get environment variables
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		data[pair[0]] = pair[1]
	}

	if len(data) == 0 {
		fmt.Fprintf(w, "<h1>Environment variables not found</h1>")
		return
	}

	keys := make([]string, 0, len(data))

	for k := range data {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	result = `
<html>
<head>
<style>
table, th, td {
  border: 1px solid black;
  border-collapse: collapse;
}

th, td {
	padding: 10px;
 }

tr:nth-child(even) {background-color: #f2f2f2;}
</style>
</head>
<body>

<h1>Environment Variables</h1>
<table>
<tr>
<th>Name</th>
<th>Value</th>
</tr>
  `

	for _, k := range keys {
		result += "<tr>"
		result += "<td>" + k + "</td>"

		if len(data[k]) >= trunc {
			// Truncate the value
			val := data[k]
			result += "<td>" + val[:trunc] + "...</td>"
		} else {
			result += "<td>" + data[k] + "</td>"
		}

		result += "</tr>"
	}

	result += "</table></html>"

	fmt.Fprintf(w, result)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/health", healthCheck)

	log.Println("Server listening or port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
