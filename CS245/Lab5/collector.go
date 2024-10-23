package main

import (
	"fmt"
	"os"
	"strconv"
	"math/rand"
	"time"
)

func collect(c int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < c ; i++ {
		r := r1.Intn(100)
		if (i != r) {
			r, i = i, r
		}
	}
	return r
}

func main() {
	n, err := strconv.Atoi(os.Args[0])
	if err == nil {
		fmt.Printf("%d\n", n)
	} else {
		fmt.Printf("ERR: %v\n", err)
	}
	trials := 1000000
	sum := 0
	for i := 0 ; i < trials; i++ {
		sum += collect(n)
	}
	fmt.Println(sum/trials)
}
