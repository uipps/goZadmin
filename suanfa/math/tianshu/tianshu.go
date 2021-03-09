/**

go run tianshu.go -b 8 -n 10
go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\math\tianshu\tianshu.go

go run ~/develope/go/go_code/src/github.com/uipps/goZadmin/suanfa/math/tianshu/tianshu.go

填数游戏，如下不同汉字代表一个数字  79365 * 7 = 555555

  算法描述题
x        算
------------
题题题题题题



    7 9 3 6 5
X           7
-------------

  5 5 5 5 5 5

*/

package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now().UnixNano()
	fmt.Printf("startTime：%d, %s\n", startTime/1e3, time.Unix(0, startTime).Format("2006-06-06 06:06:06"))

	fmt.Println("\n\n2位数字：")
	numeric2()
	fmt.Println("\n\n3位数字：")
	numeric3()
	fmt.Println("\n\n4位数字：")
	numeric4()
	fmt.Println("\n\n5位数字：")
	numeric5()
	fmt.Println("\n\n6位数字：")
	numeric6()
	fmt.Println("\n\n7位数字：")
	numeric7()
	fmt.Println("\n\n8位数字：")
	numeric8()
	fmt.Println("\n\n3个数字，6位数：")
	numeric3_2()

	// 执行时间计算
	endTime := time.Now().UnixNano()
	fmt.Printf("  endTime：%d, %s\n", endTime/1e3, time.Unix(0, endTime).Format("2006-06-06 06:06:06"))
	nanoSeconds := float64(endTime-startTime) / 1e3
	fmt.Println("spendTime：", nanoSeconds)
}

func numeric5() {
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
}

// 2个数字, 未找到
/*

    算题
x     算
------------
  题题题
*/
func numeric2() {
	i1, i2 := 1.0, 1.0

	mul, rlt := 0.0, 0.0

	for i1 = 1; i1 <= 9; i1++ { // i1、i2不能为0
		for i2 = 1; i2 <= 9; i2++ {
			mul = (i1*10 + i2) * i1
			rlt = i2 * (100 + 10 + 1)
			if mul == rlt {
				fmt.Printf("\n%5.f%2.f\n", i1, i2)
				fmt.Printf("X%6.f\n", i1)
				fmt.Printf("---------\n")
				fmt.Printf("\n%3.f%2.f%2.f\n", i2, i2, i2)
				fmt.Println("\n")
			}
		}
	}
}

// 3个数字, 无结果
/*

    算法题
x      算
------------
  题题题题
*/
func numeric3() {
	i1, i2, i3 := 1.0, 0.0, 1.0

	mul, rlt := 0.0, 0.0

	for i1 = 1; i1 <= 9; i1++ { // i1、i5不能为0，其他可以为0
		for i2 = 0; i2 <= 9; i2++ {
			for i3 = 1; i3 <= 9; i3++ {
				mul = (i1*100 + i2*10 + i3) * i1
				rlt = i3 * (1000 + 100 + 10 + 1)
				if mul == rlt {
					fmt.Printf("\n%5.f%2.f%2.f\n", i1, i2, i3)
					fmt.Printf("X%8.f\n", i1)
					fmt.Printf("-------------\n")
					fmt.Printf("\n%3.f%2.f%2.f%2.f\n", i3, i3, i3, i3)
					fmt.Println("\n")
				}
			}
		}
	}
}

// 4个数字, 无结果
/*

   算法例题
x       算
------------
 题题题题题
*/
func numeric4() {
	i1, i2, i3, i4 := 1.0, 0.0, 0.0, 1.0

	mul, rlt := 0.0, 0.0

	for i1 = 1; i1 <= 9; i1++ { // i1、i5不能为0，其他可以为0
		for i2 = 0; i2 <= 9; i2++ {
			for i3 = 0; i3 <= 9; i3++ {
				for i4 = 0; i4 <= 9; i4++ {
					mul = (i1*1000 + i2*100 + i3*10 + i4) * i1
					rlt = i4 * (10000 + 1000 + 100 + 10 + 1)
					if mul == rlt {
						fmt.Printf("\n%5.f%2.f%2.f%2.f\n", i1, i2, i3, i4)
						fmt.Printf("X%10.f\n", i1)
						fmt.Printf("-------------\n")
						fmt.Printf("\n%3.f%2.f%2.f%2.f%2.f\n", i4, i4, i4, i4, i4)
						fmt.Println("\n")
					}
				}
			}
		}
	}
}

