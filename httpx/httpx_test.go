package httpx_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/acoderup/boost/httpx"
)

func TestRequestRequestFunc(t *testing.T) {
	body, e := httpx.Request(func(c *http.Client) (*http.Request, error) {
		return http.NewRequest("POST", "https://www.baidu.com", bytes.NewBuffer([]byte("hello")))
	})
	if e != nil {
		t.Fatal(e)
	}

	t.Log(string(body))
}

func TestRequestResponseFunc(t *testing.T) {
	body, e := httpx.Request(func(c *http.Client) (*http.Response, error) {
		return c.Post("https://www.baidu.com", "application/json", bytes.NewBuffer([]byte("hello")))
	})
	if e != nil {
		t.Fatal(e)
	}

	t.Log(string(body))
}

func TestRequestWithDumpRequestFunc(t *testing.T) {
	body, dumpData, e := httpx.RequestWithDump(func(c *http.Client) (*http.Request, error) {
		return http.NewRequest("POST", "https://www.baidu.com", bytes.NewBuffer([]byte("hello")))
	})
	if e != nil {
		t.Fatal(e)
	}

	t.Log(string(body))

	s, e := json.MarshalIndent(dumpData, "", "  ")
	if e != nil {
		t.Fatal(e)
	}

	t.Log(string(s))
}
