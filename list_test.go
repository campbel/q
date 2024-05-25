package q

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListPushPop(t *testing.T) {
	assert := assert.New(t)
	l := NewList[int]()
	l.Push(1, 2, 3)
	assert.Equal(3, l.Pop())
	assert.Equal(2, l.Pop())
	assert.Equal(1, l.Pop())
	assert.Zero(l.Pop())
}

func TestListPushPopLeft(t *testing.T) {
	assert := assert.New(t)
	l := NewList[int]()
	l.PushLeft(1, 2, 3)
	assert.Equal(3, l.PopLeft())
	assert.Equal(2, l.PopLeft())
	assert.Equal(1, l.PopLeft())
	assert.Zero(l.PopRight())
}

func TestListLength(t *testing.T) {
	assert := assert.New(t)
	l := NewList[int]()
	l.PushRight(1, 2, 3)
	assert.Equal(int64(3), l.Len64())
	assert.Equal(3, l.Len())
}

func TestSort(t *testing.T) {
	assert := assert.New(t)
	list := NewList[int]()
	count := 10000
	for i := 0; i < count; i++ {
		list.PushRight(rand.Int())
	}
	assert.False(list.IsSorted(LessInt))
	list = list.Sort(LessInt)
	assert.Equal(count, list.Len())
	assert.True(list.IsSorted(LessInt))
}

func LessInt(a, b int) bool {
	return a < b
}

func TestListReverse(t *testing.T) {
	assert := assert.New(t)
	l := NewList[int]()
	l.PushRight(1, 2, 3)
	l.Reverse()
	assert.Equal(3, l.PopLeft())
	assert.Equal(2, l.PopLeft())
	assert.Equal(1, l.PopLeft())
	assert.Zero(l.PopLeft())
}

func TestListJoin(t *testing.T) {
	assert := assert.New(t)
	l1 := NewList[int]()
	l1.PushRight(1, 2, 3)
	l2 := NewList[int]()
	l2.PushRight(4, 5, 6)
	l3 := Join[int](l1, l2)
	assert.Equal(6, l3.Len())
	assert.Equal(1, l3.PopLeft())
	assert.Equal(2, l3.PopLeft())
	assert.Equal(3, l3.PopLeft())
	assert.Equal(4, l3.PopLeft())
	assert.Equal(5, l3.PopLeft())
	assert.Equal(6, l3.PopLeft())
	assert.Zero(l3.PopLeft())
}

func TestReduce(t *testing.T) {
	assert := assert.New(t)
	l := NewList[int]()
	l.PushRight(1, 2, 3)
	sum := Reduce(l, func(acc, value int) int {
		return acc + value
	}, 0)
	assert.Equal(6, sum)
}

func TestMap(t *testing.T) {
	assert := assert.New(t)
	l := NewList[int]()
	l.PushRight(1, 2, 3)
	l = Map[int, int](l, func(value int) int {
		return value * 2
	})
	assert.Equal(2, l.PopLeft())
	assert.Equal(4, l.PopLeft())
	assert.Equal(6, l.PopLeft())
	assert.Zero(l.PopLeft())
}

func TestListEach(t *testing.T) {
	assert := assert.New(t)
	l := NewList[int]()
	l.PushRight(1, 2, 3)
	sum := 0
	l.Each(func(value int) {
		sum += value
	})
	assert.Equal(6, sum)
}

func TestListFind(t *testing.T) {
	assert := assert.New(t)
	l := NewList[int]()
	l.PushRight(1, 2, 3)
	value := l.Find(func(value int) bool {
		return value == 2
	})
	assert.Equal(2, value)

	zero := l.Find(func(value int) bool {
		return false
	})
	assert.Zero(zero)
}

func TestListFilter(t *testing.T) {
	assert := assert.New(t)
	l := NewList[int]()
	l.PushRight(1, 2, 3)
	l = l.Filter(func(value int) bool {
		return value%2 == 0
	})
	assert.Equal(1, l.Len())
	assert.Equal(2, l.PopLeft())
	assert.Zero(l.PopLeft())
}

func TestListEqual(t *testing.T) {
	assert := assert.New(t)
	l1 := NewList(1, 2, 3)
	l2 := NewList(1, 2, 3)
	assert.True(Equal(l1, l2))
	l2.PushRight(4)
	assert.False(Equal(l1, l2))
	l1.PushRight(5)
	assert.False(Equal(l1, l2))
	assert.True(Equal(NewList[int](), NewList[int]()))
	assert.False(Equal(NewList[int](), nil))
	l3 := l1.Copy()
	assert.True(Equal(l1, l3))
}

func TestJoin(t *testing.T) {
	assert := assert.New(t)
	l1 := NewList(1, 2, 3)
	l2 := NewList(4, 5, 6)
	l3 := Join(l1, l2)
	assert.Equal(6, l3.Len())
	assert.Equal(1, l3.PopLeft())
	assert.Equal(2, l3.PopLeft())
	assert.Equal(3, l3.PopLeft())
	assert.Equal(4, l3.PopLeft())
	assert.Equal(5, l3.PopLeft())
	assert.Equal(6, l3.PopLeft())
	assert.Zero(l3.PopLeft())

	l4 := NewList(1)
	l5 := Join(l4)
	assert.Equal(l4, l5)

	l6 := Join[int]()
	assert.Equal(NewList[int](), l6)
}

