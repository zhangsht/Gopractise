package main

import (
	"strings"
	"fmt"
)

func main() {
	/* // int
	a := 'a'
	fmt.Printf("%d %[1]c %[1]q\n", a) */
	
	/* // float
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) */

	/* // complex
	var x complex128 = complex(1, 2)
	var y complex128 = complex(2, 4)
	fmt.Println(x*y) */

	/* var company = "hongshi"
	fmt.Println("goodbye " + company[3:]) */

	// utf8
	for i, r := range "hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	fmt.Println(string(1234567))

}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3] + "," + s[n-3:])
}
