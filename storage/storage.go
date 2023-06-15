package storage

import (
	"math/rand"
	"sync"

	"github.com/JamesHertz/webmaster/record"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

const (
	PidsK             = 10
	CidsK             = 15
	DefaultPeersSize  = 500  //1000
	DefaultCidsSize   = 1000 //10000
)

type ServerStorage struct {
	sync.RWMutex
	peers []peer.AddrInfo
	cids  []record.CidRecord
}

func NewServerStorage() *ServerStorage {
	return &ServerStorage{
		peers: make([]peer.AddrInfo, 0, DefaultPeersSize),
		cids:  make([]record.CidRecord, 0, DefaultCidsSize),
	}
}

func (st *ServerStorage) InsertAndGetPeers(peer peer.AddrInfo) []peer.AddrInfo {
	st.Lock()
	defer st.Unlock()

	length := len(st.peers)
	st.peers = append(st.peers, peer)
	return st.peers[max(0, length-PidsK):length]
}

func (st *ServerStorage) AddCidRecord(recs ...record.CidRecord) {
	st.Lock()
	st.cids = append(st.cids, recs...)
	st.Unlock()
}

func (st *ServerStorage) GetRandomRecords() []record.CidRecord {
	st.RLock()
	defer st.RUnlock()

	cids_len := len(st.cids)

	if len(st.cids) < CidsK {
		panic("someone is rushing :)")
	}

	m := make(map[record.CidRecord]struct{}, CidsK)
	for len(m) < CidsK {
		choosed := rand.Intn(cids_len)
		next := st.cids[choosed]
		m[next] = struct{}{}
	}

	res := make([]record.CidRecord, 0, CidsK)
	for cid := range m {
		res = append(res, cid)
	}

	return res
}


// some utils function
func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
