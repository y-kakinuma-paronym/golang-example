package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

	hoge := make(chan string)
	end := make(chan int)
	go func() {
		for foo := range hoge {
			fmt.Println(foo)
		}
		end <- 1
	}()
	time.Sleep(time.Second * 2)
	hoge <- "one"
	time.Sleep(time.Second * 2)
	hoge <- "two"
	time.Sleep(time.Second * 2)
	fmt.Println("close")
	close(hoge)

	<-end
	fmt.Println("end")
}
