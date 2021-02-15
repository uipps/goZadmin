/**
 go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\monkey_eat_peach\monkey_eat_peach.go -n 10

8.2.1　猴子吃桃

    一只猴子摘了一堆桃子，它每天吃掉其中的一半然后再多吃一个，直到第10天，它发现只有1个桃子了，问：它第一天摘了多少个桃子？

分析：

    猴子分10天吃桃子，最后一天只剩余1个桃子，要想求出第1天的桃子数，就先要求出第2天的桃子数……设a n 表示第n天的桃子数，则有：

a 1 =（a 2 +1）×2

a 2 =（a 3 +1）×2

…

a 9 =（a 10 +1）×2

a 10 =1

提示 　从以上算式可以看出，可以用递归算法来求解该题。



*/

package main

import (
    "flag"
    "fmt"
)

var (
    dayNum01 int
)

func init() {
    flag.IntVar(&dayNum01, "n", 10, "Usage: 10 2")
}

func main() {
    flag.Parse()

    sum := monkeyEatPeach(dayNum01)
    fmt.Printf("最初的桃子数:%d\n", sum)
}

// 第十天看成是倒数第一天，编号为1，依次为倒数第二天、倒数第三天
func peach01(n int) int {
    if n == 1 {
        return 1 // 倒数第一天（第十天），只有1个桃子
    } else {
        return (peach01(n-1) + 1) * 2
    }
}

func monkeyEatPeach(dayNum int) int {
    return peach01(dayNum)
}
