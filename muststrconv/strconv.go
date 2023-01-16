// Package muststrconv is a must-style wrapper over the standard library strconv package.
// All functions panic in response to any errors encountered.
package muststrconv

import (
	"strconv"

	"github.com/mdwhatcott/must/must"
)

func Atoi(s string) int {
	return must.Value(strconv.Atoi(s))
}
func ParseBool(str string) bool {
	return must.Value(strconv.ParseBool(str))
}
func ParseComplex(s string, bitSize int) complex128 {
	return must.Value(strconv.ParseComplex(s, bitSize))
}
func ParseFloat(s string, bitSize int) float64 {
	return must.Value(strconv.ParseFloat(s, bitSize))
}
func ParseInt(s string, base int, bitSize int) int64 {
	return must.Value(strconv.ParseInt(s, base, bitSize))
}
func ParseUint(s string, base int, bitSize int) uint64 {
	return must.Value(strconv.ParseUint(s, base, bitSize))
}
func QuotedPrefix(s string) string {
	return must.Value(strconv.QuotedPrefix(s))
}
func Unquote(s string) string {
	return must.Value(strconv.Unquote(s))
}
func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error) {
	return strconv.UnquoteChar(s, quote)
}
