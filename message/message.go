package message

import (
	"github.com/acoderup/boost/encoding"
	"github.com/acoderup/boost/route"
)

type Message struct {
	ID       uint64
	Route    route.Route
	Encoding encoding.Encoding
	Data     []byte
}
