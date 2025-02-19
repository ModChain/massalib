package massalib

import "encoding/binary"

type ChainId uint64

const (
	MainNet   ChainId = 77658377
	BuildNet  ChainId = 77658366
	SecureNet ChainId = 77658383
	LabNet    ChainId = 77658376
	Sandbox   ChainId = 77
)

func (c ChainId) Bytes() []byte {
	v, _ := c.AppendBinary(nil)
	return v
}

func (c ChainId) AppendBinary(b []byte) ([]byte, error) {
	return binary.BigEndian.AppendUint64(b, uint64(c)), nil
}
