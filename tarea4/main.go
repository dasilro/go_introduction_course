package main

import "fmt"

func main() {

	fmt.Println("Ingrese keys (finalice con '0'):")

	codes := map[string]string{
		"10": "notebook",
		"15": "tv",
		"21": "heladera",
		"27": "monitor",
		"35": "camara",
	}

	arr := []string{}
	char := ""
	for char != "0" {
		fmt.Scanf("%s\n", &char)
		if char == "0" {
			break
		}
		var value = codes[char]
		if value == "" {
			arr = append(arr, "No encontrado")
		} else {
			arr = append(arr, string(value))
		}
	}
	fmt.Println("Los valores del array son: ", arr)
}
