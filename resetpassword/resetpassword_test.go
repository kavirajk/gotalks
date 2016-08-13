package main

import (
	"testing"

	"code.launchyard.com/root/myserver/utils"

	"code.launchyard.com/root/myserver/models"
)

func TestResetPassword(t *testing.T) {
	save_password_called := false
	var user *models.User
	new_password := "jedi"

	userGetter = func(email string) *models.User { // HL
		user = &models.User{Password: "old_password", Email: email}
		return user
	}
	savePassword = func(s Saver) { // HL
		save_password_called = true
	}
	ResetPassword("test@example.com", new_password)
	expected := utils.HashString("jedi")
	if user == nil {
		t.Fatalf("userGetter not called")
	}
	if !save_password_called {
		t.Errorf("savePassword not called")
	}
	if user.Password != expected {
		t.Errorf("password not saved")
	}
}
