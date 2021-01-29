/**

go run tianshu.go -b 8 -n 10
go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\tianshu\tianshu.go


填数游戏，如下不同汉字代表一个数字  79365 * 7 = 555555

  算法描述题
x        算
------------
题题题题题题

*/

package main

import "fmt"

func main() {
	i1, i2, i3, i4, i5 := 1.0, 0.0, 0.0, 0.0, 1.0

	mul, rlt := 0.0, 0.0

	for i1 = 1; i1 <= 9; i1++ { // i1、i5不能为0，其他可以为0
		for i2 = 0; i2 <= 9; i2++ {
			for i3 = 0; i3 <= 9; i3++ {
				for i4 = 0; i4 <= 9; i4++ {
					for i5 = 1; i5 <= 9; i5++ {
						mul = (i1*10000 + i2*1000 + i3*100 + i4*10 + i5) * i1
						rlt = i5 * (100000 + 10000 + 1000 + 100 + 10 + 1)
						if mul == rlt {
							fmt.Printf("\n%5.f%2.f%2.f%2.f%2.f\n", i1, i2, i3, i4, i5)
							fmt.Printf("X%12.f\n", i1)
							fmt.Printf("-------------\n")
							fmt.Printf("\n%3.f%2.f%2.f%2.f%2.f%2.f\n", i5, i5, i5, i5, i5, i5)
							fmt.Println("\n")
						}
					}
				}
			}
		}
	}
	//

}

//
/*

    7 9 3 6 5
X           7
-------------

  5 5 5 5 5 5

*/
