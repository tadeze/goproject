//package Matrix
//package matrix
package main

import (
	"fmt"
	"math"
)

type Matrix struct {
	nRows     int
	nCols     int
	rows      [][]float64
	initValue int
}

func newMatrix(nrows int, ncols int, initValue float64) *Matrix {
	m := new(Matrix)
	m.rows = make([][]float64, nrows)
	for i := range m.rows {
		//m.rows[i] = make([]float64, ncols)
		m.rows[i] = make([]float64, ncols)

	}
	for i := 0; i < nrows; i++ {
		for j := 0; j < ncols; j++ {
			m.rows[i][j] = initValue
		}
	}
	m.nRows = nrows
	m.nCols = ncols

	return m

}
func (m *Matrix) add(m2 *Matrix) *Matrix {
	if m.nCols != m2.nCols || m.nRows != m2.nRows {
		panic("Error:The two matrix don't have the same dimension")
	}
	mat := newMatrix(m.nRows, m.nCols, 0)
	for i := 0; i < m.nRows; i++ {
		for j := 0; j < m.nCols; j++ {
			mat.rows[i][j] = m.rows[i][j] + m2.rows[i][j]
		}
	}
	return mat
}

func (m *Matrix) mult(m2 *Matrix) *Matrix {
	if m.nCols != m2.nCols || m.nRows != m2.nRows {
		panic("Error:The two matrix don't have the same dimension")
	}
	mat := newMatrix(m.nRows, m.nCols, 0)
	for i := 0; i < m.nRows; i++ {
		for j := 0; j < m.nCols; j++ {
			mat.rows[i][j] = m.rows[i][j] * m2.rows[i][j]
		}
	}
	return mat
}
func (m *Matrix) dot(m1 *Matrix) *Matrix {
	if m.nCols != m1.nRows {
		panic("Error:The two matrix don't have the same dimension")
	}
	// c = a[nxd]*b[dxm]  c = [nxm]
	mat := newMatrix(m.nRows, m1.nCols, 0)

	for i := 0; i < m.nRows; i++ {
		for j := 0; j < m1.nCols; j++ {
			for k := 0; k < m.nCols; k++ {
				mat.rows[i][j] += m.rows[i][k] * m1.rows[k][j]
			}
		}
	}
	return mat

}

func (m *Matrix) transpose() *Matrix {
	// transpose matrix m
	t := newMatrix(m.nCols, m.nRows, 0)
	for i := 0; i < m.nRows; i++ {
		for j := 0; j < m.nCols; j++ {
			t.rows[j][i] = m.rows[i][j]
		}
	}
	return t
}
func (m *Matrix) skew() bool {
	// Matrix is skew if A = -A^T
	mT := m.transpose()
	for i := 0; i < m.nRows; i++ {
		for j := 0; j < m.nCols; j++ {
			if m.rows[i][j] != -1*mT.rows[i][j] {
				return false
			}
		}
	}
	return true
}

func (m *Matrix) det() (float64, error) {
	// return determinants
	if m.nRows != m.nCols {
		return 0, fmt.Errorf("Dimension mismatch")
	}
	return _det(m.rows), nil
}
func _det(rows [][]float64) float64 {
	d := 0.0
	n := len(rows)
	if n == 2 {
		return rows[0][0]*rows[1][1] - rows[1][0]*rows[0][1]
	} else {

		for p := 0; p < n; p++ {
			//temp_mat := make([][]float64, n-1) // {}
			tempMat := [][]float64{}
			for i := 1; i < n; i++ {
				//temp_row := make([]float64, n-1) //
				tempRow := []float64{}
				for j := 0; j < n; j++ {
					if j != p {
						tempRow = append(tempRow, rows[i][j])
						//	temp_row[i] = rows[i][j]
					}
				}

				if len(tempRow) > 0 {
					tempMat = append(tempMat, tempRow)
				}
			}
			d = d + rows[0][p]*math.Pow(-1, float64(p))*_det(tempMat)
		}
		return d
	}

}

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

	// Determinants implementation
	md := newMatrix(3, 3, 3)
	md.rows = [][]float64{{1, -2.2, 3}, {3, 4, -3}, {0, -5, 14}}
	fmt.Println("Determinants of matrix ")
	d, _ := md.det()
	fmt.Println(d)
}
