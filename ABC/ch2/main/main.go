package main

import (
	"fmt"
	// "strings"
	"flag"
	// "fmt"
	"../tempconv"
)

var n = flag.Bool("n", false, "")
var sep = flag.String("s", " ", "")

func main()  {
	//// pointer
	// var x, y int
	// fmt.Println(&x == &x, &x == &y, &x == nil)

	// // flag
	// flag.Parse()
	// fmt.Print(strings.Join(flag.Args(), *sep))
	// if !*n {
	// 	fmt.Println()
	// }

	// new and var won't decide the position in memeory

	fmt.Println(tempconv.AbsoluteZeroC)
}