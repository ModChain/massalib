package massalib_test

import (
	"testing"

	"github.com/ModChain/massalib"
)

func TestAddress(t *testing.T) {
	a, err := massalib.DecodeAddress("AU128qq86hv2NzXqhowRzaeMruThQyxQQC3PgW3cgHg2ttgXMTa1A")
	if err != nil {
		t.Errorf("failed to parse addr: %s", err)
		return
	}

	if a.String() != "AU128qq86hv2NzXqhowRzaeMruThQyxQQC3PgW3cgHg2ttgXMTa1A" {
		t.Errorf("failed to reformat address: %s", err)
	}

	//log.Printf("addr = %+v", a)
}
