package q

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounterCount(t *testing.T) {
	assert := assert.New(t)
	counter := NewCounter("a", "b", "c")
	counter.Add("a", "b", "c", "d")
	assert.Equal(2, counter.Count("a"))
	assert.Equal(2, counter.Count("b"))
	assert.Equal(2, counter.Count("c"))
	assert.Equal(1, counter.Count("d"))
	assert.Equal(0, counter.Count("e"))

	assert.Equal(7, counter.Len())
	assert.Equal(4, len(counter.Elements()))
	counter.Clear()
	assert.Equal(0, counter.Len())
	assert.True(counter.IsEmpty())
}

func TestCounterString(t *testing.T) {
	assert := assert.New(t)
	counter := NewCounter("a", "b", "c")
	expected := "map[a:1 b:1 c:1]"
	assert.Equal(expected, counter.String())
}

func TestCounterEqual(t *testing.T) {
	assert := assert.New(t)
	counter1 := NewCounter("a", "b", "c")
	counter2 := NewCounter("a", "b", "c")
	assert.True(counter1.Equal(counter2))

	counter2.Add("d")
	assert.False(counter1.Equal(counter2))

	counter2.Remove("d")
	assert.True(counter1.Equal(counter2))

	counter1.Add("b")
	counter2.Add("a")
	assert.False(counter1.Equal(counter2))
}

func TestCounterAdd(t *testing.T) {
	assert := assert.New(t)
	counter := NewCounter("a", "b", "c")
	counter.Add("a", "b", "c", "d")

	expected := map[string]int{
		"a": 2,
		"b": 2,
		"c": 2,
		"d": 1,
	}

	for element, count := range expected {
		assert.Equal(count, counter.Count(element))
	}
}

func TestCounterRemove(t *testing.T) {
	assert := assert.New(t)
	counter := NewCounter("a", "b", "c")
	counter.Add("a", "b", "c", "d")

	counter.Remove("a", "b", "c", "d")
	assert.Equal(3, counter.Len())
}
