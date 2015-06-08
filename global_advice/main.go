package main

import (
	"fmt"
)

func getStuff(i int) {
	fmt.Println(i)
}

func main() {

	for i := 0; i < 10; i++ {
		getStuff(i)
	}
}
