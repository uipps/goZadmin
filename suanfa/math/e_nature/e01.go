// 计算自然对数e
//  有很多计算方法
//  1.   e = 1+ 1/2! + 1/3! + 1/4! + 1/5! + ......
//         通式规律就是： a[0] = 1
//                      a[1] = 1*1/2
//                      a[2] = 1*1/2*1/3
//                      ......
//                      a[n] = a[n-1]*1/(n+1)
//
// go run e01.go -n 1111
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\math\e_nature\e01.go -n 1111
// go run ~/develope/go/go_code/src/github.com/uipps/goZadmin/suanfa/math/e_nature/e01.go -n 1111

package main

import (
    "flag"
    "fmt"
    "github.com/uipps/goZadmin/suanfa/common"
    "runtime"
    "time"
)

var (
    xiaoshuLen int // 小数点后多少位数
    rlt_arr    []int
)

func init() {
    flag.IntVar(&xiaoshuLen, "n", 100, "Usage: 100 1000")
}

func main() {
    flag.Parse()

    startTime := time.Now().UnixNano()
    fmt.Printf("startTime：%d, %s\n", startTime/1e3, time.Unix(0, startTime).Format("2006-01-02 15:04:05"))

    fmt.Println("\n第一种计算方法")
    e01(xiaoshuLen)
    fmt.Println("\n")

    // 执行时间计算
    endTime := time.Now().UnixNano()
    fmt.Printf("  endTime：%d, %s\n", endTime/1e3, time.Unix(0, endTime).Format("2006-01-02 15:04:05"))
    nanoSeconds := float64(endTime-startTime) / 1e3
    fmt.Println("spendTime：", nanoSeconds)
}

// 任意位数的自然对数e
func e01(xiaoshuLeng int) {
    xiaoshuLeng += 2 // 十位个位占用2个; 这里也可以是10，最小是2，越大最后的数约精确
    flag01 := 1
    count := 0

    fenzi := 1
    fenmu := 1

    // 初始化
    e_arr := make([]int, xiaoshuLeng)
    temp_arr := make([]int, xiaoshuLeng)
    e_arr[1] = 1
    temp_arr[1] = 1

    _, _, line, _ := runtime.Caller(0)
    fmt.Printf("\n ---------- Line: %3d, 初始状态，count:%4d , fenzi: %3d , fenmu: %3d ---------- \n", line+1, count, fenzi, fenmu)

    // 循环计算
    for flag01 > 0 && count < 2147483646 {
        //fmt.Printf("\n\n\n   ------ Line: %3d, count:%4d , fenzi: %3d , fenmu: %3d ------ \n", line+5, count, fenzi, fenmu)
        carry := 0
        for i := xiaoshuLeng - 1; i > 0; i-- { // 从低位到高位相乘
            result := temp_arr[i]*fenzi + carry // 用每一位去乘，再加上进位
            temp_arr[i] = result % 10           // 保存个数
            carry = result / 10                 // 进位
        }

        carry = 0
        for i := 0; i < xiaoshuLeng; i++ { // 从高位到低位相除
            result := temp_arr[i] + carry*10 // 当前加上前一位的余数
            temp_arr[i] = result / fenmu     // 当前位的整数部分
            carry = result % fenmu           // 当前位的余数，累加到下一位的运算
        }

        flag01 = 0                             // 清除标记
        for i := xiaoshuLeng - 1; i > 0; i-- { // 从低位到高位, 将计算结果累加(i越小是高位，i越大是低位)
            result := e_arr[i] + temp_arr[i] // 将计算结果累加到result中
            e_arr[i] = result % 10           // 保留一位数
            e_arr[i-1] += result / 10        // 向高位进位(i越小是高位)
            flag01 |= temp_arr[i]            // 若temp中的数全部为0，退出循环
        }
        count++    // 记录大圈循环次数
        fenmu += 1 // 累加分母
    }

    // 输出数据，数字太长，因此格式化输出
    fmt.Printf("\n计算了%d次\n", count)
    common.OutPrintFmt(e_arr, xiaoshuLen, 2, 2, 1)

    return
}
