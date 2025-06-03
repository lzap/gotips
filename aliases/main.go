package main

type MyInt int

type MyAlias = int

func (MyAlias) Do() {} // HLmethod

func main() {
	var x any = 1

	a := x.(MyAlias)
	i := x.(MyInt) // HLassert

	println(a)
	println(i)
}
