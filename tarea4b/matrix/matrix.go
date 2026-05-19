package matrix

import (
	"errors"
	"fmt"
)

type Matrix struct {
	Values     [][]float64 `json:"values"`
	RowsLength int
	ColsLength int
}

func (m Matrix) Print() {
	for _, row := range m.Values {
		fmt.Println(row)
	}
	fmt.Printf("%d x %d\n", m.RowsLength, m.ColsLength)
	fmt.Printf("Cuadratic: %t", m.RowsLength == m.ColsLength)
}

func New(values ...[]float64) (matrix Matrix, err error) {
	colsLength := 0
	for rowIndex, row := range values {
		if rowIndex == 0 {
			colsLength = len(row)
		} else {
			if len(row) != colsLength {
				err = errors.New("Different columns number")
				break
			}
		}
	}

	if err != nil {
		return matrix, err
	} else {
		return Matrix{Values: values, RowsLength: len(values), ColsLength: colsLength}, err
	}
}
