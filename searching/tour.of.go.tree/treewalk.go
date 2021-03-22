package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
// In order
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	chOne := make(chan int)
	chTwo := make(chan int)
	var channelOne []int
	var channelTwo []int
	go func () {
		Walk(t1, chOne)
		close(chOne)
	}()
	go func () {
		Walk(t2, chTwo)
		close(chTwo)
	}()
	for i := range chOne {
		channelOne = append(channelOne, i)
	}
	for i := range chTwo {
		channelTwo = append(channelTwo, i)
	}
	for i,_ := range channelTwo {
		if channelOne[i] != channelTwo[i] {
			return false
		}
	}
	return true
}

func main() {
	val := Same(tree.New(1), tree.New(1))
	if val {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}