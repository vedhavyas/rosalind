package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var filePath = flag.String("file", "", "input data in the file")

func getComplement(s rune) (r rune) {
	switch s {
	case 'A':
		return 'T'
	case 'T':
		return 'A'
	case 'C':
		return 'G'
	case 'G':
		return 'C'
	}

	return r
}

func complementDNAString(dna string) string {
	var c string
	for _, s := range dna {
		c = fmt.Sprintf("%s%s", string(getComplement(s)), c)
	}

	return c
}

func complementDNAStrings(r *bufio.Reader) {
	dna, err := r.ReadString('\n')
	for err == nil {
		dnaC := complementDNAString(dna)
		fmt.Println(dnaC)
		dna, err = r.ReadString('\n')
	}
}

func main() {
	flag.Parse()

	f, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
	}

	complementDNAStrings(bufio.NewReader(f))
}
