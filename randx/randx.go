package randx

import (
	"math/rand"
	"sort"
	"time"

	"github.com/acoderup/boost/mathx"
)

type Signed interface {
	int | int8 | int16 | int32 | int64
}

type Unsigned interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	float32 | float64
}

type Number interface {
	Integer | Float
}

type Randx struct {
	*rand.Rand
}

func New(v any) *Randx {
	var r *rand.Rand
	switch v := v.(type) {
	case int64:
		r = rand.New(rand.NewSource(v))
	case rand.Source:
		r = rand.New(v)
	default:
		return nil
	}
	return &Randx{r}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// PR rands yes or no by probability
func PR(pr float64) bool {
	return rand.Float64() <= pr
}

// RandPR rands yes or no by probability
func RandPR(rand *Randx, pr float64) bool {
	return rand.Float64() <= pr
}

// PR rands yes or no by probability
func (rand *Randx) PR(pr float64) bool {
	return rand.Float64() <= pr
}

// Intn picks a random value in the specified ramge [0, n)
func Intn(n int) int {
	return rand.Intn(n)
}

// RandIntn picks a random value in the specified ramge [0, n)
func RandIntn(rand *Randx, n int) int {
	return rand.Intn(n)
}

// Intn picks a random value in the specified ramge [0, n)
func (rand *Randx) Intn(n int) int {
	return rand.Rand.Intn(n)
}

// RangeIntn picks a random value in the specified ramge [s, e]
func RangeIntn(s int, e int) int {
	return rand.Intn(e-s+1) + s
}

// RandRangeIntn picks a random value in the specified ramge [s, e]
func RandRangeIntn(rand *Randx, s int, e int) int {
	return rand.Intn(e-s+1) + s
}

// RangeIntn picks a random value in the specified ramge [s, e]
func (rand *Randx) RangeIntn(s int, e int) int {
	return rand.Intn(e-s+1) + s
}

// Int63n picks a random value in the specified ramge [0, n)
func Int63n(n int64) int64 {
	return rand.Int63n(n)
}

// RandInt63n picks a random value in the specified ramge [0, n)
func RandInt63n(rand *Randx, n int64) int64 {
	return rand.Int63n(n)
}

// Int63n picks a random value in the specified ramge [0, n)
func (rand *Randx) Int63n(n int64) int64 {
	return rand.Rand.Int63n(n)
}

// RangeInt63n picks a random value in the specified ramge [s, e]
func RangeInt63n(s int64, e int64) int64 {
	return rand.Int63n(e-s+1) + s
}

// RandRangeInt63n picks a random value in the specified ramge [s, e]
func RandRangeInt63n(rand *Randx, s int64, e int64) int64 {
	return rand.Int63n(e-s+1) + s
}

// RangeInt63n picks a random value in the specified ramge [s, e]
func (rand *Randx) RangeInt63n(s int64, e int64) int64 {
	return rand.Int63n(e-s+1) + s
}

// Float64 picks a random value in the specified ramge [0, 1)
func Float64() float64 {
	return rand.Float64()
}

// RandFloat64 picks a random value in the specified ramge [0, 1)
func RandFloat64(rand *Randx) float64 {
	return rand.Float64()
}

// Float64 picks a random value in the specified ramge [0, 1)
func (rand *Randx) Float64() float64 {
	return rand.Rand.Float64()
}

// RangeFloat64 picks a random value in the specified range [s, e)
func RangeFloat64(s float64, e float64) float64 {
	return s + rand.Float64()*(e-s)
}

// RandRangeFloat64 picks a random value in the specified range [s, e)
func RandRangeFloat64(rand *Randx, s float64, e float64) float64 {
	r := s + rand.Float64()*(e-s)
	return r
}

// RangeFloat64 picks a random value in the specified range [s, e)
func (rand *Randx) RangeFloat64(s float64, e float64) float64 {
	r := s + rand.Float64()*(e-s)
	return r
}

// Weight picks a random value in the specified slice by weight
func Weight[T Number](s []T) int {
	var weightSum float64
	for _, r := range s {
		weightSum += float64(r)
	}
	n := rand.Float64() * weightSum
	var lastKey int
	for i, r := range s {
		nReach := float64(r)
		if n <= nReach {
			return i
		}
		n -= nReach
		lastKey = i
	}
	return lastKey
}

// RandWeight picks a random value in the specified slice by weight
func RandWeight[T Number](rand *Randx, s []T) int {
	var weightSum float64
	for _, r := range s {
		weightSum += float64(r)
	}
	n := rand.Float64() * weightSum
	var lastKey int
	for i, r := range s {
		nReach := float64(r)
		if n <= nReach {
			return i
		}
		n -= nReach
		lastKey = i
	}
	return lastKey
}

// WeightMap picks a random value in the specified map by weight
func WeightMap[T1 Number, T2 Number](m map[T1]T2) T1 {
	var s = make([][]interface{}, 0, len(m))
	for k, v := range m {
		s = append(s, []interface{}{k, v})
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i][0].(T1) < s[j][0].(T1)
	})

	var keys = make([]T1, 0, len(s))
	var values = make([]T2, 0, len(s))
	for _, v := range s {
		keys = append(keys, v[0].(T1))
		values = append(values, v[1].(T2))
	}

	i := Weight(values)
	return keys[i]
}

