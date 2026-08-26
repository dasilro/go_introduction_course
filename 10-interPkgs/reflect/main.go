package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID        int64    `myLabel:"lb1" myOtherLabel:"lb2"`
	Email     string   `myLabel:"lb2"`
	FirstName string   `myLabel:"lb3"`
	LastName  string   `myLabel:"lb4"`
	Age       *float64 `myLabel:"lb5"`
	Address   Address  `myLabel:"lb6"`
}

type Address struct {
	Country string
	State   string
}

func main() {
	myInt := 5
	Examine(myInt)

	fmt.Println()

	myPnt := &myInt
	Examine(myPnt)

	fmt.Println()

	var age float64 = 32
	u := User{ID: 1, Email: "nlcostamagna@gmail.com", FirstName: "Nahuel", LastName: "Costamagna", Age: &age, Address: Address{Country: "Argentina", State: "San Isidro"}}
	Examine(u)
}

func Examine(data interface{}) {
	t := reflect.TypeOf(data)
	k := t.Kind()

	if k == reflect.Struct {
		v := reflect.ValueOf(data)
		t := reflect.TypeOf(data)

		fmt.Println("Number of fields ", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			kind := v.Field(i).Kind()
			switch kind {
			case reflect.Int, reflect.Int32, reflect.Int64:
				myVar := v.Field(i).Int()
				fmt.Printf("Field: %d type: %T value: %v\n", i, myVar, myVar)
			case reflect.String:
				myVar := v.Field(i).String()
				fmt.Printf("Field: %d type: %T value: %v\n", i, myVar, myVar)
			case reflect.Ptr:
				fmt.Printf("Field: %d value: %v\n", i, v.Field(i))
			case reflect.Struct:

				if v.Field(i).Type() == reflect.TypeOf(Address{}) {
					//Examine(v.Field(i).Interface()) recursive is an option, left the other case to check the casting
					myVar := v.Field(i).Interface().(Address)
					fmt.Println(myVar.Country)
					fmt.Println(myVar.State)
					fmt.Printf("Field: %d value: %v\n", i, myVar)
				} else {
					fmt.Println("Unsupported type ", v.Field(i).Type())
				}
			default:
				fmt.Println("Unsupported type ", v.Field(i).Type())
			}
			c := t.Field(i).Tag
			fmt.Println(c)
			fmt.Println(c.Get("myLabel"))
			fmt.Println(c.Get("myOtherLabel"))
			fmt.Println()
		}
	} else {
		v := reflect.ValueOf(data)
		fmt.Println("Type ", t)
		fmt.Println("Value ", v)
		fmt.Println("Kind ", k)
	}
}
