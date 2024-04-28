package randx

import (
	"math/rand"
	"testing"
)

func TestPcg32(t *testing.T) {
	rng := NewPcg32Stream(42, 54)
	for i := 0; i < 5; i++ {
		rng.Next()
	}

	if next := rng.Next(); next != uint32(0xcbed606e) {
		t.Errorf("Got %x, was expecting %x", next, uint64(0xcbed606e))
	}
}

func TestBounds32(t *testing.T) {
	rng := NewPcg32(0)
	test1 := rng.NextN(1)
	if test1 != 0 {
		t.Errorf("Bound of 1 did not give 0 (%x)", test1)
	}
	rand.Seed(0)
	for i := 0; i < 100000; i++ {
		bounds := rand.Uint32()
		result := rng.NextN(bounds)
		if result >= bounds {
			t.Errorf("Got %x which is outside of bound %x", result, bounds)
		}
	}
}

func TestPcg64(t *testing.T) {
	rng := NewPcg64Stream(0, 42, 0, 54)
	for i := 0; i < 5; i++ {
		rng.Next()
	}

	if next := rng.Next(); next != uint64(0x606121f8e3919196) {
		t.Errorf("Got %x, was expecting %x", next, uint64(0x606121f8e3919196))
	}
}

func TestBounds(t *testing.T) {
	rng := NewPcg64(0, 0)
	test1 := rng.NextN(1)
	if test1 != 0 {
		t.Errorf("Bound of 1 did not give 0 (%x)", test1)
	}
	rand.Seed(0)
	for i := 0; i < 100000; i++ {
		bounds := uint64(rand.Int63())
		if rand.Uint32()%2 == 1 {
			bounds |= (1 << 63)
		}
		result := rng.NextN(bounds)
		if result >= bounds {
			t.Errorf("Got %x which is outside of bound %x", result, bounds)
		}
	}
}
