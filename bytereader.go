package massalib

import (
	"bufio"
	"io"
)

type bytereader interface {
	io.Reader
	io.ByteReader
}

func asBytereader(r io.Reader) bytereader {
	if x, ok := r.(bytereader); ok {
		return x
	}
	return bufio.NewReader(r)
}
