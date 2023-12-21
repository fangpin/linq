# Linq-like syntax for Golang.

![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-%23007d9c)
[![GoDoc](https://godoc.org/github.com/fangpin/linq?status.svg)](https://pkg.go.dev/github.com/fangpin/linq)
![Build Status](https://github.com/fangpin/linq/actions/workflows/test.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/fangpin/linq)](https://goreportcard.com/report/github.com/fangpin/linq)
[![Coverage](https://img.shields.io/codecov/c/github/fangpin/linq)](https://codecov.io/gh/fangpin/linq)
[![License](https://img.shields.io/github/license/fangpin/linq)](./LICENSE)

## Install

```sh
go get github.com/fangpin/linq
```

# Supported linq operation
- aggregate
- all
- any
- average
- chunk
- contains
- count
- default_if_empty
- distinct
- distinct_by
- except
- except_by
- first
- first_or_default
- group_by
- group_join
- intersect
- intersect_by
- keys
- last
- last_or_default
- max
- max_by
- min
- min_by
- of_type
- order
- order_by
- order_by_descending
- order_descending
- prepend
- repeat
- reverse
- select
- select_many
- single
- single_or_default
- skip
- skip_last
- skip_while
- sum
- take
- take_last
- take_while
- to_dictionary
- to_hash_set
- to_lookup
- union
- union_by
- values
- where
- zero
- zip
