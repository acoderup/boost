package cast_test

import (
	"errors"
	"fmt"
	"math/big"
	"testing"

	. "github.com/frankban/quicktest"
	"github.com/tomas-qstarrs/boost/cast"
)

func runStringTest(c *C, tests []testStep, tove func(any) (any, error), tov func(any) any) {
	c.Helper()

	for i, test := range tests {
		errmsg := Commentf("i = %d, test = %#v", i, test)

		v, err := tove(test.input)
		if test.iserr {
			c.Assert(err, IsNotNil, errmsg)
			continue
		}
		c.Assert(err, IsNil, errmsg)
		switch v := v.(type) {
		case string:
			c.Assert(v, Equals, test.expect, errmsg)
		case []byte:
			c.Assert(v, DeepEquals, test.expect, errmsg)
		case fmt.Stringer:
			c.Assert(v.String(), Equals, test.expect.(fmt.Stringer).String(), errmsg)
		case error:
			c.Assert(v.Error(), Equals, test.expect.(error).Error(), errmsg)
		default:
			c.Assert(v, Equals, test.expect, errmsg)
		}

		// Non-E test:
		v = tov(test.input)
		switch v := v.(type) {
		case string:
			c.Assert(v, Equals, test.expect, errmsg)
		case []byte:
			c.Assert(v, DeepEquals, test.expect, errmsg)
		case fmt.Stringer:
			c.Assert(v.String(), Equals, test.expect.(fmt.Stringer).String(), errmsg)
		case error:
			c.Assert(v.Error(), Equals, test.expect.(error).Error(), errmsg)
		default:
			c.Assert(v, Equals, test.expect, errmsg)
		}
	}
}

func createStringTestSteps() []testStep {
	isUint := false
	t := testValues{"0", "1", "8", "-8", "8.31", "-8.31", "8.31", "-8.31"}
	return []testStep{ // positive numbers
		{int(8), t.eight, false},
		{int8(8), t.eight, false},
		{int16(8), t.eight, false},
		{int32(8), t.eight, false},
		{int64(8), t.eight, false},
		{uint(8), t.eight, false},
		{uint8(8), t.eight, false},
		{uint16(8), t.eight, false},
		{uint32(8), t.eight, false},
		{uint64(8), t.eight, false},
		{float32(8.31), t.eightpoint31_32, false},
		{float64(8.31), t.eightpoint31, false},
		{big.NewInt(8), t.eight, false},
		{big.NewFloat(8.31), t.eightpoint31, false},
		{big.NewRat(8, 1), "8/1", false},      // special case
		{complex64(8 + 0i), "(8+0i)", false},  // special case
		{complex128(8 + 0i), "(8+0i)", false}, // special case
		{true, "true", false},                 // special case
		{"8", t.eight, false},
		{[]byte{56}, t.eight, false},
		{stringer("8"), t.eight, false},
		{errors.New("8"), t.eight, false},
		{nil, "", false},
		// negative numbers
		{int(-8), t.eightnegative, isUint},
		{int8(-8), t.eightnegative, isUint},
		{int16(-8), t.eightnegative, isUint},
		{int32(-8), t.eightnegative, isUint},
		{int64(-8), t.eightnegative, isUint},
		{float32(-8.31), t.eightpoint31negative_32, isUint},
		{float64(-8.31), t.eightpoint31negative, isUint},
		{big.NewInt(-8), t.eightnegative, isUint},
		{big.NewFloat(-8.31), t.eightpoint31negative, isUint},
		{big.NewRat(-8, 1), "-8/1", isUint},      // special case
		{complex64(-8 + 0i), "(-8+0i)", isUint},  // special case
		{complex128(-8 + 0i), "(-8+0i)", isUint}, // special case
		{false, "false", false},                  // special case
		{"-8", t.eightnegative, isUint},
		{[]byte{45, 56}, t.eightnegative, isUint},
		{stringer("-8"), t.eightnegative, isUint},
		{errors.New("-8"), t.eightnegative, isUint},
		// unexpected value or types
		{"test", "test", false},
		{testing.T{}, "", true},
	}
}

func TestToStringE(t *testing.T) {
	tests := createStringTestSteps()
	runStringTest(
		New(t),
		tests,
		func(v any) (any, error) { return cast.ToStringE(v) },
		func(v any) any { return cast.ToString(v) },
	)
}

