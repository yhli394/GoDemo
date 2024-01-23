package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("start")
	defer func() {
		fmt.Println("enter defer")
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
		fmt.Println("exit defer function")
	}()
	panic(errors.New("wrong"))
	fmt.Println("exit function main.")
}
