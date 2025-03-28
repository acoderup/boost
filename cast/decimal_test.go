package cast_test

import (
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/acoderup/boost/cast"
	. "github.com/frankban/quicktest"
)

func runDecimalTest(c *C, tests []testStep, tove func(any) (any, error), tov func(any) any) {
	c.Helper()

	for i, test := range tests {
		errmsg := Commentf("i = %d, test = %#v", i, test)

		v, err := tove(test.input)
		if test.iserr {
			c.Assert(err, IsNotNil, errmsg)
			continue
		}
		c.Assert(err, IsNil, errmsg)
		t := reflect.TypeOf(v)
		if t.Implements(reflect.TypeOf((*fmt.Stringer)(nil)).Elem()) {
			c.Assert(v.(fmt.Stringer).String(), Equals, test.expect.(fmt.Stringer).String(), errmsg)
		} else {
			fmt.Println(v, test.expect)
			c.Assert(v, Equals, test.expect, errmsg)
		}

		// Non-E test:
		v = tov(test.input)
		t = reflect.TypeOf(v)
		if t.Implements(reflect.TypeOf((*fmt.Stringer)(nil)).Elem()) {
			c.Assert(v.(fmt.Stringer).String(), Equals, test.expect.(fmt.Stringer).String(), errmsg)
		} else {
			c.Assert(v, Equals, test.expect, errmsg)
		}
	}
}

func createDecimalTestSteps(t testValues) []testStep {
	isUint := isUint(t.zero)

	return []testStep{
		// positive numbers
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
		{big.NewRat(8, 1), t.eight, false},
		{complex64(8 + 0i), t.eight, false},
		{complex128(8 + 0i), t.eight, false},
		{true, t.one, false},
		{"8", t.eight, false},
		{"true", t.one, false},
		{"(8+0i)", t.eight, false},
		{"8/1", t.eight, false},
		{[]byte{56}, t.eight, false},
		{stringer("8"), t.eight, false},
		{errors.New("8"), t.eight, false},
		{nil, t.zero, false},
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
		{big.NewRat(-8, 1), t.eightnegative, isUint},
		{complex64(-8 + 0i), t.eightnegative, isUint},
		{complex128(-8 + 0i), t.eightnegative, isUint},
		{false, t.zero, false},
		{"false", t.zero, false},
		{"-8", t.eightnegative, isUint},
		{[]byte{45, 56}, t.eightnegative, isUint},
		{stringer("-8"), t.eightnegative, isUint},
		{errors.New("-8"), t.eightnegative, isUint},
		// unexpected value or types
		{"test", t.zero, true},
		{testing.T{}, t.zero, true},
	}
}

func TestToIntE(t *testing.T) {
	tests := createDecimalTestSteps(testValues{
		int(0),
		int(1),
		int(8),
		int(-8),
		int(8),
		int(-8),
		int(8),
		int(-8),
	})

	runDecimalTest(
		New(t),
		tests,
		func(v any) (any, error) { return cast.ToIntE(v) },
		func(v any) any { return cast.ToInt(v) },
	)
}

func TestToInt8E(t *testing.T) {
	tests := createDecimalTestSteps(testValues{
		int8(0),
		int8(1),
		int8(8),
		int8(-8),
		int8(8),
		int8(-8),
		int8(8),
		int8(-8),
	})

	runDecimalTest(
		New(t),
		tests,
		func(v any) (any, error) { return cast.ToInt8E(v) },
		func(v any) any { return cast.ToInt8(v) },
	)
}

func TestToInt16E(t *testing.T) {
	tests := createDecimalTestSteps(testValues{
		int16(0),
		int16(1),
		int16(8),
		int16(-8),
		int16(8),
		int16(-8),
		int16(8),
		int16(-8),
	})

	runDecimalTest(
		New(t),
		tests,
		func(v any) (any, error) { return cast.ToInt16E(v) },
		func(v any) any { return cast.ToInt16(v) },
	)
}

