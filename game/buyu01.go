/**

题目：　A、B、C、D、E五个人在某天夜里合伙去捕鱼，到第二天凌晨时都疲惫不堪，于是各自找地方睡觉。日上三杆，
A第一个醒来，他将鱼分为五份，把多余的一条鱼扔掉，拿走自己的一份。
B第二个醒来，也将鱼分为五份，把多余的一条鱼扔掉，拿走自己的一份。 求至少捕了多少条鱼？

 go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\game\buyu01.go -n 1
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
	fmt.Printf("至少捕了%d条鱼！其中A拿走了%d, B拿走了%d", i, (i-1)/5, (i-1-(i-1)/5-1)/5)
}
