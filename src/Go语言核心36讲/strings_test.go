package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBufferExample(t *testing.T) {
	var bf bytes.Buffer
	contents := "Simple byte buffer for marshaling data."
	fmt.Printf("Writing contents %q ...\n", contents)
	bf.WriteString(contents)
	fmt.Printf("The length of buffer: %d\n", bf.Len())
	fmt.Printf("The capacity of buffer: %d\n", bf.Cap())
	p1 := make([]byte, 7)
	n, _ := bf.Read(p1)
	fmt.Printf("%d bytes were read. (call Read)\n", n)
	fmt.Printf("The length of buffer: %d\n", bf.Len())
	fmt.Printf("The capacity of buffer: %d\n", bf.Cap())
}
