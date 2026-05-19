package main

import "fmt"

func main() {

	fmt.Println("Ingrese numeros (finalice con '0'):")

	arr := []string{}
	char := rune(1)
	for char != '0' {
		fmt.Scanf("%c\n", &char)
		arr = append(arr, string(char))
	}
	fmt.Println("Los valores del array son: ", arr)
}
