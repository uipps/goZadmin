// 测试主程序，有main包和main函数
package main

import (
	"flag"
	"fmt"
	"math"
)

const s string = "abcde"

var (
	a int
	b string
	c []float32
)

var mode = flag.String("mode", "", "process mode")

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


	i,j,k := 1,2,"kkk"
	m := 1

	a = 3
	fmt.Println("i + j = ", i + j)
	//goto END
	fmt.Println("a = ", a)
	fmt.Println("m = ", m)
	fmt.Println("k = ", k)
	fmt.Println("go,test01")
	fmt.Println(math.Pi)

	flag.Parse()
	fmt.Println(*mode)



	//END: fmt.Println(33333);


	/*
		for i<=3 {
			fmt.Println(i);
			i=i+1
		}

		for l:=1;l<=3;l++ {
			fmt.Println(l)
		}

		for n:=1;n<=6;n++{
			if 0 == n%2 {
				continue
			}
			fmt.Println(n)
		}


		switch i {
		case 1: println("one")
		case 2: println("two")
		case 3:
			println("three")
		default:
			println("default")
		}
	*/


}
