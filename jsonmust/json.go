// Package jsonmust is a must-style wrapper over the standard library encoding/json package.
// All functions panic in response to any errors encountered.
package jsonmust

import (
	"bytes"
	"encoding/json"

	"github.com/mdwhatcott/must/must"
)

func Compact(dst *bytes.Buffer, src []byte) {
	must.Nada(json.Compact(dst, src))
}
func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) {
	must.Nada(json.Indent(dst, src, prefix, indent))
}
func Marshal(v any) []byte {
	return must.Value(json.Marshal(v))
}
func MarshalIndent(v any, prefix, indent string) []byte {
	return must.Value(json.MarshalIndent(v, prefix, indent))
}
func Unmarshal(data []byte, v any) {
	must.Nada(json.Unmarshal(data, v))
}
