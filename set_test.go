package q

import (
	"fmt"
	"slices"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	assert := assert.New(t)
	set := NewSet[int]()
	assert.NotNil(set)
	assert.Equal(0, set.Len())
}

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	set := NewSet[int]()
	set.Add(1, 2, 3)
	assert.Equal(3, set.Len())
	assert.True(set.Contains(1))
	assert.True(set.Contains(2))
	assert.True(set.Contains(3))
}

func TestRemove(t *testing.T) {
	assert := assert.New(t)
	set := NewSet[int]()
	set.Add(1, 2, 3)
	set.Remove(2)
	assert.Equal(2, set.Len())
	assert.True(set.Contains(1))
	assert.False(set.Contains(2))
	assert.True(set.Contains(3))
}

func TestContains(t *testing.T) {
	assert := assert.New(t)
	set := NewSet[string]()
	set.Add("a", "b", "c")
	assert.True(set.Contains("a"))
	assert.True(set.Contains("b"))
	assert.True(set.Contains("c"))
	assert.False(set.Contains("d"))
}

func TestLen(t *testing.T) {
	assert := assert.New(t)
	set := NewSet[int]()
	assert.Equal(0, set.Len())
	set.Add(1, 2, 3)
	assert.Equal(3, set.Len())
	set.Remove(2)
	assert.Equal(2, set.Len())
}

func TestClear(t *testing.T) {
	assert := assert.New(t)
	set := NewSet[int]()
	set.Add(1, 2, 3)
	set.Clear()
	assert.Equal(0, set.Len())
}

func TestElements(t *testing.T) {
	assert := assert.New(t)
	set := NewSet[int]()
	set.Add(1, 2, 3)
	elements := set.Elements()
	assert.ElementsMatch([]int{1, 2, 3}, elements)
}

func TestUnion(t *testing.T) {
	assert := assert.New(t)
	set1 := NewSet[int]()
	set1.Add(1, 2, 3)
	set2 := NewSet[int]()
	set2.Add(3, 4, 5)
	unionSet := set1.Union(set2)
	assert.Equal(5, unionSet.Len())
	assert.True(unionSet.Contains(1))
	assert.True(unionSet.Contains(2))
	assert.True(unionSet.Contains(3))
	assert.True(unionSet.Contains(4))
	assert.True(unionSet.Contains(5))
}

func TestIntersection(t *testing.T) {
	assert := assert.New(t)
	set1 := NewSet[int]()
	set1.Add(1, 2, 3)
	set2 := NewSet[int]()
	set2.Add(3, 4, 5)
	intersectionSet := set1.Intersection(set2)
	assert.Equal(1, intersectionSet.Len())
	assert.True(intersectionSet.Contains(3))
}

func TestDifference(t *testing.T) {
	assert := assert.New(t)
	set1 := NewSet[int]()
	set1.Add(1, 2, 3)
	set2 := NewSet[int]()
	set2.Add(3, 4, 5)
	differenceSet := set1.Difference(set2)
	assert.Equal(2, differenceSet.Len())
	assert.True(differenceSet.Contains(1))
	assert.True(differenceSet.Contains(2))
}

// Example_Set_Add demonstrates how to add elements to a set.
func ExampleSet_Add() {
	set := NewSet[int]()
	set.Add(1, 2, 3)
	elements := set.Elements()
	slices.Sort(elements)
	fmt.Println(elements)
	// Output: [1 2 3]
}

// Example_Set_Contains demonstrates how to check if an element exists in a set.
func ExampleSet_Contains() {
	set := NewSet[string]()
	set.Add("apple", "banana", "orange")
	fmt.Println(set.Contains("banana"))
	fmt.Println(set.Contains("grape"))
	// Output:
	// true
	// false
}

// Example_Set_Union demonstrates how to find the union of two sets.
func ExampleSet_Union() {
	set1 := NewSet[int]()
	set1.Add(1, 2, 3)
	set2 := NewSet[int]()
	set2.Add(3, 4, 5)
	unionSet := set1.Union(set2)
	sorted := sort.IntSlice(unionSet.Elements())
	sorted.Sort()
	fmt.Println(sorted)
	// Output: [1 2 3 4 5]
}
