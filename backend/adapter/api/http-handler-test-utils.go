package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func ServeTestRequest(router *gin.Engine, method, uri string, data []byte) *httptest.ResponseRecorder {
	respRecorder := httptest.NewRecorder()
	req, _ := http.NewRequest(method, uri, bytes.NewReader(data))
	router.ServeHTTP(respRecorder, req)

	return respRecorder
}
