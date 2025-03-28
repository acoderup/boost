package dogfish_test

import (
	"testing"

	. "github.com/acoderup/boost/dogfish"
)

func TestCompress(t *testing.T) {
	CompressWith(CompressTypeSnappy, 10)
	plain := "1234567890123456789012345678901234567890123456789012345678901234567890"
	compressed := Compress(plain)
	decompressed := Decompress(compressed)
	if plain != decompressed {
		t.Errorf("plain=%s, compressed=%s, decompressed=%s", plain, compressed, decompressed)
	}
}
