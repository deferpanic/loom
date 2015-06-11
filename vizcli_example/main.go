package main

import (
	"fmt"
	"strconv"
)

func blah() {
	fmt.Println(strconv.Itoa(2))
}

func main() {

	fmt.Println(strconv.Itoa(44))

	blah()
}
