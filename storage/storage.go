package storage

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/JamesHertz/webmaster/record"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

const K = 10

type ServerStorage struct {
	peers []peer.AddrInfo
	cids  []record.CidRecord
	lck   sync.RWMutex
}

func NewServerStorage() *ServerStorage {
	return &ServerStorage{
		peers: []peer.AddrInfo{},
		cids:  []record.CidRecord{},
		lck:   sync.RWMutex{},
	}
}

func (st *ServerStorage) InsertAndGetPeers(peer peer.AddrInfo) []peer.AddrInfo {
	st.lck.Lock()
	defer st.lck.Unlock()

	length   := len(st.peers)
	st.peers  = append(st.peers, peer)
	return st.peers[max(0,length-K):length]
}

func (st *ServerStorage) AddCidRecord(recs ...record.CidRecord) {
	st.lck.Lock()
	st.cids = append(st.cids, recs...)
	st.lck.Unlock()
}


func MarshalPeers(peers []peer.AddrInfo) []byte {
	aux := make([]string, len(peers))

	for i, peer := range peers {
		aux[i] = fmt.Sprintf("%s/p2p/%s",peer.Addrs[0].String(), peer.ID.Pretty())
	}

	data, _ := json.Marshal(aux)
	return data
}


// some utils function
func max(x, y int) int{
	if x >= y {
		return x
	}
	return y
}