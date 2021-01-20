package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var random = rand.Int31()
	if random%2 == 0 {
		fmt.Println(random, "is even")
	} else {
		fmt.Println(random, "is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is odd")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}
