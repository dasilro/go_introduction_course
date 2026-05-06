package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	var myIntVar int
	myIntVar = -12
	fmt.Println("My variable is: ", myIntVar)

	var myUintVar uint
	myUintVar = 12
	fmt.Println("My variable is: ", myUintVar)

	var myStringVar string
	myStringVar = "My string variable"
	fmt.Println("My variable is: ", myStringVar)

	var myBooleanVar bool
	myBooleanVar = true
	fmt.Println("My variable is: ", myBooleanVar)

	fmt.Println("My variable address is: ", &myStringVar)

	myIntVar2 := 12
	fmt.Println("My variable is: ", myIntVar2)

	myStringVar2 := "My second string variable"
	fmt.Println("My variable is: ", myStringVar2)

	fmt.Println()

	const myFirstConst = "a12"
	fmt.Println("My constant is: ", myFirstConst)

	const myIntConst int = 12
	fmt.Println("My constant is: ", myIntConst)

	const myStringConst string = "My string constant"
	fmt.Println("My constant is: ", myStringConst)

	fmt.Println()

	var my8BitsIntVar int8
	fmt.Printf("Int default value is: %d\n", my8BitsIntVar)

	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my8BitsIntVar, unsafe.Sizeof(my8BitsIntVar), unsafe.Sizeof(my8BitsIntVar)*8)

	var my16BitsIntVar int16
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my16BitsIntVar, unsafe.Sizeof(my16BitsIntVar), unsafe.Sizeof(my16BitsIntVar)*8)

	var my32BitsIntVar int32
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my32BitsIntVar, unsafe.Sizeof(my32BitsIntVar), unsafe.Sizeof(my32BitsIntVar)*8)

	var my64BitsIntVar int64
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my64BitsIntVar, unsafe.Sizeof(my64BitsIntVar), unsafe.Sizeof(my64BitsIntVar)*8)

	var myIntVar3 int
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myIntVar3, unsafe.Sizeof(myIntVar3), unsafe.Sizeof(myIntVar3)*8)

	fmt.Println()

	var my8BitsUIntVar uint8
	fmt.Printf("Uint default value is: %d\n", my8BitsUIntVar)

	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my8BitsUIntVar, unsafe.Sizeof(my8BitsUIntVar), unsafe.Sizeof(my8BitsUIntVar)*8)

	var my16BitsUIntVar uint16
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my16BitsUIntVar, unsafe.Sizeof(my16BitsUIntVar), unsafe.Sizeof(my16BitsUIntVar)*8)

	var my32BitsUIntVar uint32
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my32BitsUIntVar, unsafe.Sizeof(my32BitsUIntVar), unsafe.Sizeof(my32BitsUIntVar)*8)

	var my64BitsUIntVar uint64
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my64BitsUIntVar, unsafe.Sizeof(my64BitsUIntVar), unsafe.Sizeof(my64BitsUIntVar)*8)

	var myUintVar3 uint
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myUintVar3, unsafe.Sizeof(myUintVar3), unsafe.Sizeof(myUintVar3)*8)

	fmt.Println()

	var myFloat32Var float32
	fmt.Printf("Float32 default value is: %f\n", myFloat32Var)

	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)

	var myFloat64Var float64
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myFloat64Var, unsafe.Sizeof(myFloat64Var), unsafe.Sizeof(myFloat64Var)*8)

	fmt.Println()

	var myStringVar3 string
	fmt.Printf("String default value is: %s\n", myStringVar3)

	myStringVar5 := `My string variable in golang
	with multiple
	lines`
	fmt.Printf("String default value is: %s\n", myStringVar5)

	{

		fmt.Println()
		floatVar := 33.11
		fmt.Printf("type: %T, value: %f\n", floatVar, floatVar)
		floatStrVar := fmt.Sprintf("%.2f", floatVar)
		fmt.Printf("type: %T, value: %s\n", floatStrVar, floatStrVar)

		intVar := 22
		fmt.Printf("type: %T, value: %d\n", intVar, intVar)
		intStrVar := fmt.Sprintf("%d", intVar)
		fmt.Printf("type: %T, value: %s\n", intStrVar, intStrVar)

		intVal1, err := strconv.ParseInt("1333", 0, 64)
		fmt.Println(err)
		fmt.Printf("type: %T, value: %d\n", intVal1, intVal1)

		intVal2, err := strconv.ParseInt("aa122", 0, 64)
		fmt.Println(err)
		fmt.Printf("type: %T, value: %d\n", intVal2, intVal2)

		floatVar1, _ := strconv.ParseFloat("-11.22", 64)
		fmt.Printf("type: %T, value: %f\n", floatVar1, floatVar1)
	}
}
