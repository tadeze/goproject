package main

import (
	"fmt"
	"os"
	"strconv"
)

func PrintFactors(n int) int {
	cnt := 0
	for i := 1; i*i < n; i++ {
		if n%i == 0 {
			cnt++
		}
		if i != n/i {
			cnt++
		}

	}
	return cnt

}

func main() {
	fmt.Println("Enter the number to display factors")
	arg := os.Args
	if len(arg) < 2 {
		os.Exit(1)
	}
	n, err := strconv.Atoi(arg[1])
	if err != nil {
		os.Exit(2)
	}
	fmt.Println(PrintFactors(n))
}
