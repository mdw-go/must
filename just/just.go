package just

func Value[T any](out T, _ error) T                             { return out }
func Values[T1 any, T2 any](out1 T1, out2 T2, _ error) (T1, T2) { return out1, out2 }
func Defer(f func() error) func()                               { return func() { Ignore(f()) } }
func Ignore(error)                                              {}
