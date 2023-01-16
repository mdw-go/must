// Package mustio is a must-style wrapper over the standard library io package.
// All functions panic in response to any errors encountered.
package mustio

import (
	"io"

	"github.com/mdwhatcott/must/must"
)

func Copy(dst io.Writer, src io.Reader) int64 {
	return must.Value(io.Copy(dst, src))
}
func CopyBuffer(dst io.Writer, src io.Reader, buf []byte) int64 {
	return must.Value(io.CopyBuffer(dst, src, buf))
}
func CopyN(dst io.Writer, src io.Reader, n int64) int64 {
	return must.Value(io.CopyN(dst, src, n))
}
func ReadAll(r io.Reader) []byte {
	return must.Value(io.ReadAll(r))
}
func ReadAtLeast(r io.Reader, buf []byte, min int) int {
	return must.Value(io.ReadAtLeast(r, buf, min))
}
func ReadFull(r io.Reader, buf []byte) int {
	return must.Value(io.ReadFull(r, buf))
}
func WriteString(w io.Writer, s string) int {
	return must.Value(io.WriteString(w, s))
}
