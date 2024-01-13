package main

import "fmt"

func main() {
	var i = 6
	//&运算符读作and运算符
	//指针是保存地址的变量
	fmt.Println(&i)
	f(&i)
	g(i)
}

//p的值就是i的地址
func f(p *int) {
	fmt.Println(p)
	//*：获取指针指向变量的值
	fmt.Println(*p)
	//*p就代表了变量i
	*p = 26
}

func g(i int) {
	fmt.Println(i)
}

//指针应用场景：①交换两个数；②函数返回多个值：传入的参数实际上是需要保存带回的结果的变量