func createBytesTestSteps() []testStep {
	isUint := false
	t := testValues{[]byte("0"), []byte("1"), []byte("8"), []byte("-8"), []byte("8.31"), []byte("-8.31"), []byte("8.31"), []byte("-8.31")}
	return []testStep{ // positive numbers
		{int(8), t.eight, false},
		{int8(8), t.eight, false},
		{int16(8), t.eight, false},
		{int32(8), t.eight, false},
		{int64(8), t.eight, false},
		{uint(8), t.eight, false},
		{uint8(8), t.eight, false},
		{uint16(8), t.eight, false},
		{uint32(8), t.eight, false},
		{uint64(8), t.eight, false},
		{float32(8.31), t.eightpoint31_32, false},
		{float64(8.31), t.eightpoint31, false},
		{big.NewInt(8), t.eight, false},
		{big.NewFloat(8.31), t.eightpoint31, false},
		{big.NewRat(8, 1), []byte("8/1"), false},      // special case
		{complex64(8 + 0i), []byte("(8+0i)"), false},  // special case
		{complex128(8 + 0i), []byte("(8+0i)"), false}, // special case
		{true, []byte("true"), false},                 // special case
		{"8", t.eight, false},
		{[]byte{56}, t.eight, false},
		{stringer("8"), t.eight, false},
		{errors.New("8"), t.eight, false},
		{nil, []byte{}, false},
		// negative numbers
		{int(-8), t.eightnegative, isUint},
		{int8(-8), t.eightnegative, isUint},
		{int16(-8), t.eightnegative, isUint},
		{int32(-8), t.eightnegative, isUint},
		{int64(-8), t.eightnegative, isUint},
		{float32(-8.31), t.eightpoint31negative_32, isUint},
		{float64(-8.31), t.eightpoint31negative, isUint},
		{big.NewInt(-8), t.eightnegative, isUint},
		{big.NewFloat(-8.31), t.eightpoint31negative, isUint},
		{big.NewRat(-8, 1), []byte("-8/1"), isUint},      // special case
		{complex64(-8 + 0i), []byte("(-8+0i)"), isUint},  // special case
		{complex128(-8 + 0i), []byte("(-8+0i)"), isUint}, // special case
		{false, []byte("false"), false},                  // special case
		{"-8", t.eightnegative, isUint},
		{[]byte{45, 56}, t.eightnegative, isUint},
		{stringer("-8"), t.eightnegative, isUint},
		{errors.New("-8"), t.eightnegative, isUint},
		// unexpected value or types
		{"test", []byte("test"), false},
		{testing.T{}, nil, true},
	}
}

func TestToBytesE(t *testing.T) {
	tests := createBytesTestSteps()
	runStringTest(
		New(t),
		tests,
		func(v any) (any, error) { return cast.ToBytesE(v) },
		func(v any) any { return cast.ToBytes(v) },
	)
}

type stringer string

func (s stringer) String() string {
	return string(s)
}

func createStringerTestSteps() []testStep {
	isUint := false
	t := testValues{stringer("0"), stringer("1"), stringer("8"), stringer("-8"), stringer("8.31"), stringer("-8.31"), stringer("8.31"), stringer("-8.31")}
	return []testStep{ // positive numbers
		{int(8), t.eight, false},
		{int8(8), t.eight, false},
		{int16(8), t.eight, false},
		{int32(8), t.eight, false},
		{int64(8), t.eight, false},
		{uint(8), t.eight, false},
		{uint8(8), t.eight, false},
		{uint16(8), t.eight, false},
		{uint32(8), t.eight, false},
		{uint64(8), t.eight, false},
		{float32(8.31), t.eightpoint31_32, false},
		{float64(8.31), t.eightpoint31, false},
		{big.NewInt(8), t.eight, false},
		{big.NewFloat(8.31), t.eightpoint31, false},
		{big.NewRat(8, 1), stringer("8/1"), false},      // special case
		{complex64(8 + 0i), stringer("(8+0i)"), false},  // special case
		{complex128(8 + 0i), stringer("(8+0i)"), false}, // special case
		{true, stringer("true"), false},                 // special case
		{"8", t.eight, false},
		{[]byte{56}, t.eight, false},
		{stringer("8"), t.eight, false},
		{errors.New("8"), t.eight, false},
		{nil, nil, false},
		// negative numbers
		{int(-8), t.eightnegative, isUint},
		{int8(-8), t.eightnegative, isUint},
		{int16(-8), t.eightnegative, isUint},
		{int32(-8), t.eightnegative, isUint},
		{int64(-8), t.eightnegative, isUint},
		{float32(-8.31), t.eightpoint31negative_32, isUint},
		{float64(-8.31), t.eightpoint31negative, isUint},
		{big.NewInt(-8), t.eightnegative, isUint},
		{big.NewFloat(-8.31), t.eightpoint31negative, isUint},
		{big.NewRat(-8, 1), stringer("-8/1"), isUint},      // special case
		{complex64(-8 + 0i), stringer("(-8+0i)"), isUint},  // special case
		{complex128(-8 + 0i), stringer("(-8+0i)"), isUint}, // special case
		{false, stringer("false"), false},                  // special case
		{"-8", t.eightnegative, isUint},
		{[]byte{45, 56}, t.eightnegative, isUint},
		{stringer("-8"), t.eightnegative, isUint},
		{errors.New("-8"), t.eightnegative, isUint},
		// unexpected value or types
		{"test", stringer("test"), false},
		{testing.T{}, nil, true},
	}
}

