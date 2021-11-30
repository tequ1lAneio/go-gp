package main

import (
	"fmt"
	"time"
)

func main() {
	//basicFor()
	//multiVariableFor()
	//forSliceRange()
	//forStringRange()
	//forMapRange()
	//forChannelRange()
	//forWithLabel()
	//forInGoRoutine()
	//forInSlice()
	forInMap()
}

func basicFor() {
	var sum int
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)

	// no conditional judgement expression
	var sum1 int
	for i := 0; i < 10; {
		i++
		sum += 1
	}
	fmt.Println(sum1)

	// no post-loop statement
	var sum2 int
	j := 0
	for ; j < 10; j++ {
		sum2 += 1
	}
	fmt.Println(sum2)

	// no conditional judgement expression & no post-loop statement
	var sum3 int
	k := 0
	for ; k < 10; {
		k++
		sum3 += sum3
	}
	fmt.Println(sum3)

	// no semicolon
	var sum4 int
	z := 0
	for z < 10 {
		z++
		sum4 += sum4
	}
	fmt.Println(sum4)
}

func multiVariableFor() {
	var sum int
	for i, j, k := 0, 0, 0; i < 12 && j < 10 && k < 8; i, j, k = i+1, j+1, k+1 {
		sum += 1
	}
	fmt.Println(sum)
}

func forSliceRange() {
	sl := []int{1, 3, 4, 5, 6}
	for i, v := range sl {
		fmt.Printf("sl[%d] = %d\n", i, v)
	}

	for i := range sl {
		println(i)
	}

	for _, v := range sl {
		println(v)
	}

	//for _, _ := range sl {} // this form is legal but is no commended.
	for range sl {
	}
}

func forStringRange() {
	var s = "中国人"
	for i, v := range s { // "v" is a unicode code point
		fmt.Printf("%d %s 0x%x\n", i, string(v), v)
	}
}

func forMapRange() {
	var myMap = map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}
	for k, v := range myMap {
		fmt.Printf("%d, %d\n", k, v)
	}
}

func forChannelRange() {
	var c = make(chan int)
	for v := range c {
		fmt.Println(v)
	}
}

func forWithLabel() {
	var sum int
	var sl = []int{1, 2, 3, 4, 5}
outerloop:
	for i := 0; i < len(sl); i++ {
		for j := 0; j < 5; j++ {
			if sl[i]%2 == 0 {
				continue outerloop
			}
		}
		sum += sl[i]
	}
	fmt.Println(sum)

	slice := [][]int{
		{1, 34, 26, 35, 78},
		{3, 45, 13, 24, 99},
		{101, 13, 38, 7, 127},
		{54, 27, 40, 83, 81},
	}
outerloop1:
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			if slice[i][j] == 38 {
				fmt.Printf("found gold at [%d, %d]\n", i, j)
				break outerloop1
			}
		}
	}
}

func forInGoRoutine() {
	var s = []int{1, 2, 3, 4, 5}

	for i, v := range s {
		go func() {
			time.Sleep(time.Second * 3)
			fmt.Println(i, v)
		}()
	}

	time.Sleep(time.Second * 10)

	for i, v := range s {
		go func(i, v int) {
			time.Sleep(time.Second * 3)
			fmt.Println(i, v)
		}(i, v)
	}
}

func forInSlice() {
	var s = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("original s = ", s)

	//for i, v := range s {
	for i, v := range s[:] {
		if i == 0 {
			s[1] = 12
			s[2] = 13
		}
		r[i] = v
	}

	fmt.Println("after range s = ", s)
	fmt.Println("after range r = ", r)
}

func forInMap() {
	var m = map[string]int {
		"tony": 21,
		"tom": 22,
		"jim": 23,
	}

	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "tony")
		}
		if counter == 2 {
			m["lucy"] = 24
		}
		counter++
		fmt.Println(k, v)
	}

	fmt.Println("counter is", counter)
}
