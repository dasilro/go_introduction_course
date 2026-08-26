package main

import (
	"fmt"
	"time"
)

func main() {
	go myProcess("A")
	go myProcess("B")

	time.Sleep(time.Millisecond * 3000)

	myFirstChannel := make(chan string)
	go myProcessWithChannel("C", myFirstChannel)
	result := <-myFirstChannel

	fmt.Println(result)
	close(myFirstChannel)

	channelD := make(chan string)
	channelE := make(chan string)

	go myProcessWithChannel("D", channelD)
	go myOtherProcessWithChannel("E", channelE)

	resultD := <-channelD
	resultE := <-channelE

	fmt.Println(resultD)
	close(channelD)

	fmt.Println(resultE)
	close(channelE)

}

func myProcess(p string) {
	i := 0
	for i < 15 {
		time.Sleep(time.Millisecond * 150)
		i++
		fmt.Printf("process: %s - num: %d\n", p, i)
	}
}

func myProcessWithChannel(p string, c chan string) {
	i := 0
	for i < 20 {
		time.Sleep(time.Millisecond * 150)
		i++
		fmt.Printf("process: %s - num: %d\n", p, i)
	}
	c <- "ok"
}

func myOtherProcessWithChannel(p string, c chan string) {
	i := 0
	for i < 10 {
		time.Sleep(time.Millisecond * 150)
		i++
		fmt.Printf("process: %s - num: %d\n", p, i)
	}
	c <- "ok2"
}
