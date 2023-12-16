package linq

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

type person struct {
	name string
	age  int
}

func TestGroupBy(t *testing.T) {
	people := []person{
		{"Alice", 5},
		{"Allen", 6},
		{"Bob", 7},
		{"Borris", 8},
		{"Cindy", 9}}
	// group by first letter of name
	groups := GroupBy(people,
		func(x person) string {
			return string(x.name[0])
		},
		func(p []person) float64 {
			return Average(Select(p, func(x person) int { return x.age }))
		})

	assert.Equal(t, len(groups), 3)
	assert.Equal(t, groups["A"], 5.5)
	assert.Equal(t, groups["B"], 7.5)
	assert.Equal(t, groups["C"], 9.0)
}
