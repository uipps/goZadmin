// 魔术方阵：方阵中的每行、每列或对角线位置的数各自相加的和均相等
//    采用最笨的方法，但是时间复杂度太大 O(n!)
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\suanfa\moshu_fangzhen\fangzhen.go -n 3 -b 1
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\suanfa\moshu_fangzhen\fangzhen.go
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
    "time"
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

    startTime := time.Now().UnixNano()
    fmt.Printf("startTime：%d, %s\n", startTime/1e3, time.Unix(0, startTime).Format("2006-01-02 15:04:05"))

    // 采用最笨方法先做3阶，遍历所有情况 (n^2)!
    if (4 == argN) {
        // 耗时太长，暂不能计算，采用4k规律计算比较好
        fmt.Printf("耗时太久，暂不支持运行\n")
        // fangzhen4jie(argN)
        time.Sleep(200 * time.Microsecond) // 200微秒
    } else {
        fangzhen3jie(argN)
    }

    // 执行时间计算
    endTime := time.Now().UnixNano()
    fmt.Printf("  endTime：%d, %s\n", endTime/1e3, time.Unix(0, endTime).Format("2006-01-02 15:04:05"))
    nanoSeconds := float64(endTime-startTime) / 1e3
    fmt.Println("spendTime：", nanoSeconds)
}

