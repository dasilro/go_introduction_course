package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println(os.Getenv("MY_ENV1"))
	fmt.Println(os.Getenv("MY_ENV2"))
	fmt.Println(os.Getenv("MY_ENV3"))

	if err := godotenv.Load("otherFolder/.var"); err != nil {
		fmt.Println(err)
	}

	fmt.Println(os.Getenv("MY_ENV1"))
	fmt.Println(os.Getenv("MY_ENV2"))
	fmt.Println(os.Getenv("MY_ENV3"))

	myEnv, err := godotenv.Read("otherFolder/.var")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myEnv)

	// It loads the .env by default
	if err := godotenv.Overload(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(os.Getenv("MY_ENV1"))
	fmt.Println(os.Getenv("MY_ENV2"))
	fmt.Println(os.Getenv("MY_ENV3"))

}
