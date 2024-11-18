// Filename: monasync.go
// Description: constructors for monadic asynchronous π-calculus

package monasync

import (
	. "piencodings/chanprinter"
	"piencodings/stats"
)

// 0
func Nil() {
	Print("  0")
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
	stats.GlobalStats.LogChannel()
	Print("  (ν %s)", a)

	f(a)
}

// u<v>
func Send[T any](u chan T, v T) {
	Print("  %s<%s>", u, v)
	u <- v
	stats.GlobalStats.LogSend() // instrumentation
}

// u(x).P
func Recv[T any](u chan T, f func(T)) {
	x := <-u
	stats.GlobalStats.LogRecv() // instrumentation
	Print("  %s(%s)", u, x)
	f(x)
}
