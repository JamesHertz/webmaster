package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/JamesHertz/webmaster/record"
	"github.com/libp2p/go-libp2p/core/peer"
	// peer "github.com/libp2p/go-libp2p/core/peer"
	// cidlib "github.com/ipfs/go-cid"
)

const PORT = 8080
var (
	PEERS_END_POINT = "/peers"
	CIDS_END_POINT  = "/cids"
)

func main(){
    log.Printf("Starting webmaster on http://localhost:%d/", PORT)

	st := newServerStorage()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprintf(w, "server running :)")
		} else {
			http.NotFound(w, r)
		}
	})

	http.HandleFunc(PEERS_END_POINT, func(w http.ResponseWriter, r *http.Request) {

		switch r.Method{
		case http.MethodPost:
		    body, _ := io.ReadAll(r.Body)
		    pi, err := peer.AddrInfoFromString(string(body))
		    if err != nil {
		    	http.Error(w, "400 Bad Request", http.StatusBadRequest)
		    	return
		    }

		    res, _ := json.Marshal(st.InsertAndGetPids(*pi)) // by now :)

		    // todo: marshall res and send it back :)
		    // todo: check if this is really okay
		    fmt.Fprint(w, res)
		default:
			http.NotFound(w, r)
		}


	})

	http.HandleFunc(CIDS_END_POINT, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case http.MethodGet:
			// TODO: implement this
			http.Error(w, "501 Not Implemented", http.StatusNotImplemented)
		case http.MethodPost:
			body, _  := io.ReadAll(r.Body)
			rec, err := record.Unmarshall(body)
			if err != nil {
				http.Error(w, "400 Bad Request", http.StatusBadRequest)
				return;
			}
			st.AddCidRecord(*rec)
		default:
			http.NotFound(w, r)

		}
	})


	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil); err != nil {
		log.Fatal(err)
	}
}
