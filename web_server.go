package main

import (
    "fmt"
    "log"
	"net/http"
)

func web_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "blah %s", r.URL.Path[1:])
}

func web_main() {
	http.HandleFunc("/", web_handler)
	http.HandleFunc("/debug.txt", lookup_debug_handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
