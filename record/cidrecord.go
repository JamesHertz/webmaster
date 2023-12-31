package record

import (
	"encoding/json"
	"errors"
	"fmt"

	cidlib "github.com/ipfs/go-cid"
)

type IpfsMode int

const (
	NONE        IpfsMode = -1 
	NORMAL_IPFS IpfsMode = iota + 1
	SECURE_IPFS
)

type CidRecord struct {
	Cid          cidlib.Cid `json:"cid"`
	ProviderType IpfsMode   `json:"provtype"`
}

var ErrInvalidIpfsMode = errors.New("Invalid ipfs mode.")

func NewCidRecord(cid string, ptype IpfsMode) (*CidRecord, error) {
	value, err := cidlib.Decode(cid)
	if err != nil {
		return nil, err
	}

	if err := ptype.Validate(); err != nil {
		return nil, err
	}

	return &CidRecord{
		Cid:          value,
		ProviderType: ptype,
	}, nil
}

func (rec *CidRecord) Marshall() ([]byte, error) {
	return json.Marshal(rec)
}

func (rec *CidRecord) UnmarshalJSON(data []byte) error {
	obj := struct {
		Cid      string
		Provtype int
	}{}

	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}

	aux, err := NewCidRecord(obj.Cid, IpfsMode(obj.Provtype))
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

func (mode IpfsMode) Validate() error {
	switch mode {
	case SECURE_IPFS, NORMAL_IPFS, NONE:
		return nil
	default:
		return ErrInvalidIpfsMode
	}
}

func (mode IpfsMode) String() string {
	switch mode {
	case SECURE_IPFS:
		return "secure"
	case NORMAL_IPFS:
		return "normal"
	case NONE:
		return "none"
	default:
		panic(fmt.Sprintf("Invalid ipfs mode: %d", mode))
	}
}
