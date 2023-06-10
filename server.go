package main

import (
	"fmt"
	"log"
	"net/http"
)

var PORT = ":8080"

func main(){
	log.Printf("Starting webmaster on http://localhost%s/", PORT)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprintf(w, "You get it :)")
			return;
		}
		http.NotFound(w, r)
	})

	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal(err)
	}
}