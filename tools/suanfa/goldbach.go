// ：任一大于2的偶数都可写成两个质数之和
// 	go run goldbach.go -n 30 -o 1
//  go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\suanfa\goldbach.go -n 2880800
package main

import (
	"flag"
	"fmt"
)

var (
	argN    int
	onlyOne int
)

func init() {
	flag.IntVar(&argN, "n", 0, "Usage: 10 100 13")
	flag.IntVar(&onlyOne, "o", 0, "Usage: 0 1")
}

func main() {
	flag.Parse()

	//fmt.Printf( " print only one %d !\n" , onlyOne)

	if argN < 6 {
		fmt.Printf(" %d is error num!", argN)
		return
	}

	// 从6开始，验证各个偶数
	if 0 == onlyOne {
		evenLianXu(argN)
	} else {
		procOne(argN)
	}

	return
}

func procOne(N int) int {
	i := 0
	j := 0
	isFlag := 1

	i = N
	for j = 2; j <= i/2; j++ {
		if 0 == j%2 || (0 == (i-j)%2) {
			continue
		}
		if 1 == PrimeNum(j) && 1 == PrimeNum(i-j) {
			fmt.Printf(" %d = %d + %d \n", i, j, i-j)
			isFlag = 0
			//break
		}
	}
	if isFlag == 1 {
		fmt.Printf(" 找到一个不符合要求的偶数：%d \n", j)
	}

	return isFlag
}

func evenLianXu(N int) int {
	i := 0
	j := 0
	isFlag := 1

	// 从6开始，验证各个偶数
	for i = 6; i <= N; i += 2 {
		isFlag = 1
		for j = 2; j <= i/2; j++ {
			if 0 == j%2 || (0 == (i-j)%2) {
				continue
			}
			if 1 == PrimeNum(j) && 1 == PrimeNum(i-j) {
				fmt.Printf(" %d = %d + %d \n", i, j, i-j)
				isFlag = 0
				break
			}
		}
		if isFlag == 1 {
			fmt.Printf(" 找到一个不符合要求的偶数：%d \n", j)
		}
	}
	return isFlag
}

// 1. 质数（素数）：在大于1的自然数中，除了1和它本身以外不再有其他因数的自然数。2是最小的质数，并且也是质数中唯一的偶数
// 	1.1 j*j < n 使得时间复杂度为O(log2N)
func PrimeNum(n int) int {
	isPrime := 1

	if n < 2 {
		return 0
	} else if n <= 3 {
		return 1
	}

	for j := 2; j*j <= n; j++ { // 4不是质数, 所以用"<="
		//fmt.Printf(" func primeNum, %d \n", j)
		if 0 == n%j {
			isPrime = 0
			break
		}
	}

	return isPrime
}
