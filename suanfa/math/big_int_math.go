/*
// 大整数的加减乘除
//  TODO 大数加减还未实现小数部分

go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\math\big_int_math.go -n 99988 -m 999
go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\math\big_int_math.go -n 999 -m 99988 -o sub
go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\math\big_int_math.go -n 99988 -m 999 -o sub
go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\math\big_int_math.go -n 999888 -m 999 -o sub
go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\math\big_int_math.go -n 1000888 -m 999 -o sub

*/
package main

import (
    "flag"
    "fmt"
    "github.com/uipps/goZadmin/suanfa/common"
    "math"
    "strings"
)

var (
    bcN      string // 参与运算数据之一
    bcM      string // 参与运算数据之二
    bcOption string // 何种运算, 加减乘除： + - * / , 暂时只支持加减，乘除有专门的程序！
)

func init() {
    flag.StringVar(&bcN, "n", "100000000000000000000", "Usage: 9223372036854775807 10000000000000000000")
    flag.StringVar(&bcM, "m", "129223372036854775807", "Usage: 9223372036854775807 1")
    flag.StringVar(&bcOption, "o", "add", "Usage: add sub mul div")
}

func main() {
    flag.Parse()
    rlt := ""

    if (!common.Is_numeric(bcN)) {
        fmt.Printf("-n 数据必须是数字\n")
        return
    }

    if (!common.Is_numeric(bcM)) {
        fmt.Printf("-m 数据必须是数字\n")
        return
    }
    // TODO 包括小数点

    bc_method := strings.ToLower(bcOption) // 转化为小写
    // 只处理加减，
    if "sub" == bc_method {
        // 大数的加减法
        rlt = bcSub(bcN, bcM)
        fmt.Printf("\n  大数减法计算，计算结果： %s - %s = %s \n\n", bcN, bcM, rlt)

        fmt.Printf("\n  使用PHP的bcsub方法，结果为：\n")
        command_str := fmt.Sprintf("echo bcsub('%s', '%s');", bcN, bcM)
        rlt = common.ExecPHP(command_str)
        fmt.Println(rlt)
        fmt.Println("\n")
    } else {
        rlt = bcAdd(bcN, bcM)
        fmt.Printf("\n  大数加法计算，结果为： %s + %s = %s \n\n", bcN, bcM, rlt)

        fmt.Printf("\n  使用PHP的bcadd方法，结果为：\n")
        command_str := fmt.Sprintf("echo bcadd('%s', '%s');", bcN, bcM)
        rlt = common.ExecPHP(command_str)
        fmt.Println(rlt)
        fmt.Println("\n")
    }
}

// TODO 暂不包含小数点
func bcAdd(a string, b string) string {
    // 1. 为了后面程序方便，保持长数放在前面a，短数在后b
    a_len := len(a)
    b_len := len(b)
    len_max_float := math.Max(float64(a_len), float64(b_len)) // 最长的那个
    len_max := int(len_max_float)                             // 类型转换
    if (len_max != a_len) {
        a, b = b, a // 交换一下，保持a为长，b为短
    }
    b_len = len(b)

    // 2. 再将2个数字反转
    a_r := common.ReverseStr(a)
    b_r := common.ReverseStr(b)

    // 3. 逐个位数进行加法，顺便处理进位
    c := make([]byte, len_max+1) // 存放计算结果，按照最长的，顶多进1位
    for i := 0; i < len_max; i++ {
        var rlt_c byte
        if (i < b_len) {
            rlt_c = c[i] + (a_r[i] - '0') + (b_r[i] - '0')
        } else {
            rlt_c = c[i] + (a_r[i] - '0')
        }
        c[i] = rlt_c % 10
        c[i+1] = c[i+1] + rlt_c/10 // 顺便处理最后一位的进位
    }
    // 上面已经处理过最后一位的进位问题

    // 4. 对结果反转，拼装字符串，顺便删除前导0
    // 这样得到的是乱码
    //fmt.Println(c)
    //fmt.Println(string(c))
    //rlt := common.ReverseStr(string(c))
    //fmt.Println(rlt)
    rlt := ""
    for i := len_max; i >= 0; i-- {
        if (i == len_max && c[len_max] == 0) {
            continue
        }
        rlt += string(c[i] + '0')
    }

    return rlt
}

// TODO 暂不支持小数
func bcSub(a string, b string) string {
    // 1. 比较长度，确定结果的符号；数据重排，大数在前，小数在后
    a_orig := a
    a, b = sortTwoNum(a, b) // 将大的放前面，小的放后面

    plus_minus := 1         // 结果符号，默认正号+
    if (a != a_orig) {
        plus_minus = -1
    }

    // 2. 再将2个数字反转
    a_r := common.ReverseStr(a)
    b_r := common.ReverseStr(b)
    a_len := len(a)
    b_len := len(b)

    // 3. 逐个位数进行减法，顺便处理借位
    c := make([]int8, a_len) // 存放计算结果，按照最长的
    for i := 0; i < a_len; i++ {
        if (i < b_len) {
            if (int8(a_r[i] - '0') + c[i] < int8(b_r[i] - '0')) {
                // 需要借位
                c[i+1] = c[i+1] - 1
                c[i] = int8(a_r[i] - '0') + c[i] + 10 - int8(b_r[i] - '0')
            } else {
                // 无需借位
                c[i] = int8(a_r[i] - '0') + c[i] - int8(b_r[i] - '0')
            }
        } else {
            if (int8(a_r[i] - '0') + c[i] < 0) {
                // 需要借位
                c[i+1] = c[i+1] - 1
                c[i] = int8(a_r[i] - '0') + c[i] + 10
            } else {
                c[i] = int8(a_r[i] - '0') + c[i]
            }
        }
    }

    // 4. 对结果反转，拼装字符串，顺便删除前导0
    rlt := ""
    for i := a_len - 1; i >= 0; i-- {
       if (i == a_len - 1 && c[a_len - 1] == 0) {
           continue
       }
       rlt += string(c[i] + '0')
    }

    if plus_minus < 0 {
        rlt = "-" + rlt
    }

    return rlt
}

// 两个正整数排序，哪个大，哪个在前面
func sortTwoNum(a, b string) (string, string) {
    a_len := len(a)
    b_len := len(b)

    if (a_len > b_len) {
        return a, b
    } else if (a_len == b_len) {
        // 长度相等的情况，逐个比较大小
        for i := 0; i < a_len; i++ {
            // 从高位开始逐个比较
            if (a[i] > b[i]) {
                return a, b
            } else if (a[i] < b[i]) {
                return b, a
            }
        }
        // 相等的情况
        return a, b
    } else {
        // b长的情况
        return b, a
    }
}
