package main

import (
	"fmt"
)

func main() {
	a := 'a'
	fmt.Printf("%d %[1]c %[1]q\n", a)
}
