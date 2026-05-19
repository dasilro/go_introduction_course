package main

import (
	"fmt"

	"github.com/dasilro/go_introduction_course/tarea4b/matrix"
)

func main() {
	m, err := matrix.New([]float64{2, 7, 8}, []float64{4, 4, 7}, []float64{5, 6, 1})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	m.Print()
}
