package effective_go

import "fmt"

// T hello
type T struct {
	name  string //name of sky
	value int    // age
}

func Format() {
	var slice []string
	for i := range slice {
		fmt.Println(i)
	}
}
