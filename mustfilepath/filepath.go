// Package mustfilepath is a must-style wrapper over the standard library filepath package.
// All functions panic in response to any errors encountered.
package mustfilepath

import (
	"io/fs"
	"path/filepath"

	"github.com/mdwhatcott/must/must"
)

func Abs(path string) string                 { return must.Value(filepath.Abs(path)) }
func EvalSymlinks(path string) string        { return must.Value(filepath.EvalSymlinks(path)) }
func Glob(pattern string) []string           { return must.Value(filepath.Glob(pattern)) }
func Match(pattern, name string) bool        { return must.Value(filepath.Match(pattern, name)) }
func Rel(basepath, targpath string) string   { return must.Value(filepath.Rel(basepath, targpath)) }
func Walk(root string, fn filepath.WalkFunc) { must.Nada(filepath.Walk(root, fn)) }
func WalkDir(root string, fn fs.WalkDirFunc) { must.Nada(filepath.WalkDir(root, fn)) }
