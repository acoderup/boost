package stringx_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomas-qstarrs/boost/stringx"
)

func TestCompareVersion(t *testing.T) {
	assert.Equal(t, stringx.CompareVersion("1.0.0", "1.0.0"), 0)
	assert.Equal(t, stringx.CompareVersion("1.0.0", "1.0.1"), -1)
	assert.Equal(t, stringx.CompareVersion("1.0.0", ""), 1)
}
