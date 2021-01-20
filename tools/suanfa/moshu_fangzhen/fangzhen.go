// 魔术方阵：方阵中的每行、每列或对角线位置的数各自相加的和均相等
//    采用最笨的方法，但是时间复杂度太大 O(n!)
/**

请输入矩阵行数(奇数):3
 8    1    6
 3    5    7
 4    9    2

*/

package main

import (
    "flag"
    "fmt"
)

var (
    argN     int
    argBase  int   // 从几开始的基数，默认是从1开始
    arrYiWei []int // 用1维数组存放数字好了，TODO 以后改成二维数组
)

func init() {
    flag.IntVar(&argN, "n", 3, "Usage: 5 6 7")
    flag.IntVar(&argBase, "b", 1, "Usage: 1")
}

func main() {
    flag.Parse()

    // 采用最笨方法先做3阶，遍历所有情况 (n^2)!
    fangzhen3jie(argN)
}

/**
   a[1]    a[2]    a[3]
   a[4]    a[5]    a[6]
   a[7]    a[8]    a[9]
a[1]的取值，从1~9遍历 (base+0 ~ 8+base), 其他则从剩下的数据进行便利

*/
func fangzhen3jie(n int) {
    //var a,b,c,d,e,f,g,h int

    // 初始化数组
    // 三阶，总共9个数，逐个试探
    total := n * n
    for i := 0; i <= total; i++ {
        arrYiWei = append(arrYiWei, argBase + i - 1) // 下标从0开始，但是0下标不用，从1开始，使得a[1]=base; 并让数组长度增加1位
    }
    fmt.Println(arrYiWei)

    // arrYiWei[0]，由于习惯问题， 下标0不用，从1开始编号

    for a := 0; a < total; a++ {
        arrYiWei[1] = argBase + a
        for b := 0; b < total; b++ {
            if (a == b) {
                // 数字不能相同
                continue
            }
            arrYiWei[2] = argBase + b
            for c := 0; c < total; c++ {
                if (c == a || c == b) {
                    continue
                }
                arrYiWei[3] = argBase + c
                for d := 0; d < total; d++ {
                    if (d == a || d == b || d == c) {
                        continue
                    }
                    arrYiWei[4] = argBase + d
                    for e := 0; e < total; e++ {
                        if (e == a || e == b || e == c || e == d) {
                            continue
                        }
                        arrYiWei[5] = argBase + e
                        for f := 0; f < total; f++ {
                            if (f == a || f == b || f == c || f == d || f == e) {
                                continue
                            }
                            arrYiWei[6] = argBase + f
                            for g := 0; g < total; g++ {
                                if (g == a || g == b || g == c || g == d || g == e || g == f) {
                                    continue
                                }
                                arrYiWei[7] = argBase + g
                                for h := 0; h < total; h++ {
                                    if (h == a || h == b || h == c || h == d || h == e || h == f || h == g) {
                                        continue
                                    }
                                    arrYiWei[8] = argBase + h
                                    for i := 0; i < total; i++ {
                                        if (i == a || i == b || i == c || i == d || i == e || i == f || i == g || i == h) {
                                            continue
                                        }
                                        arrYiWei[9] = argBase + i
                                        // 每行、每列、对角线分别都相等，则打印出此时的数据
                                        t_he := arrYiWei[1]+arrYiWei[2]+arrYiWei[3]
                                        if (t_he == arrYiWei[4]+arrYiWei[5]+arrYiWei[6] &&
                                            t_he == arrYiWei[7]+arrYiWei[8]+arrYiWei[9] &&
                                            t_he == arrYiWei[1]+arrYiWei[4]+arrYiWei[7] &&
                                            t_he == arrYiWei[2]+arrYiWei[5]+arrYiWei[8] &&
                                            t_he == arrYiWei[3]+arrYiWei[6]+arrYiWei[9] &&
                                            t_he == arrYiWei[1]+arrYiWei[5]+arrYiWei[9] &&
                                            t_he == arrYiWei[3]+arrYiWei[5]+arrYiWei[7]) {
                                            print_juzheng(arrYiWei, n)
                                            fmt.Println("\n")   // 增加2个换行
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

    return
}

// 格式化输出矩阵，通过一维数组输出
func print_juzheng(arr []int, n int) {
    length := len(arr)
    for i := 1; i < length; i++ {
        if 0 == (i-1)%3 {
            fmt.Println("") // 换行
        }
        fmt.Printf("%3d", arr[i])
    }
}
