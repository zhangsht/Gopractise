// learn go ABC_ch1
package main

import (
	"sync"
	"log"
	"io/ioutil"
	"time"
	"strings"
	"io"
	"net/http"
	"fmt"
	"os"
	"bufio"
)

func main() {
	var a int = 1
	fmt.Println("hello, World!", a)

	// closure
	nextInt := intSeq()
	for a < 10 {
		fmt.Println(nextInt())
		a++
	}

	// goroutine
	testGoroutine("direct")

	go testGoroutine("goroutine")

	go func (msg string)  {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println(input, " done")

	// channel
	messages := make(chan string)
	go func ()  {
		messages <- "ping"
	}()

	msg := <- messages
	fmt.Println(msg)

	// os
	for i, arg := range os.Args[1:] {
		// fmt.Println(os.Args[i])
		// fmt.Printf("type of arg is %T\n", os.Args[i])
		fmt.Printf("index is: %d, arg is: %s\n", i, arg)
	}
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
	fmt.Println("program file name is: ", os.Args[0])

	bufio
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "open file error: %v\n", err)
				continue
			}

			countLines(f, counts)
			f.Close()
		}
	}

	// map
	counts := make(map[string]int)
	for _, fileame := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "readFile error: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

	//net
	for _, url := range os.Args[1:] {
		urlPrefix := "http://www."
		if !strings.HasPrefix(url,urlPrefix) {
			url = urlPrefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch error： %v\n", err)
			os.Exit(1)
		} else {
			fmt.Printf("http resp code：%s\n", resp.Status)
		}

		f, err := os.Create("./writeFile.txt")
		check(err)

		defer f.Close()
		// b, err := ioutil.ReadAll(resp.Body)
		_, err = io.Copy(f, resp.Body)
		f.Sync()
	
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch errot: %v\n", err)
			os.Exit(1)
		}
		// write file
		f, err := os.Create("./writeFile.txt")
		check(err)

		defer f.Close()
		io.Copy(f, resp.Body)
		f.Sync()
		f.Write(b)
		f.Sync()
	}

	// goroutine
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	// web service
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

	// mutex & lock
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
var mu sync.Mutex
var visitedCount int

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	visitedCount++;
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	fmt.Fprintf(w, "count: %d\n", visitedCount)
	mu.Unlock()
}

// closure
func intSeq() func() int {
	i := 0
	return func () int {
		i += 1
		return i
	}
}

// goroutine
func testGoroutine(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, " : ", i)
	}
}

// map
func countLines(f *os.File, counts map[string]int)  {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

// panic
func check(e error)  {
	if e != nil {
		panic(e)
	}
}

// net
func fetch(url string, ch chan<- string)  {
	start := time.Now()
	urlPrefix := "http://www."
	if !strings.HasPrefix(url,urlPrefix) {
		url = urlPrefix + url
	}
	resp, err := http.Get(url)
	resp, err = http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("read: %s, error: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\t%7d\t%s\n", secs, nbytes, url)
}