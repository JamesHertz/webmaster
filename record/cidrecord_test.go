package record

import (
	"encoding/json"
	"testing"

	cidlib "github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/require"
)

func TestNewCid(t *testing.T) {

	_, err := NewCidRecord("whatever", 0)
	require.NotNil(t, err)

	cid := newCidFailOnError(t, "a jocke")

	for _, ptype := range []IpfsMode{NORMAL_IPFS, SECURE_IPFS} {
		rec, err := NewCidRecord(cid.String(), ptype)
		require.Nil(t, err)
		require.True(t, rec.Cid.Equals(cid))
	}

	_, err = NewCidRecord(cid.String(), 100)
	require.Equal(t, ErrInvalidProviderType, err)

}

func TestMarshall(t *testing.T) {
	cid := newCidFailOnError(t, "marshalling")

	rec, err := NewCidRecord(cid.String(), SECURE_IPFS)
	require.Nil(t, err)

	val, err := rec.Marshall()
	require.Nil(t, err)

	bk, err := unmarshall([]byte(val))
	require.Nil(t, err)

	require.Equal(t, rec.ProviderType, bk.ProviderType)
	require.True(t, rec.Cid.Equals(bk.Cid))

	// ....
	for _, ct := range []string{
		`whatever`, `{"keys": "missing"}`,
		`{"cid"` + cid.String() + `","provtype":"invalid"}`,
		`{"cid"` + cid.String() + `"}`,
		`{"provtype":0}`,
	} {
		_, err := unmarshall([]byte(ct))
		require.NotNil(t, err)
	}
}

func newCidFailOnError(t *testing.T, content string) cidlib.Cid {
	mh, err := multihash.Sum([]byte(content), multihash.SHA2_256, -1)
	require.Nil(t, err)
	return cidlib.NewCidV0(mh)
}

func unmarshall(data []byte) (*CidRecord, error) {
	rec := CidRecord{}
	if err := json.Unmarshal(data, &rec); err != nil {
		return nil, err
	}
	return &rec, nil
}
