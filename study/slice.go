//
package main

import "fmt"
import "os"

func main() {

	arr := [3]int{}
	fmt.Println(arr)

	var slice01 []int
	fmt.Println(slice01)

	slice02 := []int{1, 2, 3}
	fmt.Println(slice02)
	os.Exit(0)
}
