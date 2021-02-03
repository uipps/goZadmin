// 计算除法，n/m小数点后任意位数（m>n>0的正整数）
//
//        如果是循环小数，打印第一个循环节，用括号
//         1/7 = 0.(142857)
//         1/19 = 0.(052631578947368421)
//         1/95 = 0.(0105263157894736842)

// go run chufa.go -n 1 -m 7 -l 100
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\pai\chufa.go -n 1 -m 7 -l 100
// go run ~/develope/go/go_code/src/github.com/uipps/goZadmin/suanfa/pai/chufa.go -n 201 -m 101 -l 1111

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

    if fenMu1 <= 1 {
        fmt.Println("m分母数据有误，必须是大于1的正整数")
        return
    }

    fmt.Printf("\n  计算 %d/%d 的值， 显示小数点后%d位\n\n", fenZi1, fenMu1, xiaoshuLeng1)
    chuFa01(fenZi1, fenMu1, xiaoshuLeng1)
    fmt.Println("\n")

    fmt.Printf("\n  计算 %d/%d 的值， 记录循环位置\n", fenZi1, fenMu1)
    chuFa02(fenZi1, fenMu1)

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

// 两个正整数相除，如果是循环小数，则记录循环小数的循环节起止位置；使用两个数组进行记录
func chuFa02(fenzi_orig int, fenmu int) {
    fenzi := fenzi_orig
    isXunhuan := false
    iBegin := 0 // 循环节点开始位置
    iEnd := 0   // 循环节点开始位置
    iYouxian := 0
    arr_length := fenmu + 1 // 余数结果数组的长度，
    //arr_length := 100

    // 用数组存放结果
    shang_arr := make([]int, arr_length) // 保存结果，因为不知道循环节的长度，所以后面需要用append
    yushu_arr := make([]int, arr_length) // 保存余数，可能浪费空间!

    for i := 0; fenzi != 0; i++ {
        if i > fenmu {
            shang_arr = append(shang_arr, fenzi/fenmu) // 长度继续增加
        } else {
            shang_arr[i] = fenzi / fenmu
        }

        yushu := fenzi % fenmu // 余数

        // 检查余数是否存在，存在表示一轮循环结束
        if 0 == yushu {
            // 余数为0，说明有限小数
            iYouxian = i
            break
        }
        if (i == 0 && shang_arr[0] == 0) || i > 0 { // i==1之后就是小数部分
            key, in_or_not := common.Int_in_array(yushu_arr, yushu)
            if in_or_not {
                isXunhuan = true
                iBegin = key
                iEnd = i
            } else {
                isXunhuan = false
            }
        }

        // 注入余数数组
        if i > fenmu {
            yushu_arr = append(yushu_arr, yushu) // 数组长度继续增加
        } else {
            yushu_arr[i] = yushu
        }

        if isXunhuan {
            // 循环小数，则提前退出
            break
        }

        fenzi = yushu * 10 // 作为下个循环的分子
    }
    //fmt.Println(yushu_arr)

    // 非循环小数
    if !isXunhuan {
        fmt.Printf(" %d/%d=", fenzi_orig, fenmu)
        for i := 0; i <= iYouxian; i++ {
            if i == 1 {
                fmt.Printf(".") // 小数部分是从1以后开始都放到.后面
            }
            fmt.Printf("%d", shang_arr[i])
        }
    } else {
        // 先输出非循环节部分；循环部分使用()括起来
        fmt.Printf("无限循环小数: %d/%d=%d.", fenzi_orig, fenmu, shang_arr[0])

        for i := 1; i <= iEnd; i++ {
            if iBegin == 0 && i == 1 {
                fmt.Printf("(") // 输出时循环节用 '()' 括起
            }
            fmt.Printf("%d", shang_arr[i])
            if iBegin == i && i >= 1 {
                fmt.Printf("(") // 输出时循环节用 '()' 括起
            }
        }
        fmt.Printf(")") // 输出时循环节用 '()' 括起

        fmt.Printf("从小数点后第%d位开始循环, 第%d位结束，循环节长度%d", iBegin+1, iEnd, iEnd-iBegin)
    }
    fmt.Printf("\n")
}
