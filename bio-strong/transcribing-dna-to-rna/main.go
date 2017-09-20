package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var file = flag.String("file", "./dataset.txt", "Data set file path")

func transcribeDNAToRNA(rd *bufio.Reader) {
	s, err := rd.ReadString('\n')

	for err == nil {
		fmt.Println(strings.Replace(s, "T", "U", -1))
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
	transcribeDNAToRNA(r)
}
