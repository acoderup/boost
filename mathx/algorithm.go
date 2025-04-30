package mathx

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
)

func Max[T Number](x, y T) T {
	return T(math.Max(float64(x), float64(y)))
}

func Min[T Number](x, y T) T {
	return T(math.Min(float64(x), float64(y)))
}

func Div[T1 Number, T2 Number](a T1, b T2) float64 {
	r, err := DivE(a, b)
	if err != nil {
		panic(err)
	}
	return r
}

func SafeDiv[T1 Number, T2 Number](a T1, b T2) float64 {
	r, err := DivE(a, b)
	if err != nil {
		return 0
	}
	return r
}

func DivE[T1 Number, T2 Number](a T1, b T2) (float64, error) {
	if b == 0 {
		return 0, errors.New("denominator is 0")
	}
	return float64(a) / float64(b), nil
}

func Precision[T1 Float, T2 Integer](target T1, prec T2) float64 {
	fmtStr := "%." + strconv.FormatInt(int64(prec), 10) + "f"
	result, err := strconv.ParseFloat(fmt.Sprintf(fmtStr, target), 64)
	if err != nil {
		panic(err)
	}
	return result
}

// FloatEqual guarantees n of effective figure
func FloatEqual(f1, f2 float64, n int) bool {
	min := math.Pow10(-1 * n)
	for math.Abs(f1) > 1 {
		f1 /= 10.0
		f2 /= 10.0
	}
	if f1 > f2 {
		return math.Dim(f1, f2) < min
	} else {
		return math.Dim(f2, f1) < min
	}
}

// FloatEqual2 guarantees diff of two number smaller equal than the defined small number
func FloatEqual2(f1, f2 float64, min float64) bool {
	if f1 > f2 {
		return math.Dim(f1, f2) < min
	} else {
		return math.Dim(f2, f1) < min
	}
}

func FastFind[T1 Number](n T1, s []T1) int {
	var (
		min    = 0
		max    = len(s) - 1
		middle int
	)

	for {
		middle = (min + max) / 2
		if max == middle { // finish it: case min == max
			return middle
		} else if min == middle { // finish it: case min + 1 == max
			if n <= s[max] && s[max-1] < n {
				return max
			} else {
				return min
			}
		} else {
			if n <= s[middle-1] {
				max = middle - 1
			} else if s[middle] < n {
				min = middle + 1
			} else {
				return middle
			}
		}
	}
}

func In[T comparable](v T, s []T) bool {
	for _, n := range s {
		if v == n {
			return true
		}
	}

	return false
}

func Replace[T comparable](s []T, o T, n T) []T {
	for i := 0; i < len(s); i++ {
		if s[i] == o {
			s[i] = n
		}
	}
	return s
}

func ReplaceArray[T comparable](s [][]T, o T, n T) [][]T {
	for i, ts := range s {
		for i2, t := range ts {
			if t == o {
				s[i][i2] = n
			}
		}
	}
	return s
}

func Index[T comparable](s []T, v T) int {
	for i, n := range s {
		if n == v {
			return i
		}
	}
	return -1
}

func Sum[T Number](s []T) T {
	var sum T
	for _, n := range s {
		sum += n
	}
	return sum
}

func Positions[T comparable](v T, s []T) []int {
	pos := make([]int, 0)
	for p, n := range s {
		if n == v {
			pos = append(pos, p)
		}
	}
	return pos
}

func Count[T comparable](v T, s []T) int {
	return len(Positions(v, s))
}

func ContinuousPositions[T comparable](v T, s []T) [][]int {
	rt := make([][]int, 0)
	for index := 0; index < len(s); {
		if s[index] == v {
			rtt := make([]int, 0)
			rtt = append(rtt, index)
			tmp := index + 1
			for tmp < len(s) && s[tmp] == v {
				rtt = append(rtt, tmp)
				tmp++
			}
			index = tmp
			rt = append(rt, rtt)
		} else {
			index++
		}
	}
	return rt
}

// MaxContinuousCount returns max continuous count in slice  [start end) index
func MaxContinuousCount[T comparable](v T, s []T) (int, int, int) {
	maxCount := 0
	start := 0
	end := 0
	for index := 0; index < len(s); {
		if s[index] == v {
			tmp := index + 1
			for tmp < len(s) && s[tmp] == v {
				tmp++
			}
			if tmp-index > maxCount {
				maxCount = tmp - index
				start = index
				end = tmp
			}
			index = tmp
		} else {
			index++
		}
	}
	return maxCount, start, end
}

// UniqueSorted sorts and removes duplicates from a slice of comparable elements.
// It preserves the original slice and returns a new sorted, unique slice.
// Supported types: int, int64, float64, string, and other comparable types.
func UniqueSorted[T comparable](input []T) []T {
	// Handle empty slice
	if len(input) == 0 {
		return nil
	}

	// Handle single-element slice
	if len(input) == 1 {
		return []T{input[0]}
	}

	// Create a working copy only if needed (will be modified during sorting)
	work := make([]T, len(input))
	copy(work, input)

	// Sort using type-specific comparison
	sort.Slice(work, func(i, j int) bool {
		switch v := any(work[i]).(type) {
		case int:
			return v < any(work[j]).(int)
		case int64:
			return v < any(work[j]).(int64)
		case float64:
			return v < any(work[j]).(float64)
		case string:
			return v < any(work[j]).(string)
		default:
			// Fallback for other comparable types (no ordering guaranteed)
			return false
		}
	})

	// Pre-count unique elements to allocate result slice with exact capacity
	uniqueCount := 1
	for i := 1; i < len(work); i++ {
		if work[i-1] != work[i] {
			uniqueCount++
		}
	}

	// Build result with exact size
	result := make([]T, 0, uniqueCount)
	result = append(result, work[0])
	for i := 1; i < len(work); i++ {
		if work[i-1] != work[i] {
			result = append(result, work[i])
		}
	}

	return result
}

// CoinToFloat 将整数表示的货币金额(放大10000倍存储)转换为浮点数表示的常规金额
// 例如：将存储的 998866800 (表示 99886.68) 转换为 99886.68
// 参数:
//   storedCoin - 以整数形式存储的金额(实际值 ×10000)
// 返回值:
//   浮点数表示的实际金额，保留2位小数
func CoinToFloat(storedCoin int64) float64 {
	// 1. 先除以10000还原实际值
	// 2. 乘以100保留2位小数精度
	// 3. 使用math.Round四舍五入
	// 4. 再除以100得到最终结果
	return math.Round(float64(storedCoin)/10000*100) / 100
}

// FloatToCoin 将浮点数表示的金额转换为整数存储形式(放大10000倍)
// 例如：将 99886.68 转换为 998866800 存储
// 参数:
//   amount - 浮点数表示的金额
// 返回值:
//   以整数形式存储的金额(实际值 ×10000)
func FloatToCoin(amount float64) int64 {
	// 1. 先乘以10000转换为整数表示
	// 2. 加0.5实现四舍五入而非截断
	return int64(amount*10000 + 0.5)
}
