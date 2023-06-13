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

	for i := 0 ; i < len(peers); i++ {
		key := fmt.Sprintf("keyword-%d", i)
		fake_pid := genMhFailOnError(t, key)

		peer_str    := fmt.Sprintf("/ip4/127.0.0.1/tcp/%d/p2p/%s", 4001+i, fake_pid)
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

func TestServerStorageInsertAndGet(t *testing.T) {
	st    := NewServerStorage()
	peers := make([]string, 20)

	for i := 0; i < len(peers); i++ {
		fake_pid := genMhFailOnError(t, fmt.Sprintf("myfakepid-%d", i))
		peer := fmt.Sprintf( "/ip4/1.2.3.4/udp/%d/quic/p2p/%s", 5000+i, fake_pid)

		aux, err := peerlib.AddrInfoFromString(peer)
		require.Nil(t, err)

		res := st.InsertAndGetPeers(*aux)

		if i == 0 {
			require.Empty(t, res)
		} else {

			// start = max(0, i - 10)
			start := i - 10
			if start < 0 {
				start = 0
			}

			expected := peers[start:i]

			var back []string
			err = json.Unmarshal(MarshalPeers(res), &back)

			require.Nil(t, err)
			require.EqualValues(t, expected, back)
		}

		peers[i] = peer
	}
}

func genMhFailOnError(t * testing.T, key string) string {
	pid, err := multihash.Sum([]byte(key), multihash.SHA2_256, -1)
	require.Nil(t, err)
	return pid.B58String()
}
