package main

import (
	"sync"

	peer "github.com/libp2p/go-libp2p/core/peer"
	cidlib "github.com/ipfs/go-cid" 
)    

const K = 10



type ServerStorage struct {
	pids []peer.ID
	cids []cidlib.Cid
	lck sync.RWMutex
}

func newPidStorage() *ServerStorage{
	return &ServerStorage{
		pids: []peer.ID{},
		cids: []cidlib.Cid{},
		lck: sync.RWMutex{},
	}
}

func (st * ServerStorage) InsertAndGetPids(pid string) ([]peer.ID, error){
	new_pid, err := peer.Decode(pid)
	if err != nil {
		return nil, err
	}

	st.lck.Lock()
	defer st.lck.Unlock()
	
	st.pids = append(st.pids, new_pid)
	lastIdx := len(st.pids) -1
	st.pids[0], st.pids[lastIdx] = st.pids[lastIdx], st.pids[0]
	return st.pids[1:K+1], nil
}


