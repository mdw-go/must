// Package must is inspired by, but much simpler than, github.com/xiegeo/must
package must

func Value[T any](out T, err error) T {
	if err != nil {
		panic(err)
	}
	return out
}
func Values[T1 any, T2 any](out1 T1, out2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return out1, out2
}
func Nada(err error) {
	if err != nil {
		panic(err)
	}
}
func Defer(f func() error) func() {
	return func() { Nada(f()) }
}
