[![Go Reference](https://pkg.go.dev/badge/github.com/campbel/q.svg)](https://pkg.go.dev/github.com/campbel/q)
[![Go](https://github.com/campbel/q/actions/workflows/go.yml/badge.svg)](https://github.com/campbel/q/actions/workflows/go.yml)

# q

This repository contains implementations of various data structures in Go, including:

- List
- Heap
- Set

```go
import "github.com/campbel/q"

func main() {
	list := q.NewList(1, 2, 3)
	fmt.Println(list)

	heap := q.NewHeap(func(a, b int) bool {
		return a < b
	}, 1, 2, 3)
	fmt.Println(heap)

	set := q.NewSet(1, 2, 3)
	fmt.Println(set)
}
```
