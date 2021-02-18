package main

import (
	"fmt"
	"time"
)

func main() {
	var numShard int64 = 11
	n := 10000
	s := make([]int, numShard)
	for i := 0; i < n; i++ {
		// m := (time.Now().UnixNano() / 1000) % numShard
		// s[m]++
		fmt.Println(time.Now().UnixNano())
	}

	fmt.Println(s)
}
