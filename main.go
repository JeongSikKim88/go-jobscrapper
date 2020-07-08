package main

import (
	"fmt"
	"time"
)

func main() {
	go sexyCount("nico")
	sexyCount("jack")
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
