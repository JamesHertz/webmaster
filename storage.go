package main

import (
	"sync"

	"github.com/JamesHertz/webmaster/record"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

const K = 10

type ServerStorage struct {
	peers []peer.AddrInfo
	cids  []record.CidRecord
	lck  sync.RWMutex
}

func newServerStorage() *ServerStorage {
	return &ServerStorage{
		peers: []peer.AddrInfo{},
		cids: []record.CidRecord{},
		lck:  sync.RWMutex{},
	}
}

func (st *ServerStorage) InsertAndGetPids(peer peer.AddrInfo) []peer.AddrInfo {
	st.lck.Lock()
	defer st.lck.Unlock()


	st.peers = append(st.peers, peer)
	lastIdx := len(st.peers) - 1
	st.peers[0], st.peers[lastIdx] = st.peers[lastIdx], st.peers[0]
	return st.peers[1 : K+1]
}

func (st *ServerStorage) AddCidRecord(rec record.CidRecord) {
	st.lck.Lock()
	st.cids = append(st.cids, rec)
	st.lck.Unlock()
}
