package main

import "fmt"

func quickSort(a []int, l int, h int) {
	if l < h {
		p := partition(a, l, h)
		quickSort(a, l, p-1)
		quickSort(a, p+1, h)
	}

}

func partition(a []int, l int, h int) int {

	p := a[h]
	i := l - 1
	for j := l; j < h; j++ {
		if a[j] < p {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[h] = a[h], a[i+1]
	return i + 1
}

func main() {

	a := []int{2, 13, 8, 4, 10, 7}
	fmt.Println("Sorted array")
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
