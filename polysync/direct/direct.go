// Filename: direct.go
// Description: direct simulation of polyadic synchronous π-calculus
//              using monadic asynchronous π-calculus

package direct

import (
	"piencodings/monasync"
	. "piencodings/polysync/polyadicity"
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

// helper function for Gen
// f takes a list of N - i channels
func genStep[T any](i int, s string, f func([]chan T)) {
	if i == N {
		f([]chan T{}) // empty list of channels
	} else {
		monasync.Gen( // (ν ai)
			s,
			func(c chan T) {
				genStep(i+1, s, func(cs []chan T) { // next step
					// progressively build up the list of channels passed to f
					f(append([]chan T{c}, cs...))
				})
			},
		)
	}
}

// encoding [[(ν a)P]] = (ν a)[[P]]
func GenVec[T any](s string, f func([]chan T)) {
	genStep(0, s, f)
}

func GenChan[T any](s string, f func(chan T)) {
	monasync.Gen(s, f)
}

// helper function for Send
// f takes a list of N - i channels
func sendStep[T any](i int, c chan chan T, v []T, f func()) {
	if i == N {
		f()
	} else {
		monasync.Recv(c, func(yi chan T) { // c(yi)
			monasync.Par( // |
				func() { monasync.Send(yi, v[i]) }, // yi<vi>
				func() { sendStep(i+1, c, v, f) },  // next step
			)
		})
	}
}

// encoding [[u<v1, ..., vn>.P]] = (v c)(u<c>|c(y1).(y1<v1>|c(y2).(y2<v2>|...|c(yn).(yn<vn>|[[P]])...)))
func Send[T any](u chan chan chan T, v []T, f func()) {
	monasync.Gen("c", func(c chan chan T) { // (ν c)
		monasync.Par( // |
			func() { monasync.Send(u, c) },  // u<c>
			func() { sendStep(0, c, v, f) }, // c(y1).(y1<v1>|c(y2).(y2<v2>|...|c(yn).(yn<vn>|[[P]])...))
		)
	})
}

// helper function for Recv
// f takes a list of N - i objects
func recvStep[T any](i int, y chan chan T, d chan T, f func([]T)) {
	if i == N {
		f([]T{}) // empty list of objects
	} else {
		monasync.Par( // |
			func() { monasync.Send(y, d) }, // y<d>
			func() {
				monasync.Recv(d, func(xi T) { // d(xi)
					recvStep(i+1, y, d, func(xs []T) { // next step
						// progressively build up the list of objects passed to f
						f(append([]T{xi}, xs...))
					})
				})
			},
		)
	}
}

// encoding: [[u(x1, ..., xn).P]] = u(y).(v d)(y<d>|d(x1).(y<d>|d(x2).(...(y<d>|d(xn).[[P]])...)))
func Recv[T any](u chan chan chan T, f func([]T)) {
	monasync.Recv(u, func(y chan chan T) { // u(y)
		monasync.Gen("d", func(d chan T) { // (ν d)
			recvStep(0, y, d, f) // y<d>|d(x1).(y<d>|d(x2).(...(y<d>|d(xn).[[P]])...))
		})
	})
}
