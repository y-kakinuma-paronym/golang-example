package main

import (
	"fmt"
	"time"
)

func main() {
	a := 100
	go func() {
		a += 50
	}()
	go func() {
		a -= 50
	}()

	time.Sleep(time.Second * 2)
	fmt.Println(a)
}
