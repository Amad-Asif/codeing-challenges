package main

import (
	"amadasif/coding-test-examples/examples"
	"fmt"
)

func main() {
	bribes, isChaos := examples.MinimumBribes([]int32{1, 2, 5, 3, 4, 7, 8, 6})
	if isChaos {
		fmt.Println("Too chaotic")
	} else {
		fmt.Println(bribes)
	}
}
