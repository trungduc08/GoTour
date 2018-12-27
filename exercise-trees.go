package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {

	Walk2(t, ch)
	close(ch)

}
func Walk2(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	if t.Left != nil {
		Walk2(t.Left, ch)
	}

	if t.Right != nil {
		Walk2(t.Right, ch)
	}
	ch <- t.Value
}
func CheckTwoTree(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	// for i := range ch1 {

	// }
	return false

}
func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}

}
