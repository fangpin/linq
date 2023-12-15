This project is to provide linq-like syntax to Golang.

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