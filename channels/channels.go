package main

import "fmt"

func main() {
	messages := make(chan string)
	go func() {
		messages <- "ping"
		messages <- "hoge"
		messages <- "foo"
	}()

	msg := <-messages
	fmt.Println(msg)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
