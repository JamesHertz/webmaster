package record

import (
	"encoding/json"
	"errors"
	"fmt"

	cidlib "github.com/ipfs/go-cid"
)

const (
	NORMAL_IPFS uint = iota + 1
	SECURE_IPFS
)

type CidRecord struct {
	Cid          cidlib.Cid `json:"cid"`
	ProviderType uint       `json:"provtype"`
}

var ErrInvalidProviderType = errors.New("Invalid provider type.")

func NewCidRecord(cid string, ptype uint) (*CidRecord, error) {
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
		return nil, ErrInvalidProviderType
	}
}

func (rec *CidRecord) Marshall() ([]byte, error) {
	return json.Marshal(rec)
}

func Unmarshall(data []byte) (*CidRecord, error) {
	res := CidRecord{}
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (rec *CidRecord) UnmarshalJSON(data []byte) error {
	obj := struct {
		Cid      string
		Provtype uint
	}{}

	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}

	aux, err := NewCidRecord(obj.Cid, obj.Provtype)
	if err != nil {
		return err
	}

	rec.Cid, rec.ProviderType = aux.Cid, aux.ProviderType
	return nil
}

func (rec CidRecord) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(
		`{"cid":"%s","provtype":%d}`, rec.Cid.String(), rec.ProviderType),
	), nil
}