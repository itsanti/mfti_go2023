package test

import (
	"os/exec"
	"strings"
	"testing"
)

func TestTrue(t *testing.T) {
	b, err := exec.Command("go", "run", "../cmd/main.go", "../data/true.txt").CombinedOutput()

	if err != nil {
		t.Error(err)
	}

	out := strings.TrimSpace(string(b))
	dups := len(strings.Split(out, "\n"))

	if dups < 15 {
		t.Errorf("not enough duplicates found (%v) or data file is corrupt; want 15", dups)
	}
}

func TestFalse(t *testing.T) {
	b, err := exec.Command("go", "run", "../cmd/main.go", "../data/false.txt").CombinedOutput()

	if err != nil {
		t.Error(err)
	}

	o := string(b)

	if len(o) > 0 {
		dups := len(strings.Split(o, "\n"))
		t.Errorf("to much duplicates found (%v) or data file is corrupt; want 0", dups)
	}
}
