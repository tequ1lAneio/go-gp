package main

import "fmt"

func f() int {
	return -1
}

func h() int {
	return 2
}

func main() {
	if a, c := f(), h(); a > 0 {
		fmt.Println(a)
	} else if b := f(); b > 0 {
		fmt.Println(a, b)
	} else if n := 12; n >100 {
		fmt.Println(2)
	} else {
		fmt.Println(a, b, c, n)
	}
	// -1, -1, 2
}
