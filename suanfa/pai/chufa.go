// 计算除法，1/m小数点后任意位数（m>=2的正整数）
//
// go run chufa.go -n 1 -m 7 -l 100
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\suanfa\pai\chufa.go -n 1 -m 7 -l 100
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\suanfa\pai\chufa.go -n 201 -m 101 -l 1111

package main

import (
    "flag"
    "fmt"
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

    // 初始化，全部置为0， 增加2位精度
    for i := 0; i < xiaoshuLeng+2; i++ {
        k_arr1 = append(k_arr1, 0)
    }

    ///// 开始计算
    // 余数乘以10，然后再继续除以除数，不断往后推算即可
    for i := 0; i <= xiaoshuLeng; i++ {
        k_arr1[i] = fenzi / fenmu
        fenzi = fenzi % fenmu * 10
    }

    // 输出数据，数字太长，因此格式化输出
    fmt.Printf("\t---第1-1000位小数---\n")
    fmt.Printf("%d/%d=\n%d.", fenzi_orig, fenmu, k_arr1[0])
    // 小数部分要循环输出
    n := 1 // 小数点开始的序号
    for i := n; i < xiaoshuLeng; i++ {
        if i > n && (i-n)%10 == 0 { // 每十位输入一个空格
            fmt.Print(" ")
        }
        if i > n && (i-n)%50 == 0 { // 每50位换行
            fmt.Println("")
        }
        if i > n && (i-n)%1000 == 0 { // 每1000位, 显示一个提示
            fmt.Printf("\t---显示第%d-%d位小数---\n", (i-n)+1, i-n+1000)
        }
        fmt.Printf("%d", k_arr1[i]) // 输出一位小数
    }

    return
}
