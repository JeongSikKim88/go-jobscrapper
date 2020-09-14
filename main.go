package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}

	// you have two way for receving two chan msg

	// #1
	// result := <-c
	// fmt.Println(result)

	// #2
	fmt.Println("Waiting for a message")
	resultOne := <-c
	resultTwo := <-c
	fmt.Println("Received this message : ", resultOne)
	fmt.Println("Received second message : ", resultTwo)
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + " is Sexy"
}
