package fsmust

import (
	"io/fs"

	"github.com/mdw-go/must/must"
)

func Glob(f fs.FS, pattern string) (matches []string) { return must.Value(fs.Glob(f, pattern)) }
func ReadFile(f fs.FS, name string) []byte            { return must.Value(fs.ReadFile(f, name)) }
func WalkDir(f fs.FS, root string, fn fs.WalkDirFunc) { must.Void(fs.WalkDir(f, root, fn)) }
func ReadDir(f fs.FS, name string) []fs.DirEntry      { return must.Value(fs.ReadDir(f, name)) }
func Sub(f fs.FS, dir string) fs.FS                   { return must.Value(fs.Sub(f, dir)) }
func Stat(f fs.FS, name string) fs.FileInfo           { return must.Value(fs.Stat(f, name)) }
