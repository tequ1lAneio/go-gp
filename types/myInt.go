package main

type MyInt int64

var m int = 34215
var n int32 = 6
var a MyInt = MyInt(m)
var b MyInt = MyInt(n)

type myHeiInt = int64

var x int64 = 434
var z myHeiInt = x