func TestToStringerE(t *testing.T) {
	tests := createStringerTestSteps()
	runStringTest(
		New(t),
		tests,
		func(v any) (any, error) { return cast.ToStringerE(v) },
		func(v any) any { return cast.ToStringer(v) },
	)
}

func createErrorTestSteps() []testStep {
	isUint := false
	t := testValues{errors.New("0"), errors.New("1"), errors.New("8"), errors.New("-8"), errors.New("8.31"), errors.New("-8.31"), errors.New("8.31"), errors.New("-8.31")}
	return []testStep{ // positive numbers
		{int(8), t.eight, false},
		{int8(8), t.eight, false},
		{int16(8), t.eight, false},
		{int32(8), t.eight, false},
		{int64(8), t.eight, false},
		{uint(8), t.eight, false},
		{uint8(8), t.eight, false},
		{uint16(8), t.eight, false},
		{uint32(8), t.eight, false},
		{uint64(8), t.eight, false},
		{float32(8.31), t.eightpoint31_32, false},
		{float64(8.31), t.eightpoint31, false},
		{big.NewInt(8), t.eight, false},
		{big.NewFloat(8.31), t.eightpoint31, false},
		{big.NewRat(8, 1), errors.New("8/1"), false},      // special case
		{complex64(8 + 0i), errors.New("(8+0i)"), false},  // special case
		{complex128(8 + 0i), errors.New("(8+0i)"), false}, // special case
		{true, errors.New("true"), false},                 // special case
		{"8", t.eight, false},
		{[]byte{56}, t.eight, false},
		{stringer("8"), t.eight, false},
		{errors.New("8"), t.eight, false},
		{nil, nil, false},
		// negative numbers
		{int(-8), t.eightnegative, isUint},
		{int8(-8), t.eightnegative, isUint},
		{int16(-8), t.eightnegative, isUint},
		{int32(-8), t.eightnegative, isUint},
		{int64(-8), t.eightnegative, isUint},
		{float32(-8.31), t.eightpoint31negative_32, isUint},
		{float64(-8.31), t.eightpoint31negative, isUint},
		{big.NewInt(-8), t.eightnegative, isUint},
		{big.NewFloat(-8.31), t.eightpoint31negative, isUint},
		{big.NewRat(-8, 1), errors.New("-8/1"), isUint},      // special case
		{complex64(-8 + 0i), errors.New("(-8+0i)"), isUint},  // special case
		{complex128(-8 + 0i), errors.New("(-8+0i)"), isUint}, // special case
		{false, errors.New("false"), false},                  // special case
		{"-8", t.eightnegative, isUint},
		{[]byte{45, 56}, t.eightnegative, isUint},
		{stringer("-8"), t.eightnegative, isUint},
		{errors.New("-8"), t.eightnegative, isUint},
		// unexpected value or types
		{"test", errors.New("test"), false},
		{testing.T{}, nil, true},
	}
}

func TestToErrorE(t *testing.T) {
	tests := createErrorTestSteps()
	runStringTest(
		New(t),
		tests,
		func(v any) (any, error) { return cast.ToErrorE(v) },
		func(v any) any { return cast.ToError(v) },
	)
}
