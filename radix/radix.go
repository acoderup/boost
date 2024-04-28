package radix

import (
	"math"
)

type Radix struct {
	*Charset
	endian  Endian
	base    int
	maxSize int // -1 means no limit
}

func New(endian Endian, maxSize int, c *Charset) *Radix {
	r := &Radix{
		Charset: c,
		endian:  endian,
		base:    c.Size(),
		maxSize: maxSize,
	}
	return r
}

func (r *Radix) Encode(n uint64) []byte {
	n = r.validate(n)
	var data []byte
	if r.maxSize > 0 {
		data = make([]byte, 0, r.maxSize)
	} else {
		data = make([]byte, 0)
	}
	for n > 0 {
		data = append(data, r.charset[n%uint64(r.base)])
		n /= uint64(r.base)
	}
	if r.endian == BigEndian {
		r.reverseEndian(data)
	}
	return data
}

func (r *Radix) Decode(data []byte) uint64 {
	if r.endian == BigEndian {
		var copyData = make([]byte, len(data))
		copy(copyData, data)
		r.reverseEndian(copyData)
		data = copyData
	}

	var n uint64
	for i := len(data) - 1; i >= 0; i-- {
		n = n*uint64(r.base) + uint64(r.charsetMap[data[i]])
	}
	return r.validate(n)
}

func (r *Radix) validate(i uint64) uint64 {
	if r.maxSize > 0 {
		return i % uint64(math.Pow(float64(r.base), float64(r.maxSize)))
	}
	return i
}

func (r *Radix) reverseEndian(data []byte) {
	for i := 0; i < len(data)/2; i++ {
		data[i], data[len(data)-1-i] = data[len(data)-1-i], data[i]
	}
}
