package main

import "fmt"

func getRabbitCountAfter(n, k, d int) int {
	p1, p2 := 0, 1
	dc := map[int]int{
		d: 1,
	}

	for i := 2; i < n; i++ {
		p1, p2 = p2*k, p1+p2
		p2 -= dc[i]
		dc[i+d] = p1
	}

	return p1 + p2
}

func main() {
	fmt.Println(getRabbitCountAfter(84, 1, 17))
}
