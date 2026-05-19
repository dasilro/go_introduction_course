package operator

func Div(x, y float64) float64 {
	// defer func() {
	// 	fmt.Println("In my div function defer")
	// }()

	if y <= 0 {
		panic("divisor must be a positive value")
	}
	return x / y
}
