package massalib

import (
	"bufio"
	"crypto/sha256"
	"io"
	"slices"

	"github.com/KarpelesLab/cryptutil"
	"github.com/ModChain/base58"
	"lukechampine.com/blake3"
)

// See: https://docs.massa.net/docs/learn/operation-format-execution#address

type Address struct {
	Category uint64 // 0 for User, 1 for Smart contract
	Version  uint64
	Hash     []byte // Underlying blake3 hash
}

func (a *Address) Thread() byte {
	// Blockclique sharding thread of an externally owned account address is computed by taking the first 5 bits of its underlying hash.
	return a.Hash[0] & 0x1f
}

// SetValue computes the hash of the provided buffer and sets it as address Hash
func (a *Address) SetValue(in []byte) {
	h := blake3.Sum256(in)
	a.Hash = h[:]
}

// Bytes encodes the address in massa byte format
func (a *Address) Bytes() []byte {
	return slices.Concat(EncodeProtobufVarint(a.Category), EncodeProtobufVarint(a.Version), a.Hash)
}

// MarshalBinary encodes the address in massa byte format (for compatibility)
func (a *Address) MarshalBinary() ([]byte, error) {
	return a.Bytes(), nil
}

// ReadFrom reads a binary address from a given source, which is made of two varint (category, version) followed by 32 bytes of hash
func (a *Address) ReadFrom(r io.Reader) (int64, error) {
	var b *bufio.Reader
	var ok bool

	if b, ok = r.(*bufio.Reader); !ok {
		b = bufio.NewReader(r)
	}

	v1, n1, err := ReadProtobufVarint(b)
	if err != nil {
		return n1, err
	}
	v2, n2, err := ReadProtobufVarint(b)
	if err != nil {
		return n1 + n2, err
	}
	buf := make([]byte, 32)
	n3, err := io.ReadFull(b, buf)
	if err != nil {
		return n1 + n2 + int64(n3), err
	}

	a.Category = v1
	a.Version = v2
	a.Hash = buf
	return n1 + n2 + int64(n3), nil
}

// String encodes the address as a string following encoding conventions
func (a *Address) String() string {
	addrBytes := slices.Concat(EncodeProtobufVarint(a.Version), a.Hash)
	// Note that the base58check encoding is the same one as the one used by bitcoin but without version number as it is already included in the underlying binary serialization of the operation ID.
	cksum := cryptutil.Hash(addrBytes, sha256.New, sha256.New)
	suffix := base58.Bitcoin.Encode(slices.Concat(addrBytes, cksum[:4]))

	switch a.Category {
	case 0:
		return "AU" + suffix
	case 1:
		return "AS" + suffix
	default:
		return "A?" + suffix
	}
}
