package q_test

import (
	"fmt"

	"github.com/campbel/q"
)

func Example() {
	list := q.NewList(1, 2, 3)
	fmt.Println(list)

	heap := q.NewHeap(func(a, b int) bool {
		return a < b
	}, 1, 2, 3)
	fmt.Println(heap)

	set := q.NewSet(1, 2, 3)
	fmt.Println(set)
}
