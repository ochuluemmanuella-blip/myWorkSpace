package main

import (
	"fmt"
)

func ma() {
	// i := 1
	// for i < 10 {
	// 	i = i + 1
	// 	fmt.Println(i)
	// }
	for k := 0; k < 3; k++ {
		fmt.Println(k)
	}
	for i := range 3 {
		fmt.Println("range", i)
	}
	for {
		fmt.Println("loop")
		break
	}
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	fmt.Println("hello")
}
