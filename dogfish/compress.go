package dogfish

import (
	"fmt"
	"strings"

	"github.com/golang/snappy"
)

type CompressType string

const (
	CompressTypeEmpty  CompressType = ""
	CompressTypeSnappy CompressType = "snappy"
)

var (
	compressType    CompressType
	compressMinSize int
)

func CompressWith(typ CompressType, minSize int) {
	compressType = typ
	compressMinSize = minSize
}

func Compress(s string) string {
	if len(s) < compressMinSize {
		return s
	}
	if strings.HasPrefix(s, "(") {
		return s
	}
	if compressType == CompressTypeSnappy {
		b := snappy.Encode(nil, []byte(s))
		s = fmt.Sprintf("(%s)%s", CompressTypeSnappy, string(b))
	}
	return s
}

func Decompress(s string) string {
	if len(s) == 0 {
		return s
	}
	if !strings.HasPrefix(s, "(") {
		return s
	}
	if strings.HasPrefix(s, fmt.Sprintf("(%s)", CompressTypeSnappy)) {
		b := []byte(s)
		b = b[len(CompressTypeSnappy)+2:]
		b, err := snappy.Decode(nil, b)
		if err != nil {
			panic(fmt.Errorf("%s, Hashtree decompress failed, value=%#v",
				err.Error(), s))
		}
		s = string(b)
	}
	return s
}
