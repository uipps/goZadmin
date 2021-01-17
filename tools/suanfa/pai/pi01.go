// 计算圆周率π
//  有很多计算方法
//  1.   割圆法：分成很多菱形
//  2.   π/2 = 1+ 1/3 + 1/3*2/5 + 1/3*2/5*3/7 + 1/3*2/5*3/7*4/9 +......
//         通式规律就是： a[0] = 1
//                      a[1] = 1*1/3
//                      a[2] = 1/3*2/5
//                      ......
//                      a[n] = a[n-1]*n/(2n+1)

//  3.   π/2 = 2/1*2/3 * 4/3*4/5 * 6/5*6/7 * .....
//  4.   任意位数的π

// go run pi01.go -n 1111
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\suanfa\pai\pi01.go -n 1111

package main

import (
    "flag"
    "fmt"
    "math"
    "time"
)

var (
    numQieGe   int // 切割次数
    xiaoshuLen int // 小数点后多少位数
    pi_arr     []int
    temp_arr   []int // 临时计算结果
)

func init() {
    flag.IntVar(&numQieGe, "q", 10, "Usage: 1 3")
    flag.IntVar(&xiaoshuLen, "n", 100, "Usage: 100 1000")
}

func main() {
    flag.Parse()

    startTime := time.Now().UnixNano()
    //fmt.Println("startTime：", startTime)

    fmt.Println("\n第四种计算方法，任意位")
    pai04(xiaoshuLen)
    fmt.Println("\n")

    //fmt.Println("\n第一种计算方法")
    //geyuanfa01(numQieGe) // 割圆法
    //fmt.Println("\n第二种计算方法")
    //pai02(1e-16)
    //pai02(1e-100) // 没啥差别

    // 执行时间计算
    endTime := time.Now().UnixNano()
    fmt.Printf("startTime：%d, %s\n", startTime/1e3, time.Unix(0, startTime).Format("2006-01-02 15:04:05"))
    fmt.Printf("  endTime：%d, %s\n", endTime/1e3, time.Unix(0, endTime).Format("2006-01-02 15:04:05"))
    nanoSeconds := float64(endTime-startTime) / 1e3
    fmt.Println("spendTime：", nanoSeconds)
}

// 任意位数的PI π
func pai04(xiaoshuLeng int) {
    xiaoshuLeng += 2 // 增加2位精度

    flag01 := 1
    count := 0

    fenzi := 1
    fenmu := 3

    // 初始化
    for i := 0; i < xiaoshuLeng; i++ {
        pi_arr = append(pi_arr, 0)
        temp_arr = append(temp_arr, 0)
        pi_arr[i] = 0
        temp_arr[i] = 0
    }
    pi_arr[1] = 2
    temp_arr[1] = 2
    //fmt.Println(pi_arr)
    //fmt.Println(temp_arr)

    // 循环计算
    for flag01 > 0 && count < 2147483646 {
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
            result := pi_arr[i] + temp_arr[i] // 将计算结构累加到result中
            pi_arr[i] = result % 10           // 保留一位数
            pi_arr[i-1] += result / 10        // 向高位进位(i越小是高位)
            flag01 |= temp_arr[i]             // 若temp中的数全部为0，退出循环
        }
        count++    // 记录大圈循环次数
        fenzi++    // 累加分子
        fenmu += 2 // 累加分母
    }
    fmt.Println(pi_arr)
    fmt.Println(temp_arr) // 最后全部0

    // 输出数据，数字太长，因此格式化输出
    fmt.Printf("\n计算了%d次\n", count)
    fmt.Printf("\t---第1-1000位小数---\n")
    fmt.Printf("PI=%d.", pi_arr[1])
    // 小数部分要循环输出
    for i := 2; i < xiaoshuLeng; i++ {
        if i > 2 && (i-2)%10 == 0 { // 每十位输入一个空格
            fmt.Print(" ")
        }
        if i > 2 && (i-2)%50 == 0 { // 每50位换行
            fmt.Println("")
        }
        if i > 2 && (i-2)%1000 == 0 { // 每1000位, 显示一个提示
            fmt.Printf("\t---显示第%d-%d位小数---\n", (i-2)/1000*1000+1, ((i-2)/1000+1)*1000)
        }
        fmt.Printf("%d", pi_arr[i]) // 输出一位小数
    }

    return
}

// 公式2计算方式, 两边边都乘以2，π = 2+ 2/3 + 2/3*2/5 + 2/3*2/5*3/7 + 2/3*2/5*3/7*4/9 +......
func pai02(jindu float64) float64 {
    pi := 2.0
    temp := 2.0

    fenzi := 1.0
    fenmu := 3.0

    for temp > jindu {
        temp = temp * fenzi / fenmu
        pi += temp
        fenzi++
        fenmu += 2
    }

    fmt.Printf("fenzi: %4.f, PI=%.18f\n", fenzi, pi)
    return pi
}

// 割圆法,num越大，值越准
func geyuanfa01(num int) {
    i := 0 // 切割次数，初始值
    //numQieGe       // 切割次数

    k := 3.0
    y2 := 1.0 // 内接6边形的边长
    s := 6    // 初始内接多边形的边数

    for i < num {
        fmt.Printf("第%d次切割，为%d边形，PI=%.24f\n", i, s, k*math.Sqrt(y2))
        s *= 2
        y2 = 2 - math.Sqrt(4-y2) // 玄长
        i++
        k *= 2.0
    }

    return
}
