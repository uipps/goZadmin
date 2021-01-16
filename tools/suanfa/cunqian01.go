// 父亲为小龙的4年大学生活一次性在银行储蓄一笔钱，使用整存零取的方式，控制小龙每月月底子还能提取1000，
// 假设银行一年整存零取的年利息为1.71% , 请问编程计算父亲至少需要一次性存款多少
//  4年就是48个月，倒推的方式进行计算
// go run cunqian01.go -f 1000 -r 0.0171 -m 48
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\suanfa\cunqian01.go -f 1000 -r 0.0171 -m 2
package main

import (
	"flag"
	"fmt"
)

var (
	fetch01            float64 // 本金
	month01            int64   // 贷款月数，20年就是240月
	interestRateYear01 float64 // 贷款年利率
	continueDo01       int     // 继续计算还是退出 1-继续 0-退出
)

func init() {
	flag.Float64Var(&fetch01, "f", 1000, "Usage: 1000 1500")
	flag.Int64Var(&month01, "m", 48, "Usage: 12 48")
	flag.Float64Var(&interestRateYear01, "r", 0.0171, "Usage: 0.05")
}

func main() {
	flag.Parse()

	fangshi01(fetch01, interestRateYear01, month01)

	//continueDo01 = 1
	// 继续还是退出
	fmt.Println("\n")
	fmt.Println("继续计算还是退出？输入1表示继续，输入其他数字或字符表示退出")
	_, err := fmt.Scan(&continueDo01)
	if err != nil {
		fmt.Println("输入有误！\n")
	}

	// do-while循环
	for {
		// 检测全局变量的值
		if 1 != continueDo01 {
			//os.Exit(0)
			break
		}
		fmt.Println("\n\n")
		scanData01()
		fmt.Println("\n\n")
	}

	return
}

func fangshi01(fetch01 float64, interestRateYear01 float64, month01 int64) {

	var money = make([]float64, month01+1) // 切片，变长数组
	money[month01] = fetch01

	for i := month01 - 1; i > 0; i-- {
		//money[i] = (money[i+1] + fetch01) / (1 + interestRateYear01/12)	// 书上的算法，我感觉有误，应该用下面的
		money[i] = (money[i+1])/(1+interestRateYear01/12) + fetch01 // 自己觉得这样才是对的
	}

	for i := month01; i > 0; i-- {
		fmt.Printf("第%2d个月末本利合计： %9.2f\n", i, money[i])
	}
}

func scanData01() {
	var (
		lFetch            float64
		lMonth            int64
		lInterestRateYear float64
	)

	// 循环判断
	for lFetch <= 0 {
		fmt.Println("\n")
		// 请输入每月提取金额
		fmt.Println("请输入每月提取金额")
		fmt.Scan(&lFetch) // 字符串被强制转成了float64，其值为0
	}

	for lMonth <= 0 {
		// 请输入月数，如4年就是48个月
		fmt.Println("\n")
		fmt.Println("请输入月数")
		fmt.Scan(&lMonth)
	}

	for lInterestRateYear <= 0 || lInterestRateYear >= 1 {
		// 请输入存款年利率
		fmt.Println("\n")
		fmt.Println("请输入存款年利率，在0~1之间，例如2%，则输入0.02")
		fmt.Scan(&lInterestRateYear)
	}

	fangshi01(lFetch, lInterestRateYear, lMonth)

	fmt.Println("\n")
	fmt.Println("继续计算还是退出？输入1表示继续，输入其他数字或字符表示退出")
	_, err := fmt.Scan(&continueDo01)
	if err != nil {
		fmt.Println("输入有误！\n")
	}
}

