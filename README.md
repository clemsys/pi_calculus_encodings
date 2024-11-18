# Encoding of polyadic synchronous π-calculus in monadic asynchronous π-calculus

Author: [Clément Chapot](mailto:contact@clementchapot.com)<br/>
Description: Code submission for practical of the CS course "Distributed Processes, Types and Programming" at Oxford

## Running

```bash
$ go run .
```

## Project structure

For all encodings, terms are represented using Go functions.

```
.
├── stats/stats.go              # instrumentation utils
├── chanprinter/chanprinter.go  # utils to print terms
├── main.go                     # usage examples
├── monasync/monasync.go        # simulation of monadic async π-calculus using Go channels
├── monsync/monsync.go          # simulation of monadic sync π-calculus using monadic async π-calculus
└── polysync                    # simulation of polyadic sync π-calculus...
    ├── direct/direct.go          # directly using monadic async π-calculus
    └── indirect/indirect.go      # indirectly using sync π-calculus
```

There are two different paths to simulate asynchronous π-calculus:
1. Direct encoding to monadic asynchronous π-calculus, using `monasync.go`
2. Indirect encoding to monadic synchronous π-calculus using `monsync.go`, which itself uses `monasync.go`

An effort has been made to only use functions from `monasync.go` in both `monsync.go` and `direct.go`, 
and to only use functions from `monsync.go` in `indirect.go`.

The style of the code is also mostly functionnal, to show clearly that the successive encodings
are obtained from the previous ones without using any additionnal features from Go.

## Why the direct encoding requires less messages and channels than the direct encoding

With the **direct encoding**, encoding a `Send` or a `Recv` causes one additional channel to be generated: there is only one nu in each of the two following formulae:
- `[[u(x1, ..., xn).P]] = u(y).(v d)(y<d>|d(x1).(y<d>|d(x2).(...(y<d>|d(xn).[[P]])...)))`
- `[[u<v1, ..., vn>.P]] = (v c)(u<c>|c(y1).(y1<v1>|c(y2).(y2<v2>|...|c(yn).(yn<vn>|[[P]])...)))`

With the **indirect encoding** however, encoding a `Send` causes `N+2` additional channels to be generated (1 by `[[.]]_poly` and `N+1` by `[[.]]_sync`) and encoding a `Recv` causes `N+1` additional channels to be generated (all by `[[.]]_sync`).

Since we are considering polyadic **synchronous** π-calculus, a term always contain as many `Recv` as `Send`. 
Therefore, overall, the direct encoding causes `2 * n_send` additional channels to be generated whereas the indirect one causes `(2N + 3) * n_send` additional channels to be generated, which is obviously more.
