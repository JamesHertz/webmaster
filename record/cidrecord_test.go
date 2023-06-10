package record

import (
	"testing"

	cidlib "github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/require"
)

func TestNewCid(t *testing.T) {

	_, err := NewCidRecord("whatever", 0)
	require.NotNil(t, err)

	cid := newCidFailOnError(t, "a jocke");

	for _, ptype := range []uint{NORMAL_IPFS, SECURE_IPFS} {
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

	bk, err := Unmarshall([]byte(val))
	require.Nil(t, err)

	require.Equal(t, rec.ProviderType, bk.ProviderType)
	require.True(t, rec.Cid.Equals(bk.Cid))

	// ....
}

func newCidFailOnError(t * testing.T,  content string) cidlib.Cid {
	mh, err := multihash.Sum([]byte(content), multihash.SHA2_256, -1)
	require.Nil(t, err)
	return cidlib.NewCidV0(mh)
}