// 6个数字, 都没有结果。
/*

  算法描述的题
x          算
------------
题题题题题题题


*/
func numeric6() {
	i1, i2, i3, i4, i5, i6 := 1.0, 0.0, 0.0, 0.0, 0.0, 1.0

	mul, rlt := 0.0, 0.0

	for i1 = 1; i1 <= 9; i1++ { // i1、i5不能为0，其他可以为0
		for i2 = 0; i2 <= 9; i2++ {
			for i3 = 0; i3 <= 9; i3++ {
				for i4 = 0; i4 <= 9; i4++ {
					for i5 = 0; i5 <= 9; i5++ {
						for i6 = 1; i6 <= 9; i6++ {
							mul = ((i1*10000+i2*1000+i3*100+i4*10+i5)*10 + i6) * i1
							rlt = i6 * (1000000 + 100000 + 10000 + 1000 + 100 + 10 + 1)
							if mul == rlt {
								fmt.Printf("\n%5.f%2.f%2.f%2.f%2.f%2.f\n", i1, i2, i3, i4, i5, i6)
								fmt.Printf("X%14.f\n", i1)
								fmt.Printf("-----------------\n")
								fmt.Printf("\n%3.f%2.f%2.f%2.f%2.f%2.f%2.f\n", i6, i6, i6, i6, i6, i6, i6)
								fmt.Println("\n")
							}
						}
					}
				}
			}
		}
	}
}

// 7个数字, 都没有结果。
/*

  算法描述的例题
x            算
--------------
题题题题题题题题
*/
func numeric7() {
	i1, i2, i3, i4, i5, i6, i7 := 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0

	mul, rlt := 0.0, 0.0

	for i1 = 1; i1 <= 9; i1++ { // i1、i5不能为0，其他可以为0
		for i2 = 0; i2 <= 9; i2++ {
			for i3 = 0; i3 <= 9; i3++ {
				for i4 = 0; i4 <= 9; i4++ {
					for i5 = 0; i5 <= 9; i5++ {
						for i6 = 0; i6 <= 9; i6++ {
							for i7 = 1; i7 <= 9; i7++ {
								mul = (i1*1000000 + i2*100000 + i3*10000 + i4*1000 + i5*100 + i6*10 + i7) * i1
								rlt = i7 * (10000000 + 1000000 + 100000 + 10000 + 1000 + 100 + 10 + 1)
								if mul == rlt {
									fmt.Printf("\n%5.f%2.f%2.f%2.f%2.f%2.f%2.f\n", i1, i2, i3, i4, i5, i6, i7)
									fmt.Printf("X%16.f\n", i1)
									fmt.Printf("-----------------\n")
									fmt.Printf("\n%3.f%2.f%2.f%2.f%2.f%2.f%2.f%2.f\n", i7, i7, i7, i7, i7, i7, i7, i7)
									fmt.Println("\n")
								}
							}
						}
					}
				}
			}
		}
	}
}

// 8个数字, 都没有结果。将
/*

  算法描述综合习题
x             算
----------------
题题题题题题题题题
*/
func numeric8() {
	i1, i2, i3, i4, i5, i6, i7, i8 := 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0

	mul, rlt := 0.0, 0.0

	for i1 = 1; i1 <= 9; i1++ { // i1、i5不能为0，其他可以为0
		for i2 = 0; i2 <= 9; i2++ {
			for i3 = 0; i3 <= 9; i3++ {
				for i4 = 0; i4 <= 9; i4++ {
					for i5 = 0; i5 <= 9; i5++ {
						for i6 = 0; i6 <= 9; i6++ {
							for i7 = 0; i7 <= 9; i7++ {
								for i8 = 1; i8 <= 9; i8++ {
									mul = (i1*10000000 + i2*1000000 + i3*100000 + i4*10000 + i5*1000 + i6*100 + i7*10 + i8) * i1
									rlt = i8 * (100000000 + 10000000 + 1000000 + 100000 + 10000 + 1000 + 100 + 10 + 1)
									if mul == rlt {
										fmt.Printf("\n%5.f%2.f%2.f%2.f%2.f%2.f%2.f%2.f\n", i1, i2, i3, i4, i5, i6, i7, i8)
										fmt.Printf("X%18.f\n", i1)
										fmt.Printf("-----------------\n")
										fmt.Printf("\n%3.f%2.f%2.f%2.f%2.f%2.f%2.f%2.f%2.f\n", i8, i8, i8, i8, i8, i8, i8, i8, i8)
										fmt.Println("\n")
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

/*

3个数字
  题题题题题题
x          做
------------
我做做做做做题


    5 5 5 5 5 5
X             9
---------------
  4 9 9 9 9 9 5


    4 4 4 4 4 4
X             6
-----------------

  2 6 6 6 6 6 4

*/
func numeric3_2() {
	i1, i2, i3 := 1.0, 0.0, 1.0

	mul, rlt := 0.0, 0.0

	for i1 = 1; i1 <= 9; i1++ { // i1、i5不能为0，其他可以为0
		for i2 = 1; i2 <= 9; i2++ {
			for i3 = 1; i3 <= 9; i3++ {
				mul = i3 * (100000 + 10000 + 1000 + 100 + 10 + 1) * i2
				rlt = i1*1000000 + i2*(100000+10000+1000+100+10) + i3
				if mul == rlt {
					fmt.Printf("\n%5.f%2.f%2.f%2.f%2.f%2.f\n", i3, i3, i3, i3, i3, i3)
					fmt.Printf("X%14.f\n", i2)
					fmt.Printf("-----------------\n")
					fmt.Printf("\n%3.f%2.f%2.f%2.f%2.f%2.f%2.f\n", i1, i2, i2, i2, i2, i2, i3)
					fmt.Println("\n")
				}
			}
		}
	}
}
