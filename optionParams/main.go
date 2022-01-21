package main

import "fmt"

type Foo struct {
	name string
	id   int
	age  int

	db interface{}
}

type FooOption func(foo *Foo)

func WithName(name string) FooOption {
	return func(foo *Foo) {
		foo.name = name
	}
}

func WithAge(age int) FooOption {
	return func(foo *Foo) {
		foo.age = age
	}
}

func WithDB(db interface{}) FooOption {
	return func(foo *Foo) {
		foo.db = db
	}
}

func NewFoo(id int, options ...FooOption) *Foo {
	foo := &Foo{
		name: "default",
		id:   id,
		age:  10,
		db:   nil,
	}

	for _, option := range options {
		option(foo)
	}
	return foo
}

func main() {
	foo := NewFoo(1, WithAge(15), WithAge(15), WithName("foo"))
	fmt.Println(foo)
}