// 1. 书上的计算方法 money[i] = (money[i+1] + fetch01) / (1 + interestRateYear01/12)
/*
第48个月末本利合计：   1000.00
第47个月末本利合计：   1997.15
第46个月末本利合计：   2992.89
第45个月末本利合计：   3987.21
第44个月末本利合计：   4980.11
第43个月末本利合计：   5971.60
第42个月末本利合计：   6961.68
第41个月末本利合计：   7950.35
第40个月末本利合计：   8937.62
第39个月末本利合计：   9923.47
第38个月末本利合计：  10907.93
第37个月末本利合计：  11890.99
第36个月末本利合计：  12872.64
第35个月末本利合计：  13852.90
第34个月末本利合计：  14831.77
第33个月末本利合计：  15809.24
第32个月末本利合计：  16785.32
第31个月末本利合计：  17760.01
第30个月末本利合计：  18733.32
第29个月末本利合计：  19705.24
第28个月末本利合计：  20675.77
第27个月末本利合计：  21644.93
第26个月末本利合计：  22612.71
第25个月末本利合计：  23579.11
第24个月末本利合计：  24544.13
第23个月末本利合计：  25507.78
第22个月末本利合计：  26470.06
第21个月末本利合计：  27430.97
第20个月末本利合计：  28390.52
第19个月末本利合计：  29348.70
第18个月末本利合计：  30305.51
第17个月末本利合计：  31260.96
第16个月末本利合计：  32215.06
第15个月末本利合计：  33167.79
第14个月末本利合计：  34119.17
第13个月末本利合计：  35069.20
第12个月末本利合计：  36017.87
第11个月末本利合计：  36965.20
第10个月末本利合计：  37911.17
第 9个月末本利合计：  38855.81
第 8个月末本利合计：  39799.09
第 7个月末本利合计：  40741.04
第 6个月末本利合计：  41681.64
第 5个月末本利合计：  42620.90
第 4个月末本利合计：  43558.83
第 3个月末本利合计：  44495.43
第 2个月末本利合计：  45430.69
第 1个月末本利合计：  46364.62

*/

// 2. 自己的算法 money[i] = (money[i+1]) / (1 + interestRateYear01/12) + fetch01
/*
第48个月末本利合计：   1000.00
第47个月末本利合计：   1998.58
第46个月末本利合计：   2995.73
第45个月末本利合计：   3991.47
第44个月末本利合计：   4985.79
第43个月末本利合计：   5978.70
第42个月末本利合计：   6970.19
第41个月末本利合计：   7960.27
第40个月末本利合计：   8948.94
第39个月末本利合计：   9936.21
第38个月末本利合计：  10922.07
第37个月末本利合计：  11906.53
第36个月末本利合计：  12889.59
第35个月末本利合计：  13871.24
第34个月末本利合计：  14851.51
第33个月末本利合计：  15830.37
第32个月末本利合计：  16807.85
第31个月末本利合计：  17783.93
第30个月末本利合计：  18758.62
第29个月末本利合计：  19731.93
第28个月末本利合计：  20703.85
第27个月末本利合计：  21674.39
第26个月末本利合计：  22643.55
第25个月末本利合计：  23611.33
第24个月末本利合计：  24577.73
第23个月末本利合计：  25542.76
第22个月末本利合计：  26506.41
第21个月末本利合计：  27468.69
第20个月末本利合计：  28429.60
第19个月末本利合计：  29389.15
第18个月末本利合计：  30347.33
第17个月末本利合计：  31304.15
第16个月末本利合计：  32259.60
第15个月末本利合计：  33213.70
第14个月末本利合计：  34166.43
第13个月末本利合计：  35117.82
第12个月末本利合计：  36067.85
第11个月末本利合计：  37016.52
第10个月末本利合计：  37963.85
第 9个月末本利合计：  38909.83
第 8个月末本利合计：  39854.46
第 7个月末本利合计：  40797.75
第 6个月末本利合计：  41739.69
第 5个月末本利合计：  42680.30
第 4个月末本利合计：  43619.57
第 3个月末本利合计：  44557.50
第 2个月末本利合计：  45494.09
第 1个月末本利合计：  46429.36
*/
