package main

import (
	"fmt"
)




func main() {
	m := newMatrix(3, 3, 2)
	fmt.Println(m.rows)
	// Add two matrix
	m2 := newMatrix(3, 3, 3)
	m = m.add(m2)
	fmt.Println(m.rows)
	// Multiplication of two matrix
	m3 := newMatrix(3, 6, 3)
	fmt.Println("Multiplication two matrix")
	fmt.Println(m.dot(m3).rows)
	// matrix operations
	// Transpose
	fmt.Println(m3.transpose())
}
