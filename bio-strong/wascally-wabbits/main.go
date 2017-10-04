package main

import (
	"fmt"
)

func getRabbitCountAfter(n, k int) int {
	nmRbs, mRbs := 0, 1

	for i := 2; i < n; i++ {
		nmRbs, mRbs = mRbs*k, nmRbs+mRbs
	}

	return nmRbs + mRbs
}

func main() {
	fmt.Println(getRabbitCountAfter(34, 4))
}
