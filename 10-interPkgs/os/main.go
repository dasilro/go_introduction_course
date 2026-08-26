package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("myFile.txt")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	data := make([]byte, 100)
	c, err := file.Read(data)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("read %d bytes: %q\n", c, data[:c])

	env1 := os.Getenv("HOME")
	fmt.Println(env1)

	os.Setenv("MI_ENV", "mi valor")
	fmt.Println(os.Getenv("MI_ENV"))

	dir, _ := os.Getwd()
	fmt.Println(dir)

	os.Setenv("ENV_EXISTS", "")
	env2 := os.Getenv("ENV_EXISTS")
	env3 := os.Getenv("ENV_DOESNT_EXIST")
	fmt.Println(env2)
	fmt.Println(env3)

	env4, ok4 := os.LookupEnv("ENV_EXISTS")
	fmt.Println(env4, ok4)

	env5, ok5 := os.LookupEnv("ENV_DOESNT_EXIST")
	fmt.Println(env5, ok5)

	os.Setenv("DB_USERNAME", "nahuel")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "27018")
	os.Setenv("DB_NAME", "users")

	// ${var} or $var
	dbUrl := os.ExpandEnv("mongodb://${DB_USERNAME}:${DB_PASSWORD}@$DB_HOST:$DB_PORT/$DB_NAME")
	fmt.Println(dbUrl)
}
