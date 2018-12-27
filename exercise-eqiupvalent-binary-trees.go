package main

func Walk(t *Tree, ch chan int) {
	if t.Left != nil {
		Walk(t, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	close(ch)

}
func main() {
	tree := Tree{}
	ch := make(chan int)
	go Walk(tree.New(1), ch)
}
