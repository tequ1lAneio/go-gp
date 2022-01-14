package main

type T struct {
	a int
}

func (t T) M1() {
	t.a = 10
}

func (t *T) M2() {
	t.a = 11
}


func main() {
	var t1 T
	println(t1.a)
	t1.M1()
	println(t1.a)
	t1.M2()
	println(t1.a)

	var t2 = &T{}
	println(t2.a)
	t2.M1()
	println(t2.a)
	t2.M2()
	println(t2.a)
}
