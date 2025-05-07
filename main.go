package main

import (
	"fmt"
	controlstructre "my-app/controlstructure"
	plays "my-app/plays"
)

/*
Pass by reference:
string,slice,map,channel,pointers

pass by value:
array
*/

func init() {
	fmt.Println("this func runs before main")
}

func main() {
	fmt.Println("Hello World")

	//arrays are passed by value, thus elements are copied
	var a [5]int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a[i])
	}
	j := 0
	for {
		if j > 4 {
			break
		}
		fmt.Printf("%v ", a[j])
		j += 1
	}
	var b [5]int = a //pass by value
	fmt.Println()
	b[0] = 10
	j = 0
	for {
		if j > 4 {
			break
		}
		fmt.Printf("%v ", b[j])
		j += 1
	}
	fmt.Println()
	j = 0
	for {
		if j > 4 {
			break
		}
		fmt.Printf("%v ", a[j])
		j += 1
	}
	fmt.Println()

	//string are immutable
	//string slice point to the same array of character
	//string are passed by reference, thus they aren't copied
	var s string = "Rishav"
	s1 := s[0:3]
	s2 := s[2:4]
	r := s[:5] + "b" //c
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(r)
	fmt.Println(&s)
	fmt.Println(&r)

	//slices are passed by reference
	//slices have variable length,backed by some array,they are copied
	//slice is descriptor which has {pointer,length,capacity} similar to string descriptor {pointer,length}

	var x []int //nil
	var y = []int{1, 2}

	x = append(x, 1)
	y = append(y, 3)

	x = y

	t := []byte("string")
	fmt.Println(len(t), t)
	fmt.Println(t[:2])
	fmt.Println(t[2:])

	// map m[[]bookings]string // slices cannot be used as map key because it's not comparable(==) unlike arrays

	//maps are dictionaries: indexed by key,returing a value
	//maps are passed by reference
	//type used for the key must have == and != defined(not slices,maps, or funcs)
	//map also is a type of descriptor

	var m map[string]int      //nil,no storage
	p := make(map[string]int) //non-nil but empty

	a1 := p["the"] //returns 0
	b1 := m["the"] //same thing
	//m["and"] = 1   //panic - nil map
	m = p      //pass by ref
	m["and"]++ // Ok, same map as p now
	c1 := p["and"]
	fmt.Println(a1, " ", b1, " ", c1)

	b2, ok := p["the"]
	fmt.Println(b2, " ", ok)
	p["the"]++
	c2, ok := p["the"]
	fmt.Println(c2, " ", ok)

	if w, ok := p["the"]; ok {
		fmt.Println(w)
	}

	delete(m, "the")
	w2, ok := p["the"]
	fmt.Println(w2, " ", ok)

	var rk = []int{1, 2, 3, 4, 5}
	rk2 := []int{1, 2, 3}
	fmt.Print(rk)
	fmt.Print(rk2)

	var slice = []string{"I", "am", "Rishav", "I", "live", "in", "Pune", "Pune", "is", "a", "metro", "city", "Rishav", "loves", "go", "Rishav"}
	plays.FindLengthUniqueWords(slice)

	controlstructre.Display()
}
