package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEcho(t *testing.T) {
	arg := "hello"
	b, err := exec.Command("go", "run", "echo.go", arg).CombinedOutput()

	if err != nil {
		t.Error(err)
	}

	out := strings.TrimSpace(string(b))

	if !strings.EqualFold(out, arg) {
		t.Errorf("Invalid command output: %s", out)
	}

}
