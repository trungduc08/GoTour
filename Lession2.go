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
func SumGo(s []int, c chan int) chan int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
	// fmt.Println(c)
	return c
}
func GoFibonaci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}

}
func main() {
	// c := make(chan int)
	// quit := make(chan int)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(<-c)
	// 	}
	// 	quit <- 0
	// }()
	// GoFibonaci(c, quit)
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
