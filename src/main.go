package main

import (
	"fmt"
	"net/http"
	"os"

	"migrations"
)

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, migrations.SomeFunc()+": %s\n", r.URL.Path)
	})
	http.ListenAndServe(":"+PORT, nil)
}
