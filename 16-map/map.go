package main

import "fmt"

func defineMaps() {
	var m1 map[string]int
	//m1["key"] = 1 	// panic: assignment to entry in nil map

	m2 := map[int]string{}

	m3 := map[int][]string{
		1: []string{"val1_1", "val1_2"},
		3: []string{"val3_1", "val3_2", "val3_3"},
		7: []string{"val7_1"},
	}

	type Position struct {
		x float64
		y float64
	}
	m4 := map[Position]string {
		Position{29.935523, 52.568915}: "school",
		Position{25.352594, 113.304361}: "shopping-mall",
		Position{73.224455, 111.804306}: "hospital",
	}
	m5 := map[Position]string {
		{29.935523, 52.568915}: "school",
		{25.352594, 113.304361}: "shopping-mall",
		{73.224455, 111.804306}: "hospital",
	}

	m6 := make(map[int]string)
	m7 := make(map[int]string, 5)

	fmt.Println(m1, m2, m3, m4)
}

func addKey() {
	m1 := make(map[int]string)
	m1[1] = "value1"
	m1[2] = "value2"
	m1[3] = "value3"

	m2 := map[string]int {
		"key1": 1,
		"key2": 2,
	}
	m2["key1"] = 11
	m2["key3"] = 666
}

func compareTypes() {
	s1 := make([]int, 1)
	s2 := make([]int, 2)
	f1 := func() error { return nil }
	f2 := func() error { return nil }
	m1 := make(map[int]string)
	m2 := make(map[int]string)
	//fmt.Println(s1 == s2)
	//fmt.Println(f1 == f2)
	//fmt.Println(m1 == m2)
	fmt.Println(s1, s2, f1(), f2(), m1, m2)
}

func main() {
	defineMaps()
}
