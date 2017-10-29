package main

import (
	"github.com/davecgh/go-spew/spew"
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
	assert.NotEqual(t, 2, Add(1, 1))
}

func TestProductList(t *testing.T) {
	request, _ := http.NewRequest("GET", "/productList", nil)
	response := httptest.NewRecorder()

	//	same
	//	engine := SetupGin()
	//	engine.ServeHTTP(response, request)
	SetupGin().ServeHTTP(response, request)

	p := response.Body.String()
	t.Log(spew.Sdump(p))

	assert.Equal(t, 200, response.Code, "OK response is expected")
	assert.JSONEq(t, `{"products":["ApplePie", "BananaPie"]}`, p)

}
