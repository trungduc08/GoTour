package main

import "fmt"

type Vector struct {
	X int
	Y int
}
type Vertex struct {
	Lat, Long float64
}

// var m map[string]Vertex

var (
	v1 = Vector{1, 2}  // has type Vertex
	v2 = Vector{X: 1}  // Y:0 is implicit
	v3 = Vector{}      // X:0 and Y:0
	p  = &Vector{1, 2} // has type *Vertex
)
var lstMap = map[string]Vertex{
	"map1": {10.10, 20.20},
	"map2": {30.30, 40.40},
}

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer1"] = 48
	fmt.Println("The value:", m["Answer1"])

	m["Answer2"] = 49
	fmt.Println("The value:", m["Answer2"])

	m["Answer6"] = 1
	fmt.Println("The value:", m["Answer6"])

	// delete(m, "Answer")
	// fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer6"]
	fmt.Println("The value:", v, "Present?", ok)
}
