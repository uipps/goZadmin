// 测试主程序，有main包和main函数
package main

import "fmt"

const s string = "abcde"

func main() {
	var str01 string = "haha"
	var int01, int02 int = 2, 3
	const C = 3e20
	fmt.Println(str01)
	fmt.Println(int64(int01), int02, C)
	fmt.Println(s)
	fmt.Println("go,test01")
	fmt.Println("1+1=", 1+1)
	fmt.Println(true)
}
