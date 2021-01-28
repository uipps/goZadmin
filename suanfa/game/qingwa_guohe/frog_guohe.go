// 青蛙过河:
//		-> -> -> O <- <- <-
//	go run frog_guohe.go -n 3
// 	go run /Users/cf/develope/go/go_code/src/github.com/uipps/goZadmin/suanfa/game/qingwa_guohe/frog_guohe.go -n 3

package main

import (
	"flag"
	"fmt"
)

var (
	argNum int // 每边青蛙数量
)

func init() {
	flag.IntVar(&argNum, "n", 3, "Usage: 3 4 5") // 每一边青蛙数量
}

func main() {
	flag.Parse()

	if argNum < 1 {
		fmt.Println("青蛙数量不能小于1\n")
		return
	}

	// 自动注入初始数据，左侧为-1，右侧为1，空白位置为0; 转化成数学问题
	var arrTest []int
	for i := 0; i < argNum*2+1; i++ {
		if i < argNum {
			arrTest = append(arrTest, -1)
		} else if i == argNum {
			arrTest = append(arrTest, 0)
		} else {
			arrTest = append(arrTest, 1)
		}
	}
	//arrTest = []int{-1, -1, -1, 0, 1, 1, 1};
	//fmt.Println(arrTest)
	//return

	print_road(arrTest, argNum*2+1)
	move_frog(arrTest, argNum*2+1, argNum)

	return
}

// left side the number is less than 0
// right side the number is bigger than 0
// empty grid is 0
func move_frog(arrRoad []int, nLen int, nEmptyGrid int) {
	bMove := true
	nSuccessCount := 0
	//nLen := len(arrRoad)

	// move the left side
	for nSuccessCount < nLen-1 {
		bMove = true
		for i := 0; i < nEmptyGrid && bMove; i++ {
			if (arrRoad[nEmptyGrid-2] < 0) && (arrRoad[nEmptyGrid-1] > 0) {
				// 交换swap(&arrRoad[nEmptyGrid - 2], &arrRoad[nEmptyGrid]);
				arrRoad[nEmptyGrid-2], arrRoad[nEmptyGrid] = arrRoad[nEmptyGrid], arrRoad[nEmptyGrid-2]
				if nEmptyGrid == nLen-1 {
					nSuccessCount++
				}
				nEmptyGrid -= 2
				print_road(arrRoad, nLen)
				bMove = false
				break
			}
		}

		for i := nLen - 1; i > nEmptyGrid && bMove; i-- {
			if (arrRoad[nEmptyGrid+2] > 0) && (arrRoad[nEmptyGrid+1] < 0) {
				//swap(&arrRoad[nEmptyGrid + 2], &arrRoad[nEmptyGrid]);
				arrRoad[nEmptyGrid+2], arrRoad[nEmptyGrid] = arrRoad[nEmptyGrid], arrRoad[nEmptyGrid+2]
				if nEmptyGrid == 0 {
					nSuccessCount++
				}

				print_road(arrRoad, nLen)
				nEmptyGrid += 2
				bMove = false
				break
			}
		}

		for i := 0; i < nEmptyGrid && bMove; i++ {
			if (arrRoad[nEmptyGrid-1] < 0) && (nEmptyGrid < nLen-2 && (arrRoad[nEmptyGrid-2]) != arrRoad[nEmptyGrid+1]) {
				//swap(&arrRoad[nEmptyGrid - 1], &arrRoad[nEmptyGrid]);
				arrRoad[nEmptyGrid-1], arrRoad[nEmptyGrid] = arrRoad[nEmptyGrid], arrRoad[nEmptyGrid-1]
				if nEmptyGrid == nLen-1 {
					nSuccessCount++
				}

				print_road(arrRoad, nLen)
				nEmptyGrid -= 1
				bMove = false
				break
			}
		}

		for i := nLen - 1; i > nEmptyGrid && bMove; i-- {
			if (arrRoad[nEmptyGrid+1] > 0) && (nEmptyGrid > 0 && arrRoad[nEmptyGrid+2] != arrRoad[nEmptyGrid-1]) {
				//swap(&arrRoad[nEmptyGrid + 1], &arrRoad[nEmptyGrid]);
				arrRoad[nEmptyGrid+1], arrRoad[nEmptyGrid] = arrRoad[nEmptyGrid], arrRoad[nEmptyGrid+1]
				if nEmptyGrid == 0 {
					nSuccessCount++
				}

				print_road(arrRoad, nLen)
				nEmptyGrid += 1
				bMove = false
				break
			}
		}
	}
}

// 每跳一次，打印一下个人位置情况
func print_road(arrRoad []int, nLen int) {
	//nLen := len(arrRoad)

	for i := 0; i < nLen; i++ {
		if arrRoad[i] < 0 {
			fmt.Printf("-> ")
		} else if arrRoad[i] > 0 {
			fmt.Printf("<- ")
		} else {
			fmt.Printf("O ")
		}
	}
	fmt.Println("\n")
}
