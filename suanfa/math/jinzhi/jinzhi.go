// 任意进制之间的转换： argB2 只能是2~76，argB2==1则一直循环，

// go run jinzhi.go -b 8 -n 10
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\math\jinzhi\jinzhi.go -b 2 -n 16

package main

import (
    "flag"
    "fmt"
    "github.com/uipps/goZadmin/suanfa/common/jinzhiToAny"
)

var (
    argN2 int
    argB2 int // 需要展示的进制
)

func init() {
    flag.IntVar(&argN2, "n", 100, "Usage: 10 100 13")
    flag.IntVar(&argB2, "b", 16, "Usage: 2 8 16")
}

func main() {
    flag.Parse()

    // argB2 只能是2~76，argB2==1则一直循环，
    // 参数校验
    if argB2 < 2 || argB2 > 76 {
        fmt.Printf("转换成的进制base：%2d, 只能是2~76之间的正整数!\n", argB2)
    }
    signStr := "" // 正负号
    if argN2 < 0 {
        signStr = "-"
        argN2 = -1 * argN2 // 转成正数进行计算
    }

    str := ""

    jinzhiToAny.DecimalToAnyDigui(&str, argN2, argB2)
    fmt.Printf("十进制数据:%9d, 转换成的进制base：%2d, 结果是: %s ------ (%s)\n", argN2, argB2, signStr+str, "DecimalToAnyDigui")

    //println("  按照另外的方法计算的数值")
    fmt.Printf("十进制数据:%9d, 转换成的进制base：%2d, 结果是: %s ------ (D十进制转换)\n", argN2, argB2, signStr+jinzhiToAny.D十进制转换(argN2, argB2))
    fmt.Printf("十进制数据:%9d, 转换成的进制base：%2d, 结果是: %s ------ (DecimalToAny)\n", argN2, argB2, signStr+jinzhiToAny.DecimalToAny(argN2, argB2))

    return
}
