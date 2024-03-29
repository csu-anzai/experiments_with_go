package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func Walk(t *Tree, c chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, c)
	c <- t.Value
	Walk(t.Right, c)
}

func Walker(t *Tree) <-chan int {
	ch := make(chan int)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}

func Compare(t1 *Tree, t2 *Tree) bool {
	c1, c2 := Walker(t1), Walker(t2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			break
		}
	}
	return false
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		//fmt.Printf("Inserting %d\n", v)
		return &Tree{
			Left:  nil,
			Right: nil,
			Value: v,
		}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func New(n, k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(n) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func inOrder2(root *Tree) {
	if root == nil {
		return
	} else {
		inOrder2(root.Left)
		fmt.Println(root.Value)
		inOrder2(root.Right)
	}

}

func main() {
	t := New(12, 1)
	//inOrder2(t)
	fmt.Println(Compare(New(12, 2), t))
	fmt.Println(Compare(New(12, 1), t))
	fmt.Println(Compare(New(13, 1), t))
}
