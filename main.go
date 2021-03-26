package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	var result string

	// Get environment variables stargs with DEMO_WEBSERVER_
	re := regexp.MustCompile(`^DEMO_WEBSERVER_`)

	for _, e := range os.Environ() {
		if re.MatchString(e) {
			pair := strings.SplitN(e, "=", 2)
			data[pair[0]] = pair[1]
		}

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

	for _, k := range keys {
		result += "<h1>" + k + ": " + data[k] + "</h1>"
	}

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
