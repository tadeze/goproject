package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func printPrime(n int) {
	for i := 0; i < n; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	printPrime(13)
	f("direct")
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)

	}("going")
	time.Sleep(time.Second)
	fmt.Println("done")
}
