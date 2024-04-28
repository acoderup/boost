package regexp

import (
	"sync"

	"github.com/dlclark/regexp2"
)

type Regexp struct {
	sync.Map
}

func New() *Regexp {
	return &Regexp{}
}

func (r *Regexp) MatchString(pattern string, str string) (bool, error) {
	v, ok := r.Load(pattern)
	if !ok {
		v = regexp2.MustCompile(pattern, regexp2.RE2)
		r.Store(pattern, v)
	}

	re := v.(*regexp2.Regexp)
	return re.MatchString(str)
}

func (r *Regexp) ReplaceAllStringFunc(pattern string, str string, repl func(string) string) (string, error) {
	v, ok := r.Load(pattern)
	if !ok {
		v = regexp2.MustCompile(pattern, regexp2.RE2)
		r.Store(pattern, v)
	}

	re := v.(*regexp2.Regexp)
	return re.ReplaceFunc(str, func(m regexp2.Match) string {
		return repl(m.String())
	}, -1, -1)
}
