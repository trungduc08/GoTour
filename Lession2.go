package main

import "fmt"

type Vector struct {
	X int
	Y int
}

var (
	v1 = Vector{1, 2}  // has type Vertex
	v2 = Vector{X: 1}  // Y:0 is implicit
	v3 = Vector{}      // X:0 and Y:0
	p  = &Vector{1, 2} // has type *Vertex
)

func main() {
	// s := []int{2, 3, 5, 7, 11, 13}
	// s = s[:0]

	// fmt.Println(len(s), s)
	// fmt.Println(cap(s), s)
	// a := make([]int, 5)
	// printSlice("a", a)

	// b := make([]int, 0, 5)
	// printSlice("b", b)

	// c := b[:2]
	// printSlice("c", c)

	// d := c[2:5]
	// printSlice("d", d)
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
		// fmt.Printf("%d\n", pow[i])
	}
	fmt.Println("----------------------------")
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

}
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
