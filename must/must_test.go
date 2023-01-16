package must_test

import (
	"errors"
	"testing"

	"github.com/mdwhatcott/must/must"
)

func TestNada(t *testing.T)   { defer catch(t); must.Nada(nada()) }
func TestValue(t *testing.T)  { defer catch(t); _ = must.Value(value()) }
func TestValues(t *testing.T) { defer catch(t); _, _ = must.Values(values()) }
func catch(t *testing.T) {
	if r := recover(); r.(error).Error() != err.Error() {
		t.Error("should have panicked on err")
	}
}

var err = errors.New("err")

func nada() error               { return err }
func value() (int, error)       { return 0, err }
func values() (int, int, error) { return 0, 0, err }
