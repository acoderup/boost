package randx

import (
	. "github.com/davidminor/uint128"
)

type Permute64x32 func(uint64) uint32

type Pcg64x32 struct {
	lcg     *Lcg64
	permute Permute64x32
}

func NewPcg64x32(lcg *Lcg64, permuteFunc Permute64x32) *Pcg64x32 {
	// this is how the reference library sets the initial state
	lcg.State += lcg.Stream
	return &Pcg64x32{lcg, permuteFunc}
}

// Get the next random uint32 value
func (pcg *Pcg64x32) Next() uint32 {
	return pcg.permute(pcg.lcg.Next())
}

// Get a random uint32 value evenly distributed in [0,bounds)
func (pcg *Pcg64x32) NextN(bounds uint32) uint32 {
	threshold := (0 - bounds) % bounds
	for {
		result := pcg.Next()
		if result >= threshold {
			return result % bounds
		}
	}
}

// Set the stream of this PCG
func (rng *Pcg64x32) Stream(stream uint64) {
	// stream must be odd for LCG, so shift left 1 and turn on the 1 bit
	rng.lcg.Stream = stream<<1 | 1
}

// Pcg32 uses the top 37 bits of its 64 bit LCG, XOR'ing the highest half
// with the lowest, and then randomly rotating the lower 32 of them
// (which are returned)
type Pcg32 struct {
	*Pcg64x32
}

// Create a new Pcg32 with the given seed
func NewPcg32(seed uint64) Pcg32 {
	lcg := NewLcg64(seed)
	return Pcg32{NewPcg64x32(lcg, XshRr)}
}

// Create a new Pcg32 with the given seed and stream
func NewPcg32Stream(seed, stream uint64) Pcg32 {
	lcg := NewLcg64Stream(seed, stream)
	return Pcg32{NewPcg64x32(lcg, XshRr)}
}

// Take the highest 37 bits, xor the top half with the bottom,
// then use the top 5 to randomly rotate the next 32 (which we return)
func XshRr(state uint64) uint32 {
	rot := uint32(state >> 59)
	shift := uint32(((state >> 18) ^ state) >> 27)
	return (shift >> rot) | (shift << (32 - rot))
}

type Permute128x64 func(Uint128) uint64

type Pcg128x64 struct {
	lcg     *Lcg128
	permute Permute128x64
}

func NewPcg128x64(lcg *Lcg128, permuteFunc Permute128x64) *Pcg128x64 {
	// this is how the reference library sets the initial state
	lcg.State = lcg.State.Add(lcg.Stream)
	// reference library uses "post-advance" state of LCG, so we do an
	// initial advance to match its output
	lcg.Next()
	return &Pcg128x64{lcg, permuteFunc}
}

// Get the next random uint64 value
func (pcg *Pcg128x64) Next() uint64 {
	return pcg.permute(pcg.lcg.Next())
}

// Get a random uint64 value evenly distributed in [0,bounds)
func (pcg *Pcg128x64) NextN(bounds uint64) uint64 {
	threshold := (0 - bounds) % bounds
	for {
		result := pcg.Next()
		if result >= threshold {
			return result % bounds
		}
	}
}

// Set the stream of this PCG
func (rng *Pcg128x64) Stream(streamH, streamL uint64) {
	// stream must be odd for LCG, so shift left 1 and turn on the 1 bit
	rng.lcg.Stream = Uint128{streamH, streamL}.ShiftLeft(1)
	rng.lcg.Stream.L |= 1
}

// Pcg64 uses XOR of high and low bits combined with random shift
type Pcg64 struct {
	*Pcg128x64
}

// Create a new Pcg64 with the given high and low bits of seed
func NewPcg64(seedH, seedL uint64) Pcg64 {
	lcg := NewLcg128(Uint128{seedH, seedL})
	return Pcg64{NewPcg128x64(lcg, XslRr)}
}

// Create a new Pcg64 with the given high and low bits of seed and stream
func NewPcg64Stream(seedH, seedL, streamH, streamL uint64) Pcg64 {
	lcg := NewLcg128Stream(Uint128{seedH, seedL},
		Uint128{streamH, streamL})
	return Pcg64{NewPcg128x64(lcg, XslRr)}
}

// Permute functions

// Xor the state's top bits with the bottom and randomly rotate them
// based on the highest 6 bits.
func XslRr(state Uint128) uint64 {
	h, l := state.H, state.L
	shift := l ^ h
	rot := h >> 58
	return (shift >> rot) | (shift << (64 - rot))
}
