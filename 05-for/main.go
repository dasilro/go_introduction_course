package main

func main() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum++
	}
	println(sum)

	sum = 1
	for sum < 1000 {
		sum++
	}
	println(sum)

	sum = 0

	for {
		if sum > 1000 {
			break
		}
		sum++
	}
	println(sum)

	arr := []int{1, 2, 3, 4, 5}

	for i := range arr {
		println("Index:", i, " - Value:", arr[i])
	}

	println()
	for i, v := range arr {
		println("Index:", i, " - Value:", v)
	}

	println()

	map2 := map[string]float64{
		"A": 12.3,
		"B": 23.1,
		"C": 34,
	}

	println()

	for k, v := range map2 {
		println("Key:", k, " - Value:", v)
	}

	println()

	map3 := map[string][]int{
		"A": nil,
		"B": {2, 34, 1, 2, 4},
		"C": {4, 5, 3, 2, 1},
	}

	for k, value := range map3 {
		println("Key:", k)
		for _, v := range value {
			println("	Value:", v)
		}
	}
}
