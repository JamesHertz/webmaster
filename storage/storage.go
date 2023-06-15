package storage

import (
	"encoding/json"
	"sync"
	"math/rand"

	"github.com/JamesHertz/webmaster/record"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

const (
	PIDS_K = 10
	CIDS_K = 15
	DEFAULT_PEERS_SIZE = 500 //1000
	DEFAULT_CIDS_SIZE  = 1000 //10000
)

type ServerStorage struct {
	sync.RWMutex
	peers []peer.AddrInfo
	cids  []record.CidRecord
}

func NewServerStorage() *ServerStorage {
	return &ServerStorage{
		peers:  make([]peer.AddrInfo, 0, DEFAULT_PEERS_SIZE),
		cids:   make([]record.CidRecord, 0, DEFAULT_CIDS_SIZE),
	}
}

func (st *ServerStorage) InsertAndGetPeers(peer peer.AddrInfo) []peer.AddrInfo {
	st.Lock()
	defer st.Unlock()

	length   := len(st.peers)
	st.peers  = append(st.peers, peer)
	return st.peers[max(0,length-PIDS_K):length]
}

func (st *ServerStorage) AddCidRecord(recs ...record.CidRecord) {
	st.Lock()
	st.cids = append(st.cids, recs...)
	st.Unlock()
}

func (st *ServerStorage) GetRandomRecords() []record.CidRecord{
	st.RLock()
	defer st.RUnlock()

	cids_len := len(st.cids)

	if len(st.cids) < CIDS_K {
		panic("someone is rushing :)")
	}

	m := make(map[record.CidRecord]struct{}, CIDS_K)
	for len(m) < CIDS_K {
		choosed :=  rand.Intn( cids_len )
		next := st.cids[ choosed ]
		m[next] = struct{}{}
	}

	res := make([]record.CidRecord, 0, CIDS_K)
	for cid := range m {
		res = append(res, cid)
	}

	return res
}


func MarshalPeers(peers []peer.AddrInfo) []byte {
	data, _ := json.Marshal(peers)
	return data
}


// some utils function
func max(x, y int) int{
	if x >= y {
		return x
	}
	return y
}