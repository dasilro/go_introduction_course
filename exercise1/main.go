package main

import "fmt"

func main() {

	array := [5]int{4, 2, 5, 6, 7}

	// realizar la funcionalidad
	for i := range array {
		array[i] = array[i] + 20
	}

	fmt.Println("Los valores del array son: ", array)
}
