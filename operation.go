package massalib

import (
	"bytes"
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"slices"

	"lukechampine.com/blake3"
)

type OperationType uint32

const (
	OpTransaction OperationType = iota // 0
	OpRollBuy                          // 1
	OpRollSell                         // 2
	OpExecuteSC                        // 3
	OpCallSC                           // 4
)

var (
	_ = OperationBody(&BodyTransaction{})
)

type OperationBody interface {
	Bytes() []byte
	UnmarshalBinary(data []byte) error
	Type() OperationType
	ReadFrom(r io.Reader) (int64, error)
}

type Operation struct {
	Fee    Amount
	Expire uint64 // expire_period, typically current period + 10
	Body   OperationBody
}

// Bytes returns a binary representation of the operation
func (o *Operation) Bytes() []byte {
	// Fee, Expire, Type & body
	return slices.Concat(o.Fee.Bytes(), EncodeProtobufVarint(o.Expire), EncodeProtobufVarint(uint64(o.Body.Type())), o.Body.Bytes())
}

// Sign signs the given operation and returns a serialized operation
func (o *Operation) Sign(chainId ChainId, key crypto.Signer) ([]byte, error) {
	pub, ok := key.Public().(ed25519.PublicKey)
	if !ok {
		return nil, errors.New("invalid key passed to Sign")
	}
	h := o.Hash(chainId, pub)
	sig, err := key.Sign(rand.Reader, h, crypto.Hash(0))
	if err != nil {
		return nil, err
	}
	// See: https://docs.massa.net/docs/learn/operation-format-execution
	return slices.Concat(EncodeProtobufVarint(0), sig, EncodeProtobufVarint(0), pub, o.Bytes()), nil

}

// Hash returns the operation contents hash
func (o *Operation) Hash(chainId ChainId, pubKey ed25519.PublicKey) []byte {
	// chainid + pubkey vers[0] + pubkey + Bytes
	buf := slices.Concat(chainId.Bytes(), EncodeProtobufVarint(0), pubKey, o.Bytes())
	h := blake3.Sum256(buf)
	return h[:]
}

func (o *Operation) UnmarshalBinary(b []byte) error {
	_, err := o.ReadFrom(bytes.NewReader(b))
	return err
}

func (o *Operation) ReadFrom(r io.Reader) (int64, error) {
	b := asBytereader(r)
	// read fee, expire, type
	fee, n1, err := ReadProtobufVarint(b)
	if err != nil {
		return n1, err
	}
	expire, n2, err := ReadProtobufVarint(b)
	if err != nil {
		return n1 + n2, err
	}
	typ, n3, err := ReadProtobufVarint(b)
	if err != nil {
		return n1 + n2 + n3, err
	}

	switch OperationType(typ) {
	case OpTransaction:
		o.Body = &BodyTransaction{}
	default:
		return n1 + n2 + n3, fmt.Errorf("unsupported tx type %d", typ)
	}

	n4, err := o.Body.ReadFrom(b)
	if err != nil {
		return n1 + n2 + n3 + n4, err
	}

	o.Fee = Amount(fee)
	o.Expire = expire
	return n1 + n2 + n3 + n4, nil
}

type BodyTransaction struct {
	Destination *Address
	Amount      Amount
}

func (tx *BodyTransaction) Bytes() []byte {
	return slices.Concat(tx.Destination.Bytes(), tx.Amount.Bytes())
}

func (tx *BodyTransaction) Type() OperationType {
	return OpTransaction
}

func (tx *BodyTransaction) UnmarshalBinary(buf []byte) error {
	_, err := tx.ReadFrom(bytes.NewReader(buf))
	return err
}

func (tx *BodyTransaction) ReadFrom(r io.Reader) (int64, error) {
	b := asBytereader(r)

	addr := &Address{}
	n1, err := addr.ReadFrom(b)
	if err != nil {
		return n1, err
	}
	var amt Amount
	n2, err := amt.ReadFrom(b)
	if err != nil {
		return n1 + n2, err
	}

	tx.Destination = addr
	tx.Amount = amt
	return n1 + n2, nil
}

type BodyRollBuy struct {
	Rolls uint64
}

type BodyRollSell struct {
	Rolls uint64
}

type DataStoreItem struct {
	Key   []byte
	Value []byte
}

type BodyExecuteSC struct {
	MaxGas    uint64
	MaxCoins  Amount
	Bytecode  []byte           // Raw bytes of bytecode to execute (up to 10MB)
	Datastore []*DataStoreItem // Concatenated datastore items. (keyval, keys up to
}

type BodyCallSC struct {
	MaxGas   uint64
	MaxCoins Amount
	Target   *Address
	Function string // Name of the function to call encoded as UTF-8 string without null termination
	Param    []byte
}
