// tests/user_controller_test.go

package tests

import (
	"GoBoard/controllers"
	"GoBoard/tools"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"

	"testing"
)

func getLoginPOSTPayload() string {
	params := url.Values{}
	params.Add("username", "user1")
	params.Add("password", "pass1")

	return params.Encode()
}

func getRegistrationPOSTPayload() string {
	params := url.Values{}
	params.Add("username", "u1")
	params.Add("password", "p1")

	return params.Encode()
}

func TestShowRegistrationPageUnauthenticated(t *testing.T) {
	r := tools.GetRouter(true)

	r.GET("/u/register", controllers.ShowRegistrationPage)

	req, _ := http.NewRequest("GET", "/u/register", nil)

	tools.CheckHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Register</title>") > 0

		return statusOK && pageOK
	})
}

func TestRegisterUnauthenticated(t *testing.T) {
	tools.SaveLists()
	w := httptest.NewRecorder()

	r := tools.GetRouter(true)

	r.POST("/u/register", controllers.Register)

	registrationPayload := getRegistrationPOSTPayload()
	req, _ := http.NewRequest("POST", "/u/register", strings.NewReader(registrationPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(registrationPayload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "<title>Successful registration &amp; Login</title>") < 0 {
		t.Fail()
	}
	tools.RestoreLists()
}

func TestRegisterUnauthenticatedUnavailableUsername(t *testing.T) {
	tools.SaveLists()
	w := httptest.NewRecorder()

	r := tools.GetRouter(true)

	r.POST("/u/register", controllers.Register)

	registrationPayload := getLoginPOSTPayload()
	req, _ := http.NewRequest("POST", "/u/register", strings.NewReader(registrationPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(registrationPayload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
	tools.RestoreLists()
}
