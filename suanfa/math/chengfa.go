// 计算乘法，n*m, 任意大数或小数，任意小数点位
//
// go run chengfa.go -n 123 -m 987 -l 100
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\math\chengfa.go -n 201 -m 134
// go run ~/develope/go/go_code/src/github.com/uipps/goZadmin/suanfa/math/chengfa.go -n 201 -m 134

package main

import (
    "flag"
    "fmt"
    "github.com/syyongx/php2go"
    "github.com/uipps/goZadmin/suanfa/common"
    "strconv"
    "strings"
    "time"
)

var (
    chengshu1 string // 乘数1
    chengshu2 string // 乘数2
)

func init() {
    flag.StringVar(&chengshu1, "n", "123456", "Usage: 1")
    flag.StringVar(&chengshu2, "m", "87654321", "Usage: 7 13")
}

func forexample01() {
    var a, b, c float64
    a = 1.69
    b = 1.7
    c = a * b
    fmt.Println(c) // 2.8729999999999998

    a = 1.69 * 100
    b = 1.7 * 10
    c = a * b / (100 * 10)

    fmt.Println(c) // 2.873， 正确结果
}

func main() {
    flag.Parse()

    // 参数校验，是否数字(包括小数)
    if (!common.Is_numeric(chengshu1) || !common.Is_numeric(chengshu2)) {
        fmt.Println("乘数必须都是数字!")
        return
    }

    startTime := time.Now().UnixNano()
    fmt.Printf("startTime：%d, %s\n", startTime/1e3, time.Unix(0, startTime).Format("2006-01-02 15:04:05"))

    fmt.Printf("\n  计算 %s x %s 的值 \n\n", chengshu1, chengshu2)
    chengFa01(chengshu1, chengshu2)
    fmt.Println("\n")

    // 执行时间计算
    endTime := time.Now().UnixNano()
    fmt.Printf("  endTime：%d, %s\n", endTime/1e3, time.Unix(0, endTime).Format("2006-01-02 15:04:05"))
    nanoSeconds := float64(endTime-startTime) / 1e3
    fmt.Println("spendTime：", nanoSeconds)
}

// 任意位数乘法, 包括小数相乘
func chengFa01(chengshuA string, chengshuB string) {
    // 1. 确定正负值，顺便去掉负号
    plus_minus_1 := php2go.Strpos(chengshuA, "-", 0)
    plus_minus_2 := php2go.Strpos(chengshuB, "-", 0)
    if (-1 == plus_minus_1) {
        plus_minus_1 = 1 // 不包含，则为正数
    } else {
        // 含有-负号
        plus_minus_1 = -1
        chengshuA = strings.Replace(chengshuA, "-", "", -1) // 去掉负号
    }
    if (-1 == plus_minus_2) {
        plus_minus_2 = 1 // 不包含，则小数位的长度重置为0
    } else {
        // 含有-负号
        plus_minus_2 = -1
        chengshuB = strings.Replace(chengshuB, "-", "", -1) // 去掉负号
    }

    // 2. 确定小数位数，顺便去掉小数点
    xiaoshuLen1 := php2go.Strpos(chengshuA, ".", 0) // 返回-1表示不包含小数点
    xiaoshuLen2 := php2go.Strpos(chengshuB, ".", 0)

    //fmt.Println(xiaoshuLen1)
    //fmt.Println(xiaoshuLen2)

    if (-1 == xiaoshuLen1) {
        xiaoshuLen1 = 0 // 不包含，则小数位的长度重置为0
    } else {
        // 含有小数点，记录小数位长度，然后删除小数点
        xiaoshuLen1 = len(chengshuA) - xiaoshuLen1 - 1      // 字符串总长度减去(整数和小数点位数)，假如小数点后全为0也不影响
        chengshuA = strings.Replace(chengshuA, ".", "", -1) // 去掉小数点
    }
    if (-1 == xiaoshuLen2) {
        xiaoshuLen2 = 0 // 不包含，则小数位的长度重置为0
    } else {
        // 含有小数点，记录小数位长度，然后删除小数点
        xiaoshuLen2 = len(chengshuB) - xiaoshuLen2 - 1
        chengshuB = strings.Replace(chengshuB, ".", "", -1) // 去掉小数点
    }

    // 2. 为方便进位，需将数位反转，由原来的从高到低位 反转成 从低到高位
    chengshuA_r := common.ReverseStr(chengshuA)
    chengshuB_r := common.ReverseStr(chengshuB)
    //chengshuA_r := chengshuA
    //chengshuB_r := chengshuB

    fmt.Println(chengshuA_r)
    fmt.Println(chengshuB_r)

    // 3. 大数相乘，不考虑小数点
    a_length := len(chengshuA_r)
    b_length := len(chengshuB_r)
    c := make([]int, a_length+b_length) // 存放结果的数组，初始每位都是0
    for j := 0; j < b_length; j++ {
        for i := 0; i < a_length; i++ {
            //if (0 == chengshuB_r[j] || 0 == chengshuA_r[i]) {
            //    continue
            //}

            //fmt.Printf(chengshuA_r[i])
            fmt.Printf("%d, %c, %T", chengshuB_r[j], chengshuB_r[j], chengshuB_r[j])
            fmt.Printf(" | %d, %c, %T -- chengji: \n", chengshuA_r[i], chengshuA_r[i], chengshuA_r[i])
            //return

            tm1, _ := strconv.Atoi(string(chengshuB_r[j]))
            tm2, _ := strconv.Atoi(string(chengshuA_r[i]))

            fmt.Printf("tm1: %d\n", tm1)
            fmt.Printf("tm2: %d\n", tm2)

            //ji := chengshuB_r[j] * chengshuA_r[i] // 可能进位了
            ji := tm1 * tm2 // 可能进位了
            fmt.Printf("ji: %d\n", ji)
            //ji := int(tm1) * int(tm2) // 可能进位了
            c[i+j] = c[i+j] + ji%10 // 这里是10进制
            c[i+j+1] = c[i+j+1] + ji/10
        }
    }

    // 4. 对结果加小数点和符号 TODO

    fmt.Println(c)

    // 5. 将结果反转
    for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
        c[i], c[j] = c[j], c[i]
    }

    fmt.Println(c)

    // 6. 删除前导0

    //
    if (plus_minus_1*plus_minus_1 < 0) {
        fmt.Println("-") // 输出负号
    }

    // 输出数据，如果数字太长，则格式化输出
    //common.OutPrintFmt(k_arr1, xiaoshuLeng, 1, fenzi_orig, fenmu)

    return
}
