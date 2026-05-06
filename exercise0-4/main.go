package main

import "fmt"

func main() {

	license := true
	age := 15

	if license && age >= 15 {
		fmt.Println("Puede seguir avanzando")
	}

	if age <= 15 || !license {
		fmt.Println("No puede seguir circulando")
	}

}
