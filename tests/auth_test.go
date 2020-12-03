// tests/auth_test.go

package tests

import (
	auth "GoBoard/middlewares"
	"GoBoard/tools"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestEnsureLoggedInUnauthenticated(t *testing.T) {
	r := tools.GetRouter(false)
	r.GET("/", setLoggedIn(false), auth.EnsureLoggedIn(), func(c *gin.Context) {
		t.Fail()
	})

	tools.TestMiddlewareRequest(t, r, http.StatusUnauthorized)
}

func setLoggedIn(b bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("is_logged_in", b)
	}
}

func TestEnsureLoggedInAuthenticated(t *testing.T) {
	r := tools.GetRouter(false)
	r.GET("/", setLoggedIn(true), auth.EnsureLoggedIn(), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	tools.TestMiddlewareRequest(t, r, http.StatusOK)
}

func TestEnsureNotLoggedInAuthenticated(t *testing.T) {
	r := tools.GetRouter(false)
	r.GET("/", setLoggedIn(true), auth.EnsureNotLoggedIn(), func(c *gin.Context) {
		t.Fail()
	})

	tools.TestMiddlewareRequest(t, r, http.StatusUnauthorized)
}

func TestEnsureNotLoggedInUnauthenticated(t *testing.T) {
	r := tools.GetRouter(false)
	r.GET("/", setLoggedIn(false), auth.EnsureNotLoggedIn(), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	tools.TestMiddlewareRequest(t, r, http.StatusOK)
}

func TestSetUserStatusAuthenticated(t *testing.T) {
	r := tools.GetRouter(false)
	r.GET("/", auth.SetUserStatus(), func(c *gin.Context) {
		loggedInInterface, exists := c.Get("is_logged_in")
		if !exists || !loggedInInterface.(bool) {
			t.Fail()
		}
	})

	w := httptest.NewRecorder()

	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header = http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}

	r.ServeHTTP(w, req)
}

func TestSetUserStatusUnauthenticated(t *testing.T) {
	r := tools.GetRouter(false)
	r.GET("/", auth.SetUserStatus(), func(c *gin.Context) {
		loggedInInterface, exists := c.Get("is_logged_in")
		if exists && loggedInInterface.(bool) {
			t.Fail()
		}
	})

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/", nil)

	r.ServeHTTP(w, req)
}
