// 计算圆周率π
//  有很多计算方法, 任意位 (此方法收敛太慢，不合适)
//  5.   π/4 = 1- 1/3 + 1/5 - 1/7 + 1/9 - 1/11......
//         通式规律就是： a[0] = 1
//                      a[1] = pow((-1), 1)*1/3
//                      a[2] = pow((-1), 2)*1/5
//                      ......
//                      a[n] = pow((-1), n)*1/(2n+1)
//      也可简化为：
//       π/4 = 2/(1*3) + 2/(5*7) + 2/(9*11) + 2/(13*15) ......
//         π = 8/(1*3) + 8/(5*7) + 8/(9*11) + 8/(13*15) ...... 就按照这个计算好了
//
// go run pi02.go -n 1111
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\pai\pi02.go -n 1111
// go run ~/develope/go/go_code/src/github.com/uipps/goZadmin/suanfa/pai/pi02.go -n 1111 -c 1000000

package main

import (
	"flag"
	"fmt"
	"github.com/uipps/goZadmin/suanfa/common"
	"time"
)

var (
	xiaoshuLen2 int // 小数点后多少位数
	xunHuanCi2 int 	// 分母n(n+2)的n值，也是多少项
)

func init() {
	flag.IntVar(&xiaoshuLen2, "n", 100, "Usage: 100 1000")
	flag.IntVar(&xunHuanCi2, "c", 1000000, "Usage: 100 1000")
}

func main() {
	flag.Parse()

	startTime := time.Now().UnixNano()
	fmt.Printf("startTime：%d, %s\n", startTime/1e3, time.Unix(0, startTime).Format("2006-01-02 15:04:05"))

	fmt.Println("\n第五种计算方法，任意位")
	pai05(xiaoshuLen2)
	fmt.Println("\n")

	//fmt.Println("\n第5-2种计算方法，精度有限，不推荐")
	//pai05_2()
	//fmt.Println("\n")

	// 执行时间计算
	endTime := time.Now().UnixNano()
	fmt.Printf("  endTime：%d, %s\n", endTime/1e3, time.Unix(0, endTime).Format("2006-01-02 15:04:05"))
	nanoSeconds := float64(endTime-startTime) / 1e3
	fmt.Println("spendTime：", nanoSeconds)
}

// 任意位数的PI π , 计算公式π = 8/(1*3) + 8/(5*7) + 8/(9*11) + 8/(13*15) ......
func pai05(xiaoshuLeng int) {
	xiaoshuLeng += 2 // 十位个位占用2个; 这里也可以是10，最小是2，越大最后的数约精确

	pi_arr := make([]int, xiaoshuLeng)   // 存放结果
	temp_arr := make([]int, xiaoshuLeng) // 每项计算结果：a[i] = 8/(n^n-2n)
	pi_arr[0] = 2	// 个位初始值

	count := 0 // 循环次数计数

	fenzi := 8 // 分子初始值
	fenmu := 3 // 分母初始值

	// 循环计算
	flag01 := 1 // 用于提前退出循环
	for flag01 > 0 && count < xunHuanCi2 {
		// 计算每项的值，采用任意位数组的方式存放到temp_arr, 如8/3，则temp_arr[0] = 2,temp_arr[1]=6.....
		fenzi = 8
		fen_mu_new := fenmu * (fenmu - 2)  // 分母通式：n(n-2), 初始n=3，以后逐渐+4； 暂未超过整数最大范围2^63, 改进为2次除耗时更长
		//fen_mu_new := fenmu
		for i := 0; i < xiaoshuLeng; i++ { // 从高位到低位，记录每项的整数和小数值（小数点位数到xiaoshuLeng位）
			temp_arr[i] = fenzi / fen_mu_new
			fenzi = (fenzi % fen_mu_new) *10
		}
		// 在上次基础上，除以fenmu - 2 ; 两次相除耗时更长，还是用上面一次相除
		//fen_mu_new = fenmu - 2
		//yushu := 0
		//for i := 0; i < xiaoshuLeng; i++ {
		//    fenzi = temp_arr[i] + yushu	* 10		// 对上次结果的数组进行除法运算，
		//	temp_arr[i] = fenzi / fen_mu_new
		//	yushu = (fenzi % fen_mu_new)
		//}
		// temp_arr[i]可能全部为0，当fenmu足够大的时候

		flag01 = 0                             // 清除标记
		for i := xiaoshuLeng - 1; i > 0; i-- { // 从低位到高位, 将计算结果累加(i越小是高位，i越大是低位)
			result := pi_arr[i] + temp_arr[i] // 将计算结果累加到pi_arr中，对应位相加，可能有进位
			pi_arr[i] = result % 10           // 保留一位数
			pi_arr[i-1] += result / 10        // 向高位进位(i越小是高位)
			flag01 |= temp_arr[i]             // 若temp中的数全部为0，退出循环
		}
		count++    // 记录大圈循环次数
		fenmu += 4 // 累加分母
	}
	//fmt.Println(pi_arr)
	//fmt.Println(temp_arr) // 最后全部0

	// 输出数据，数字太长，因此格式化输出
	fmt.Printf("\n计算了%d次\n", count)
	common.OutPrintFmt(pi_arr, xiaoshuLen2, 1, 3, 1)

	return
}

// 直接累加，但是精度有限，不推荐
func pai05_2() {
	var (
		f float64
		n float64
		p float64
	)
	f = 1.0
	n = 1.0
	p = 0.0

	for n < 10000000 {
		p = p + 4.0*f/n
		f = -1 * f
		n += 2
	}
	fmt.Printf("%.40f\n", p)

	return
}
