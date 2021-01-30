/*

假设有一种彩票，每注由7个1~29的数字组成，且这7个数字不能相同，编写程序生成所有的号码组合。循环嵌套程序，
    1  2  3  4  5  6  7
    7  6  5  4  3  2  1
是一样的

-- 大乐透
go run caipiao.go -e daletou -a 35 -c 12

-- 双色球
go run caipiao.go -e shuangseqiu -a 32 -c 16


go run caipiao.go -l 7 -a 35 -c 12
go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\suanfa\caipiao.go -l 3 -a 5

*/

package main

import (
    "flag"
    "fmt"
    "time"
)

var (
    lottery01  []int // 彩票 1 2 3 4 5
    numSlice01 []int // 可用数字 1 ~ 29，用切片存放

    max01     = 5
    min01     = 1
    leng01    = 2
    minBlue01 = 1
    maxBlue01 = 16
    leixing01 = "daletou"
)

func init() {
    flag.IntVar(&min01, "i", 1, "Usage: 1 3")
    flag.IntVar(&max01, "a", 32, "Usage: 32 29")
    flag.IntVar(&leng01, "l", 7, "Usage: 7 3")

    flag.IntVar(&minBlue01, "b", 1, "Usage: 1")
    flag.IntVar(&maxBlue01, "c", 16, "Usage: 16")
    flag.StringVar(&leixing01, "e", "daletou", "Usage: shuangseqiu daletou")
}

func main() {
    flag.Parse()

    startTime := time.Now().UnixNano()
    //fmt.Println("startTime：", startTime)

    if max01 < min01 {
        fmt.Println("最大数不能小于最小数")
        return
    }
    if leng01 > (max01 - min01 + 1) {
        fmt.Println("玩法球数不能大于可选数字数")
        return
    }

    // 初始化随机数-切片
    for i := 0; i < max01-min01+1; i++ {
        numSlice01 = append(numSlice01, min01+i)
    }
    //fmt.Println(numSlice01)
    // 初始化-彩票
    for i := 0; i < leng01; i++ {
        lottery01 = append(lottery01, 0)
    }
    //fmt.Println(lottery01)

    //caipiao01(min01, max01)
    //caipiao02(min01, max01, leng01)
    if "daletou" == leixing01 {
        daletou01(min01, max01, minBlue01, maxBlue01)
    } else {
        shuangseqiu01(min01, max01, minBlue01, maxBlue01)
    }

    // 执行时间计算
    endTime := time.Now().UnixNano()
    fmt.Printf("startTime：%d, %s\n", startTime/1e3, time.Unix(0, startTime).Format("2006-01-02 15:04:05"))
    fmt.Printf("  endTime：%d, %s\n", endTime/1e3, time.Unix(0, endTime).Format("2006-01-02 15:04:05"))
    nanoSeconds := float64(endTime-startTime) / 1e3
    fmt.Println("spendTime：", nanoSeconds)
}

//
func caipiao02(min int, max int, leng int) {
    for i := max - min + 1; i > 0; i-- {
        lottery01[leng-1] = numSlice01[i-1]
        if leng > 1 {
            caipiao02(min, max-1, leng-1)
        } else {
            for j := leng01 - 1; j >= 0; j-- {
                fmt.Printf("%3d", lottery01[j])
            }
            fmt.Println("")
        }
    }
}

// https://github.com/chenhg5/collection
//         a := []int{2, 3, 4, 5, 6, 7}
//        fmt.Println(Collect(a).Contains(3))
//
func inArray(needle int) bool {
    return false
}

// 体彩大乐透：前区号码和后区号码，前区号码范围为 1～35，后区号码范围为 1～12。顺序不限.
// 1  2  3  4  5  6  7 和 7  6  5  4  3  2  1 是一样的，因此就按照从小到大的顺序好了
func daletou01(min int, max int, minBlue int, maxBlue int) {
    i1, i2, i3, i4, i5, i6, i7 := 1, 1, 1, 1, 1, 1, 1

    for i1 = min; i1 <= max; i1++ {
        for i2 = i1 + 1; i2 <= max; i2++ {
            for i3 = i2 + 1; i3 <= max; i3++ {
                for i4 = i3 + 1; i4 <= max; i4++ {
                    for i5 = i4 + 1; i5 <= max; i5++ {
                        for i6 = minBlue; i6 <= maxBlue; i6++ {
                            for i7 = i6 + 1; i7 <= maxBlue; i7++ {
                                // 输出
                                fmt.Printf("%3d%3d%3d%3d%3d%3d%3d\n", i1, i2, i3, i4, i5, i6, i7)
                                time.Sleep(200 * time.Microsecond) // 200微秒
                            }
                        }
                    }
                }
            }
        }
    }
}

