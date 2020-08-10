package main

import "fmt"

func reverse(arr []int, i int, j int) {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func main() {
	// Todo three reverse of the array
	a := []int{1, 2, 3, 4, 5, 6, 7}
	k, n := 4, len(a)-1

	reverse(a, n-k+1, n)
	reverse(a, 0, n-k)
	reverse(a, 0, n)
	fmt.Println(a)

}
