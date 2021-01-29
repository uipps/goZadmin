/**

填运算符
  下面再演示一个用枚举算法解决问题的实例，具体的题目是：在下面的算式中适当地添加“＋、－、×、÷”运算符，使等式成立（不使用括号）。
5　5　5　5　5=5

go run tian_fuhao04.go -n "5 5 5 5 5" -r 5
go run tian_fuhao04.go -n 5 -r 5
go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\tianshu\tian_fuhao04.go -n "5 5 5 5 5" -r 5
go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\tianshu\tian_fuhao04.go -n 5 -r 5


二、 在五个8之间填上合适的运算符号,使计算结果 等于1至8（8和8要连加减乘除)  TODO 暂未支持括号
(8*8+8)/8-8=1
(8+8)/8+8-8=2
8/8+(8+8)/8=3
8/(8/8+8/8)=4
8-(8+8+8)/8=5
8-8/8-8/8=6
8+8-8-8/8=7
8+8-8+8-8=8

*/

package main

import (
    "flag"
    "fmt"
    "strconv"
    "strings"
)

var (
    arg5Num04      string  // 5个数字，用空格分隔
    argNumResult04 float64 // 结果
)

func init() {
    flag.StringVar(&arg5Num04, "n", "5 5 5 5 5", "Usage: 5 5 5 5 5") // 左边的多个数字
    flag.Float64Var(&argNumResult04, "r", 5, "Usage: 5")             // 右边结果数字
}

func main() {
    flag.Parse()

    tianfuhao04(arg5Num04, argNumResult04)
}

func tianfuhao04(arg_5Num string, arg_NumResult float64) {
    // 解析字符串
    //fmt.Println(arg_5Num)
    arr2 := strings.Split(arg_5Num, " ") // 重置
    //fmt.Println(arr2)

    numTotal := 5               // 左边数字个数，5个
    blankTotal := numTotal - 1  // 左边数字之间的空白数，5个数字4个空白可以填符号

    oper := [5]string{" ", "+", "-", "*", "/"}; // 运算符, 1表示+ ，2 表示-,3 表示* ，4 表示/
    blankI := make([]int, blankTotal+1)     // 循环变量，数组i用来表示需要填充的blankTotal个运算符 oper[blankI[j]]

    // 数字字符串转成float类型
    var arr1 []float64
    for _, val := range arr2 {
        if "" == val {
            continue // 过滤空白
        }
        i, _ := strconv.ParseFloat(val, 64) // 字符串转float64类型
        arr1 = append(arr1, i)
    }
    num := []float64{0.0}
    // 参数判断
    if (len(arr1) > 1 && len(arr1) < numTotal) {
        fmt.Printf("参数错误，-n请提供%d个数字，用空格分隔\n", numTotal)
        return
    } else if (1 == len(arr1)) {
        // 如果arg_Num只有一个数字，则5个自动填充
        for i := 0; i < numTotal; i++ {
            num = append(num, arr1[0])
        }
    } else {
        // 只需要5个
        for i := 0; i < numTotal; i++ {
            num = append(num, arr1[i])
        }
    }
    //num = []float64{0.0, 5.0, 5.0, 5.0, 5.0, 5.0}; // 保存操作数
    result := arg_NumResult // 保存运算式的结果值

    sign := 0                // 累加运算时的符号
    count := 0;              // 计数器，统计符合条件的方案
    left, right := 0.0, 0.0; // 保存中间结果

    for blankI[1] = 1; blankI[1] <= 4; blankI[1]++ { // 循环4种运算符，1表示+ ，2 表示-,3 表示* ，4 表示/
        if ((blankI[1] == 4) && (num[2] == 0)) { // 运算符若是/, 则第二个运算数不能为0
            continue
        }
        for blankI[2] = 1; blankI[2] <= 4; blankI[2]++ {
            if ((blankI[2] == 4) && (num[3] == 0)) {
                continue
            }
            for blankI[3] = 1; blankI[3] <= 4; blankI[3]++ {
                if ((blankI[3] == 4) && num[4] == 0) {
                    continue
                }
                for blankI[4] = 1; blankI[4] <= 4; blankI[4]++ {
                    if ((blankI[4] == 4) && (num[numTotal] == 0)) {
                        continue
                    }
                    left = 0;
                    right = num[1];
                    sign = 1;
                    for j := 1; j <= blankTotal; j++ {
                        switch oper[blankI[j]] {
                        case "+":
                            left = left + float64(sign)*right;
                            sign = 1;
                            right = num[j+1];
                            break;
                        case "-":
                            left = left + float64(sign)*right;
                            sign = -1;
                            right = num[j+1];
                            break; // 通过f=-1实现减法
                        case "*":
                            right = right * num[j+1];
                            break; // 实现乘法
                        case "/":
                            right = right / num[j+1]; // 实现除法
                            break;
                        }
                    }
                    if (left+float64(sign)*right == result) {
                        count++;
                        fmt.Printf("%3d：", count);
                        for j := 1; j <= blankTotal; j++ {
                            fmt.Printf("%.f%s", num[j], oper[blankI[j]]);
                        }

                        fmt.Printf("%.f=%.f\n", num[numTotal], result);
                    }
                }
            }
        }
    }
    if (count == 0) {
        fmt.Println("没有符合要求的方法！");
    }
    return
}