// 双色球：每注投注号码由6个红色球号码和1个蓝色球号码组成。红色球号码从1--33中选择；蓝色球号码从1--16中选择。
// 1  2  3  4  5  6  7 和 7  6  5  4  3  2  1 是一样的，因此就按照从小到大的顺序好了
func shuangseqiu01(min int, max int, minBlue int, maxBlue int) {
    i1, i2, i3, i4, i5, i6, i7 := 1, 1, 1, 1, 1, 1, 1

    for i1 = min; i1 <= max; i1++ {
        for i2 = i1 + 1; i2 <= max; i2++ {
            for i3 = i2 + 1; i3 <= max; i3++ {
                for i4 = i3 + 1; i4 <= max; i4++ {
                    for i5 = i4 + 1; i5 <= max; i5++ {
                        for i6 = i5 + 1; i6 <= max; i6++ {
                            for i7 = minBlue; i7 <= maxBlue; i7++ {
                                // 输出
                                fmt.Printf("%3d%3d%3d%3d%3d%3d%3d\n", i1, i2, i3, i4, i5, i6, i7)
                                time.Sleep(200 * time.Microsecond) // 200微秒
                            }
                        }
                    }
                }
            }
        }
    }
}

// 最基本的方法
// 1  2  3  4  5  6  7 和 7  6  5  4  3  2  1 是一样的，因此就按照从小到大的顺序好了
func caipiao01(min int, max int) {
    i1, i2, i3, i4, i5, i6, i7 := 1, 1, 1, 1, 1, 1, 1

    for i1 = min; i1 <= max; i1++ {
        for i2 = i1 + 1; i2 <= max; i2++ {
            for i3 = i2 + 1; i3 <= max; i3++ {
                for i4 = i3 + 1; i4 <= max; i4++ {
                    for i5 = i4 + 1; i5 <= max; i5++ {
                        for i6 = i5 + 1; i6 <= max; i6++ {
                            for i7 = i6 + 1; i7 <= max; i7++ {
                                // 输出
                                fmt.Printf("%3d%3d%3d%3d%3d%3d%3d\n", i1, i2, i3, i4, i5, i6, i7)
                            }
                        }
                    }
                }
            }
        }
    }
}

// 1  2  3  4  5  6  7 和 7  6  5  4  3  2  1 是不一样的情况
func caipiao03(min int, max int) {
    i1, i2, i3, i4, i5, i6, i7 := 1, 1, 1, 1, 1, 1, 1

    for i1 = min; i1 <= max; i1++ {
        for i2 = min; i2 <= max; i2++ {
            if i2 == i1 {
                continue
            }
            for i3 = min; i3 <= max; i3++ {
                if i3 == i1 || i3 == i2 {
                    continue
                }
                for i4 = min; i4 <= max; i4++ {
                    if i4 == i1 || i4 == i2 || i4 == i3 {
                        continue
                    }
                    for i5 = min; i5 <= max; i5++ {
                        if i5 == i1 || i5 == i2 || i5 == i3 || i5 == i4 {
                            continue
                        }
                        for i6 = min; i6 <= max; i6++ {
                            if i6 == i1 || i6 == i2 || i6 == i3 || i6 == i4 || i6 == i5 {
                                continue
                            }
                            for i7 = min; i7 <= max; i7++ {
                                if i7 == i1 || i7 == i2 || i7 == i3 || i7 == i4 || i7 == i5 || i7 == i6 {
                                    continue
                                }
                                // 输出
                                fmt.Printf("%3d%3d%3d%3d%3d%3d%3d\n", i1, i2, i3, i4, i5, i6, i7)
                            }
                        }
                    }
                }
            }
        }
    }
}
