/**

题目：　A、B、C、D、E五个人在某天夜里合伙去捕鱼，到第二天凌晨时都疲惫不堪，于是各自找地方睡觉。日上三杆，
A第一个醒来，他将鱼分为五份，把多余的一条鱼扔掉，拿走自己的一份。
B第二个醒来，也将鱼分为五份，把多余的一条鱼扔掉，拿走自己的一份。 求至少捕了多少条鱼？

 go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\math\buyu_fenyu\buyu01.go -n 1
*/

package main

import (
    "flag"
    "fmt"
    "time"
)

var (
    argYu int
)

func init() {
    flag.IntVar(&argYu, "n", 0, "Usage: 10 100 13")
}

func main() {
    flag.Parse()

    buyu01()
    fmt.Printf("网上另一个方法，似乎不对：")
    buyu02() // TODO 网上另一个方法，似乎不对
}

func buyu01() {
    i := 0 // 最少有几条鱼
    oneFen := 0

    for i = argYu; i < 1000000; i++ {
        // 减去1能被5整除
        if 0 == (i-1)%5 {
            oneFen = (i - 1) / 5 // A拿走的一份

            if (i-1-oneFen-1)%5 == 0 {
                break
            }
        }
        time.Sleep(100 * time.Microsecond)
    }
    fmt.Printf("至少捕了%d条鱼！其中A拿走了%d, B拿走了%d\n", i, (i-1)/5, (i-1-(i-1)/5-1)/5)
}

// 这个程序似乎不对
func buyu02() {
    n, fg, x, i, t := 1, 1, 1, 1, 1

    for n = 6; 1 == t; n++ {
        for i = 0; i < 5; i++ {
            if i == 0 {
                x = n
                fg = 0
            }
            if (x-1)%5 == 0 {
                x = (x - 1) * 4 / 5
                fg++
            }
            if fg == 5 { // fg == 2 也可以啊？
                fmt.Printf("至少捕了%d条鱼！", n)
                t = 0
            }
        }
    }
}
