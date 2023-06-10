package record

import (
	"errors"
	"encoding/json"

	cidlib "github.com/ipfs/go-cid"
)

const (
	NORMAL_IPFS ProvType = iota
	SECURE_IPFS
)

type ProvType int

type CidRecord struct {
	Cid          cidlib.Cid `json:"cid"`
	ProviderType ProvType   `json:"provtype"`
}

var InvalidProviderType = errors.New("Invalid provider type.")

func NewCidRecord(cid string, ptype ProvType) (*CidRecord, error) {
	value, err := cidlib.Decode(cid)
	if err != nil {
		return nil, err
	}

	switch ptype {
	case SECURE_IPFS, NORMAL_IPFS:
		return &CidRecord{
			Cid:          value,
			ProviderType: ptype,
		}, nil
	default:
		return nil, InvalidProviderType
	}
}


func (rec * CidRecord) Marshall() ([]byte, error) {
	return json.Marshal(rec)
}