package redirect_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/acoderup/boost/redirect"
)

func TestRedirect(t *testing.T) {
	var buf bytes.Buffer
	r := redirect.NewRedirector(redirect.WithWriter(&buf), redirect.WithDuplicate())
	if err := r.Stdout(func() {
		fmt.Print("hello")
	}); err != nil {
		t.Fatal(err)
	}

	if buf.String() != "hello" {
		t.Fatal("buf.String() != \"hello\"")
	}
}
