package main

import "fmt"

func switchWithNoExpr() {
	switch {
	case true:
		fmt.Println("case 1")
	case false:
		fmt.Println("case 2")
	}
}

func switchWithMultiCases(a int) {
	switch a {
	case 1, 2, 3, 4, 5:
		println("it is a work day")
	case 6, 7:
		println("it is a weekend day")
	default:
		println("you are live on earth")
	}
}

func getValue() int {
	return 1
}

func switchWithFallThrough() {
	switch 1 {
	case 1:
		println("case 1.1")
		fallthrough
	case 2:
		println("case 2.1")
		fallthrough
	case getValue():
		println("case 3.1")
		fallthrough
	default:
		println("default")
	}
}

func typeSwitch() {
	var x interface{} = 13
	switch x.(type) {		// "x" must be an interface variable.
	case nil:
		print("x is nil")
	case int:
		println("the type of x is int")
	case string:
		println("the type of x is string")
	case bool:
		println("the type of x is boolean")
	default:
		println("doesn't support the type")
	}

	switch v := x.(type) {
	case nil:
		println("v is nil")
	case int:
		fmt.Printf("the type of v is int, x = %T v = %d", x, v)
	case string:
		println("the type of v is string, v = ", v)
	case bool:
		println("the type of v is bool, v = ", v)
	default:
		println("doesn't support the type")
	}
}

func main() {
	//switchWithNoExpr()
	//switchWithMultiCases(88)
	//switchWithFallThrough()
	typeSwitch()
}
