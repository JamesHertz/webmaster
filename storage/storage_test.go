package storage

import (
	"encoding/json"
	"fmt"
	"testing"

	peerlib "github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/require"
)


func TestMarshalPeers(t *testing.T){
	peers    := make([]peerlib.AddrInfo, 10)
	expected := make([]string,  10)

	for i := 0 ; i < 10; i++ {
		key := fmt.Sprintf("keyword-%d", i)

		fake_pid, err := multihash.Sum([]byte(key), multihash.SHA2_256, -1)
		require.Nil(t, err)

		peer_str    := fmt.Sprintf("/ip4/127.0.0.1/tcp/%d/p2p/%s", 4001+i, fake_pid.B58String())
		peer, err   := peerlib.AddrInfoFromString( peer_str )

		require.Nil(t, err)
		peers[i]    = *peer
		expected[i] = peer_str
	}

	res := MarshalPeers(peers)

	var back []string
	require.Nil(t, json.Unmarshal(res, &back))
	require.EqualValues(t, back, expected)
}
