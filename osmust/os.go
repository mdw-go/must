// Package osmust is a must-style wrapper over the standard library os package.
// All functions panic in response to any errors encountered.
package osmust

import (
	"os"
	"time"

	"github.com/mdw-go/must/must"
)

func Chdir(dir string)                                     { must.Void(os.Chdir(dir)) }
func Chmod(name string, mode os.FileMode)                  { must.Void(os.Chmod(name, mode)) }
func Chown(name string, uid, gid int)                      { must.Void(os.Chown(name, uid, gid)) }
func Chtimes(name string, atime, mtime time.Time)          { must.Void(os.Chtimes(name, atime, mtime)) }
func Create(name string) *os.File                          { return must.Value(os.Create(name)) }
func CreateTemp(dir, pattern string) *os.File              { return must.Value(os.CreateTemp(dir, pattern)) }
func Executable() string                                   { return must.Value(os.Executable()) }
func FindProcess(pid int) *os.Process                      { return must.Value(os.FindProcess(pid)) }
func Getgroups() []int                                     { return must.Value(os.Getgroups()) }
func Getwd() (dir string)                                  { return must.Value(os.Getwd()) }
func Hostname() (name string)                              { return must.Value(os.Hostname()) }
func Lchown(name string, uid, gid int)                     { must.Void(os.Lchown(name, uid, gid)) }
func Link(oldname, newname string)                         { must.Void(os.Link(oldname, newname)) }
func Lstat(name string) os.FileInfo                        { return must.Value(os.Lstat(name)) }
func Mkdir(name string, perm os.FileMode)                  { must.Void(os.Mkdir(name, perm)) }
func MkdirAll(path string, perm os.FileMode)               { must.Void(os.MkdirAll(path, perm)) }
func MkdirTemp(dir, pattern string) string                 { return must.Value(os.MkdirTemp(dir, pattern)) }
func NewSyscallError(syscall string, err error)            { must.Void(os.NewSyscallError(syscall, err)) }
func Open(name string) *os.File                            { return must.Value(os.Open(name)) }
func Pipe() (r *os.File, w *os.File)                       { return must.Values(os.Pipe()) }
func ReadDir(name string) []os.DirEntry                    { return must.Value(os.ReadDir(name)) }
func ReadFile(name string) []byte                          { return must.Value(os.ReadFile(name)) }
func Readlink(name string) string                          { return must.Value(os.Readlink(name)) }
func Remove(name string)                                   { must.Void(os.Remove(name)) }
func RemoveAll(path string)                                { must.Void(os.RemoveAll(path)) }
func Rename(oldpath, newpath string)                       { must.Void(os.Rename(oldpath, newpath)) }
func Setenv(key, value string)                             { must.Void(os.Setenv(key, value)) }
func Stat(name string) os.FileInfo                         { return must.Value(os.Stat(name)) }
func Symlink(oldname, newname string)                      { must.Void(os.Symlink(oldname, newname)) }
func Truncate(name string, size int64)                     { must.Void(os.Truncate(name, size)) }
func Unsetenv(key string)                                  { must.Void(os.Unsetenv(key)) }
func UserCacheDir() string                                 { return must.Value(os.UserCacheDir()) }
func UserConfigDir() string                                { return must.Value(os.UserConfigDir()) }
func UserHomeDir() string                                  { return must.Value(os.UserHomeDir()) }
func WriteFile(name string, data []byte, perm os.FileMode) { must.Void(os.WriteFile(name, data, perm)) }
func OpenFile(name string, flag int, perm os.FileMode) *os.File {
	return must.Value(os.OpenFile(name, flag, perm))
}
func StartProcess(name string, argv []string, attr *os.ProcAttr) *os.Process {
	return must.Value(os.StartProcess(name, argv, attr))
}
