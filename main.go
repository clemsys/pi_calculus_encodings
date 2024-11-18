// Filename: main.go
// Usage: go run .
// Description: Usage examples of the different encodings

package main

import (
	"fmt"
	"piencodings/chanprinter"
	"piencodings/monasync"
	"piencodings/monsync"
	"piencodings/polysync/direct"
	"piencodings/polysync/indirect"
	"piencodings/stats"
	"time"
)

func main() {

	print("[Monadic asynchronous]\n  π-process: (ν a)(ν b)(a<b>|a(x).0)\n  Calls:")
	chanprinter.ChanTblReset()
	stats.GlobalStats.Reset()
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
	fmt.Print("\n")
	stats.GlobalStats.PrintStats()

	print("\n\n[Monadic asynchronous]\n  π-process: (ν a)(ν b)(ν c)(a<b>|b<c>|a(x).x(y).0)\n  Calls:")
	chanprinter.ChanTblReset()
	stats.GlobalStats.Reset()
	monasync.Gen(
		"a",
		func(a chan chan chan any) {
			monasync.Gen(
				"b",
				func(b chan chan any) {
					monasync.Gen(
						"c",
						func(c chan any) {
							monasync.Par(
								func() {
									monasync.Send(a, b)
								},
								func() {
									monasync.Par(
										func() {
											monasync.Send(b, c)
										},
										func() {
											monasync.Recv(a, func(x chan chan any) {
												monasync.Recv(x, func(y chan any) {
													monasync.Nil()
												})
											})
										},
									)
								},
							)
						},
					)
				},
			)
		},
	)
	time.Sleep(100 * time.Millisecond)
	fmt.Print("\n")
	stats.GlobalStats.PrintStats()

	print("\n\n[Monadic synchronous]\n  simulated using monadic asynchronous π-calculus\n")
	print("  π-process (ν a)(ν b)(a<b>.0|a(x).0)\n  Calls:")
	chanprinter.ChanTblReset()
	stats.GlobalStats.Reset()
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
	fmt.Print("\n")
	stats.GlobalStats.PrintStats()

	print("\n\n[Polyadic (N=3) synchronous][Direct]\n  simulated using monadic asynchronous π-calculus\n")
	print("  π-process: (ν a)(ν b)(a<b>.0|a(x).0)\n  Calls:")
	chanprinter.ChanTblReset()
	stats.GlobalStats.Reset()
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
	fmt.Print("\n")
	stats.GlobalStats.PrintStats()

	print("\n\n[Polyadic (N=3) synchronous][Indirect]\n  simulated using monadic asynchronous π-calculus\n")
	print("  π-process: (ν a)(ν b)(a<b>.0|a(x).0)\n  Calls:")
	chanprinter.ChanTblReset()
	stats.GlobalStats.Reset()
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
	fmt.Print("\n")
	stats.GlobalStats.PrintStats()

}
