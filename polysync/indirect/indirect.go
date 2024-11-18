// Filename: indirect.go
// Description: indirect simulation of polyadic synchronous π-calculus
//              using monadic synchronous π-calculus

package indirect

import (
	"piencodings/monsync"
	. "piencodings/polysync/polyadicity"
)

type stats struct {
	sent     int
	received int
}

// 0
func Nil() {
	monsync.Nil()
}

// P|Q
// Caution: the order of f and g matter
// Guarantee: Par returns when g returns
func Par(f func(), g func()) {
	monsync.Par(f, g)
}

// helper function for Gen
// f takes a list of N - i channels
func genStep[T any](i int, s string, f func([]chan T)) {
	if i == N {
		f([]chan T{}) // empty list of channels
	} else {
		monsync.Gen( // (ν ai)
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
	monsync.Gen(s, f)
}

// helper function for Send
func sendStep[T any](i int, c chan chan chan T, v []T, f func()) {
	if i == N {
		f()
	} else {
		monsync.Send(c, v[i], func() { sendStep(i+1, c, v, f) }) // c<vi>...
	}
}

// encoding [[u<v1, ..., vn>.P]] = (v c)u<c>.c<v1>...c<vn>.[[P]]
func Send[T any](u chan chan chan chan chan chan T, v []T, f func()) {
	monsync.Gen("c", func(c chan chan chan T) { // (ν c)
		monsync.Send(u, c, func() { // u<c>
			sendStep(0, c, v, f) // c<v1>...c<vn>.[[P]]
		})
	})
}

// helper function for Recv
// the function f takes a list of N-i values
func recvStep[T any](i int, y chan chan chan T, f func([]T)) {
	if i == N {
		f([]T{}) // empty list of values
	} else {
		monsync.Recv(y, func(xi T) { // y(xi)
			recvStep(i+1, y, func(xs []T) {
				// progressively build up the list of values passed to f
				f(append([]T{xi}, xs...))
			})
		})
	}
}

// encoding: [[u(x1, ..., xn).P]] = u(y).y(x1)...y(xn).[[P]]
func Recv[T any](u chan chan chan chan chan chan T, f func([]T)) {
	monsync.Recv(u, func(y chan chan chan T) { // u(y).
		recvStep(0, y, f) // y(x1)...y(xn).[[P]]
	})
}
