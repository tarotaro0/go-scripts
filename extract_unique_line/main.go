package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const filename = "./input.txt"

func main() {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("failed to open file", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	d := map[string]bool{}

	for s.Scan() {
		t := s.Text()
		if !d[t] {
			d[t] = true
		}
	}

	if err = s.Err(); err != nil {
		log.Fatal("failed to scan", err)
	}

	for k := range d {
		fmt.Println(k)
	}
}
