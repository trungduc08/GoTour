package main

import (
	"fmt"
	"time"
)

type I interface {
	M()
}
type T struct {
	S string
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}
func Describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func Do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v x 2 = %v\n", v, v*2)
	case string:
		fmt.Printf("%T\n", i)
	default:
		fmt.Printf("default")
	}
}

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v %v", p.name, p.age)
}

type MyError struct {
	when time.Time
	what string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.when, e.what)
}
func run() error {
	return &MyError{
		time.Now(),
		"did not work",
	}
}

// func main() {
// 	r := strings.NewReader("Hello, Reader!")
// 	b := make([]byte, 8)

// 	// for {
// 	// 	n, err := r.Read(b)

// 	// 	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
// 	// 	if err == io.EOF {
// 	// 		// fmt.Printf("%v \n", n)
// 	// 		break
// 	// 	}
// 	// }
// 	for i := 0; i < len(b); i++ {
// 		n, err := r.Read(b)
// 		fmt.Printf("n = %v err = %v b = (%c)\n", n, err, b)
// 		fmt.Printf("b[:n] = %q\n", b[:n])
// 		if err == io.EOF {
// 			break
// 		}
// 	}

// }
