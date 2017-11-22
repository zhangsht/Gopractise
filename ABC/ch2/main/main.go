package main

import (
	"log"
	"os"
	"fmt"
	"strings"
	"flag"
	"../tempconv"
)

var n = flag.Bool("n", false, "")
var sep = flag.String("s", " ", "")

var pc [256]byte
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	count := 0
	var i uint
	for ; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

var cwd string

func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os Getwd failed: %v", err)
	}
}

func main()  {
	// pointer
	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil)

	// flag
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

	//new and var won't decide the position in memeory

	// import other package
	fmt.Println(tempconv.AbsoluteZeroC)

	// scanf
	var a int
	var f float64
	fmt.Scanf("%d%f", &a, &f)
	fmt.Println(a, f)

	// init
	for i, v := range pc {
		fmt.Printf("num %d has %d 1 bits\n", i, v)
	}
	fmt.Println(PopCount(64))

	// global var
	fmt.Println(cwd)
}