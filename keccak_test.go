package keccak

import (
	"reflect"
	"testing"
)

var vector = []byte{0xc5, 0xd2, 0x46, 0x01, 0x86, 0xf7, 0x23, 0x3c,
	0x92, 0x7e, 0x7d, 0xb2, 0xdc, 0xc7, 0x03, 0xc0,
	0xe5, 0x00, 0xb6, 0x53, 0xca, 0x82, 0x27, 0x3b,
	0x7b, 0xfa, 0xd8, 0x04, 0x5d, 0x85, 0xa4, 0x70,
}

func TestKeccak(t *testing.T) {
	h := New()
	sum := h.Sum(nil)
	if !reflect.DeepEqual(sum, vector) {
		t.Errorf("\"\": want % x, got % x", vector, sum)
	}
}

func Benchmark256(b *testing.B) {
	var tmp [Size]byte
	h := New()
	b.SetBytes(BlockSize)
	for i := 0; i < b.N; i ++ {
		h.Sum(tmp[:])
	}
}
