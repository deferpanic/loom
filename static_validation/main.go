package main

import (
	"fmt"
)

func blah(stuff string) bool {
	return true
}

func main() {
	fmt.Println("hi")

	fmt.Println(blah("stuff"))
	fmt.Println(blah("otherstuff"))

}
