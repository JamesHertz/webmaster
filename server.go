package main

import (
	"fmt"
	"log"
	"net/http"
    // peer "github.com/libp2p/go-libp2p/core/peer"
	// cidlib "github.com/ipfs/go-cid" 
)

var PORT = 8080

func main(){
    log.Printf("Starting webmaster on http://localhost:%d/", PORT)

    /*
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprintf(w, "You get it :)")
			return;
		}
		http.NotFound(w, r)
	})
    */

	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil); err != nil {
		log.Fatal(err)
	}
}
