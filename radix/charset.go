package radix

import (
	"math/rand"
)

type Charset struct {
	charset    []byte
	charsetMap map[byte]int
}

func NewCharset(charset []byte) *Charset {
	c := &Charset{
		charset:    make([]byte, len(charset)),
		charsetMap: make(map[byte]int),
	}
	copy(c.charset, charset)

	for i, b := range c.charset {
		c.charsetMap[b] = i
	}

	return c
}

func (c *Charset) Shuffle(seed int64) *Charset {
	r := rand.New(rand.NewSource(seed))
	r.Shuffle(len(c.charset), func(i, j int) {
		c.charset[i], c.charset[j] = c.charset[j], c.charset[i]
	})

	for i, b := range c.charset {
		c.charsetMap[b] = i
	}

	return c
}

func (c *Charset) Size() int {
	return len(c.charset)
}
