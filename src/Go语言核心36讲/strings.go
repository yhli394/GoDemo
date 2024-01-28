package main

import (
	"bytes"
	"fmt"
	"strings"
)

func Func1(value string) {
	var sb strings.Builder
	var a string = "hello"
	sb.WriteString(a)
	sb.WriteString(value)
	fmt.Println(sb.String())
}

func BufferExample(str string) {
	var bf bytes.Buffer
	contents := "Simple byte buffer for marshaling data."
	fmt.Printf("Writing contents %q ...\n", contents)
	bf.WriteString(contents)
	fmt.Printf("The length of buffer: %d\n", bf.Len())
	fmt.Printf("The capacity of buffer: %d\n", bf.Cap())
}
