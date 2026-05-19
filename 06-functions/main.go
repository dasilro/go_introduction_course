package main

import (
	"fmt"

	"github.com/dasilro/go_introduction_course/06-functions/function"
)

func main() {
	fmt.Println(function.Add(3, 4))

	fmt.Println(function.RepeatString(10, "as"))

	v, error := function.Calc(function.SUM, 3, 6)
	if error != nil {
		fmt.Println(error.Error())
	} else {
		fmt.Println("Value: ", v)
	}

	v, error = function.Calc(function.DIV, 3, 0)
	if error != nil {
		fmt.Println(error.Error())
	} else {
		fmt.Println("Value: ", v)
	}

	x, y := function.Split(20)
	fmt.Println("Value1: ", x, " Value2: ", y)

	v = function.MSum(23, 12, 32, 12, 3, 4, 5)
	fmt.Println("Multi sum: ", v)

	v, error = function.MOperations(function.SUM, 1, 2, 3, 4, 0, 9)
	if error != nil {
		fmt.Println(error.Error())
	} else {
		fmt.Println("Multi operations: ", v)
	}

	fn := function.FactoryOperation(function.SUM)
	println("Operation result: ", fn(2, 3))

	fn = function.FactoryOperation(function.MUL)
	println("Operation result: ", fn(2, 3))

}
