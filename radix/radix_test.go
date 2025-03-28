package radix_test

import (
	"testing"

	"github.com/acoderup/boost/radix"
)

func TestRadix(t *testing.T) {
	c := radix.NewCharset([]byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_")).Shuffle(0)
	if c.Size() != 64 {
		t.Error("radix charset size error")
	}

	r := radix.New(radix.BigEndian, 8, c)
	t.Log(r.Encode(123456789))

	if r.Decode(r.Encode(123456789)) != 123456789 {
		t.Error("radix decode error")
	}

	r2 := radix.New(radix.LittleEndian, 8, c)
	t.Log(r2.Encode(123456789))

	if r2.Decode(r2.Encode(123456789)) != 123456789 {
		t.Error("radix decode error")
	}
}
