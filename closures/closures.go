package main

import "fmt"

func intSeq(num int) func() int {
	i := num
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq(3)

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq(9)
	fmt.Println(newInts())
}
