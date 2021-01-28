// 青蛙过河: -> -> -> O <- <- <-
// (1)青蛙只能向前移动,不能向后移动。
// (2)一次只能有一只青蛙向前移动。
// (3)青蛙可以向前面的空位中移动,也可以跳过对方一只青蛙进入前面的一个空位。
// (4)不能一次跳过两个位置。

// 分析：下面首先总结一下规律。由于只有一个空位,而且跳跃最多间隔一个青蛙,每次移动青蛙时需要防止出现阻塞现象,
//      所谓“阻塞”现象是：在移动青蛙的过程中,两个尚未到位的同向青蛙连接在一起,使其他青蛙无法继续移动
//
//     go run frog_guohe.go -n 3
//     go run /Users/cf/develope/go/go_code/src/github.com/uipps/goZadmin/suanfa/game/qingwa_guohe/frog_guohe.go -n 3
//     go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\game\qingwa_guohe\frog_guohe.go -n 3 -m 4

package main

import (
    "flag"
    "fmt"
)

var (
    argNumLeft  int  // 左边青蛙数量
    argNumRight int  // 右边青蛙数量，左右两边数量可以不等
    number      = -1 // 移动步数计数
)

func init() {
    flag.IntVar(&argNumLeft, "n", 3, "Usage: 3 4 5")  // 每一边青蛙数量
    flag.IntVar(&argNumRight, "m", 0, "Usage: 3 4 5") // 每一边青蛙数量
}

func main() {
    flag.Parse()

    // 参数校验，左、右侧青蛙数量
    if argNumLeft < 1 {
        fmt.Println("青蛙数量不能小于1\n")
        return
    }

    // 右侧默认跟左侧青蛙数量一样
    if argNumRight < 1 {
        argNumRight = argNumLeft
    }

	move_frog2(argNumLeft, argNumRight)
    //move_frog(argNumLeft, argNumRight) // 此方法来自网上，有问题，不可用

    return
}

// 位置如图：
//        -> -> -> O <- <- <-
//        -1 -1 -1 0 1  1  1

// left side the number is less than 0 (-1)
// right side the number is bigger than 0 (1)
// empty grid is 0 (0)
func move_frog(left_num int, right_num int) {
    arrRoad := initFrogArrByNum(left_num, right_num) // 构造青蛙数组

    // 初始状态
    bMove := true
    nSuccessCount := 0
    nEmptyGrid := left_num
    nLen := len(arrRoad)

    print_road(arrRoad)

    // move the left side
    for nSuccessCount < nLen-1 {
        // TODO 退出条件：两个同向挨到一起了，或都走完了。

        // 空白的左边向右边跳动(blank-2左移2位)
        bMove = true
        for i := 0; i < nEmptyGrid && bMove; i++ {
            // 左侧距离中间空白-1位置是对方的青蛙，-2位置是己方青蛙的情况
            if (nEmptyGrid >= 2 && arrRoad[nEmptyGrid-2] < 0) && (nEmptyGrid >= 1 && arrRoad[nEmptyGrid-1] > 0) {
                // 交换距离空白位置2的位置和空白，即实现了左侧向中间空白的跳跃
                arrRoad[nEmptyGrid-2], arrRoad[nEmptyGrid] = arrRoad[nEmptyGrid], arrRoad[nEmptyGrid-2]
                if nEmptyGrid == nLen-1 {
                    nSuccessCount++
                }
                nEmptyGrid -= 2
                print_road(arrRoad)
                bMove = false
                break
            }
        }

        // 空白的右边向左边跳动(blank++右移1位) : 右侧距离中间空白1位置是对方的青蛙，2位置是己方青蛙的情况
        for i := nLen - 1; i > nEmptyGrid && bMove; i-- {
            if (arrRoad[nEmptyGrid+2] > 0) && (arrRoad[nEmptyGrid+1] < 0) {
                //swap(&arrRoad[nEmptyGrid + 2], &arrRoad[nEmptyGrid]);
                arrRoad[nEmptyGrid+2], arrRoad[nEmptyGrid] = arrRoad[nEmptyGrid], arrRoad[nEmptyGrid+2]
                if nEmptyGrid == 0 {
                    nSuccessCount++
                }

                print_road(arrRoad)
                nEmptyGrid += 2
                bMove = false
                break
            }
        }

        // 空白的左边向右边跳动(blank--左移1位): 跳过去后，相连两个不能相等 [blank-2] != [blank+1]
        for i := 0; i < nEmptyGrid && bMove; i++ {
            if (nEmptyGrid>=1 && arrRoad[nEmptyGrid-1] < 0) && (nEmptyGrid < nLen-2 && (nEmptyGrid >= 2 && (arrRoad[nEmptyGrid-2]) != arrRoad[nEmptyGrid+1])) {
                //swap(&arrRoad[nEmptyGrid - 1], &arrRoad[nEmptyGrid]);
                arrRoad[nEmptyGrid-1], arrRoad[nEmptyGrid] = arrRoad[nEmptyGrid], arrRoad[nEmptyGrid-1]
                if nEmptyGrid == nLen-1 {
                    nSuccessCount++
                }

                print_road(arrRoad)
                nEmptyGrid -= 1
                bMove = false
                break
            }
        }

        // 空白的右边向左边跳动(blank++右移1位):
        for i := nLen - 1; i > nEmptyGrid && bMove; i-- {
            if (arrRoad[nEmptyGrid+1] > 0) && (nEmptyGrid > 0 && (nEmptyGrid>=1 && arrRoad[nEmptyGrid+2] != arrRoad[nEmptyGrid-1]) ) {
                //swap(&arrRoad[nEmptyGrid + 1], &arrRoad[nEmptyGrid]);
                arrRoad[nEmptyGrid+1], arrRoad[nEmptyGrid] = arrRoad[nEmptyGrid], arrRoad[nEmptyGrid+1]
                if nEmptyGrid == 0 {
                    nSuccessCount++
                }

                print_road(arrRoad)
                nEmptyGrid += 1
                bMove = false
                break
            }
        }
    }
}