func TestListExtend(t *testing.T) {
	assert := assert.New(t)
	l1 := NewList(1, 2, 3)
	l2 := NewList(4, 5, 6)
	l1.Extend(l2)
	assert.Equal(6, l1.Len())
	assert.Equal(1, l1.PopLeft())
	assert.Equal(2, l1.PopLeft())
	assert.Equal(3, l1.PopLeft())
	assert.Equal(4, l1.PopLeft())
	assert.Equal(5, l1.PopLeft())
	assert.Equal(6, l1.PopLeft())
	assert.Zero(l1.PopLeft())

	l3 := NewList(1)
	l4 := NewList(2)
	l3.Extend(l4)
	assert.Equal(2, l3.Len())
	assert.Equal(1, l3.PopLeft())
	assert.Equal(2, l3.PopLeft())
	assert.Zero(l3.PopLeft())

	l5 := NewList(1)
	l6 := NewList[int]()
	l5.Extend(l6)
	assert.Equal(1, l5.Len())
	assert.Equal(1, l5.PopLeft())
	assert.Zero(l5.PopLeft())

	l7 := NewList[int]()
	l8 := NewList(1)
	l7.Extend(l8)
	assert.Equal(1, l7.Len())
	assert.Equal(1, l7.PopLeft())
	assert.Zero(l7.PopLeft())
}

func TestAll(t *testing.T) {
	assert := assert.New(t)
	l := NewList(1, 2, 3)
	assert.True(l.All(func(value int) bool {
		return value > 0
	}))
	assert.False(l.All(func(value int) bool {
		return value > 1
	}))
}

func TestAny(t *testing.T) {
	assert := assert.New(t)
	l := NewList(1, 2, 3)
	assert.True(l.Any(func(value int) bool {
		return value > 2
	}))
	assert.False(l.Any(func(value int) bool {
		return value > 3
	}))
}

func TestIndexOf(t *testing.T) {
	assert := assert.New(t)
	l := NewList(1, 2, 3)
	assert.Equal(1, IndexOf(l, 2))
	assert.Equal(-1, IndexOf(l, 4))
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)
	l := NewList(1, 2, 3)
	Remove(l, 2)
	assert.Equal(2, l.Len())
	assert.Equal(1, l.PopLeft())
	assert.Equal(3, l.PopLeft())
	assert.Zero(l.PopLeft())

	l = NewList(1, 2, 3)
	Remove(l, 1)
	assert.Equal(2, l.Len())
	assert.Equal(2, l.PopLeft())
	assert.Equal(3, l.PopLeft())
	assert.Zero(l.PopLeft())

	l = NewList(1, 2, 3)
	Remove(l, 3)
	assert.Equal(2, l.Len())
	assert.Equal(1, l.PopLeft())
	assert.Equal(2, l.PopLeft())
	assert.Zero(l.PopLeft())
}

func TestSlice(t *testing.T) {
	assert := assert.New(t)
	l := NewList(1, 2, 3, 4, 5).Elements()
	assert.Equal([]int{1, 2, 3, 4, 5}, l)
}

func TestString(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("[1,2,3]", NewList(1, 2, 3).String())
	assert.Equal("[]", NewList[int]().String())
	assert.Equal("[a,b,c]", NewList("a", "b", "c").String())
	assert.Equal("[a]", NewList("a").String())
	assert.Equal("[a,b]", NewList("a", "b").String())
	assert.Equal("[true,false]", NewList(true, false).String())
	assert.Equal("[map[1:2 3:4]]", NewList(map[int]int{1: 2, 3: 4}).String())
}

// ExampleList_Push demonstrates how to add elements to the list.
func ExampleList_Push() {
	list := NewList[int]()
	list.Push(1, 2, 3)
	fmt.Println(list)
	// Output: [1,2,3]
}

// ExampleList_Pop demonstrates how to remove and return the last element from the list.
func ExampleList_Pop() {
	list := NewList[int](1, 2, 3)
	last := list.Pop()
	fmt.Println(last)
	fmt.Println(list)
	// Output:
	// 3
	// [1,2]
}

// ExampleList_Reverse demonstrates how to reverse the order of elements in the list.
func ExampleList_Reverse() {
	list := NewList[string]("apple", "banana", "orange")
	list.Reverse()
	fmt.Println(list)
	// Output: [orange,banana,apple]
}

// ExampleList_Each demonstrates how to apply a callback function to each element in the list.
func ExampleList_Each() {
	list := NewList[int](1, 2, 3, 4, 5)
	list.Each(func(value int) {
		fmt.Print(value*2, " ")
	})
	// Output: 2 4 6 8 10
}

// ExampleList_Find demonstrates how to find the first element that satisfies a condition.
func ExampleList_Find() {
	list := NewList[int](1, 2, 3, 4, 5)
	even := list.Find(func(value int) bool {
		return value%2 == 0
	})
	fmt.Println(even)
	// Output: 2
}

// ExampleList_Sort demonstrates how to sort the list using a custom less function.
func ExampleList_Sort() {
	list := NewList[int](5, 2, 4, 1, 3)
	sorted := list.Sort(func(a, b int) bool {
		return a < b
	})
	fmt.Println(sorted)
	// Output: [1,2,3,4,5]
}
