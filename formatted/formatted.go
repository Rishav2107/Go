package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	fmt.Printf("%v\n", s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%#v\n", s)

	a := [3]rune{'a', 'b', 'c'}
	fmt.Printf("%v\n", a)
	fmt.Printf("%q\n", a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)

	m := map[string]string{"name": "rishav", "city": "pune"}
	fmt.Printf("%v\n", m)
	fmt.Printf("%T\n", m)
	fmt.Printf("%#v\n", m)

	str := "a string"
	fmt.Printf("%v\n", str)
	fmt.Printf("%T\n", str)
	fmt.Printf("%#v\n", str)

}
