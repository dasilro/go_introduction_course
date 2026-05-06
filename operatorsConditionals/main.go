package main

import "fmt"

func main() {

	yearsOld := 32

	fmt.Println("Operators")
	fmt.Println(yearsOld > 30)
	fmt.Println(yearsOld < 32)
	fmt.Println(yearsOld <= 32)
	fmt.Println(yearsOld >= 40)
	fmt.Println(yearsOld == 32)

	fmt.Println()

	fmt.Println("Conditionals")

	// OR
	fmt.Println(yearsOld < 32 || yearsOld == 32)
	fmt.Println(yearsOld < 32 || yearsOld == 33)
	fmt.Println(yearsOld < 40 || yearsOld == 32)

	fmt.Println()

	// AND
	fmt.Println(yearsOld < 32 && yearsOld == 32)
	fmt.Println(yearsOld < 32 && yearsOld == 33)
	fmt.Println(yearsOld < 40 && yearsOld == 32)

	fmt.Println()

	// NOT
	fmt.Println(true)
	fmt.Println(!true)
	fmt.Println(yearsOld < 40)
	fmt.Println(!(yearsOld < 40))

	fmt.Println(yearsOld < 25 && yearsOld == 32 || yearsOld < 40)
	fmt.Println(yearsOld < 25 && (yearsOld == 32 || yearsOld < 40))

	yearsOld = 20

	if yearsOld > 18 {
		fmt.Printf("%d is higher than 18\n", yearsOld)
	}

	boolVal := true

	if boolVal {
		fmt.Println("This is true")
	} else {
		fmt.Println("This is false")
	}

	if value := true; value {
		fmt.Println("is true")
	}

	number := 3

	if number == 1 {
		fmt.Println("one")
	} else if number == 2 {
		fmt.Println("two")
	} else if number == 3 {
		fmt.Println("three")
	}

	switch number {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("undefined number")
	}

	switch number := 1; number {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("undefined number")
	}

	switch {
	case number > 0:
		fmt.Println("positive")
	case number < 0:
		fmt.Println("negative")
	case number == 0:
		fmt.Println("zero")
	}
}
