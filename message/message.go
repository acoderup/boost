package message

import (
	"github.com/tomas-qstarrs/boost/encoding"
	"github.com/tomas-qstarrs/boost/route"
)

type Message struct {
	ID       uint64
	Route    route.Route
	Encoding encoding.Encoding
	Data     []byte
}
