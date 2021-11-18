package main

import (
	"fmt"
	"unsafe"
)

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func createArrays() {
	var arr1 [5]int
	arr2 := [5]int{}
	arr3 := [5]int{1,2,3,4,5}
	arr4 := [5]int{1,2,3}
	arr5 := [...]int{1,2,3,4,5}
	arr6 := [5]string{0: "first", 2: "third", 1: "second"}

	fmt.Println(arr1, arr2, arr3, arr4, arr5, arr6)
}

func createSlices() {
	var slice1 []int
	slice2 := []int{1, 2, 3, 4, 5}
	slice3 := make([]int, 5)
	slice4 := make([]int, 5, 10)
	slice5 := make([]int, 0, 10)
	slice6 := []string{0: "first", 2: "third", 1: "second"}

	fmt.Println(slice1, slice2, slice3, slice4, slice5, slice6)

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice7 := arr[3:7:9]

	fmt.Println(slice7)
	slice1 = append(slice7, 3,4,5,6,7,8,9,9,123)
	fmt.Println(slice7)
}

func main() {
	arr := [3]int{
		2: 4,
	}
	arr[0] = 666
	fmt.Println(arr)
	createSlices()
}