func TestToInt32E(t *testing.T) {
	tests := createDecimalTestSteps(testValues{
		int32(0),
		int32(1),
		int32(8),
		int32(-8),
		int32(8),
		int32(-8),
		int32(8),
		int32(-8),
	})

	runDecimalTest(
		New(t),
		tests,
		func(v any) (any, error) { return cast.ToInt32E(v) },
		func(v any) any { return cast.ToInt32(v) },
	)
}

func TestToInt64E(t *testing.T) {
	tests := createDecimalTestSteps(testValues{
		int64(0),
		int64(1),
		int64(8),
		int64(-8),
		int64(8),
		int64(-8),
		int64(8),
		int64(-8),
	})

	runDecimalTest(
		New(t),
		tests,
		func(v any) (any, error) { return cast.ToInt64E(v) },
		func(v any) any { return cast.ToInt64(v) },
	)
}

func TestToUintE(t *testing.T) {
	tests := createDecimalTestSteps(testValues{
		uint(0),
		uint(1),
		uint(8),
		uint(0),
		uint(8),
		uint(8),
		uint(8),
		uint(8),
	})

	runDecimalTest(
		New(t),
		tests,
		func(v any) (any, error) { return cast.ToUintE(v) },
		func(v any) any { return cast.ToUint(v) },
	)
}

func TestToUint8E(t *testing.T) {
	tests := createDecimalTestSteps(testValues{
		uint8(0),
		uint8(1),
		uint8(8),
		uint8(0),
		uint8(8),
		uint8(8),
		uint8(8),
		uint8(8),
	})

	runDecimalTest(
		New(t),
		tests,
		func(v any) (any, error) { return cast.ToUint8E(v) },
		func(v any) any { return cast.ToUint8(v) },
	)
}

func TestToUint16E(t *testing.T) {
	tests := createDecimalTestSteps(testValues{
		uint16(0),
		uint16(1),
		uint16(8),
		uint16(0),
		uint16(8),
		uint16(8),
		uint16(8),
		uint16(8),
	})

	runDecimalTest(
		New(t),
		tests,
		func(v any) (any, error) { return cast.ToUint16E(v) },
		func(v any) any { return cast.ToUint16(v) },
	)
}

func TestToUint32E(t *testing.T) {
	tests := createDecimalTestSteps(testValues{
		uint32(0),
		uint32(1),
		uint32(8),
		uint32(0),
		uint32(8),
		uint32(8),
		uint32(8),
		uint32(8),
	})

	runDecimalTest(
		New(t),
		tests,
		func(v any) (any, error) { return cast.ToUint32E(v) },
		func(v any) any { return cast.ToUint32(v) },
	)
}

func TestToUint64E(t *testing.T) {
	tests := createDecimalTestSteps(testValues{
		uint64(0),
		uint64(1),
		uint64(8),
		uint64(0),
		uint64(8),
		uint64(8),
		uint64(8),
		uint64(8),
	})

	runDecimalTest(
		New(t),
		tests,
		func(v any) (any, error) { return cast.ToUint64E(v) },
		func(v any) any { return cast.ToUint64(v) },
	)
}

func TestToFloat32E(t *testing.T) {
	tests := createDecimalTestSteps(testValues{
		float32(0),
		float32(1),
		float32(8),
		float32(-8),
		float32(8.31),
		float32(-8.31),
		float32(8.31),
		float32(-8.31),
	})

	runDecimalTest(
		New(t),
		tests,
		func(v any) (any, error) { return cast.ToFloat32E(v) },
		func(v any) any { return cast.ToFloat32(v) },
	)
}

func TestToFloat64E(t *testing.T) {
	tests := createDecimalTestSteps(testValues{
		float64(0),
		float64(1),
		float64(8),
		float64(-8),
		float64(8.31),
		float64(-8.31),
		float64(float32(float64(8.31))),
		float64(float32(float64(-8.31))),
	})

	runDecimalTest(
		New(t),
		tests,
		func(v any) (any, error) { return cast.ToFloat64E(v) },
		func(v any) any { return cast.ToFloat64(v) },
	)
}

