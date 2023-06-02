package just_test

import (
	"errors"
	"testing"

	"github.com/mdwhatcott/must/just"
)

func TestNoErrIgnore(t *testing.T) { defer assertErrorIgnored(t); just.Ignore(nada()) }
func TestNoErrValue(t *testing.T)  { defer assertErrorIgnored(t); _ = just.Value(value()) }
func TestNoErrValues(t *testing.T) { defer assertErrorIgnored(t); _, _ = just.Values(values()) }
func assertErrorIgnored(t *testing.T) {
	if r := recover(); r != nil {
		t.Error("should NOT have panicked on err")
	}
}

var err = errors.New("err")

func nada() error               { return err }
func value() (int, error)       { return 1, err }
func values() (int, int, error) { return 2, 3, err }

func TestValue(t *testing.T) {
	v := just.Value(value())
	if v != 1 {
		t.Error("wrong value:", v)
	}
}
func TestValue2(t *testing.T) {
	v, w := just.Values(values())
	if v != 2 {
		t.Error("wrong value:", v)
	}
	if w != 3 {
		t.Error("wrong value:", w)
	}
}
