package dogfish

import (
	"fmt"
	"math/big"
	"reflect"
	"strconv"

	"github.com/shopspring/decimal"
)

func strconvFromString(s string, i interface{}) error {
	switch v := i.(type) {
	case *int:
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		*v = int(n)
	case *int8:
		n, err := strconv.ParseInt(s, 10, 8)
		if err != nil {
			return err
		}
		*v = int8(n)
	case *int16:
		n, err := strconv.ParseInt(s, 10, 16)
		if err != nil {
			return err
		}
		*v = int16(n)
	case *int32:
		n, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return err
		}
		*v = int32(n)
	case *int64:
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		*v = n
	case *uint:
		n, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return err
		}
		*v = uint(n)
	case *uint8:
		n, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			return err
		}
		*v = uint8(n)
	case *uint16:
		n, err := strconv.ParseUint(s, 10, 16)
		if err != nil {
			return err
		}
		*v = uint16(n)
	case *uint32:
		n, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			return err
		}
		*v = uint32(n)
	case *uint64:
		n, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return err
		}
		*v = n
	case *bool:
		t, err := strconv.ParseBool(s)
		if err != nil {
			return err
		}
		*v = t
	default:
		return fmt.Errorf("invalid type from string:%+v, %+v", s, reflect.TypeOf(v))
	}
	return nil
}

func decimalFromString(s string, i interface{}) error {
	switch v := i.(type) {
	case *float32:
		d, err := decimal.NewFromString(s)
		if err != nil {
			return err
		}
		f, _ := d.Float64()
		*v = float32(f)
	case *float64:
		d, err := decimal.NewFromString(s)
		if err != nil {
			return err
		}
		f, _ := d.Float64()
		*v = f
	default:
		return fmt.Errorf("invalid type from string:%+v, %+v", s, reflect.TypeOf(v))
	}
	return nil
}

func fmtFromString(s string, i interface{}) error {
	switch v := i.(type) {
	case *complex64:
		// New version of strconv required
		// n, err := strconv.ParseComplex(s, 64)
		// if err != nil {
		// 	return err
		// }
		// *v = complex64(n)

		// Use fmt.Sscan instead
		_, err := fmt.Sscan(s, v)
		if err != nil {
			return err
		}
	case *complex128:
		// new version of strconv required
		// n, err := strconv.ParseComplex(s, 128)
		// if err != nil {
		// 	return err
		// }
		// *v = complex128(n)

		// Use fmt.Sscan instead
		_, err := fmt.Sscan(s, v)
		if err != nil {
			return err
		}
	case *uintptr:
		_, err := fmt.Sscan(s, v)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid type from string:%+v, %+v", s, reflect.TypeOf(v))
	}
	return nil
}

func bigFromString(s string, i interface{}) error {
	switch v := i.(type) {
	case *big.Int:
		b, ok := new(big.Int).SetString(s, 10)
		if !ok {
			return fmt.Errorf("big.Int set string %s failed", s)
		}
		*v = *b
	case *big.Rat:
		b, ok := new(big.Rat).SetString(s)
		if !ok {
			return fmt.Errorf("big.Rat set string %s failed", s)
		}
		*v = *b
	case *big.Float:
		b, ok := new(big.Float).SetString(s)
		if !ok {
			return fmt.Errorf("big.Float set string %s failed", s)
		}
		*v = *b
	default:
		return fmt.Errorf("invalid type from string:%+v, %+v", s, reflect.TypeOf(v))
	}
	return nil
}

// FromString convert string to some type
func FromString(s string, i interface{}) error {
	switch v := i.(type) {
	case *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *bool:
		return strconvFromString(s, i)
	case *float32, *float64:
		return decimalFromString(s, i)
	case *complex64, *complex128, *uintptr:
		return fmtFromString(s, i)
	case *big.Int, *big.Rat, *big.Float:
		return bigFromString(s, i)
	case *string:
		*v = s
	case *[]byte:
		*v = []byte(s)
	case *[]rune:
		*v = []rune(s)
	default:
		return fmt.Errorf("invalid type from string:%+v, %+v", s, reflect.TypeOf(v))
	}
	return nil
}

func strconvToString(i interface{}) (string, error) {
	var s string
	switch v := i.(type) {
	case int:
		s = strconv.FormatInt(int64(v), 10)
	case int8:
		s = strconv.FormatInt(int64(v), 10)
	case int16:
		s = strconv.FormatInt(int64(v), 10)
	case int32:
		s = strconv.FormatInt(int64(v), 10)
	case int64:
		s = strconv.FormatInt(v, 10)
	case uint:
		s = strconv.FormatUint(uint64(v), 10)
	case uint8:
		s = strconv.FormatUint(uint64(v), 10)
	case uint16:
		s = strconv.FormatUint(uint64(v), 10)
	case uint32:
		s = strconv.FormatUint(uint64(v), 10)
	case uint64:
		s = strconv.FormatUint(v, 10)
	case bool:
		s = strconv.FormatBool(v)
	default:
		return "", fmt.Errorf("invalid type to string:%+v, %+v", i, reflect.TypeOf(i))
	}
	return s, nil
}

func decimalToString(i interface{}) (string, error) {
	var s string
	switch v := i.(type) {
	case float32:
		// Use decimal to fix precision issue, FormatFloat is instable.
		s = decimal.NewFromFloat32(v).String()
	case float64:
		// Use decimal to fix precision issue, FormatFloat is instable.
		s = decimal.NewFromFloat(v).String()
	default:
		return "", fmt.Errorf("invalid type to string:%+v, %+v", i, reflect.TypeOf(i))
	}
	return s, nil
}

func fmtToString(i interface{}) (string, error) {
	var s string
	switch v := i.(type) {
	case complex64:
		// New version of strconv required
		// s = strconv.FormatComplex(v, 10)

		// Use fmt.Sprint instead
		s = fmt.Sprint(v)
	case complex128:
		// New version of strconv required
		// s = strconv.FormatComplex(v, 10)

		// Use fmt.Sprint instead
		s = fmt.Sprint(v)
	case uintptr:
		s = fmt.Sprint(v)
	default:
		return "", fmt.Errorf("invalid type to string:%+v, %+v", i, reflect.TypeOf(i))
	}
	return s, nil
}

func bigToString(i interface{}) (string, error) {
	var s string
	switch v := i.(type) {
	case big.Int:
		s = v.String()
	case big.Rat:
		s = v.String()
	case big.Float:
		s = v.String()
	case *big.Int:
		s = v.String()
	case *big.Rat:
		s = v.String()
	case *big.Float:
		s = v.String()
	default:
		return "", fmt.Errorf("invalid type to string:%+v, %+v", i, reflect.TypeOf(i))
	}
	return s, nil
}

// ToString convert some type to string
func ToString(i interface{}) (string, error) {
	var s string
	switch v := i.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, bool:
		return strconvToString(i)
	case float32, float64:
		return decimalToString(i)
	case complex64, complex128, uintptr:
		return fmtToString(i)
	case big.Int, big.Rat, big.Float, *big.Int, *big.Rat, *big.Float:
		return bigToString(i)
	case string:
		s = v
	case []byte:
		s = string(v)
	case []rune:
		s = string(v)
	default:
		s = fmt.Sprint(v)
	}
	return s, nil
}
