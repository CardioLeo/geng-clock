package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	n1 := 1
	data, err := json.Marshal(n1)
	if err != nil {
		log.Fatal(err)
	}

	var n2 interface{}
	if err := json.Unmarshal(data, &n2); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("n1 is %T, n2 is %T\n", n1, n2)
}

// output:
// n1 is int, n2 is float64

// This is the expected output, and the output I got;
// the code tells me the types of the two different vars
