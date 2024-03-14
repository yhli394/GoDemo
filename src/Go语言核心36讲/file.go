package main

import (
	"fmt"
	"os"
)

func main() {
	create, err := os.Create("D:\\hello.txt")
	if err != nil {
		fmt.Println("create success, create：[%v]", create)
	}
}
