// int8 int32 int int64
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\study\datatype.go -m 7 -n 10

package main

import (
	"fmt"
)

func main() {
	var a0 int8 = -128
	//var au uint8 = 2147483647
	var b0 int16 = 32767
	//var bu uint16 = 2147483647
	var c0 int32 = 2147483647
	//var cu uint32 = 2147483647
	var d0 int64 = 9223372036854775807 // 9223372036854775808
	//var du uint32 = 2147483647
	var e0 int = 9223372036854775806
	//var eu uint = 2147483647

	var f0 byte = 255        // uint8
	var g0 rune = 2147483647 // int32

	fmt.Println("a0:", a0)
	fmt.Println("a0-1:", a0-1)
	fmt.Println("b0:", b0)
	fmt.Println("b0+1:", b0+1)
	fmt.Println("c0:", c0)
	fmt.Println("c0+1:", c0+1)
	fmt.Println("d0:", d0)
	fmt.Println("d0+1:", d0+1)
	fmt.Println("e0:", e0)
	fmt.Println("e0+1:", e0+1)
	fmt.Println("f0:", f0)
	fmt.Println("f0+1:", f0+1)
	fmt.Println("g0:", g0)
	fmt.Println("g0+1:", g0+1)
}