// WeightMap picks a random value in the specified map by weight
func RandWeightMap[T1 Number, T2 Number](rand *Randx, m map[T1]T2) T1 {
	var s = make([][]interface{}, 0, len(m))
	for k, v := range m {
		s = append(s, []interface{}{k, v})
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i][0].(T1) < s[j][0].(T1)
	})

	var keys = make([]T1, 0, len(s))
	var values = make([]T2, 0, len(s))
	for _, v := range s {
		keys = append(keys, v[0].(T1))
		values = append(values, v[1].(T2))
	}

	i := RandWeight(rand, values)
	return keys[i]
}

// Unrepeated picks unrepeated random values in the specified object by weight
func Unrepeated[T Number](s []T, count int) []int {
	result := make([]int, 0)
	for len(result) < count {
		var index int
		for index = Weight(s); mathx.In(index, result); {
			index = Weight(s)
		}
		result = append(result, index)
	}
	return result
}

// RandUnrepeated picks unrepeated random values in the specified object by weight
func RandUnrepeated[T Number](rand *Randx, s []T, count int) []int {
	result := make([]int, 0)
	for len(result) < count {
		var index int
		for index = RandWeight(rand, s); mathx.In(index, result); {
			index = RandWeight(rand, s)
		}
		result = append(result, index)
	}
	return result
}

// Shuffle shuffles the specified slice
func Shuffle[T comparable](s []T) []T {
	length := int64(len(s))
	for i := length; i > 0; i-- {
		pos := rand.Int63n(i)
		s[i-1], s[pos] = s[pos], s[i-1]
	}
	return s
}

// RandShuffle shuffles the specified slice
func RandShuffle[T comparable](rand *Randx, s []T) []T {
	length := int64(len(s))
	for i := length; i > 0; i-- {
		pos := rand.Int63n(i)
		s[i-1], s[pos] = s[pos], s[i-1]
	}
	return s
}

// RandShuffleWeights randomly shuffles weights and returns n indices
func RandShuffleWeights[T Number](rand *Randx, n int, weights []T) []int {
	// Calculate the number of non-zero weights
	var noZero int
	var zeroValue T
	for _, weight := range weights {
		if weight != zeroValue {
			noZero++
		}
	}

	// If the number of elements to select is less than the number of non-zero weights, panic
	if noZero < n {
		panic("noZero < n in RandShuffleWeights")
	}

	// Create a copy of the weights to avoid modifying the original array
	weightsCopy := make([]T, len(weights))
	copy(weightsCopy, weights)

	// Array to store the randomly selected indices
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		// Randomly select an index based on the weights
		index := RandWeight(rand, weightsCopy)
		indices[i] = index
		// Set the weight of the selected index to zero to prevent it from being selected again
		weightsCopy[index] = zeroValue
	}

	return indices
}
