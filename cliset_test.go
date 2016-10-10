package cliset

import (
	// "fmt"
	"os"
	"testing"
)

// func TestWarn(t *testing.T) {
// 	Warn("Error we cannot detect the form")
// }

func TestIsPiped(t *testing.T) {
	t.Log(IsPiped())
}

func TestPipeValue(t *testing.T) {
	s, err := PipeValue()
	if err != nil {
		t.Error(err)
	}
	t.Log(s)
}

func TestGetMsg(t *testing.T) {
	s := GetMsg()
	t.Log(s)
}

func TestSubcmd(t *testing.T) {
	cmd := []string{"A", "B", "c", "d", "e", "f", "g", "h", "i", "j", "k", "test"}
	os.Args = append(os.Args, cmd...)
	b, n := Subcmd("test")
	t.Log(b, n)
}
