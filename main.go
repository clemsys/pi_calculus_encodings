package main

import (
	. "piencodings/chanprinter"
	"piencodings/monasync"
	"piencodings/monsync"
	"piencodings/polysync/direct"
	"piencodings/polysync/indirect"
	"time"
)

func init() {
	ChanTbl.Tbl = make(map[any]any)
}

func main() {

	print("Monadic asynchronous π-process: (ν a)(ν b)(a<b>|a(x).0)\n")
	monasync.Gen(
		"a",
		func(a chan chan any) {
			monasync.Gen(
				"b",
				func(b chan any) {
					monasync.Par(
						func() {
							monasync.Send(a, b)
						},
						func() {
							monasync.Recv(a, func(chan any) { monasync.Nil() })
						},
					)
				},
			)
		},
	)
	time.Sleep(100 * time.Millisecond)

	print("\nSimulation of monadic synchronous π-calculus using monadic asynchronous π-calculus\n")
	print("π-process (ν a)(ν b)(a<b>.0|a(x).0)\n")
	monsync.Gen(
		"a",
		func(a chan chan chan chan any) {
			monsync.Gen(
				"b",
				func(b chan any) {
					monsync.Par(
						func() {
							monsync.Send(a, b, monsync.Nil)
						},
						func() {
							monsync.Recv(a, func(chan any) { monsync.Nil() })
						},
					)
				},
			)
		},
	)
	time.Sleep(100 * time.Millisecond)

	print("\nSimulation of polyadic (N=3) synchronous π-calculus using monadic asynchronous π-calculus\n")
	print("With the direct encoding\n")
	print("π-process: (ν a)(ν b)(a<b>.0|a(x).0)\n")
	direct.GenChan(
		"a",
		func(a chan chan chan chan any) {
			direct.GenVec(
				"b",
				func(b []chan any) {
					direct.Par(
						func() {
							direct.Send(a, b, direct.Nil)
						},
						func() {
							direct.Recv(a, func([]chan any) { direct.Nil() })
						},
					)
				},
			)
		},
	)
	time.Sleep(100 * time.Millisecond)

	print("\nSimulation of polyadic (N=3) synchronous π-calculus using monadic asynchronous π-calculus\n")
	print("With the indirect encoding\n")
	print("π-process: (ν a)(ν b)(a<b>.0|a(x).0)\n")
	indirect.GenChan(
		"a",
		func(a chan chan chan chan chan chan chan any) {
			indirect.GenVec(
				"b",
				func(b []chan any) {
					indirect.Par(
						func() {
							indirect.Send(a, b, direct.Nil)
						},
						func() {
							indirect.Recv(a, func([]chan any) { direct.Nil() })
						},
					)
				},
			)
		},
	)
	time.Sleep(100 * time.Millisecond)

}
