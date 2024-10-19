package fmtmust

import (
	f "fmt"
	"io"

	"github.com/mdw-go/must/must"
)

func Fprint(w io.Writer, a ...any) (n int)              { return must.Value(f.Fprint(w, a...)) }
func Fprintf(w io.Writer, fmt string, a ...any) (n int) { return must.Value(f.Fprintf(w, fmt, a...)) }
func Fprintln(w io.Writer, a ...any) (n int)            { return must.Value(f.Fprintln(w, a...)) }

func Fscan(r io.Reader, a ...any) (n int)              { return must.Value(f.Fscan(r, a...)) }
func Fscanf(r io.Reader, fmt string, a ...any) (n int) { return must.Value(f.Fscanf(r, fmt, a...)) }
func Fscanln(r io.Reader, a ...any) (n int)            { return must.Value(f.Fscanln(r, a...)) }

func Print(a ...any) (n int)              { return must.Value(f.Print(a...)) }
func Printf(fmt string, a ...any) (n int) { return must.Value(f.Printf(fmt, a...)) }
func Println(a ...any) (n int)            { return must.Value(f.Println(a...)) }

func Scan(a ...any) (n int)              { return must.Value(f.Scan(a...)) }
func Scanf(fmt string, a ...any) (n int) { return must.Value(f.Scanf(fmt, a...)) }
func Scanln(a ...any) (n int)            { return must.Value(f.Scanln(a...)) }

func Sscan(str string, a ...any) (n int)              { return must.Value(f.Sscan(str, a...)) }
func Sscanf(str string, fmt string, a ...any) (n int) { return must.Value(f.Sscanf(str, fmt, a...)) }
func Sscanln(str string, a ...any) (n int)            { return must.Value(f.Sscanln(str, a...)) }
