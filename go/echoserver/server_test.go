package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestRead(t *testing.T) {
	var (
		e   *echo.Echo                 = echo.New()
		req *http.Request              = httptest.NewRequest(http.MethodGet, "/", nil)
		rec *httptest.ResponseRecorder = httptest.NewRecorder()
		c   echo.Context               = e.NewContext(req, rec)
	)

	err := getAllUsers(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != http.StatusOK {
		t.Error("invalid status code")
	}

	respBytes, err := ioutil.ReadAll(rec.Result().Body)
	if err != nil {
		t.Error(err.Error())
	}
	var respData []User
	err = json.Unmarshal(respBytes, &respData)
	if err != nil {
		t.Error(err.Error())
	}
}
