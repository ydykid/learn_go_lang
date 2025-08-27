package main

import "fmt"

func sliceUnbindOrigArray() {
	u := []int{11, 12, 13, 14}
	fmt.Println("arrary:", u)
	s := u[1:3]
	fmt.Println("slice:", s)
	s = append(s, 24)
	fmt.Println("after append 24, arrary:", u)
	fmt.Println("after append 24, slice:", s)

	s = append(s, 25)
	fmt.Println("after append 25, arrary:", u)
	fmt.Println("after append 25, slice:", s)

	s = append(s, 26)
	fmt.Println("after append 26, arrary:", u)
	fmt.Println("after append 26, slice:", s)

	s[0] = 22
	fmt.Println("after reassign slice[0] to 22, arrary:", u)
	fmt.Println("after reassign slice[0] to 22, slice:", len(s), cap(s), s)

}
