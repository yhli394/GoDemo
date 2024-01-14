package main

import "fmt"

/*func switch1() {
	value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch 1 + 3 {
	case value1[0], value1[1]:
		fmt.Println("0 or 1")
	case value1[2], value1[3]:
		fmt.Println("2 or 3")
	case value1[4], value1[5], value1[6]:
		fmt.Println("4 or 5 or 6")
	}
}*/

func switch2() {
	value2 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value2[4] {
	case 0, 1:
		fmt.Println("0 or 1")
	case 2, 3:
		fmt.Println("2 or 3")
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	}
}

//case表达式中有值相同的常量子表达式，编译不会通过
/*func switch3() {
	value3 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value3[4] {
	case 0, 1, 2:
		fmt.Println("0 or 1")
	case 2, 3:
		fmt.Println("2 or 3")
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	}
}*/
