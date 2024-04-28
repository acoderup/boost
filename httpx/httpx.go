package httpx

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/tomas-qstarrs/boost/cast"
)

type ResponseFunc = func(*http.Client) (*http.Response, error)
type RequestFunc = func(*http.Client) (*http.Request, error)

type ClientConfig struct {
	Retry   int64
	Timeout time.Duration

	// Transport
	Proxy                 string
	DialerTimeout         time.Duration
	DialerKeepAlive       time.Duration
	ForceAttemptHTTP2     bool
	MaxIdleConns          int
	MaxIdleConnsPerHost   int
	MaxConnsPerHost       int
	IdleConnTimeout       time.Duration
	TLSHandshakeTimeout   time.Duration
	ResponseHeaderTimeout time.Duration
	ExpectContinueTimeout time.Duration
}

type Client struct {
	Config ClientConfig
	*http.Client
}

type DumpData struct {
	// Request
	RequestHead      string
	RequestHeadError string
	RequestBody      string
	RequestBodyError string

	// Do
	DoError string

	// Response
	ResponseHead      string
	ResponseHeadError string
	ResponseBody      string
	ResponseBodyError string
}

var DefaultClientConfig = ClientConfig{
	Retry:   1,
	Timeout: 60 * time.Second,

	// Transport
	Proxy:                 "",
	DialerTimeout:         30 * time.Second,
	DialerKeepAlive:       30 * time.Second,
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          0,
	MaxIdleConnsPerHost:   2,
	MaxConnsPerHost:       0,
	TLSHandshakeTimeout:   10 * time.Second,
	ResponseHeaderTimeout: 10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
	IdleConnTimeout:       180 * time.Second,
}

func Default() *Client {
	return NewClient(DefaultClientConfig)
}

func NewClient(config ClientConfig) *Client {
	c := &Client{
		Config: config,
		Client: &http.Client{
			Transport: &http.Transport{
				Proxy: nil,
				DialContext: (&net.Dialer{
					Timeout:   config.DialerTimeout,
					KeepAlive: config.DialerKeepAlive,
				}).DialContext,
				ForceAttemptHTTP2:     config.ForceAttemptHTTP2,
				MaxIdleConns:          config.MaxIdleConns,
				MaxIdleConnsPerHost:   config.MaxIdleConnsPerHost,
				MaxConnsPerHost:       config.MaxConnsPerHost,
				TLSHandshakeTimeout:   config.TLSHandshakeTimeout,
				ResponseHeaderTimeout: config.ResponseHeaderTimeout,
				ExpectContinueTimeout: config.ExpectContinueTimeout,
				IdleConnTimeout:       config.IdleConnTimeout,
			},
		},
	}

	if c.Config.Proxy != "" {
		proxyURL, err := url.Parse(c.Config.Proxy)
		if err != nil {
			log.Panicln(err)
		}

		c.Client.Transport.(*http.Transport).Proxy = http.ProxyURL(proxyURL)
	}

	c.Client.Timeout = c.Config.Timeout

	return c
}

func Request(f any) ([]byte, error) {
	return Default().Request(f)
}

func (c *Client) Request(f any) (data []byte, err error) {
	var times int64
	for times < c.Config.Retry {
		times++
		data, err = c.doFunc(f)
		if err == nil {
			return
		}
	}

	return
}

func (c *Client) doFunc(f any) (respBody []byte, err error) {
	var resp *http.Response

	switch f := f.(type) {
	case RequestFunc:
		var req *http.Request
		req, err = f(c.Client)
		if err != nil {
			return
		}

		resp, err = c.Client.Do(req)
		if err != nil {
			return
		}
	case ResponseFunc:
		resp, err = f(c.Client)
		if err != nil {
			return
		}
	default:
		err = fmt.Errorf("unknown type %T", f)
		return
	}

	defer func() {
		_, _ = io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()

	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http server returns status code %d", resp.StatusCode)
		return
	}

	return
}

func RequestWithDump(f RequestFunc) ([]byte, *DumpData, error) {
	return Default().RequestWithDump(f)
}

func (c *Client) RequestWithDump(f RequestFunc) (data []byte, dumpData *DumpData, err error) {
	dumpData = new(DumpData)
	var times int64
	for times < c.Config.Retry {
		times++
		data, err = c.doFuncWithDump(f, dumpData)
		if err == nil {
			return
		}
	}

	return
}

func (c *Client) doFuncWithDump(f RequestFunc, dumpData *DumpData) ([]byte, error) {
	// Get request
	req, err := f(c.Client)
	if err != nil {
		return nil, err
	}

	// Dump request head
	dumpRequest, err := httputil.DumpRequest(req, false)
	dumpData.RequestHead = string(dumpRequest)
	dumpData.RequestHeadError = cast.ToString(err)

	// Dump request body
	reqBodyBytes, err := io.ReadAll(req.Body)
	dumpData.RequestBody = string(reqBodyBytes)
	dumpData.RequestBodyError = cast.ToString(err)

	// Reset request body
	_, _ = io.Copy(io.Discard, req.Body)
	req.Body.Close()
	req.Body = io.NopCloser(bytes.NewBuffer(reqBodyBytes))

	// Do
	resp, err := c.Client.Do(req)
	if err != nil {
		// Dump do
		dumpData.DoError = cast.ToString(err)
		return nil, err
	}

	// Dump response head
	dumpResponse, err := httputil.DumpResponse(resp, false)
	dumpData.ResponseHead = string(dumpResponse)
	dumpData.ResponseHeadError = cast.ToString(err)

	// Dump response body
	respBodyBytes, err := io.ReadAll(resp.Body)
	dumpData.ResponseBody = string(respBodyBytes)
	dumpData.ResponseBodyError = cast.ToString(err)

	// Reset response body
	_, _ = io.Copy(io.Discard, resp.Body)
	resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return respBodyBytes, fmt.Errorf("http server (%s) returns status code %d", resp.Request.URL.Host, resp.StatusCode)
	}

	return respBodyBytes, nil
}
