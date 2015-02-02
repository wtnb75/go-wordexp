package wordexp

import (
	"os"
	"path/filepath"
	"testing"
)

var homedir string

func init() {
	homedir = os.Getenv("HOME")
}

func TestWE1(t *testing.T) {
	res, err := WordExp("~/go", 0)
	if err != nil {
		t.Errorf("WE1 error1", res, err)
	}
	if len(res) != 1 {
		t.Errorf("WE1 error2", res, err)
	}
	if res[0] != filepath.Join(homedir, "go") {
		t.Errorf("WE1 error3", res, err)
	}
}

func TestFM1(t *testing.T) {
	err := FnMatch("abc[de]fg.go", "abcdfg.go", 0)
	if err != nil {
		t.Errorf("FM1 error1")
	}
	err = FnMatch("abc[de]fg.go", "abcffg.go", 0)
	if err == nil {
		t.Errorf("FM1 error2")
	}
	if err.Error() != "String does not match" {
		t.Errorf("FM1 error3", err)
	}
}
