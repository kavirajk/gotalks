package tricky

import (
	"os"
	"os/exec"
	"testing"
)

func TestMustEnv(t *testing.T) {
	// Inspired from the great talk by Andrew Gerrand (Core Go team)
	var value string
	key := "DB_DRIVER"

	// Unset env variable before testing it if any, and rollback once test is done
	value = os.Getenv(key)
	if value != "" {
		os.Unsetenv(key)
		defer func() { os.Setenv(key, value) }()
	}
	if os.Getenv("TEST_MUST_ENV") == "1" {
		mustEnv(key)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestMustEnv")
	cmd.Env = append(os.Environ(), "TEST_MUST_ENV=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("mustEnv passed even without setting env. expected: fail")
}
