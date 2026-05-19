package main

import (
	"errors"
	"fmt"

	"github.com/dasilro/go_introduction_course/09-errors/operator"
)

func main() {
	var err error
	err = errors.New("my new error")
	fmt.Println(err)
	fmt.Println(err.Error())

	err2 := fmt.Errorf("my formatted err, string: %s, number: %d", "my string", 23)
	fmt.Println(err2)
	fmt.Println(err2.Error())

	defer func() {
		fmt.Println("In main defer")
		r := recover()
		if r != nil {
			fmt.Println("There is no result")
			fmt.Println("Recovered in ", r)
		}
	}()

	fmt.Println(err2)

	z := operator.Div(4, 0)

	fmt.Println("Result")
	fmt.Printf("z is : %f\n", z)

}
