// 大整数的加减乘除
//  TODO 乘除已经实现，加减还未实现
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\math\big_int_math.go -m 3 -n 41

package main

import (
    "flag"
    "fmt"
    "github.com/uipps/goZadmin/suanfa/common"
    "math"
    "strings"
)

type bigint struct {
    minus  string // 长整数数组
    num    string // 符号，正还是负数
    length int    // 保存该数的位数
}

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
    len_min := len(b)

    // 2. 再将2个数字反转
    a_r := common.ReverseStr(a)
    b_r := common.ReverseStr(b)

    // 3. 逐个位数进行加法，顺便处理进位
    c := make([]byte, len_max+1) // 存放计算结果，按照最长的，顶多进1位
    for i := 0; i < len_max; i++ {
        var rlt_c byte
        if (i < len_min) {
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

func bcSub(n string, m string) string {
    rlt := ""

    return rlt
}
