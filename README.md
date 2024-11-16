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
are obtained from the previous ones without really using additionnal features from Go.
