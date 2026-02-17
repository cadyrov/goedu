package main

import "fmt"

func main() {
	x := "ПЖ"

	for i := range x {
		fmt.Println(x[i])
	}

	fmt.Println("___________________________")

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

}
