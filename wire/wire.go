package wire

import (
	"bufio"
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

type Wire struct {
	engine *gin.Engine
}

func New(engine *gin.Engine) *Wire {
	return &Wire{
		engine: engine,
	}
}

func (s *Wire) Init() {
}

func (s *Wire) Invoke(_, wireReq string) (wireRsp string) {
	request, err := http.ReadRequest(bufio.NewReader(strings.NewReader(wireReq)))
	if err != nil {
		panic(err)
	}

	var rspRecorder = httptest.NewRecorder()

	s.engine.ServeHTTP(rspRecorder, request)

	response := rspRecorder.Result()

	var buf = bytes.Buffer{}
	if err = response.Write(&buf); err != nil {
		panic(err)
	}

	wireRsp = buf.String()

	return
}

func (s *Wire) Close() {}
