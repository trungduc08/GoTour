package main

import (
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

	ch <- t.Value
	if t.Right != nil {
		Walk2(t.Right, ch)
	}

}
func CheckSame(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		item1, valid1 := <-ch1
		item2, _ := <-ch2
		if !valid1 {
			break
		}
		if item1 != item2 {
			return false
			break
		}

	}
	return true

}

// func main() {
// 	ch := make(chan int)
// 	treeTest := tree.New(1)
// 	go Walk(treeTest, ch)
// 	fmt.Println(treeTest.String())
// 	fmt.Println("----------")
// 	s := rand.Perm(10)
// 	fmt.Println(s)
// 	for {
// 		i, valid := <-ch
// 		if !valid {
// 			break
// 		}
// 		fmt.Println(i)
// 	}
// 	fmt.Println(CheckSame(tree.New(1), tree.New(2)))
// 	// ch := make(chan int, 2)
// 	// ch <- 1
// 	// ch <- 2
// 	// fmt.Println(<-ch)
// 	// fmt.Println(<-ch)

// }
