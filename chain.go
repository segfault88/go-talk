package main

import (
	"fmt"
	"time"
)

// START OMIT
func f(left, right chan int) { // type inference, don't need to type left
	left <- 1 + <-right
}

func main() {
	start := time.Now()
	const n = 10000
	leftmost := make(chan int)

	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	go func(c chan int) { c <- 0 }(right)

	fmt.Println(<-leftmost, time.Since(start))
}
