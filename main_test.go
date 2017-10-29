package main

import (
	"github.com/stretchr/testify/assert"
	//	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddOK(t *testing.T) {
	assert.Equal(t, 3, Add(1, 2))
}

func TestAddFail(t *testing.T) {
	assert.Equal(t, 2, Add(1, 1))
}

func TestProductList(t *testing.T) {
	request, _ := http.NewRequest("GET", "/productList", nil)
	response := httptest.NewRecorder()

	engine := SetupGin()
	engine.ServeHTTP(response, request)

	p := response.Body.String()

	t.Log(p)

	assert.Equal(t, 200, response.Code, "OK response is expected")
}
