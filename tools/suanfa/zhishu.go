// 输出1000（n可变值）以内的所有质数。并输出计算次数进行对比
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\suanfa\zhishu.go -n 1000
package main

import (
	"flag"
	"fmt"

	"github.com/uipps/goZadmin/tools/jinzhiToAny"
)

var (
	argN int
)

func init() {
	flag.IntVar(&argN, "n", 100, "Usage: 10 100 13")
}

func main() {
	flag.Parse()

	fmt.Printf(" 第一种方法primeNum01：\n")
	primeNum01(argN)
	fmt.Printf(" 第二种方法primeNum02：\n")
	primeNum02(argN)

	return
}

// 1.1 质数的定义进行遍历，时间复杂度最大的
func primeNum01(n int) int {
	runCounter := 0
	counter := 0
	i := 0

	for no := 2; no < n; no++ {
		for i = 2; i < no; i++ {
			runCounter++
			if 0 == no%i {
				break
			}
		}
		if no == i {
			counter++
			// 直到最后也没有找出
			fmt.Printf("%8d, %8o, %8X, %20b \n", no, no, no, no)
			//fmt.Printf("Line %4d: %8d, %8o, %8X, %20b \n", counter, no, no, no, no)
			//fmt.Printf("%d\n", no)
		}
	}

	fmt.Printf("运算次数 %d\n", runCounter)
	fmt.Printf("%d 以内的质数个数 %d\n", n, counter)
	return 0
}

// 1.2 除的时候，只需要no开方以下即可，不需要尝试到no-1; 17只需要尝试4以下的，大于4的不需要尝试
func primeNum02(n int) int {
	runCounter := 0
	counter := 0
	i := 0
	isPrime := 1 // 质数
	jinzhiTo3 := ""
	jinzhiTo5 := ""
	jinzhiTo7 := ""

	// 2也是质数
	no := 2
	counter++
	//jinzhiTo7 = jinzhiToAny.DecimalToAny(no, 7)
	fmt.Printf("%8d, \n", no)
	//fmt.Printf("%8d, %8s, %8o, %8X, %20b \n", no, jinzhiTo7, no, no, no)
	//fmt.Printf("%8d, %8o, %8X, %20b \n", no, no, no, no)

	// 大于2的偶数不可能，所以只在奇数中找
	for no = 3; no < n; no += 2 {
		isPrime = 1
		for i = 2; i*i <= no; i++ {
			runCounter++
			if 0 == no%i {
				isPrime = 0 // 合数
				break
			}
		}
		if 1 == isPrime {
			counter++
			// 直到最后也没有找出
			jinzhiTo3 = jinzhiToAny.DecimalToAny(no, 3)
			jinzhiTo5 = jinzhiToAny.DecimalToAny(no, 5)
			jinzhiTo7 = jinzhiToAny.D十进制转换(no, 7)
			fmt.Printf("%8d, %8s, %8s, %8s, %8o, %8X, %20b \n", no, jinzhiTo3, jinzhiTo5, jinzhiTo7, no, no, no)
			//fmt.Printf("Line %4d: %8d, %8o, %8X, %20b \n", counter, no, no, no, no)
			//fmt.Printf("%d\n", no)
		}
	}

	fmt.Printf("运算次数 %d\n", runCounter)
	fmt.Printf("%d 以内的质数个数 %d\n", n, counter)
	return 0
}
