package main

//package main

import "fmt"

func BinarySearch(n int, arr []int) bool {
	l := 0
	u := len(arr) - 1
	flag := false
	for l <= u {
		m := (l + u) / 2
		if n == arr[m] {
			flag = true
			break
		} else if n < arr[m] {
			u = m - 1
		} else {
			l = m + 1
		}
	}
	return flag
}

/* // Binary search approach
func FindDuplicates(arr1 []int, arr2 []int) []int {
duplicates := []int{}
	for i := 0; i < len(arr1); i++ {
		// do binary search on the large array.
		if BinarySearch(arr1[i], arr2) {
			duplicates = append(duplicates, arr1[i])
		}
	}
	return duplicates
}
*/
// The second approach using o(n)
func FindDuplicates(arr1 []int, arr2 []int) []int {
	duplicates := []int{}
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] == arr2[j] {
			duplicates = append(duplicates, arr1[i])
			i++
			j++
		} else if arr1[i] < arr2[j] {
			i++
		} else {
			j++
		}

	}

	return duplicates
}

func main() {
	arr1 := []int{1, 2, 3, 5, 6, 7}
	arr2 := []int{3, 6, 7, 8, 20}

	fmt.Println(FindDuplicates(arr1, arr2))
}
