package main

import (
	"fmt"
)

func main () {
	jack := map[string]string{"name":"jack","age":"33"}
	for key, value :=range jack{
		fmt.Println(key,value)
	}
}
