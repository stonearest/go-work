package main

import (
	"fmt"
	"strconv"

	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello"))
	fmt.Println(reverse.String("Hello"), Int(24601))
}

// Int returns the decimal reversal of the integer i.
func Int(i int) int {
	i, _ = strconv.Atoi(reverse.String(strconv.Itoa(i)))
	return i
}
