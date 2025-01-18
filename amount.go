package massalib

import (
	"bufio"
	"io"
)

// See: https://docs.massa.net/docs/learn/operation-format-execution#coin-amounts

// While Massa documentation specifies uint64 values, there are limits to that and it could yield to
// unexpected problems (overflow, etc).
//
// Massa uses integer values, that match values with 9 decimals.
// As such, the maximum amount that can be normally represented is: 18,446,744,073.709551615

type Amount uint64

func (a Amount) Bytes() []byte {
	return EncodeProtobufVarint(uint64(a))
}

func (a *Amount) ReadFrom(r io.Reader) (int64, error) {
	if b, ok := r.(*bufio.Reader); ok {
		v, ln, err := ReadProtobufVarint(b)
		if err != nil {
			return ln, err
		}
		*a = Amount(v)
		return ln, nil
	}
	v, ln, err := ReadProtobufVarint(bufio.NewReader(r))
	if err != nil {
		return ln, err
	}
	*a = Amount(v)
	return ln, nil
}
