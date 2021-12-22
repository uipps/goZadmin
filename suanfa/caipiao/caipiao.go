/*

假设有一种彩票，每注由7个1~29的数字组成，且这7个数字不能相同，编写程序生成所有的号码组合。循环嵌套程序，
    1  2  3  4  5  6  7
    7  6  5  4  3  2  1
是一样的

-- 大乐透（体彩）：5+2，前区1-35|后区1-12
go run caipiao.go -e daletou -a 35 -c 12

-- 双色球（福彩）: 6+1，6红1蓝， 红球1-33|蓝球1-16
go run caipiao.go -e shuangseqiu -a 33 -c 16


go run caipiao.go -l 7 -a 35 -c 12
go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\caipiao\caipiao.go -l 3 -a 5

一、 福彩双色球中奖概率
  红球1-33|蓝球1-16
        6 + 1

<一> 一等奖 6红+1蓝
    1) 排列组合公式 C33-6 * C16-1 :
33*32*31*30*29*28 * 16/(6*5*4*3*2*1 * 1) = 17721088 （种）组合

    2) 概率表示也可以：
6/33 * 5/32 * 4/31 * 3/30 * 2/29 * 1/28  |  1/16
=  1/17721088
其实就是所有的组合有 17721088种，用程序也可以输出这个排列组合数。

<二> 二等奖 6红+0蓝
也就是一等奖的最后那位有15种可能，对应的中奖组合是15种，所以中奖概率是：
15/17721088 = 1/1181405.86666667 ~= 1/1181406

<三> 三等奖 5红+1蓝
只有一个红的不中，6个中奖红球任选一个红球，这个不中的红球从余下的27个球中取，所以有 6*27 = 162种情况；
6*27/17721088 = 1/109389.432 ~= 1/109389
或者可以这么理解：假设红色中奖号码1，2，3，4，5，6，蓝色中奖号码1；那么某个红球不中的排列组合情况有多少种？剩余27个红球可以顶替1的位置，
                 也可以顶替2的位置......也可以顶替6的位置...... 所以一共有6*27种可能。

<四> 四等奖 5红+0蓝 或 4红1蓝
  (1) 5红+0蓝：从6个中奖中选1个不中的就有C6-1种，第6个球有27种，另外16个篮球选1个不中奖就有15种情况：
      C6-1 * C27-1 * C15-1 = 2430
  (2) 4红1蓝： 6个中奖选2个不中奖有C6-2=30种情况，而2种不中奖是从27个不中奖球中选择C27-2， C6-2*C27-2=5265
  合并两种情况，一共存在的组合是 2430 + 5265 = 7695 种
  所以中奖概率为： 7695/17721088 = 1/2302.9354 ~= 1/2303

<五> 五等奖 4红+0蓝 或 3红1蓝
  (1) 4红+0蓝: C6-2 * C27-2 * C15-1 = 78975
  (2) 3红1蓝： C6-3 * C27-3 = 58500
  合并两种情况：(78975+58500)/17721088 = 1/128.904077 ~= 1/129

<六> 六等奖 1蓝 + （0红 or 1红 or 2红）
  也就是蓝的中，但红的可中可不中，红的中奖不超过2个（超过2个就是五等奖）
  (1) 2红+1蓝: C6-4 * C27-4
  (2) 1红+1蓝: C6-5 * C27-5
  (3) 0红+1蓝: C6-6 * C27-6
  合并三种情况：(C6-4 * C27-4 + C6-5 * C27-5 + C6-6 * C27-6)/17721088 = 104364/17721088 = 1/16.98 ~= 1/17


二、 体彩大乐透中奖概率
    前区1-35 | 后区1-12
           5 + 2

<一> 一等奖 5前区+2后区
一等奖的排列组合有多少种？C35-5 * C12-2 :
35*34*33*32*31 * 12*11 /(5*4*3*2*1 * 2*1)
= 21425712 （种）组合

<二> 二等奖 5前+1后
也就是一等奖的后区2个中了1个，C10-1 * C2-1 = 20，20种可能，所以中奖概率
20/21425712 = 1/1071285.6 ~= 1/1071286

<三> 三等奖 5前+0后 + 4前+2后
  (1) 5前+0后：C10-2
  (2) 4前+2后: C5-1*C30-1
  两种合并：(C10-2 + C5-1*C30-1)/21425712 = 195/21425712 = 1/109875.44 ~= 1/109875

<四> 四等奖 4前+1后 + 3前+2后
  (1) 4前+1后：C5-1 * C30-1 * C2-1*C10-1
  (2) 3前+2后: C5-2*C30-2
  两种合并：7350/21425712 = 1/2915.06 ~= 1/2915

<五> 五等奖 4前+0后 + 3前+1后 + 2前+2后

  (C5-1*C30-1*C10-2 + C5-2*C30-2*C2-1*C10-1 + C5-3*C30-3)/21425712 = 1/159.4768 ~= 1/159

<六> 六等奖 3前+0后 + 2前+1后 + 1前+2后 + 0前+2后

  = 1/16.644 ~= 1/17




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
