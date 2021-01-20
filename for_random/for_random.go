package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		var r = rand.Intn(32)
		fmt.Println(r)
	}
}
