package main

import (
	"fmt"
	"time"
)

func main() {
	// timeout := 3
	// first replacement version follows below
	var timeout2 time.Duration = 180
	// second replacement version next
	const timeout3 = 180
	// third replacement version
        fmt.Printf("before ")
        time.Sleep(timeout2 * time.Millisecond)
        fmt.Println("after")
	fmt.Printf("before ")
	time.Sleep(time.Duration(timeout3) * time.Millisecond)
	fmt.Println("after")
}

// output:
// # command-line-arguments
// ./puzz-6-there-yet.go:11:13: invalid operation: timeout * time.Millisecond (mismatched types int and time.Duration)

