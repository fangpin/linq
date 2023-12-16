package linq

import (
	"sort"
	"testing"

	"github.com/magiconair/properties/assert"
)

type student struct {
	id   int
	name string
	std  int
}

type standard struct {
	std     int
	stdName string
}

type ret struct {
	stdName  string
	students []student
}

func TestGroupJoin(t *testing.T) {
	students := []student{
		{1, "Alice", 1},
		{2, "Bob", 1},
		{3, "Cidy", 2},
		{4, "David", 2},
		{5, "Eva", 3},
	}
	standardList := []standard{
		{1, "standard 1"},
		{2, "standard 2"},
		{3, "standard 3"},
	}
	target := GroupJoin(
		standardList,
		students,
		func(s standard) int {
			return s.std
		},
		func(stu student) int {
			return stu.std
		},
		func(std standard, groupedStudents []student) ret {
			return ret{
				stdName:  std.stdName,
				students: groupedStudents,
			}
		})

	sort.Slice(target, func(i, j int) bool {
		return target[i].stdName < target[j].stdName
	})
	assert.Equal(t, len(target), 3)
	assert.Equal(t, target[0].stdName, "standard 1")
	assert.Equal(t, target[1].stdName, "standard 2")
	assert.Equal(t, target[2].stdName, "standard 3")
	assert.Equal(t, len(target[0].students), 2)
	assert.Equal(t, len(target[1].students), 2)
	assert.Equal(t, len(target[2].students), 1)
}
