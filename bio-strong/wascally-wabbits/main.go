package main

import (
	"flag"
	"fmt"
)

var file = flag.String("file", "", "path to data set")

func getRabbitCountAfter(n, k int) int {
	p1, p2 := 0, 1

	for i := 2; i < n; i++ {
		p1, p2 = p2*k, p1+p2
	}

	return p1 + p2
}

func main() {
	fmt.Println(getRabbitCountAfter(34, 4))
}
