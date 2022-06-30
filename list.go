package main

import (
	"fmt"
)

// In addition to generic functions, Go also supports generic types. A type can be parameterized with a type parameter, which could be useful for implementing generic data structures.
// This example demonstrates a simple type declaration for a singly-linked list holding any type of value.
// As an exercise, add some functionality to this list implementation.

// List represents a singly-linked list that holds values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// function
func traverse[T any](l *List[T]) {
	for l != nil {
		fmt.Println(l.val)
		l = l.next
	}
}

// method
func (l *List[T]) traverse() {
	for l != nil {
		fmt.Println(l.val)
		// Assignment to the method receiver propagates only to callees but not to callers
		l = l.next
	}
}

func main() {
	node5 := List[int]{nil, 5}
	node4 := List[int]{&node5, 4}
	node3 := List[int]{&node4, 3}
	node2 := List[int]{&node3, 2}
	node1 := List[int]{&node2, 1}

	traverse(&node1)
	node1.traverse()
	// Assignment to the method receiver propagates only to callees but not to callers.
	fmt.Println(node1.val)
}
