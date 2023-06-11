package main

import (
	"sync"

	"github.com/JamesHertz/webmaster/record"
	peer "github.com/libp2p/go-libp2p/core/peer"
)

const K = 10

type ServerStorage struct {
	pids []peer.ID
	cids []record.CidRecord
	lck  sync.RWMutex
}

func newServerStorage() *ServerStorage {
	return &ServerStorage{
		pids: []peer.ID{},
		cids: []record.CidRecord{},
		lck:  sync.RWMutex{},
	}
}

func (st *ServerStorage) InsertAndGetPids(pid peer.ID) []peer.ID {
	st.lck.Lock()
	defer st.lck.Unlock()

	st.pids = append(st.pids, pid)
	lastIdx := len(st.pids) - 1
	st.pids[0], st.pids[lastIdx] = st.pids[lastIdx], st.pids[0]
	return st.pids[1 : K+1]
}

func (st *ServerStorage) AddCidRecord(rec record.CidRecord) {
	st.lck.Lock()
	st.cids = append(st.cids, rec)
	st.lck.Unlock()
}
