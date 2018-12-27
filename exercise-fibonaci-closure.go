package main

func fibonacci() func() int {
	x, y := 1, 0
	var temp int
	return func() int {
		temp = x
		x = y
		y = y + temp
		return x
	}
}

// func main() {
// 	f := fibonacci()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(f())
// 	}
// }
