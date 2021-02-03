// 计算圆周率π
//  有很多计算方法
//  1.   割圆法：分成很多菱形
//  2.   π/2 = 1+ 1/3 + 1/3*2/5 + 1/3*2/5*3/7 + 1/3*2/5*3/7*4/9 +......
//         通式规律就是： a[0] = 1
//                      a[1] = 1*1/3
//                      a[2] = 1/3*2/5
//                      ......
//                      a[n] = a[n-1]*n/(2n+1)
//        两边乘以2得：π = 2 + 2/3 + 2/3*2/5 + 2/3*2/5*3/7 + 2/3*2/5*3/7*4/9 +......

//  3.   π/2 = 2/1*2/3 * 4/3*4/5 * 6/5*6/7 * .....
//  4.   任意位数的π，依然利用上面2的公式进行任意位数的计算
//  5.   π/4 = 1- 1/3 + 1/5 - 1/7 + 1/9 - 1/11......   (pi02.go里面有计算)

// go run pi01.go -n 1111
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\pai\pi01.go -n 1111
// go run ~/develope/go/go_code/src/github.com/uipps/goZadmin/suanfa/pai/pi01.go -n 1111

package main

import (
    "flag"
    "fmt"
    "math"
    "runtime"
    "time"
    "github.com/uipps/goZadmin/suanfa/common"
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
    fmt.Printf("startTime：%d, %s\n", startTime/1e3, time.Unix(0, startTime).Format("2006-01-02 15:04:05"))

    fmt.Println("\n第四种计算方法，任意位")
    pai04(xiaoshuLen)
    fmt.Println("\n")

    fmt.Println("\n第一种计算方法，割圆法")
    geyuanfa01(numQieGe) // 割圆法
    fmt.Println("\n第二种计算方法")
    pai02(1e-16)
    pai02(1e-100) // 没啥差别

    // 执行时间计算
    endTime := time.Now().UnixNano()
    fmt.Printf("  endTime：%d, %s\n", endTime/1e3, time.Unix(0, endTime).Format("2006-01-02 15:04:05"))
    nanoSeconds := float64(endTime-startTime) / 1e3
    fmt.Println("spendTime：", nanoSeconds)
}

// 任意位数的PI π
func pai04(xiaoshuLeng int) {
    xiaoshuLeng += 2 // 十位个位占用2个; 这里也可以是10，最小是2，越大最后的数约精确

    flag01 := 1
    count := 0

    fenzi := 1
    fenmu := 3

    // 初始化
    pi_arr := make([]int, xiaoshuLeng)
    temp_arr := make([]int, xiaoshuLeng)
    pi_arr[1] = 2
    temp_arr[1] = 2
    //fmt.Println(pi_arr)
    //fmt.Println(temp_arr)
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

            //fmt.Printf("\n      --- Line:%3d, count:%d , i: %d --- \n", line+12, count, i)
            //fmt.Println(result)
            //fmt.Printf("temp_arr: ")
            //fmt.Println(temp_arr)
            //fmt.Println(carry)
        }

        carry = 0
        for i := 0; i < xiaoshuLeng; i++ { // 从高位到低位相除
            result := temp_arr[i] + carry*10 // 当前加上前一位的余数
            temp_arr[i] = result / fenmu     // 当前位的整数部分
            carry = result % fenmu           // 当前位的余数，累加到下一位的运算

            //fmt.Printf("\n      --- Line: %3d, count:%d , i: %d --- \n", line+25, count, i)
            //fmt.Println(result)
            //fmt.Printf("temp_arr: ")
            //fmt.Println(temp_arr)
            //fmt.Println(carry)
        }

        flag01 = 0                             // 清除标记
        for i := xiaoshuLeng - 1; i > 0; i-- { // 从低位到高位, 将计算结果累加(i越小是高位，i越大是低位)
            result := pi_arr[i] + temp_arr[i] // 将计算结果累加到result中
            pi_arr[i] = result % 10           // 保留一位数
            pi_arr[i-1] += result / 10        // 向高位进位(i越小是高位)
            flag01 |= temp_arr[i]             // 若temp中的数全部为0，退出循环

            //fmt.Printf("\n      --- Line: %3d, count:%d , i: %d --- \n", line+39, count, i)
            //fmt.Println(result)
            //fmt.Printf("pi_arr: ")
            //fmt.Println(pi_arr)
            //fmt.Printf("temp_arr: ")
            //fmt.Println(temp_arr)
            //fmt.Println(flag01)
        }
        count++    // 记录大圈循环次数
        fenzi++    // 累加分子
        fenmu += 2 // 累加分母
    }
    //fmt.Println(pi_arr)
    //fmt.Println(temp_arr) // 最后全部0

    // 输出数据，数字太长，因此格式化输出
    fmt.Printf("\n计算了%d次\n", count)
    common.OutPrintFmt(pi_arr, xiaoshuLen, 2, 3, 1)

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

// 割圆法,num越大，值越准，初始是6边形；第一次切割变成12边形；第二次切割变成24边形
/*
 弦长的计算通过半角公式结合弦长xian = 2r*sinA (r半径为1) 推导：

        xian = 2*sinA       (1) // 初始弦长
    xian_new = 2*sin(A/2)   (2) // 切割一次后的弦长
  sin^2(A/2) = (1-cosA)/2   (3)
cosA = sqrt(1 - (sinA)^2)   (4)

  A) 推导过程
xian_new 和 xian的关系可以推导出来，将(1),(2)分别平方,
    xian^2 = 4 * sin^2A     (5)
xian_new^2 = 4 * sin^2(A/2) (6)

将(3)带入(6)得：
xian_new^2 = 4 * (1-cosA)/2 = 2 * (1-cosA) = 2 - 2cosA      (7)
将(4),(5)分别代入上式(7)
最后得出，xian_new平方 和 xian平方的关系：
xian_new^2 = 2 - sqrt(4 - xian^2)

因为涉及到开方运算，因此精度有限。而且切割25次以后，反而不准了,xian_square为0
*/
func geyuanfa01(num int) {
    i := 0 // 切割次数，初始值

    s := 6.0           // 初始内接多边形的边数，初始6边形，以后逐次翻倍
    xian_square := 1.0 // 注意：该变量是弦长的平方，初始6边形的弦长(等于半径长度1)，1的平方等于1，弦长以后逐渐减少
    for i < num {
        fmt.Printf("第%2d次切割，为%d边形，PI=%.24f\n", i, int(s), s/2*math.Sqrt(xian_square))
        s *= 2.0
        xian_square = 2 - math.Sqrt(4-xian_square) // 切割一次，弦长变短，利用三角函数sin(A/2)
        i++                                        // 切割次数，逐次自增
    }

    return
}