/**
   a[1]    a[2]    a[3]
   a[4]    a[5]    a[6]
   a[7]    a[8]    a[9]
a[1]的取值，从1~9遍历 (base+0 ~ 8+base), 其他则从剩下的数据进行便利

*/
func fangzhen3jie(n int) {
    //var a,b,c,d,e,f,g,h int
    n = 3   // 强制为3
    conunt := 0

    // 初始化数组
    // 三阶，总共9个数，逐个试探，试探次数 9! = 362880
    total := n * n
    for i := 0; i <= total; i++ {
        arrYiWei = append(arrYiWei, argBase + i - 1) // 下标从0开始，但是0下标不用，从1开始，使得a[1]=base; 并让数组长度增加1位
    }
    //fmt.Println(arrYiWei)

    // arrYiWei[0]，由于习惯问题， 下标0不用，从1开始编号

    for a := 0; a < total; a++ {
        arrYiWei[1] = argBase + a
        for b := a+1; b < total; b++ {  // 不用从0开始，直接用一个比a大的值，减少循环次数，由于最终结果是对称的，所以这个限制可以
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
                                        conunt++    // 记录多少次运算
                                        time.Sleep(100 * time.Microsecond) // 100微秒, 防止cpu太高
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
    fmt.Printf("尝试计算了: %d次\n", conunt)
    return
}

func fangzhen4jie(n int) {
    n = 4
    var conunt int64 = 0

    // 初始化数组
    // 四阶，总共16个数，逐个试探，试探次数 16! = 87178291200
    total := n * n
    for i := 0; i <= total; i++ {
        arrYiWei = append(arrYiWei, argBase + i - 1) // 下标从0开始，但是0下标不用，从1开始，使得a[1]=base; 并让数组长度增加1位
    }
    //fmt.Println(arrYiWei)

    // arrYiWei[0]，由于习惯问题， 下标0不用，从1开始编号
    // var a,b,c,d,e,f,g,h,i,j, k,l,m,n,o,p int
    for a := 0; a < total; a++ {
        arrYiWei[1] = argBase + a
        for b := a+1; b < total; b++ {  // 节省时间，直接从a+1开始，直接要求比第一项大的情况，因为结果对称，因此可以这么限制
            arrYiWei[2] = argBase + b   //
            for c := b+1; c < total; c++ {  // 节省时间，不从0开始
                arrYiWei[3] = argBase + c
                for d := a+3; d < total; d++ {          // 节省时间，直接从4开始
                    if (d == a || d == b || d == c) {
                        continue
                    }
                    arrYiWei[4] = argBase + d
                    for e := d+1; e < total; e++ {      // 节省时间，比第四位大
                        if (e == b || e == c || e == a || e == d) {
                            continue
                        }
                        arrYiWei[5] = argBase + e
                        for f := e+1; f < total; f++ {    // 节省时间，比上一位大
                            if (f == a || f == b || f == c || f == d || f == e) {
                                continue
                            }
                            arrYiWei[6] = argBase + f
                            for g := a+1; g < total; g++ {    // 节省时间，后面的初值都比第一个大
                                if (g == a || g == b || g == c || g == d || g == e || g == f) {
                                    continue
                                }
                                arrYiWei[7] = argBase + g
                                for h := a+1; h < total; h++ {
                                    if (h == a || h == b || h == c || h == d || h == e || h == f || h == g) {
                                        continue
                                    }
                                    arrYiWei[8] = argBase + h
                                    for i := a+1; i < total; i++ {
                                        if (i == a || i == b || i == c || i == d || i == e || i == f || i == g || i == h) {
                                            continue
                                        }
                                        arrYiWei[9] = argBase + i
                                        for j := a+1; j < total; j++ {
                                            if (j == a || j == b || j == c || j == d || j == e || j== f || j== g || j== h || j == i) {
                                                continue
                                            }
                                            arrYiWei[10] = argBase + j
                                            for k := a+1; k < total; k++ {
                                                if (k == a || k == b || k == c || k == d || k == e || k == f || k == g || k == h || k == i || k == j) {
                                                    continue
                                                }
                                                arrYiWei[11] = argBase + k
                                                for l := a+1; l < total; l++ {
                                                    if (l == a || l == b || l == c || l == d || l == e || l == f || l == g || l == h || l == i || l == j||
                                                        l == k) {
                                                        continue
                                                    }
                                                    arrYiWei[12] = argBase + l
                                                    for m := a+1; m < total; m++ {
                                                        if (m == a || m == b || m == c || m == d || m == e || m == f || m == g || m == h || m == i || m == j||
                                                            m == k || m == l) {
                                                            continue
                                                        }
                                                        arrYiWei[13] = argBase + m
                                                        for n := a+1; n < total; n++ {
                                                            if (n == a || n == b || n == c || n == d || n == e || n == f || n == g || n == h || n == i || n == j ||
                                                                n == k || n == l || n == m ) {
                                                                continue
                                                            }
                                                            arrYiWei[14] = argBase + n
                                                            for o := a+1; o < total; o++ {
                                                                if (o == a || o == b || o == c || o == d || o == e || o == f || o == g || o == h || o == i || o == j||
                                                                    o == k || o == l || o == m || o == n) {
                                                                    continue
                                                                }
                                                                arrYiWei[15] = argBase + o
                                                                for p := a+1; p < total; p++ {
                                                                    if (p == a || p == b || p == c || p == d || p == e || p == f || p == g || p == h || p == i || p == j ||
                                                                        p == k || p == l || p == m || p == n || p == o) {
                                                                        continue
                                                                    }
                                                                    arrYiWei[16] = argBase + p

                                                                    // 每行、每列、对角线分别都相等，则打印出此时的数据
                                                                    t_he := arrYiWei[1]+arrYiWei[2]+arrYiWei[3]+arrYiWei[4]
                                                                    if (t_he == arrYiWei[5]+arrYiWei[6]+arrYiWei[7]+arrYiWei[8] &&
                                                                        t_he == arrYiWei[9]+arrYiWei[10]+arrYiWei[11]+arrYiWei[12]&&
                                                                        t_he == arrYiWei[13]+arrYiWei[14]+arrYiWei[15]+arrYiWei[16]&&
                                                                        t_he == arrYiWei[1]+arrYiWei[5]+arrYiWei[9]+arrYiWei[13] &&
                                                                        t_he == arrYiWei[2]+arrYiWei[6]+arrYiWei[10]+arrYiWei[14] &&
                                                                        t_he == arrYiWei[3]+arrYiWei[7]+arrYiWei[11]+arrYiWei[15] &&
                                                                        t_he == arrYiWei[4]+arrYiWei[8]+arrYiWei[12]+arrYiWei[16] &&
                                                                        t_he == arrYiWei[1]+arrYiWei[6]+arrYiWei[11]+arrYiWei[16] &&
                                                                        t_he == arrYiWei[4]+arrYiWei[7]+arrYiWei[10]+arrYiWei[13]) {
                                                                        print_juzheng(arrYiWei, n)
                                                                        fmt.Println("\n")   // 增加2个换行
                                                                    }
                                                                    conunt++    // 记录多少次运算
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
                        }
                    }
                }
            }
        }
    }
    fmt.Printf("count: %d\n", conunt)

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
