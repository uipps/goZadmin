// 斐波那契数列(Fibonacci): 0 1 1 2 3 5 8 13 21
// go run fibnacci.go -n 20
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\suanfa\fibnacci.go -n 20
package main

import (
	"flag"
	"fmt"
)

var (
	argN int
)

func init() {
	flag.IntVar(&argN, "n", 0, "Usage: 10 100 13")
}

func main() {
	flag.Parse()

	for i := 0; i < argN; i++ {
		fmt.Printf("%d\n", fibonaci(i))
	}
}

func fibonaci(i int) int {
	if i <= 1 {
		return i
	}
	return fibonaci(i-1) + fibonaci(i-2)
}
