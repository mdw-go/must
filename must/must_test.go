package must_test

import (
	"errors"
	"testing"

	"github.com/mdw-go/must/must"
)

func TestNada(t *testing.T)   { defer catch(t); must.Void(nada()) }
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

func TestDeferExecuted(t *testing.T) {
	counter := 0
	defer func() {
		if counter == 0 {
			t.Error("should have incremented counter")
		}
	}()
	incrementer := func() error {
		counter++
		return nil
	}
	defer must.Defer(incrementer)()
}
func TestDeferPanics(t *testing.T) {
	defer catch(t)
	fails := func() error { return err }
	defer must.Defer(fails)()
}
