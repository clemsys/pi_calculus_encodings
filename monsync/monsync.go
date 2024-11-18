// Filename: monsync.go
// Description: simulation of monadic synchronous π-calculus
//              using monadic asynchronous π-calculus

package monsync

import (
	"piencodings/monasync"
)

// 0
func Nil() {
	monasync.Nil()
}

// P|Q
// Caution: the order of f and g matter
// Guarantee: Par returns when g returns
func Par(f func(), g func()) {
	monasync.Par(f, g)
}

// (ν a)P
func Gen[T any](s string, f func(chan T)) {
	monasync.Gen(s, f)
}

// encoding: [[u<v>.P]] = (ν c)(u<c>|c(y).(y<v>|[[P]]))
// where y not in fv(P) and c not in fn(P)
func Send[T any](u chan chan chan T, v T, f func()) {
	monasync.Gen("c", func(c chan chan T) { // (ν c)
		monasync.Par( // |
			func() { monasync.Send(u, c) }, // u<c>
			func() {
				monasync.Recv(c, func(y chan T) { // c(y).
					monasync.Par( // |
						func() { monasync.Send(y, v) }, // y<v>
						f,                              // [[P]]
					)
				})
			},
		)
	})
}

// encoding: [[u(x).P]] = u(y).(ν d)(y<d>|d(x).[[P]])
// where y not in fv(P) and d not in fn(P)
func Recv[T any](u chan chan chan T, f func(T)) {
	monasync.Recv(u, func(y chan chan T) { // u(y).
		monasync.Gen("d", func(d chan T) { // (ν d)
			monasync.Par( // |
				func() { monasync.Send(y, d) }, // y<d>
				func() { monasync.Recv(d, f) }, // d(x).[[P]]
			)
		})
	})
}
