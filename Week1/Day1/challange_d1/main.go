package main

import "fmt"

func main() {
	i := 21
	j := true
	var k float64 = 123.456
	l := 21
	m := 21
	n := 15
	s := 'Я'

	fmt.Printf("%v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%%\n")
	fmt.Printf("%v\n\n", j)

	fmt.Printf("%b\n", l)
	fmt.Printf("?\n")
	fmt.Printf("%d\n", l)
	fmt.Printf("%o\n", m)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)
	fmt.Printf("%U\n\n", s)

	fmt.Printf("%f\n", k)
	fmt.Printf("%E\n", k)
}