func TestToBigIntE(t *testing.T) {
	tests := createDecimalTestSteps(testValues{
		big.NewInt(0),
		big.NewInt(1),
		big.NewInt(8),
		big.NewInt(-8),
		big.NewInt(8),
		big.NewInt(-8),
		big.NewInt(8),
		big.NewInt(-8),
	})

	runDecimalTest(
		New(t),
		tests,
		func(v any) (any, error) { return cast.ToBigIntE(v) },
		func(v any) any { return cast.ToBigInt(v) },
	)
}

func TestToBigFloatE(t *testing.T) {
	f, _ := big.NewFloat(8.31).Float32()
	eightpoint31_32 := new(big.Float).SetFloat64(float64(f))
	f, _ = big.NewFloat(-8.31).Float32()
	eightpoint31negative_32 := new(big.Float).SetFloat64(float64(f))
	tests := createDecimalTestSteps(testValues{
		big.NewFloat(0),
		big.NewFloat(1),
		big.NewFloat(8),
		big.NewFloat(-8),
		big.NewFloat(8.31),
		big.NewFloat(-8.31),
		eightpoint31_32,
		eightpoint31negative_32,
	})

	runDecimalTest(
		New(t),
		tests,
		func(v any) (any, error) { return cast.ToBigFloatE(v) },
		func(v any) any { return cast.ToBigFloat(v) },
	)
}

func TestToBigRatE(t *testing.T) {
	f, _ := big.NewRat(0, 1).SetFloat64(8.31).Float32()
	eightpoint31_32 := new(big.Rat).SetFloat64(float64(f))
	f, _ = big.NewRat(0, 1).SetFloat64(-8.31).Float32()
	eightpoint31negative_32 := new(big.Rat).SetFloat64(float64(f))
	tests := createDecimalTestSteps(testValues{
		big.NewRat(0, 1),
		big.NewRat(1, 1),
		big.NewRat(8, 1),
		big.NewRat(-8, 1),
		big.NewRat(0, 1).SetFloat64(8.31),
		big.NewRat(0, 1).SetFloat64(-8.31),
		eightpoint31_32,
		eightpoint31negative_32,
	})

	runDecimalTest(
		New(t),
		tests,
		func(v any) (any, error) { return cast.ToBigRatE(v) },
		func(v any) any { return cast.ToBigRat(v) },
	)
}

func TestToComplex64(t *testing.T) {
	tests := createDecimalTestSteps(testValues{
		complex64(0 + 0i),
		complex64(1 + 0i),
		complex64(8 + 0i),
		complex64(-8 + 0i),
		complex64(8.31 + 0i),
		complex64(-8.31 + 0i),
		complex64(8.31 + 0i),
		complex64(-8.31 + 0i),
	})

	runDecimalTest(
		New(t),
		tests,
		func(v any) (any, error) { return cast.ToComplex64E(v) },
		func(v any) any { return cast.ToComplex64(v) },
	)
}

func TestToComplex128(t *testing.T) {
	tests := createDecimalTestSteps(testValues{
		complex128(0 + 0i),
		complex128(1 + 0i),
		complex128(8 + 0i),
		complex128(-8 + 0i),
		complex128(8.31 + 0i),
		complex128(-8.31 + 0i),
		complex128(complex64(complex128(8.31 + 0i))),
		complex128(complex64(complex128(-8.31 + 0i))),
	})

	runDecimalTest(
		New(t),
		tests,
		func(v any) (any, error) { return cast.ToComplex128E(v) },
		func(v any) any { return cast.ToComplex128(v) },
	)
}

func TestToBoolE(t *testing.T) {
	tests := createDecimalTestSteps(testValues{
		false,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
	})

	runDecimalTest(
		New(t),
		tests,
		func(v any) (any, error) { return cast.ToBoolE(v) },
		func(v any) any { return cast.ToBool(v) },
	)
}
