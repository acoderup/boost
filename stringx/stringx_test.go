package stringx_test

import (
	"testing"

	"github.com/acoderup/boost/stringx"
	"github.com/stretchr/testify/assert"
)

func TestCompareVersion(t *testing.T) {
	assert.Equal(t, stringx.CompareVersion("1.0.0", "1.0.0"), 0)
	assert.Equal(t, stringx.CompareVersion("1.0.0", "1.0.1"), -1)
	assert.Equal(t, stringx.CompareVersion("1.0.0", ""), 1)
}
