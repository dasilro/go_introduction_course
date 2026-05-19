package function

import "errors"

type Operation int

const (
	SUM Operation = iota
	SUB
	DIV
	MUL
)

func Add(x int, y int) int {
	return x + y
}

func RepeatString(increment int, s string) string {
	var result string
	for i := 0; i < increment; i++ {
		result += s
	}
	return result
}

func Calc(op Operation, x, y float64) (float64, error) {
	switch op {
	case SUM:
		return x + y, nil
	case SUB:
		return x - y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("division by zero")
		}
		return x / y, nil
	case MUL:
		return x * y, nil
	default:
		return 0, errors.New("invalid operation")
	}
}

func Split(v int) (x, y int) {
	x = v * 4 / 9
	y = v - x
	return
}

func MSum(values ...float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum
}

func MOperations(op Operation, values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("At least 1 value required")
	}

	sum := values[0]
	var err error = nil

	for _, v := range values[1:] {
		sum, err = Calc(op, sum, v)
		if err != nil {
			return 0, err
		}
	}
	return sum, err
}
