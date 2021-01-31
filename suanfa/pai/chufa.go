// 计算除法，1/m小数点后任意位数（m>=2的正整数）
//
// go run chufa.go -n 1 -m 7 -l 100
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\suanfa\pai\chufa.go -n 1 -m 7 -l 100
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\suanfa\pai\chufa.go -n 201 -m 101 -l 1111

package main

import (
    "flag"
    "fmt"
    "github.com/uipps/goZadmin/suanfa/common"
    "time"
)

var (
    fenZi1       int // 分子
    fenMu1       int // 分母
    xiaoshuLeng1 int // 小数点后多少位数
    k_arr1       []int
)

func init() {
    flag.IntVar(&fenZi1, "n", 1, "Usage: 1")
    flag.IntVar(&fenMu1, "m", 7, "Usage: 7 13")
    flag.IntVar(&xiaoshuLeng1, "l", 100, "Usage: 100 1000")
}

func main() {
    flag.Parse()

    startTime := time.Now().UnixNano()
    fmt.Printf("startTime：%d, %s\n", startTime/1e3, time.Unix(0, startTime).Format("2006-01-02 15:04:05"))

    if (fenMu1 <= 1) {
        fmt.Println("m分母数据有误，必须是大于1的正整数")
        return
    }

    fmt.Printf("\n  计算 %d/%d 的值， 显示小数点后%d位\n\n", fenZi1, fenMu1, xiaoshuLeng1)
    chuFa01(fenZi1, fenMu1, xiaoshuLeng1)
    fmt.Println("\n")

    // 执行时间计算
    endTime := time.Now().UnixNano()
    fmt.Printf("  endTime：%d, %s\n", endTime/1e3, time.Unix(0, endTime).Format("2006-01-02 15:04:05"))
    nanoSeconds := float64(endTime-startTime) / 1e3
    fmt.Println("spendTime：", nanoSeconds)
}

// 任意位数的
func chuFa01(fenzi_orig int, fenmu int, xiaoshuLeng int) {
    // xiaoshuLeng += 2
    fenzi := fenzi_orig

    // 初始化，全部置为0， 整数部分占用一位，所以实际申请 xiaoshuLeng+1 位长度
    k_arr1 = make([]int, xiaoshuLeng+1)

    ///// 开始计算
    // 余数乘以10，然后再继续除以除数，不断往后推算即可
    for i := 0; i <= xiaoshuLeng; i++ {
        k_arr1[i] = fenzi / fenmu
        fenzi = fenzi % fenmu * 10
    }
    //fmt.Println(k_arr1)

    // 输出数据，数字太长，因此格式化输出
    common.OutPrintFmt(k_arr1, xiaoshuLeng, 1, fenzi_orig, fenmu)

    return
}
