// 大整数的加减乘除
// 约瑟夫环：
//  41个人围成一个圆圈，由第一个人开始报数，报到3的人就必须自杀，再由下一个人开始报数，直到所有人都自杀为止。
//      约瑟夫和他的朋友安排在16和31的位置，最后就剩下他们两人，不符合游戏规则，所以逃过了该死亡游戏
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\suanfa\big_int_math.go -m 3 -n 41

package main

import (
    "flag"
    "fmt"
    "strconv"
)

type bigint struct {
    minus string        // 长整数数组
    num string        // 符号，正还是负数
    length int        // 保存该数的位数
}

var (
    bcN string // 参与运算数据之一
    bcM string // 参与运算数据之二
)

func init() {
    flag.StringVar(&bcN, "n", "100000000000000000000", "Usage: 9223372036854775807 10000000000000000000")
    flag.StringVar(&bcM, "m", "129223372036854775807", "Usage: 9223372036854775807 1")
}

func main() {
    flag.Parse()
    rlt := ""

    if (!IsNum(bcN)) {
        fmt.Printf("-n 数据必须是数字\n")
        return
    }

    if (!IsNum(bcM)) {
        fmt.Printf("-m 数据必须是数字\n")
        return
    }

    // 大数的加减乘除
    rlt = bcAdd(bcN, bcM)
    fmt.Printf(" %s + %s = %s \n", bcN, bcM, rlt)
}

func IsNum(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}

func bcAdd(n string, m string) string {
    rlt := ""



    return rlt
}

// 整理大整数，去掉前面多余的0
func BigIntTrim(num1 bigint) {
    //
}

func BigIntTrans(num1 bigint) {
    str1 = ""
    return str1
}