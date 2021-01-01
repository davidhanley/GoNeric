package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestTree(t *testing.T) {
	var tree *Node

	t1 := insert(tree, 10, 10)

	if t1 == nil {
		t.Error("Tree is nil")
	}

	t2 := insert(t1, 4, 4)

	if t2.left == nil {
		t.Error("Tree left is nil")
	}

	if t2.left.key != 4 {
		t.Error("wrong value as left child")
	}

	t2 = nil

	for a := 0; a < 10000; a++ {
		t2 = insert(t2, TreeMapType1(a), TreeMapType2(a))
	}

	if t2.height > 15 {
		t.Error(fmt.Sprintf("height not maintianed:%d\n", t2.height))
	}
	fmt.Printf("10000 node tree height: %d", t2.height)
}

func TestTreeIter(t *testing.T) {
	var tree *Node
	for a := 0; a < 10; a++ {
		tree = insert(tree, TreeMapType1(a), TreeMapType2(a))
	}

	c := make(chan TreeMapType1)

	go toStream(tree, c)

	//do we iterate back the right values?
	for a := 0; a < 10; a++ {
		if <-c != TreeMapType1(a) {
			t.Error("didn't get the right values back from the tree!")
		}
	}

	tree2 := tree
	//test deletion.. delete the odd numbers
	for a := 0; a < 10; a += 2 {
		tree2 = delete(tree2, TreeMapType1(a))
	}

	//now loop through and make sure it's only evens left
	c2 := make(chan TreeMapType1)

	go toStream(tree2, c2)

	for a := 1; a < 10; a += 2 {
		if <-c2 != TreeMapType1(a) {
			t.Error("didn't get the right values back from the deleted tree!")
		}
	}

	//now throw in a few new #'s

	tree2 = insert(tree2, 55, 55)
	tree2 = insert(tree2, 66, 55)

	//now see if the original 0..9 tree is still there

	c3 := make(chan TreeMapType1)

	go toStream(tree, c3)

	if size(tree) < 10 {
		t.Error(fmt.Printf("tree too smol %d\n", size(tree)))
	}
	//do we iterate back the right values?
	for a := 0; a < size(tree); a++ {
		v := <-c3
		if v != TreeMapType1(a) {
			t.Error(fmt.Sprintf("didn't get the right values back from the tree! %d %d ", a, v))
		}
	}
}

func doARound(t *testing.T) {
	var tree *Node
	ar := make([]int, 0)

	for a := 0; a < 100; a++ {
		tree = insert(tree, TreeMapType1(a), TreeMapType2(a))
		ar = append(ar, a)
	}

	rand.Shuffle(len(ar), func(i, j int) {
		ar[i], ar[j] = ar[j], ar[i]
	})

	for a := 0; a < 100; a++ {
		t2 := delete(tree, TreeMapType1(ar[a]))
		if size(tree) != 100-a {
			t.Error("old tree shrank")
		}
		if size(t2) != (100-a)-1 {
			t.Error("new tree didn't shrink")
		}
		tree = t2
	}

}

func TestTreeShuffleDel(t *testing.T) {
	for a := 0; a < 100; a++ {
		doARound(t)
	}
}

func BenchmarkTree(b *testing.B) {
	for a := 0; a < 1000; a++ {
		var tree *Node
		for i := 0; i < 10000; i++ {
			tree = insert(tree, TreeMapType1(i), TreeMapType2(i))
		}
	}
}
