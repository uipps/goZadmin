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
   a[0]    a[1]    a[2]
   a[3]    a[4]    a[5]
   a[6]    a[7]    a[8]
a[0]的取值，从1~9遍历 (1+base ~ 9+base), 其他则从剩下的数据进行便利

*/
func fangzhen3jie(n int) {
    // 初始化数组
    var a,b,c,d,e,f,g,h int

    // 三阶，总共9个数，逐个试探
    total := n * n
    for i := 0; i < total; i++ {
        arrYiWei = append(arrYiWei, argBase+i) // 下标从0开始
    }
    fmt.Println(arrYiWei)

    for a = 0; a < total; a++ {
        arrYiWei[0] = argBase + a
        for b = 0; b < total; b++ {
            if (a == b) {
                // 数字不能相同
                continue
            }
            arrYiWei[1] = argBase + b
            for c = 0; c < total; c++ {
                if (c == a || c == b) {
                    continue
                }
                arrYiWei[2] = argBase + c
                for d = 0; d < total; d++ {
                    if (d == a || d == b || d == c) {
                        continue
                    }
                    arrYiWei[3] = argBase + d
                    for e = 0; e < total; e++ {
                        if (e == a || e == b || e == c || e == d) {
                            continue
                        }
                        arrYiWei[4] = argBase + e
                        for f = 0; f < total; f++ {
                            if (f == a || f == b || f == c || f == d || f == e) {
                                continue
                            }
                            arrYiWei[5] = argBase + f
                            for g = 0; g < total; g++ {
                                if (g == a || g == b || g == c || g == d || g == e || g == f) {
                                    continue
                                }
                                arrYiWei[6] = argBase + g
                                for h = 0; h < total; h++ {
                                    if (h == a || h == b || h == c || h == d || h == e || h == f || h == g) {
                                        continue
                                    }
                                    arrYiWei[7] = argBase + h
                                    for i := 0; i < total; i++ {
                                        if (i == a || i == b || i == c || i == d || i == e || i == f || i == g || i == h) {
                                            continue
                                        }
                                        arrYiWei[8] = argBase + i
                                        // 每行、每列、对角线分别都相等，则打印出此时的数据
                                        if (arrYiWei[0]+arrYiWei[1]+arrYiWei[2] == arrYiWei[3]+arrYiWei[4]+arrYiWei[5] &&
                                            arrYiWei[0]+arrYiWei[1]+arrYiWei[2] == arrYiWei[6]+arrYiWei[7]+arrYiWei[8] &&
                                            arrYiWei[0]+arrYiWei[1]+arrYiWei[2] == arrYiWei[0]+arrYiWei[3]+arrYiWei[6] &&
                                            arrYiWei[0]+arrYiWei[1]+arrYiWei[2] == arrYiWei[1]+arrYiWei[4]+arrYiWei[7] &&
                                            arrYiWei[0]+arrYiWei[1]+arrYiWei[2] == arrYiWei[2]+arrYiWei[5]+arrYiWei[8] &&
                                            arrYiWei[0]+arrYiWei[1]+arrYiWei[2] == arrYiWei[0]+arrYiWei[4]+arrYiWei[8] &&
                                            arrYiWei[0]+arrYiWei[1]+arrYiWei[2] == arrYiWei[2]+arrYiWei[4]+arrYiWei[6]) {
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
    //var length,i int
    length := len(arr)
    for i := 0; i < length; i++ {
        if 0 == i%3 {
            fmt.Println("") // 换行
        }
        fmt.Printf("%3d", arr[i])
    }
}
