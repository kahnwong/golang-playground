package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	x := float32(3.0)
	fmt.Println(x)

	m := map[string]int{"one": 1, "two": 2}
	fmt.Println(m)

	// struct
	type database struct {
		hostname string
		username string
		password string
	}
	prodDB := database{
		hostname: "localhost",
		username: "root",
		password: "password",
	}
	fmt.Println(prodDB)

	// if-else
	stockCount := 4
	if stockCount > 5 {
		fmt.Println("all good")
	} else if stockCount < 5 {
		fmt.Println("need to restock")
	}

	// for loop
	for i := 0; i < stockCount; i++ {
		fmt.Println(i)
	}
}
