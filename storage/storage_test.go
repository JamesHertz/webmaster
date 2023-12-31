package storage

import (
	"encoding/json"
	"fmt"
	"testing"

	peerlib "github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/require"
)

func TestMarshalPeers(t *testing.T) {
	peers := make([]peerlib.AddrInfo, 10)

	for i := 0; i < len(peers); i++ {
		key := fmt.Sprintf("keyword-%d", i)
		fake_pid := genMhFailOnError(t, key)

		maddr := fmt.Sprintf("/ip4/127.0.0.1/tcp/%d/p2p/%s", 4001+i, fake_pid)
		peer, err := peerlib.AddrInfoFromString(maddr)
		require.Nil(t, err)

		peers[i] = *peer
	}

	res, err := json.Marshal(peers)
	require.Nil(t, err)

	var back []peerlib.AddrInfo
	require.Nil(t, json.Unmarshal(res, &back))
	require.EqualValues(t, peers, back)
}

func TestServerStorageInsertAndGet(t *testing.T) {
	st := NewServerStorage()
	peers := make([]peerlib.AddrInfo, 20)

	for i := 0; i < len(peers); i++ {
		fake_pid := genMhFailOnError(t, fmt.Sprintf("myfakepid-%d", i))
		maddr := fmt.Sprintf("/ip4/1.2.3.4/udp/%d/quic/p2p/%s", 5000+i, fake_pid)

		peer, err := peerlib.AddrInfoFromString(maddr)
		require.Nil(t, err)

		res := st.InsertAndGetPeers(*peer)

		if i == 0 {
			require.Empty(t, res)
		} else {

			// start = max(0, i - PIDS_K)
			start := i - PidsK
			if start < 0 {
				start = 0
			}

			expected := peers[start:i]

			require.Nil(t, err)
			require.EqualValues(t, expected, res)
		}

		peers[i] = *peer
	}
}

func genMhFailOnError(t *testing.T, key string) string {
	pid, err := multihash.Sum([]byte(key), multihash.SHA2_256, -1)
	require.Nil(t, err)
	return pid.B58String()
}
