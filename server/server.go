package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/JamesHertz/webmaster/record"
	"github.com/JamesHertz/webmaster/storage"
	"github.com/libp2p/go-libp2p/core/peer"
)

const PORT = 80

func main() {
	rand.Seed( time.Now().Unix() )

	log.Printf("Starting webmaster on http://localhost:%d/", PORT)

	st := storage.NewServerStorage()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprintf(w, "server running :)")
		} else {
			http.NotFound(w, r)
		}
	})

	http.HandleFunc("/peers", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodPost:
			log.Println("/peers POST")
			body, _ := io.ReadAll(r.Body)
			pi  := peer.AddrInfo{}
			err := json.Unmarshal(body, &pi)
			if err != nil {
				log.Printf("ERROR: Bad peerAddress: %s", string(body))
				http.Error(w, "400 Bad Request", http.StatusBadRequest)
				return
			}

			res, _ := json.Marshal(st.InsertAndGetPeers(pi))
			fmt.Fprint(w, string(res))
			log.Printf("+new peer added (peer: %s)", pi.ID)
		default:
			http.NotFound(w, r)
		}

	})

	http.HandleFunc("/cids", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			log.Printf("/cids GET")
			res := st.GetRandomRecords()
			data, _ := json.Marshal(res)	
			fmt.Fprint(w, string(data))
		case http.MethodPost:
			log.Printf("/cids POST")
			body, _ := io.ReadAll(r.Body)
			recs := []record.CidRecord{}
			err  := json.Unmarshal(body, &recs)
			if err != nil || len(recs) == 0 {
				log.Printf("ERROR: bad cids or none cids: %s", string(body))
				http.Error(w, "400 Bad Request", http.StatusBadRequest)
				return
			}
			st.AddCidRecord(recs...)
			log.Printf("+%d cid added", len(recs))
		default:
			http.NotFound(w, r)
		}
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil); err != nil {
		log.Fatal(err)
	}
}
