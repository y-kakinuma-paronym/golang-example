package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

type runtimeError struct {
	message string
}

func (e *runtimeError) Error() string {
	return "Runtime Error:" + e.message
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "cant't work with it"}
	} else if arg == 0 {
		return -1, &runtimeError{"0 Runtime Error"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{4, 42} {
		r, e := f1(i)
		if e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
		fmt.Println("r:", r)
	}
	for _, i := range []int{4, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println("argError")
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	} else if ae, ok := e.(*runtimeError); ok {
		fmt.Println("runtimeError")
		fmt.Println(ae.message)
	}
}
