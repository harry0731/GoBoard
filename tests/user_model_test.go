// tests/user_model_test.go

package tests

import (
	"GoBoard/models"
	"GoBoard/tools"
	"testing"
)

func TestUsernameAvailability(t *testing.T) {
	tools.SaveLists()

	if !models.IsUsernameAvailable("newuser") {
		t.Fail()
	}

	if models.IsUsernameAvailable("user1") {
		t.Fail()
	}

	models.RegisterNewUser("newuser", "newpass")

	if models.IsUsernameAvailable("newuser") {
		t.Fail()
	}

	tools.RestoreLists()
}

func TestValidUserRegistration(t *testing.T) {
	tools.SaveLists()

	u, err := models.RegisterNewUser("newuser", "newpass")

	if err != nil || u.Username == "" {
		t.Fail()
	}

	tools.RestoreLists()
}

func TestInvalidUserRegistration(t *testing.T) {
	tools.SaveLists()

	u, err := models.RegisterNewUser("user1", "pass1")

	if err == nil || u != nil {
		t.Fail()
	}

	u, err = models.RegisterNewUser("newuser", "")

	if err == nil || u != nil {
		t.Fail()
	}

	tools.RestoreLists()
}

func TestUserValidity(t *testing.T) {
	if !models.IsUserValid("user1", "pass1") {
		t.Fail()
	}

	if models.IsUserValid("user2", "pass1") {
		t.Fail()
	}

	if models.IsUserValid("user1", "") {
		t.Fail()
	}

	if models.IsUserValid("", "pass1") {
		t.Fail()
	}

	if models.IsUserValid("User1", "pass1") {
		t.Fail()
	}
}
