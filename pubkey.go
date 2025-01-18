package massalib

import (
	"crypto/ed25519"
	"crypto/sha256"
	"crypto/subtle"
	"errors"
	"fmt"
	"slices"

	"github.com/KarpelesLab/cryptutil"
	"github.com/ModChain/base58"
)

type PublicKey struct {
	Version uint64
	PubKey  ed25519.PublicKey
}

func (pk *PublicKey) Bytes() []byte {
	return slices.Concat(EncodeProtobufVarint(pk.Version), pk.PubKey)
}

func (pk *PublicKey) MarshalBinary() ([]byte, error) {
	return pk.Bytes(), nil
}

func (pk *PublicKey) String() string {
	buf := pk.Bytes()
	cksum := cryptutil.Hash(buf, sha256.New, sha256.New)
	return "P" + base58.Bitcoin.Encode(slices.Concat(buf, cksum[:4]))
}

func (pk *PublicKey) MarshalText() (text []byte, err error) {
	return []byte(pk.String()), nil
}

func (pk *PublicKey) UnmarshalText(text []byte) error {
	if len(text) < 18 {
		return errors.New("invalid public key text: length must be higher than 18")
	}
	if text[0] != 'P' {
		return errors.New("invalid public key text: must start with P")
	}

	text = text[1:]
	buf, err := base58.Bitcoin.Decode(string(text))
	if err != nil {
		return fmt.Errorf("invalid public key text: failed to decode: %w", err)
	}

	// check checksump
	cksum := buf[len(buf)-4:]
	buf = buf[:len(buf)-4]
	h := cryptutil.Hash(buf, sha256.New, sha256.New)
	if subtle.ConstantTimeCompare(cksum, h[:4]) == 0 {
		return errors.New("invalid public key text: bad checksum")
	}

	vers, l, err := DecodeProtobufVarint(buf)
	if err != nil {
		return fmt.Errorf("invalid public key text: failed to read version: %w", err)
	}
	buf = buf[l:]

	pk.Version = vers
	pk.PubKey = ed25519.PublicKey(buf)
	return nil
}