// 每跳一次,打印一下位置情况 : -> O <- 表示
func print_road(arrRoad []int) {
    nLen := len(arrRoad)
    for i := 0; i < nLen; i++ {
        if arrRoad[i] < 0 {
            fmt.Printf("-> ")
        } else if arrRoad[i] > 0 {
            fmt.Printf("<- ")
        } else {
            fmt.Printf("__ ")
            //fmt.Printf("O ")
        }
    }
    fmt.Println("\n")
	return
}

// 输出经过一次移动后青蛙的位置情况：◇ □ ◆ 表示
func print_road2(frog []int) {
    number++
    if (0 == number) {
        fmt.Printf("初  始：")
    } else { //输出步数
        fmt.Printf("第%2d步：", number)
    }
    //print_road(frog);return 	// 另外一种方式输出

    nLen := len(frog)
    //nLen := argNumLeft + argNumRight +1
    for i := 0; i < nLen; i++ {
        if (frog[i] == 0) { //	若为空格
            //fmt.Printf("□")
            fmt.Printf("__ ")
        } else if (frog[i] == -1) { // 向右移动的青蛙
            fmt.Printf("◇ ")
        } else {
            fmt.Printf("◆ ") // 向左移动的青蛙
        }
    }

    fmt.Println("\n")
}

// 初始化青蛙数组
func initFrogArrByNum(left_num int, right_num int) []int {
    // 自动注入初始数据,左侧为-1,右侧为1,空白位置为0; 转化成数学问题
    var arrTest []int
    for i := 0; i < left_num+right_num+1; i++ {
        if i < left_num {
            arrTest = append(arrTest, -1)
        } else if i == left_num {
            arrTest = append(arrTest, 0)
        } else {
            arrTest = append(arrTest, 1)
        }
    }
    return arrTest
}

func move_frog2(left_num int, right_num int) {
    //frog := []int{-1, -1, -1, 0, 1, 1, 1} 		// 表示青蛙的数组
    frog := initFrogArrByNum(left_num, right_num) 	// 获取初始化数组
    fg_flag := true

    fmt.Printf("初始数组：")
    fmt.Println(frog, "\n")
    print_road2(frog) // 输出初始状态

    for {
        // 退出条件, 右侧1全部到左边，左边-1全部到右侧
        left_total := 0
        right_total := 0
        for i := 0; i < right_num; i++ {
            left_total += frog[i]
        }
        for i := left_num + right_num; i > right_num; i-- { // 空白右侧累加和
            right_total += frog[i]
        }
        if (left_total == 1*right_num && right_total == -1*left_num) {
            // 左侧=3；右侧=-3说明已经全部交换了（初始左边是-3，右边是3）
            break
        }
        //

        // 下面是循环处理
        fg_flag = true //fg_flag为青蛙移动一步的标记

        for i := 0; fg_flag && i < left_num+right_num-1; i++ { //循环检查现有排列
            if (frog[i] == -1 && frog[i+1] == 1 && frog[i+2] == 0) { //若向右的青蛙可以向右跳过
                frog[i], frog[i+2] = frog[i+2], frog[i] //向右跳动
                print_road2(frog)                       //输出移动一次各青蛙的位置
                fg_flag = false
            }
        }

        for i := 0; fg_flag && i < left_num+right_num-1; i++ {
            if (frog[i] == 0 && frog[i+1] == -1 && frog[i+2] == 1) { //若向左的青蛙可以向左跳
                frog[i], frog[i+2] = frog[i+2], frog[i] //向左跳动
                print_road2(frog)                       //输出移动一次各青蛙的位置
                fg_flag = false
            }
        }

        for i := 0; fg_flag && i < left_num+right_num; i++ { //循环检查现有排列
            if (frog[i] == -1 && frog[i+1] == 0 && (i == 0 || i == left_num+right_num-1 || frog[i-1] != frog[i+2])) {
                //若向右移动青蛙不会产生阻塞
                frog[i], frog[i+1] = frog[i+1], frog[i] //向右跳动
                print_road2(frog)                       //输出移动一次各青蛙的位置
                fg_flag = false
            }
        }

        for i := 0; fg_flag && i < left_num+right_num; i++ {
            if (frog[i] == 0 && frog[i+1] == 1 && (i == 0 || i == left_num+right_num-1 || frog[i-1] != frog[i+2])) {
                //若向左移动青蛙不会产生阻塞
                frog[i], frog[i+1] = frog[i+1], frog[i] //向左跳动
                print_road2(frog)                       //输出移动一次各青蛙的位置
                fg_flag = false
            }
        }
    }
    fmt.Printf("\n结果数组：")
    fmt.Println(frog)
    return
}
