package stringx

import (
	"fmt"
	"math"
	"strings"

	"github.com/acoderup/boost/cast"
	"github.com/dlclark/regexp2"
)

func Merge(ss ...string) string {
	var b strings.Builder
	b.Grow(64)
	for _, s := range ss {
		b.WriteString(s)
	}
	return b.String()
}

func PickLast(s string, sep string) string {
	lastIndex := strings.LastIndex(s, sep)
	if lastIndex < 0 {
		return s
	}
	return s[lastIndex+len(sep):]
}

func PruneLast(s string, sep string) string {
	lastIndex := strings.LastIndex(s, sep)
	if lastIndex < 0 {
		return s
	}
	return s[:lastIndex]
}

func PickFirst(s string, sep string) string {
	firstIndex := strings.Index(s, sep)
	if firstIndex < 0 {
		return s
	}
	return s[:firstIndex]
}

func PruneFirst(s string, sep string) string {
	firstIndex := strings.Index(s, sep)
	if firstIndex < 0 {
		return s
	}
	return s[firstIndex+len(sep):]
}

func ContainsAny(s string, v ...any) bool {
	size := len(v)
	if size == 0 {
		return false
	}

	if size > 1 {
		for _, item := range v {
			if strings.Contains(s, item.(string)) {
				return true
			}
		}
		return false
	}

	switch val := v[0].(type) {
	case string:
		return strings.ContainsAny(s, val)
	case []string:
		for _, item := range val {
			if strings.Contains(s, item) {
				return true
			}
		}
		return false
	default:
		return false
	}
}

func Mod(s string, n int) int {
	var sum int
	for _, b := range []byte(s) {
		sum += int(b)
	}
	return sum % n
}

func CompareVersion(alpha string, beta string) int {
	if alpha == "" {
		alpha = cast.ToString(math.MinInt64)
	}
	if beta == "" {
		beta = cast.ToString(math.MinInt64)
	}
	alphaStrs := strings.Split(alpha, ".")
	betaStrs := strings.Split(beta, ".")

	var size int
	if len(alphaStrs) > len(betaStrs) {
		size = len(alphaStrs)
	} else {
		size = len(betaStrs)
	}

	alphaInts := make([]int, size)
	betaInts := make([]int, size)
	for i := 0; i < size; i++ {
		if i < len(alphaStrs) {
			alphaInts[i] = cast.ToInt(alphaStrs[i])
		} else {
			alphaInts[i] = math.MinInt64
		}
		if i < len(betaStrs) {
			betaInts[i] = cast.ToInt(betaStrs[i])
		} else {
			betaInts[i] = math.MinInt64
		}
	}

	for i := 0; i < size; i++ {
		if alphaInts[i] > betaInts[i] {
			return 1
		} else if alphaInts[i] < betaInts[i] {
			return -1
		}
	}

	return 0
}

func Shorten(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max]
}

var (
	reIPv4        = regexp2.MustCompile(`^(?P<ip>(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?).(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?).(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?).(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?))$`, regexp2.RE2)
	reIPv4Port    = regexp2.MustCompile(`^(?P<ip>(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?).(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?).(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?).(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)):(?P<port>([0-9]{1,5}))$`, regexp2.RE2)
	reIPv6        = regexp2.MustCompile(`^(?P<ip>([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]).){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]).){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$`, regexp2.RE2)
	reIPv6Bracket = regexp2.MustCompile(`^\[(?P<ip>([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]).){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]).){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))\]$`, regexp2.RE2)
	reIPv6Port    = regexp2.MustCompile(`^\[(?P<ip>([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]).){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]).){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))\]:(?P<port>[0-9]{1,5})$`, regexp2.RE2)
)

type IPType int

const (
	IPNil = iota
	IPv4
	IPv6
)

var ipTypeStringMap = map[IPType]string{
	IPNil: "nil",
	IPv4:  "ipv4",
	IPv6:  "ipv6",
}

func (t IPType) String() string {
	return ipTypeStringMap[t]
}

func ParseRemoteAddr(addr string) (string, uint, IPType) {
	match, _ := reIPv4Port.FindStringMatch(addr)
	if match != nil {
		return match.GroupByName("ip").String(), cast.ToUint(match.GroupByName("port").String()), IPv4
	}

	match, _ = reIPv6Port.FindStringMatch(addr)
	if match != nil {
		return match.GroupByName("ip").String(), cast.ToUint(match.GroupByName("port").String()), IPv6
	}

	match, _ = reIPv4.FindStringMatch(addr)
	if match != nil {
		return match.GroupByName("ip").String(), 0, IPv4
	}

	match, _ = reIPv6.FindStringMatch(addr)
	if match != nil {
		return match.GroupByName("ip").String(), 0, IPv6
	}

	match, _ = reIPv6Bracket.FindStringMatch(addr)
	if match != nil {
		return match.GroupByName("ip").String(), 0, IPv6
	}

	return "", 0, IPNil
}

func FormatRemoteAddr(addr string) string {
	ip, port, typ := ParseRemoteAddr(addr)
	switch typ {
	case IPv4:
		return fmt.Sprintf("%s:%d", ip, port)
	case IPv6:
		return fmt.Sprintf("[%s]:%d", ip, port)
	}

	return addr
}
