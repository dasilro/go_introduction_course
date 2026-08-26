package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	log.Println("test")

	//log.Fatal("my fatal error")

	//log.Panic("my panic error")

	log.Print("Entra")

	date := time.Now()
	file, err := os.Create(fmt.Sprintf("%d.log", date.UnixNano()))

	if err != nil {
		panic(err)
	}

	l := log.New(file, "success: ", log.LstdFlags|log.Lshortfile)
	l.Println("test 1")
	l.Println("test 2")
	l.Println("test 3")
	l.Println("test 4")

	l.SetPrefix("error: ")
	l.Println("test 5")
	l.Println("test 6")
	l.Println("test 7")
	l.Println("test 8")

}
