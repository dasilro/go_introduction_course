package function

func FactoryOperation(op Operation) func(float64, float64) float64 {
	switch op {
	case SUM:
		return sum
	case SUB:
		return sub
	case DIV:
		return div
	case MUL:
		return mul
	}
	return nil
}

func sum(x, y float64) float64 {
	return x + y
}

func sub(x, y float64) float64 {
	return x - y
}

func div(x, y float64) float64 {
	if y == 0 {
		return 0
	}
	return x / y
}

func mul(x, y float64) float64 {
	return x * y
}
