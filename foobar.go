package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(Foobar(i))
	}
}

// Foobar returns a string for a int value as foobar
func Foobar(value int) string {
	var returnValue string
	returnValue += isDivisibleBy(value, 3, "foo")
	returnValue += isDivisibleBy(value, 5, "bar")
	if returnValue != "" {
		return returnValue
	}
	return strconv.Itoa(value)
}

func isDivisibleBy(value int, divisor int, special string) string {
	if value%divisor == 0 {
		return special
	}
	return ""
}
