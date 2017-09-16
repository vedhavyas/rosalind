package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var file = flag.String("file", "./dataset.txt", "Data set file path")

func getCount(s string) (a, c, g, t int64) {
	m := make(map[rune]int64)
	for _, r := range s {
		m[r]++
	}

	return m['A'], m['C'], m['G'], m['T']

}

func writeSymbols(rd *bufio.Reader) {
	s, err := rd.ReadString('\n')

	for err == nil {
		a, c, g, t := getCount(s)
		fmt.Println(a, c, g, t)
		s, err = rd.ReadString('\n')
	}
}

func main() {
	flag.Parse()

	fh, err := os.Open(*file)
	if err != nil {
		log.Fatalf("failed to open file %s: %v\n", *file, err)
	}

	defer fh.Close()
	r := bufio.NewReader(fh)
	writeSymbols(r)
}
