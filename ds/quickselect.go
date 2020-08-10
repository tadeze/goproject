package main

import "fmt"

func partition(a []int, l, h int) int {

	i := l - 1
	for j := 0; j < h; j++ {
		if a[j] < a[h] {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[h] = a[h], a[i+1]
	return i + 1
}
func quickSelect(a []int, k, l, h int) int {

	p := partition(a, l, h)
	if p == k-1 {
		return a[p]
	} else if p > k-1 {
		return quickSelect(a, k, l, p-1)
	} else {
		return quickSelect(a, k, p+1, h)
	}

}

func main() {
	fmt.Println(" Return Kth element from unsorted array ")
	k := 3
	a := []int{3, 12, 8, 1, 6, 5}
	fmt.Println(a)
	n := quickSelect(a, k, 0, len(a)-1)
	fmt.Println(n)
}
