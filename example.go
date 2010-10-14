
//                              Tree Comparison


// From http://golang.org/

// Two binary trees may be of different shapes, 
// but have the same contents. For example:
//
//        4               6
//      2   6          4     7
//     1 3 5 7       2   5
//                  1 3
//
// Go's concurrency primitives make it easy to
// traverse and compare the contents of two trees
// in parallel.

package main

import (
	"fmt"
	"rand"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk traverses a tree depth-first, 
// sending each Value on a channel.
func Walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Walker launches Walk in a new goroutine,
// and returns a read-only channel of values.
func Walker(t *Tree) <-chan int {
	ch := make(chan int)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}

// Compare reads values from two Walkers
// that run simultaneously, and returns true
// if t1 and t2 have the same contents.
func Compare(t1, t2 *Tree) bool {
	c1, c2 := Walker(t1), Walker(t2)
	for <-c1 == <-c2 {
		if closed(c1) || closed(c1) {
			return closed(c1) == closed(c2)
		}
	}
	return false
}

// New returns a new, random binary tree
// holding the values 1k, 2k, ..., nk.
func New(k, n int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(n) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func main() {
	t1 := New(1, 100)
	fmt.Println(Compare(t1, New(1, 100)), "Same Contents")
	fmt.Println(Compare(t1, New(1, 99)), "Differing Sizes")
	fmt.Println(Compare(t1, New(2, 100)), "Differing Values")
	fmt.Println(Compare(t1, New(2, 101)), "Dissimilar")
}
