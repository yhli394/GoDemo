package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		//goroutine是一个线程，go func(x,y,z)表明开启一个新线程来执行函数func()
		go func() {
			fmt.Println(i)
		}()
	}
}
