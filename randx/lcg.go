// Copyright 2014, David Minor. All rights reserved.
// Use of this source code is governed by the MIT
// license which can be found in the LICENSE file.

package randx

import (
	. "github.com/davidminor/uint128"
)

var default_multiplier32 = uint32(747796405)
var default_stream32 = uint32(2891336453)

type Lcg32 struct {
	State, Stream, Multiplier uint32
}

func (lcg *Lcg32) Next() uint32 {
	lcg.State = lcg.State*lcg.Multiplier + lcg.Stream
	return lcg.State
}

func NewLcg32(seed uint32) *Lcg32 {
	return &Lcg32{seed, default_stream32, default_multiplier32}
}

func NewLcg32Stream(seed, stream uint32) *Lcg32 {
	stream = stream<<1 | 1
	return &Lcg32{seed, stream, default_multiplier32}
}

func (lcg *Lcg32) Int63() int64 {
	return int64(lcg.Next()>>1) | int64(lcg.Next())<<30
}

func (lcg *Lcg32) Seed(s int64) {
	lcg.State = uint32(s)
}

var default_multiplier64 = uint64(6364136223846793005)
var default_stream64 = uint64(1442695040888963407)

type Lcg64 struct {
	State, Stream, Multiplier uint64
}

func (lcg *Lcg64) Next() uint64 {
	lcg.State = lcg.State*lcg.Multiplier + lcg.Stream
	return lcg.State
}

func NewLcg64(seed uint64) *Lcg64 {
	return &Lcg64{seed, default_stream64, default_multiplier64}
}

func NewLcg64Stream(seed, stream uint64) *Lcg64 {
	stream = stream<<1 | 1
	return &Lcg64{seed, stream, default_multiplier64}
}

func (lcg *Lcg64) Int63() int64 {
	return int64(lcg.Next() >> 1)
}

func (lcg *Lcg64) Seed(s int64) {
	lcg.State = uint64(s)
}

var default_multiplier128 = Uint128{2549297995355413924, 4865540595714422341}
var default_stream128 = Uint128{6364136223846793005, 1442695040888963407}

type Lcg128 struct {
	State, Stream, Multiplier Uint128
}

func (lcg *Lcg128) Next() Uint128 {
	lcg.State = lcg.State.Mult(lcg.Multiplier).Add(lcg.Stream)
	return lcg.State
}

func NewLcg128(seed Uint128) *Lcg128 {
	return &Lcg128{seed, default_stream128, default_multiplier128}
}

func NewLcg128Stream(seed, stream Uint128) *Lcg128 {
	stream = stream.ShiftLeft(1)
	stream.L |= 1
	return &Lcg128{seed, stream, default_multiplier128}
}

func (lcg *Lcg128) Int63() int64 {
	n := lcg.Next()
	return int64((n.H ^ n.L) >> 1)
}

func (lcg *Lcg128) Seed(s int64) {
	lcg64 := NewLcg64(uint64(s))
	lcg.State.H, lcg.State.L = lcg64.Next(), lcg64.Next()
}
