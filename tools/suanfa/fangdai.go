// 房贷利息计算器，2021.1.9
// go run fangdai.go -c 480000 -m 240 -r 0.0441 -f 1
// go run fangdai.go -c 100000 -m 12 -r 0.05 -f 0

//  参考PHP代码，php E:/chengfeng/OneDrive/chengfeng/95533-建设银行/fangdai.php -c 480000 -m 240 -r 0.0441 -f 0

//	分为两种方式：
// 		1、等额本息，特点：每月的月供相同
//				公式：月供=贷款本金×[月利率×(1+月利率) ^ 还款月数]÷{[(1+月利率) ^ 还款月数]-1}
//		2、等额本金，特点：第一个还款金额最高，每月还款金额递减
//				公式：月供=(本金/还款月数)+(本金-累计已还本金)×月利率
// 举例说明：借款10万元，贷款期限为12月，贷款年利率为5%，分别用等额本息与等额本金计算房贷。
//		等额本息月供为8560.47元，还款总额为102728.99元，利息总额为2728.99元。
//		等额本金第一个月月供为8750元，第二月为8715.28元，第三个月为8680.56元……最后一个月月供为8368.06元，还款总额为102708.33元，利息总额为2708.33。

package main

import (
	"fmt"
	"flag"
	//"math"
)

var (
	capital  float64	// 本金
	month    int		// 贷款月数，20年就是240月
	interestRateYear float64	// 贷款年利率
	fangan    int		// 方案1
)

func init() {
	flag.Float64Var(&capital, "c", 480000, "Usage: 100000 480000")
	flag.IntVar(&month, "m", 240, "Usage: 12 240")
	flag.Float64Var(&interestRateYear, "r", 0.0441, "Usage: 0.05")
	flag.IntVar(&fangan, "f", 0, "Usage: 0 1 2")	// 1-等额本息 2-等额本金
}

func main()  {
	flag.Parse()

	if (1 == fangan) {
		Debx(capital, month, interestRateYear);
	} else if (2 == fangan) {
		Debj(capital, month, interestRateYear);
	} else {
		fmt.Println("等额本息方式：")
		Debx(capital, month, interestRateYear);
		fmt.Println("\n\n等额本金方式：")
		Debj(capital, month, interestRateYear);
	}
	return
}


// 1. 等额本息，特点：每月的月供相同
//      等额本息的月供公式是如何推导出来的呢？ 每一期的利息为总贷款减去累计已还的本金然后乘以月利率。每一期还款本金=x-当期利息；
//      设贷款总额为d, 每月利息为R，每月还款x
// 第几期      本金                              利息                                       还款（都是等额的x）
//第  1期  x-dR = x(还款额x)-dR（当期利息）        dR                                          x
//第  2期 (x-dR)(1+R) = x - <当期利息>      [d-(x-dR)]R = [d(1+R)-x]R                         x  (本期的利息就是 总额减去上期还的本金然后乘以月利率)
//第  3期 (x-dR)(1+R)^2 = x - <当期利息>    [d-((x-dR)+(x-dR)(1+R))]R = dR-(x-dR)(2+R)R       x  (本期的利息就是 总额减去上2期还的本金然后乘以月利率)
//第  4期 (x-dR)(1+R)^3
//第  5期 (x-dR)(1+R)^4
// ......
//第  n期 (x-dR)(1+R)^(n-1)

// 如总额480000，240月，年利息0.0441的情况
//第  1期 本金:   1249.44 利息:   1764.00 总额:   3013.44
//第  2期 本金:   1254.04 利息:   1759.40 总额:   3013.44
//第  3期 本金:   1258.65 利息:   1754.79 总额:   3013.44
//第  4期 本金:   1263.27 利息:   1750.17 总额:   3013.44
//第  5期 本金:   1267.91 利息:   1745.53 总额:   3013.44
//
//func Debx(dkTotal=480000, dkm=240, dknl=0.0441)  {
func Debx(dkTotal float64, dkm int, dknl float64)  {
	//var lx float64
	lx := 0.0
	benjin := 0.0
	lxTotal := 0.0

	fmt.Printf("  贷款总额： %.2f ，贷款周期： %d 个月，年利率： %.2f 还贷方式： 等额本息。\n", dkTotal, dkm, dknl)

	emTotal := dkTotal * dknl / 12 * pow(1 + dknl / 12, dkm) / (pow(1 + dknl / 12, dkm) - 1) //每月还款金额

	for i:=0; i < dkm; i++  {
		lx = dkTotal*dknl/12;
		benjin = emTotal - lx
		fmt.Printf("第%3d期  本金:%10.2f 利息:%10.2f 总额:%10.2f\n", i+1, benjin, lx, emTotal )
		dkTotal = dkTotal - benjin
		lxTotal = lxTotal + lx
	}
	fmt.Printf("总利息: %.2f", lxTotal)
}

// 2. 等额本金，特点：第一个还款金额最高，每月还款金额递减
//func Debj(dkTotal float64 = 480000, dkm int = 240, dknl float64 = 0.0441) {
func Debj(dkTotal float64, dkm int, dknl float64) {
	lx := 0.0
	benjin := 0.0
	lxTotal := 0.0

	fmt.Printf("  贷款总额： %.2f ，贷款周期： %d 个月，年利率： %.2f 还贷方式： 等额本金。\n", dkTotal, dkm, dknl)

	benjin = dkTotal / float64(dkm)	// 等额本金， float64(强制类型转换
	for i:=0; i < dkm; i++  {
		lx = dkTotal*dknl/12;
		fmt.Printf("第%3d期  本金:%10.2f 利息:%10.2f 总额:%10.2f\n", i+1, benjin, lx, benjin+lx )
		dkTotal = dkTotal - benjin
		lxTotal = lxTotal + lx
	}
	fmt.Printf("总利息: %.2f", lxTotal)
}

// 递归法 求x^n
func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else {
		return x * pow(x, n-1)
	}
}