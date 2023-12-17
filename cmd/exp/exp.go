package main

import (
	"fmt"
	"strings"
)

func main() {
	Demo()
	Demo(1)
	Demo(1, 2, 3)
}

func Demo(numbers ...int) {
	for _, number := range numbers {
		fmt.Print(number, " ")
	}
	fmt.Println()
}

func Join(vals ...string) string {
	var sb strings.Builder
	for i, s := range vals {
		sb.WriteString(s)
		if i < len(vals)-1 {
			sb.WriteString(", ")
		}
	}
	return sb.String()
}
