package q

import (
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

func TestListShift(t *testing.T) {
	assert := assert.New(t)
	l := NewList[int]()
	l.Push(1, 2, 3)
	assert.Equal(1, l.Shift())
	assert.Equal(2, l.Shift())
	assert.Equal(3, l.Shift())
	assert.Zero(l.Shift())
}

func TestListUnshift(t *testing.T) {
	assert := assert.New(t)
	l := NewList[int]()
	l.Unshift(1, 2, 3)
	assert.Equal(3, l.Shift())
	assert.Equal(2, l.Shift())
	assert.Equal(1, l.Shift())
	assert.Zero(l.Pop())
}

func TestListLength(t *testing.T) {
	assert := assert.New(t)
	l := NewList[int]()
	l.Push(1, 2, 3)
	assert.Equal(int64(3), l.Len64())
	assert.Equal(3, l.Len())
}

func TestSort(t *testing.T) {
	assert := assert.New(t)
	list := NewList[int]()
	count := 10000
	for i := 0; i < count; i++ {
		list.Push(rand.Int())
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
	l.Push(1, 2, 3)
	l.Reverse()
	assert.Equal(3, l.Shift())
	assert.Equal(2, l.Shift())
	assert.Equal(1, l.Shift())
	assert.Zero(l.Shift())
}

func TestListJoin(t *testing.T) {
	assert := assert.New(t)
	l1 := NewList[int]()
	l1.Push(1, 2, 3)
	l2 := NewList[int]()
	l2.Push(4, 5, 6)
	l3 := Join[int](l1, l2)
	assert.Equal(6, l3.Len())
	assert.Equal(1, l3.Shift())
	assert.Equal(2, l3.Shift())
	assert.Equal(3, l3.Shift())
	assert.Equal(4, l3.Shift())
	assert.Equal(5, l3.Shift())
	assert.Equal(6, l3.Shift())
	assert.Zero(l3.Shift())
}

func TestReduce(t *testing.T) {
	assert := assert.New(t)
	l := NewList[int]()
	l.Push(1, 2, 3)
	sum := Reduce(l, func(acc, value int) int {
		return acc + value
	}, 0)
	assert.Equal(6, sum)
}

func TestMap(t *testing.T) {
	assert := assert.New(t)
	l := NewList[int]()
	l.Push(1, 2, 3)
	l = Map[int, int](l, func(value int) int {
		return value * 2
	})
	assert.Equal(2, l.Shift())
	assert.Equal(4, l.Shift())
	assert.Equal(6, l.Shift())
	assert.Zero(l.Shift())
}

func TestListEach(t *testing.T) {
	assert := assert.New(t)
	l := NewList[int]()
	l.Push(1, 2, 3)
	sum := 0
	l.Each(func(value int) {
		sum += value
	})
	assert.Equal(6, sum)
}

func TestListFind(t *testing.T) {
	assert := assert.New(t)
	l := NewList[int]()
	l.Push(1, 2, 3)
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
	l.Push(1, 2, 3)
	l = l.Filter(func(value int) bool {
		return value%2 == 0
	})
	assert.Equal(1, l.Len())
	assert.Equal(2, l.Shift())
	assert.Zero(l.Shift())
}

func TestListEqual(t *testing.T) {
	assert := assert.New(t)
	l1 := NewList(1, 2, 3)
	l2 := NewList(1, 2, 3)
	assert.True(Equal(l1, l2))
	l2.Push(4)
	assert.False(Equal(l1, l2))
	l1.Push(5)
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
	assert.Equal(1, l3.Shift())
	assert.Equal(2, l3.Shift())
	assert.Equal(3, l3.Shift())
	assert.Equal(4, l3.Shift())
	assert.Equal(5, l3.Shift())
	assert.Equal(6, l3.Shift())
	assert.Zero(l3.Shift())

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
	assert.Equal(1, l1.Shift())
	assert.Equal(2, l1.Shift())
	assert.Equal(3, l1.Shift())
	assert.Equal(4, l1.Shift())
	assert.Equal(5, l1.Shift())
	assert.Equal(6, l1.Shift())
	assert.Zero(l1.Shift())

	l3 := NewList(1)
	l4 := NewList(2)
	l3.Extend(l4)
	assert.Equal(2, l3.Len())
	assert.Equal(1, l3.Shift())
	assert.Equal(2, l3.Shift())
	assert.Zero(l3.Shift())

	l5 := NewList(1)
	l6 := NewList[int]()
	l5.Extend(l6)
	assert.Equal(1, l5.Len())
	assert.Equal(1, l5.Shift())
	assert.Zero(l5.Shift())

	l7 := NewList[int]()
	l8 := NewList(1)
	l7.Extend(l8)
	assert.Equal(1, l7.Len())
	assert.Equal(1, l7.Shift())
	assert.Zero(l7.Shift())
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
	Delete(l, 2)
	assert.Equal(2, l.Len())
	assert.Equal(1, l.Shift())
	assert.Equal(3, l.Shift())
	assert.Zero(l.Shift())

	l = NewList(1, 2, 3)
	Delete(l, 1)
	assert.Equal(2, l.Len())
	assert.Equal(2, l.Shift())
	assert.Equal(3, l.Shift())
	assert.Zero(l.Shift())

	l = NewList(1, 2, 3)
	Delete(l, 3)
	assert.Equal(2, l.Len())
	assert.Equal(1, l.Shift())
	assert.Equal(2, l.Shift())
	assert.Zero(l.Shift())
}

func TestSlice(t *testing.T) {
	assert := assert.New(t)
	l := NewList(1, 2, 3, 4, 5).Slice()
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
