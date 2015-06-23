package main

import "fmt"

// START OMIT
func Generate(ch chan<- int) { // Send the sequence 2, 3, 4, ... to channel 'ch'.
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}
func Filter(in <-chan int, out chan<- int, prime int) {
	for { // copy from in to out, removing those divisible by 'prime'.
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i
		} // Send 'i' to 'out'.
	}
}
func main() {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}

// END OMIT
