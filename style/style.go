package style

import (
	"strings"

	"github.com/acoderup/boost/magic"
)

var abbrs = []string{
	"ACL", "API", "ASCII",
	"CPU", "CSS",
	"DNS",
	"EOF",
	"GUID",
	"HTML", "HTTP", "HTTPS",
	"ID",
	"VIP",
	"IP",
	"JSON",
	"LHS",
	"QPS",
	"RAM", "RHS", "RPC",
	"SLA", "SMTP", "SQL", "SSH",
	"TCP", "TLS", "TTL",
	"UDP", "UI", "UID", "UUID", "URI", "URL", "UTF8",
	"VM",
	"XML", "XMPP", "XSRF", "XSS",
}

var abbrMap = make(map[string]string)

var (
	googleChain = *NewChainStyle(magic.SeparatorSlash, magic.SeparatorHyphen)
	unixChain   = *NewChainStyle(magic.SeparatorPeriod, magic.SeparatorUnderscore)
)

func init() {
	for _, abbr := range abbrs {
		abbrMap[camelize(abbr)] = abbr
	}
}

func camelize(s string) string {
	s = strings.ToLower(s)
	b := []byte(s)
	if b[0] >= 'a' && b[0] <= 'z' {
		b[0] -= 32
	}
	return string(b)
}

func Standardize(s string, sep magic.SeparatorType) string {
	if s == "" || sep == magic.SeparatorLazy {
		return s
	}

	var words []string
	if sep == magic.SeparatorNone {
		words = []string{s}
	} else {
		words = strings.Split(s, sep)
	}

	var b = []byte{}
	for _, word := range words {
		word = camelize(word)
		abbr, ok := abbrMap[word]
		if ok {
			word = abbr
		}
		b = append(b, []byte(word)...)
	}
	return string(b)
}

type ChainStyle struct {
	ChainSeperator magic.SeparatorType
	WordSeparator  magic.SeparatorType
}

func NewChainStyle(chainSeparator, wordSeparator string) *ChainStyle {
	return &ChainStyle{
		ChainSeperator: chainSeparator,
		WordSeparator:  wordSeparator,
	}
}

func Chain(s string, cs ChainStyle) []string {
	return cs.Chain(s)
}

func (cs ChainStyle) Chain(s string) []string {
	chain := strings.Split(s, cs.ChainSeperator)
	for index := 0; index < len(chain); index++ {
		chain[index] = Standardize(chain[index], cs.WordSeparator)
	}
	return chain
}

func GoogleChain(s string) []string {
	return Chain(s, googleChain)
}

func UnixChain(s string) []string {
	return Chain(s, unixChain)
}
