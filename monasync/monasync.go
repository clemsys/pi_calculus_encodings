// Filename: monasync.go
// Description: constructors for monadic asynchronous π-calculus

package monasync

import (
	. "piencodings/chanprinter"
)

// 0
func Nil() {
	Print("  0\n")
}

// P|Q
// Caution: the order of f and g matter
// Guarantee: Par returns when g returns
func Par(f func(), g func()) {
	go f()
	g()
}

// (ν a)P
func Gen[T any](s string, f func(chan T)) {
	a := make(chan T) // (ν a)
	SetC(a, GenNameS(s))
	Print("  (ν %s)\n", a)

	f(a)
}

// u<v>
func Send[T any](u chan T, v T) {
	Print("  %s<%s>\n", u, v)
	u <- v
}

// u(x).P
func Recv[T any](u chan T, f func(T)) {
	x := <-u
	Print("  %s(%s)\n", u, x)
	f(x)
}
