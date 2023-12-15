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
you can see the respective test function to check how to use it.

- Select
- Take
- Where

# examples
Select
```go
source := []int{1, 2, 3}
target := Select(source, func(x int) string { return strconv.Itoa(x * x) })
assert.Equal(t, target, []string{"1", "4", "9"})
```

Take
```go
source := []int{1, 2, 3, 4}
target := Take(source, 3)
assert.Equal(t, target, []int{1, 2, 3})
```

Where
```go
source := []int{1, 2, 3, 4}
target := Where(source, func(x int) bool { return x%2 == 1 })
assert.Equal(t, target, []int{1, 3})
```