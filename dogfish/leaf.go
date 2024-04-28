package dogfish

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/tomas-qstarrs/boost/encoding"
)

var (
	json  = encoding.NewJSON()
	proto = encoding.NewProtobuf()
)

// Int is a wrapper for int.
type Int struct {
	_root  *Root
	_key   string
	_value int
}

// Get is a getter for Int
func (f *Int) Get() int {
	return f._value
}

// SafeGet is a safe getter for Int
func (f *Int) SafeGet() int {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for Int
func (f *Int) Set(value int) {
	if value == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for Int
func (f *Int) SafeSet(value int) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// Int8 is a wrapper for int8.
type Int8 struct {
	_root  *Root
	_key   string
	_value int8
}

// Get is a getter for Int8
func (f *Int8) Get() int8 {
	return f._value
}

// SafeGet is a safe getter for Int8
func (f *Int8) SafeGet() int8 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for Int8
func (f *Int8) Set(value int8) {
	if value == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for Int8
func (f *Int8) SafeSet(value int8) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// Int16 is a wrapper for int16.
type Int16 struct {
	_root  *Root
	_key   string
	_value int16
}

// Get is a getter for Int16
func (f *Int16) Get() int16 {
	return f._value
}

// SafeGet is a safe getter for Int16
func (f *Int16) SafeGet() int16 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for Int16
func (f *Int16) Set(value int16) {
	if value == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for Int16
func (f *Int16) SafeSet(value int16) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// Int32 is a wrapper for int32.
type Int32 struct {
	_root  *Root
	_key   string
	_value int32
}

// Get is a getter for Int32
func (f *Int32) Get() int32 {
	return f._value
}

// SafeGet is a safe getter for Int32
func (f *Int32) SafeGet() int32 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for Int32
func (f *Int32) Set(value int32) {
	if value == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for Int32
func (f *Int32) SafeSet(value int32) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// Int64 is a wrapper for int64.
type Int64 struct {
	_root  *Root
	_key   string
	_value int64
}

// Get is a getter for Int64
func (f *Int64) Get() int64 {
	return f._value
}

// SafeGet is a safe getter for Int64
func (f *Int64) SafeGet() int64 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for Int64
func (f *Int64) Set(value int64) {
	if value == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for Int64
func (f *Int64) SafeSet(value int64) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// Uint is a wrapper for uint.
type Uint struct {
	_root  *Root
	_key   string
	_value uint
}

// Get is a getter for Uint
func (f *Uint) Get() uint {
	return f._value
}

// SafeGet is a safe getter for Uint
func (f *Uint) SafeGet() uint {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for Uint
func (f *Uint) Set(value uint) {
	if value == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for Uint
func (f *Uint) SafeSet(value uint) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// Uint8 is a wrapper for uint8.
type Uint8 struct {
	_root  *Root
	_key   string
	_value uint8
}

// Get is a getter for Uint8
func (f *Uint8) Get() uint8 {
	return f._value
}

// SafeGet is a safe getter for Uint8
func (f *Uint8) SafeGet() uint8 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a getter for Uint8
func (f *Uint8) Set(value uint8) {
	if value == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for Uint8
func (f *Uint8) SafeSet(value uint8) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// Uint16 is a wrapper for uint16.
type Uint16 struct {
	_root  *Root
	_key   string
	_value uint16
}

// Get is a getter for Uint16
func (f *Uint16) Get() uint16 {
	return f._value
}

// SafeGet is a safe getter for Uint16
func (f *Uint16) SafeGet() uint16 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for Uint16
func (f *Uint16) Set(value uint16) {
	if value == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for Uint16
func (f *Uint16) SafeSet(value uint16) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// Uint32 is a wrapper for int8.
type Uint32 struct {
	_root  *Root
	_key   string
	_value uint32
}

// Get is a getter for Uint32
func (f *Uint32) Get() uint32 {
	return f._value
}

// SafeGet is a safe getter for Uint32
func (f *Uint32) SafeGet() uint32 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for Uint32
func (f *Uint32) Set(value uint32) {
	if value == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for Uint32
func (f *Uint32) SafeSet(value uint32) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// Uint64 is a wrapper for uint64.
type Uint64 struct {
	_root  *Root
	_key   string
	_value uint64
}

// Get is a getter for Uint64
func (f *Uint64) Get() uint64 {
	return f._value
}

// SafeGet is a safe getter for Uint64
func (f *Uint64) SafeGet() uint64 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for Uint64
func (f *Uint64) Set(value uint64) {
	if value == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for Uint64
func (f *Uint64) SafeSet(value uint64) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// Float32 is a wrapper for float32.
type Float32 struct {
	_root  *Root
	_key   string
	_value float32
}

// Get is a getter for Float32
func (f *Float32) Get() float32 {
	return f._value
}

// SafeGet is a safe getter for Float32
func (f *Float32) SafeGet() float32 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for Float32
func (f *Float32) Set(value float32) {
	if value == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for Float32
func (f *Float32) SafeSet(value float32) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// Float64 is a wrapper for float64.
type Float64 struct {
	_root  *Root
	_key   string
	_value float64
}

// Get is a getter for Float64
func (f *Float64) Get() float64 {
	return f._value
}

// SafeGet is a safe getter for Float64
func (f *Float64) SafeGet() float64 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for Float64
func (f *Float64) Set(value float64) {
	if value == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for Float64
func (f *Float64) SafeSet(value float64) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// BigInt is a wrapper for big.Int
type BigInt struct {
	_root  *Root
	_key   string
	_value string
}

// Get is a getter for BigInt
func (f *BigInt) Get() int64 {
	if f._value == "" {
		return 0
	}
	n, ok := new(big.Int).SetString(f._value, 10)
	if !ok {
		panic(fmt.Errorf("%s, Hashtree BigInt Parse failed, value=%#v",
			"big.Int SetString error", f._value))
	}
	return n.Int64()
}

// SafeGet is a safe getter for BigInt
func (f *BigInt) SafeGet() int64 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for BigInt
func (f *BigInt) Set(value int64) {
	strValue := big.NewInt(value).String()
	if strValue == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = strValue
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for BigInt
func (f *BigInt) SafeSet(value int64) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// GetBig is a getter for BigInt
func (f *BigInt) GetBig() *big.Int {
	if f._value == "" {
		return big.NewInt(0)
	}
	n, ok := new(big.Int).SetString(f._value, 10)
	if !ok {
		panic(fmt.Errorf("%s, Hashtree BigInt Parse failed, value=%#v",
			"big.Int SetString error", f._value))
	}
	return n
}

// SafeGetBig is a safe getter for BigInt
func (f *BigInt) SafeGetBig() *big.Int {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.GetBig()
}

// SetBig is a setter for BigInt
func (f *BigInt) SetBig(n *big.Int) {
	var strValue string
	if n != nil {
		strValue = n.String()
	}
	if strValue == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = strValue
	f._root._mod[f._key] = f._value
}

// SafeSetBig is a safe setter for BigInt
func (f *BigInt) SafeSetBig(n *big.Int) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.SetBig(n)
}

// BigRat is a wrapper for big.Rat
type BigRat struct {
	_root  *Root
	_key   string
	_value string
}

// Get is a getter for BigFloat
func (f *BigRat) Get() float64 {
	if f._value == "" {
		return 0
	}
	n, ok := new(big.Rat).SetString(f._value)
	if !ok {
		panic(fmt.Errorf("%s, Hashtree BigRat Parse failed, value=%#v",
			"big.Rat SetString error", f._value))
	}
	v, _ := n.Float64()
	return v
}

// SafeGet is a safe getter for BigRat
func (f *BigRat) SafeGet() float64 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for BigFloat
func (f *BigRat) Set(v float64) {
	rat, _ := big.NewFloat(v).Rat(nil)
	strValue := rat.String()
	if strValue == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = strValue
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for BigRat
func (f *BigRat) SafeSet(v float64) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(v)
}

// GetBig is a getter for BigRat
func (f *BigRat) GetBig() *big.Rat {
	if f._value == "" {
		return big.NewRat(0, 0)
	}
	n, ok := new(big.Rat).SetString(f._value)
	if !ok {
		panic(fmt.Errorf("%s, Hashtree BigRat Parse failed, value=%#v",
			"big.Rat SetString error", f._value))
	}
	return n
}

// SafeGetBig is a safe getter for BigRat
func (f *BigRat) SafeGetBig() *big.Rat {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.GetBig()
}

// SetBig is a setter for BigRat
func (f *BigRat) SetBig(n *big.Rat) {
	var strValue string
	if n != nil {
		strValue = n.String()
	}
	if strValue == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = strValue
	f._root._mod[f._key] = f._value
}

// SafeSetBig is a safe setter for BigRat
func (f *BigRat) SafeSetBig(n *big.Rat) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.SetBig(n)
}

// BigFloat is a wrapper for big.Float
type BigFloat struct {
	_root  *Root
	_key   string
	_value string
}

// Get is a getter for BigFloat
func (f *BigFloat) Get() float64 {
	if f._value == "" {
		return 0
	}
	n, ok := new(big.Float).SetString(f._value)
	if !ok {
		panic(fmt.Errorf("%s, Hashtree BigFloat Parse failed, value=%#v",
			"big.Float SetString error", f._value))
	}
	f64, _ := n.Float64()
	return f64
}

// SafeGet is a safe getter for BigFloat
func (f *BigFloat) SafeGet() float64 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for BigFloat
func (f *BigFloat) Set(value float64) {
	strValue := big.NewFloat(value).String()
	if strValue == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = strValue
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for BigFloat
func (f *BigFloat) SafeSet(value float64) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// GetBig is a getter for BigFloat
func (f *BigFloat) GetBig() *big.Float {
	if f._value == "" {
		return big.NewFloat(0)
	}
	n, ok := new(big.Float).SetString(f._value)
	if !ok {
		panic(fmt.Errorf("%s, Hashtree BigFloat Parse failed, value=%#v",
			"big.Float SetString error", f._value))
	}
	return n
}

// SafeGetBig is a safe getter for BigFloat
func (f *BigFloat) SafeGetBig() *big.Float {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.GetBig()
}

// SetBig is a setter for BigFloat
func (f *BigFloat) SetBig(n *big.Float) {
	var strValue string
	if n != nil {
		strValue = n.String()
	}
	if strValue == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = strValue
	f._root._mod[f._key] = f._value
}

// SafeSetBig is a safe setter for BigFloat
func (f *BigFloat) SafeSetBig(n *big.Float) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.SetBig(n)
}

// Bool is a wrapper for bool.
type Bool struct {
	_root  *Root
	_key   string
	_value bool
}

// Get is a getter for Bool
func (f *Bool) Get() bool {
	return f._value
}

// SafeGet is a safe getter for Bool
func (f *Bool) SafeGet() bool {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for Bool
func (f *Bool) Set(value bool) {
	if value == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for Bool
func (f *Bool) SafeSet(value bool) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// String is a wrapper for bool.
type String struct {
	_root  *Root
	_key   string
	_value string
}

// Get is a getter for String
func (f *String) Get() string {
	return f._value
}

// SafeGet is a safe getter for String
func (f *String) SafeGet() string {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for String
func (f *String) Set(value string) {
	if value == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for String
func (f *String) SafeSet(value string) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// Time is a wrapper for Unix time
type Time struct {
	_root  *Root
	_key   string
	_value string
}

// Get is a getter for Time
func (f *Time) Get() int64 {
	if f._value == "" {
		return 0
	}
	return timeStringToStamp(f._value)
}

// SafeGet is a safe getter for Time
func (f *Time) SafeGet() int64 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for Time
func (f *Time) Set(value int64) {
	var strValue string
	if value != 0 {
		strValue = timeStampToString(value)
	}
	if strValue == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = strValue
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for Time
func (f *Time) SafeSet(value int64) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// JSON is a wrapper for json
type JSON struct {
	_root  *Root
	_key   string
	_value string
}

// Get is a getter for String
func (f *JSON) Get(n interface{}) {
	value := Decompress(f._value)

	if len(value) == 0 {
		return
	}
	err := json.Unmarshal([]byte(value), n)
	if err != nil {
		panic(fmt.Errorf("%s, Hashtree JSON Unmarshal failed, value=%#v",
			err.Error(), value))
	}
}

// SafeGet is a safe getter for String
func (f *JSON) SafeGet(n interface{}) {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	f.Get(n)
}

// GetString is a getter for String
func (f *JSON) GetString() string {
	return Decompress(f._value)
}

// SafeGetString is a safe getter for String
func (f *JSON) SafeGetString() string {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.GetString()
}

// GetBytes is a getter for String
func (f *JSON) GetBytes() []byte {
	return []byte(Decompress(f._value))
}

// SafeGetBytes is a safe getter for String
func (f *JSON) SafeGetBytes() []byte {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.GetBytes()
}

// Set is a setter for String
func (f *JSON) Set(value interface{}) {
	b, err := json.Marshal(value)
	if err != nil {
		panic(fmt.Errorf("%s, Hashtree JSON Marshal failed, value=%#v",
			err.Error(), value))
	}
	strValue := string(b)
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = Compress(strValue)
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for String
func (f *JSON) SafeSet(value interface{}) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// SetString is a setter for String
func (f *JSON) SetString(value string) {
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = Compress(value)
	f._root._mod[f._key] = f._value
}

// SafeSetString is a safe setter for String
func (f *JSON) SafeSetString(value string) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.SetString(value)
}

// SetBytes is a setter for String
func (f *JSON) SetBytes(value []byte) {
	f.SetString(string(value))
}

// SafeSetBytes is a safe setter for String
func (f *JSON) SafeSetBytes(value []byte) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.SetBytes(value)
}

// Proto is a wrapper for json
type Proto struct {
	_root  *Root
	_key   string
	_value string
}

// Get is a getter for String
func (f *Proto) Get(n interface{}) {
	err := proto.Unmarshal([]byte(f._value), n)
	if err != nil {
		panic(fmt.Errorf("%s, Hashtree Proto Unmarshal failed", err.Error()))
	}
}

// SafeGet is a safe getter for String
func (f *Proto) SafeGet(n interface{}) {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	f.Get(n)
}

// GetString is a getter for String
func (f *Proto) GetString() string {
	return f._value
}

// SafeGetString is a safe getter for String
func (f *Proto) SafeGetString() string {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.GetString()
}

// GetBytes is a getter for String
func (f *Proto) GetBytes() []byte {
	return []byte(f._value)
}

// SafeGetBytes is a safe getter for String
func (f *Proto) SafeGetBytes() []byte {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.GetBytes()
}

// Set is a setter for String
func (f *Proto) Set(value interface{}) {
	b, err := proto.Marshal(value)
	if err != nil {
		panic(fmt.Errorf("%s, Hashtree Proto Marshal failed, value=%#v",
			err.Error(), value))
	}
	strValue := string(b)
	if strValue == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = strValue
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for String
func (f *Proto) SafeSet(value interface{}) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// SetString is a setter for String
func (f *Proto) SetString(value string) {
	if value == f._value {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = value
	f._root._mod[f._key] = f._value
}

// SafeSetString is a safe setter for String
func (f *Proto) SafeSetString(value string) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.SetString(value)
}

// SetBytes is a setter for String
func (f *Proto) SetBytes(value []byte) {
	f.SetString(string(value))
}

// SafeSetBytes is a safe setter for String
func (f *Proto) SafeSetBytes(value []byte) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.SetBytes(value)
}

// SliceInt is a wrapper for []int.
type SliceInt struct {
	_root  *Root
	_key   string
	_value []int
}

// Get is a getter for SliceInt
func (f *SliceInt) Get() []int {
	value := make([]int, len(f._value))
	copy(value, f._value)
	return value
}

// SafeGet is a safe getter for SliceInt
func (f *SliceInt) SafeGet() []int {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for SliceInt
func (f *SliceInt) Set(value []int) {
	if reflect.DeepEqual(value, f._value) {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]int, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for SliceInt
func (f *SliceInt) SafeSet(value []int) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// SliceInt8 is a wrapper for []int8.
type SliceInt8 struct {
	_root  *Root
	_key   string
	_value []int8
}

// Get is a getter for SliceInt8
func (f *SliceInt8) Get() []int8 {
	value := make([]int8, len(f._value))
	copy(value, f._value)
	return value
}

// SafeGet is a safe getter for SliceInt8
func (f *SliceInt8) SafeGet() []int8 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for SliceInt8
func (f *SliceInt8) Set(value []int8) {
	if reflect.DeepEqual(value, f._value) {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]int8, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for SliceInt8
func (f *SliceInt8) SafeSet(value []int8) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// SliceInt16 is a wrapper for []int16.
type SliceInt16 struct {
	_root  *Root
	_key   string
	_value []int16
}

// Get is a getter for SliceInt16
func (f *SliceInt16) Get() []int16 {
	value := make([]int16, len(f._value))
	copy(value, f._value)
	return value
}

// SafeGet is a safe getter for SliceInt16
func (f *SliceInt16) SafeGet() []int16 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for SliceInt16
func (f *SliceInt16) Set(value []int16) {
	if reflect.DeepEqual(value, f._value) {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]int16, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for SliceInt16
func (f *SliceInt16) SafeSet(value []int16) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// SliceInt32 is a wrapper for []int32.
type SliceInt32 struct {
	_root  *Root
	_key   string
	_value []int32
}

// Get is a getter for SliceInt32
func (f *SliceInt32) Get() []int32 {
	value := make([]int32, len(f._value))
	copy(value, f._value)
	return value
}

// SafeGet is a safe getter for SliceInt32
func (f *SliceInt32) SafeGet() []int32 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for SliceInt32
func (f *SliceInt32) Set(value []int32) {
	if reflect.DeepEqual(value, f._value) {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]int32, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for SliceInt32
func (f *SliceInt32) SafeSet(value []int32) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// SliceInt64 is a wrapper for []int64.
type SliceInt64 struct {
	_root  *Root
	_key   string
	_value []int64
}

// Get is a getter for SliceInt64
func (f *SliceInt64) Get() []int64 {
	value := make([]int64, len(f._value))
	copy(value, f._value)
	return value
}

// SafeGet is a safe getter for SliceInt64
func (f *SliceInt64) SafeGet() []int64 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for SliceInt64
func (f *SliceInt64) Set(value []int64) {
	if reflect.DeepEqual(value, f._value) {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]int64, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for SliceInt64
func (f *SliceInt64) SafeSet(value []int64) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// SliceUint is a wrapper for []uint.
type SliceUint struct {
	_root  *Root
	_key   string
	_value []uint
}

// Get is a getter for SliceUint
func (f *SliceUint) Get() []uint {
	value := make([]uint, len(f._value))
	copy(value, f._value)
	return value
}

// SafeGet is a safe getter for SliceUint
func (f *SliceUint) SafeGet() []uint {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for SliceUint
func (f *SliceUint) Set(value []uint) {
	if reflect.DeepEqual(value, f._value) {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]uint, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for SliceUint
func (f *SliceUint) SafeSet(value []uint) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// SliceUint is a wrapper for []uint8.
type SliceUint8 struct {
	_root  *Root
	_key   string
	_value []uint8
}

// Get is a getter for SliceUint8
func (f *SliceUint8) Get() []uint8 {
	value := make([]uint8, len(f._value))
	copy(value, f._value)
	return value
}

// SafeGet is a safe getter for SliceUint8
func (f *SliceUint8) SafeGet() []uint8 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for SliceUint8
func (f *SliceUint8) Set(value []uint8) {
	if reflect.DeepEqual(value, f._value) {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]uint8, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for SliceUint8
func (f *SliceUint8) SafeSet(value []uint8) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// SliceUint16 is a wrapper for []uint16.
type SliceUint16 struct {
	_root  *Root
	_key   string
	_value []uint16
}

// Get is a getter for SliceUint16
func (f *SliceUint16) Get() []uint16 {
	value := make([]uint16, len(f._value))
	copy(value, f._value)
	return value
}

// SafeGet is a safe getter for SliceUint16
func (f *SliceUint16) SafeGet() []uint16 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for SliceUint16
func (f *SliceUint16) Set(value []uint16) {
	if reflect.DeepEqual(value, f._value) {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]uint16, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for SliceUint16
func (f *SliceUint16) SafeSet(value []uint16) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// SliceUint32 is a wrapper for []uint32.
type SliceUint32 struct {
	_root  *Root
	_key   string
	_value []uint32
}

// Get is a getter for SliceUint32
func (f *SliceUint32) Get() []uint32 {
	value := make([]uint32, len(f._value))
	copy(value, f._value)
	return value
}

// SafeGet is a safe getter for SliceUint32
func (f *SliceUint32) SafeGet() []uint32 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for SliceUint32
func (f *SliceUint32) Set(value []uint32) {
	if reflect.DeepEqual(value, f._value) {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]uint32, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for SliceUint32
func (f *SliceUint32) SafeSet(value []uint32) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// SliceUint64 is a wrapper for []int64.
type SliceUint64 struct {
	_root  *Root
	_key   string
	_value []uint64
}

// Get is a getter for SliceUint64
func (f *SliceUint64) Get() []uint64 {
	value := make([]uint64, len(f._value))
	copy(value, f._value)
	return value
}

// SafeGet is a safe getter for SliceUint64
func (f *SliceUint64) SafeGet() []uint64 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for SliceUint64
func (f *SliceUint64) Set(value []uint64) {
	if reflect.DeepEqual(value, f._value) {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]uint64, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for SliceUint64
func (f *SliceUint64) SafeSet(value []uint64) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// SliceFloat32 is a wrapper for []float32.
type SliceFloat32 struct {
	_root  *Root
	_key   string
	_value []float32
}

// Get is a getter for SliceFloat32
func (f *SliceFloat32) Get() []float32 {
	value := make([]float32, len(f._value))
	copy(value, f._value)
	return value
}

// SafeGet is a safe getter for SliceFloat32
func (f *SliceFloat32) SafeGet() []float32 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for SliceFloat32
func (f *SliceFloat32) Set(value []float32) {
	if reflect.DeepEqual(value, f._value) {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]float32, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for SliceFloat32
func (f *SliceFloat32) SafeSet(value []float32) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// SliceFloat64 is a wrapper for []float64.
type SliceFloat64 struct {
	_root  *Root
	_key   string
	_value []float64
}

// Get is a getter for SliceFloat64
func (f *SliceFloat64) Get() []float64 {
	value := make([]float64, len(f._value))
	copy(value, f._value)
	return value
}

// SafeGet is a safe getter for SliceFloat64
func (f *SliceFloat64) SafeGet() []float64 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for SliceFloat64
func (f *SliceFloat64) Set(value []float64) {
	if reflect.DeepEqual(value, f._value) {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]float64, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for SliceFloat64
func (f *SliceFloat64) SafeSet(value []float64) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// SliceBigInt is a wrapper for []big.Int.
type SliceBigInt struct {
	_root  *Root
	_key   string
	_value []*big.Int
}

// Get is a getter for SliceBigInt
func (f *SliceBigInt) Get() []*big.Int {
	value := make([]*big.Int, len(f._value))
	for i, v := range f._value {
		n := new(big.Int)
		value[i] = n.Add(v, n)
	}
	return value
}

// SafeGet is a safe getter for SliceBigInt
func (f *SliceBigInt) SafeGet() []*big.Int {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for SliceBigInt
func (f *SliceBigInt) Set(value []*big.Int) {
	if reflect.DeepEqual(value, f._value) {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]*big.Int, len(value))
	for i, v := range value {
		n := new(big.Int)
		f._value[i] = n.Add(v, n)
	}
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for SliceBigInt
func (f *SliceBigInt) SafeSet(value []*big.Int) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// SliceBigRat is a wrapper for []big.Rat.
type SliceBigRat struct {
	_root  *Root
	_key   string
	_value []*big.Rat
}

// Get is a getter for SliceBigRat
func (f *SliceBigRat) Get() []*big.Rat {
	value := make([]*big.Rat, len(f._value))
	for i, v := range f._value {
		n := new(big.Rat)
		value[i] = n.Add(v, n)
	}
	return value
}

// SafeGet is a safe getter for SliceBigRat
func (f *SliceBigRat) SafeGet() []*big.Rat {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for SliceBigRat
func (f *SliceBigRat) Set(value []*big.Rat) {
	if reflect.DeepEqual(value, f._value) {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]*big.Rat, len(value))
	for i, v := range value {
		n := new(big.Rat)
		f._value[i] = n.Add(v, n)
	}
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for SliceBigRat
func (f *SliceBigRat) SafeSet(value []*big.Rat) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// SliceBigFloat is a wrapper for []big.Int.
type SliceBigFloat struct {
	_root  *Root
	_key   string
	_value []*big.Float
}

// Get is a getter for SliceBigFloat
func (f *SliceBigFloat) Get() []*big.Float {
	value := make([]*big.Float, len(f._value))
	for i, v := range f._value {
		n := new(big.Float)
		value[i] = n.Add(v, n)
	}
	return value
}

// SafeGet is a safe getter for SliceBigFloat
func (f *SliceBigFloat) SafeGet() []*big.Float {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for SliceBigFloat
func (f *SliceBigFloat) Set(value []*big.Float) {
	if reflect.DeepEqual(value, f._value) {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]*big.Float, len(value))
	for i, v := range value {
		n := new(big.Float)
		f._value[i] = n.Add(v, n)
	}
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for SliceBigFloat
func (f *SliceBigFloat) SafeSet(value []*big.Float) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// SliceTime is a wrapper for []Unix time
type SliceTime struct {
	_root  *Root
	_key   string
	_value []string
}

// Get is a getter for SliceTime
func (f *SliceTime) Get() []int64 {
	var ns []int64
	for _, s := range f._value {
		if s == "" {
			ns = append(ns, 0)
		} else {
			t := timeStringToStamp(s)
			ns = append(ns, t)
		}
	}
	return ns
}

// SafeGet is a safe getter for SliceTime
func (f *SliceTime) SafeGet() []int64 {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for SliceTime
func (f *SliceTime) Set(value []int64) {
	var ssValue []string
	for _, n := range value {
		if n == 0 {
			ssValue = append(ssValue, "")
		} else {
			ssValue = append(ssValue, timeStampToString(n))
		}
	}
	if reflect.DeepEqual(ssValue, f._value) {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = ssValue
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for SliceTime
func (f *SliceTime) SafeSet(ns []int64) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(ns)
}

// SliceBool is a wrapper for []bool
type SliceBool struct {
	_root  *Root
	_key   string
	_value []bool
}

// Get is a getter for SliceBool
func (f *SliceBool) Get() []bool {
	value := make([]bool, len(f._value))
	copy(value, f._value)
	return value
}

// SafeGet is a safe getter for SliceBool
func (f *SliceBool) SafeGet() []bool {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for SliceBool
func (f *SliceBool) Set(value []bool) {
	if reflect.DeepEqual(value, f._value) {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]bool, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for SliceBool
func (f *SliceBool) SafeSet(value []bool) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}

// SliceString is a wrapper for []string
type SliceString struct {
	_root  *Root
	_key   string
	_value []string
}

// Get is a getter for SliceString
func (f *SliceString) Get() []string {
	value := make([]string, len(f._value))
	copy(value, f._value)
	return value
}

// SafeGet is a safe getter for SliceString
func (f *SliceString) SafeGet() []string {
	f._root.rw.RLock()
	defer f._root.rw.RUnlock()

	return f.Get()
}

// Set is a setter for SliceString
func (f *SliceString) Set(value []string) {
	if reflect.DeepEqual(value, f._value) {
		return
	}
	_, ok := f._root._bak[f._key]
	if !ok {
		f._root._bak[f._key] = f._value
	}
	f._value = make([]string, len(value))
	copy(f._value, value)
	f._root._mod[f._key] = f._value
}

// SafeSet is a safe setter for SliceString
func (f *SliceString) SafeSet(value []string) {
	f._root.rw.Lock()
	defer f._root.rw.Unlock()

	f.Set(value)
}
