/**

填运算符
  具体的题目是：在下面的算式中适当地添加“＋、－、×、÷”运算符，使等式成立，可使用括号。
2 2 2 2=5
3 3 3 3=5

go run tian_fuhao04.go -n "3 3 3 3" -r 5
go run tian_fuhao04.go -n 3 -r 5
go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\tianshu\tian_fuhao04.go -n "3 3 3 3" -r 5
go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\tianshu\tian_fuhao04.go -n 3 -r 5

*/

package main

import (
    "flag"
    "fmt"
    "strconv"
    "strings"
)

var (
    arg4Num        string  // 3个数字，用空格分隔
    argNumResult03 float64 // 结果
)

func init() {
    flag.StringVar(&arg4Num, "n", "8 8 8 8", "Usage: 8 8 8 8") // 左边的多个数字
    flag.Float64Var(&argNumResult03, "r", 6, "Usage: 8")       // 右边结果数字
}

func main() {
    flag.Parse()

    tianfuhao03(arg4Num, argNumResult03)
}

func tianfuhao03(arg_4Num string, arg_NumResult float64) {
    // 解析字符串
    //fmt.Println(arg_4Num)
    arr2 := strings.Split(arg_4Num, " ") // 重置
    //fmt.Println(arr2)

    numTotal := 4              // 左边数字个数，4个
    blankTotal := numTotal - 1 // 左边数字之间的空白数，4个数字3个空白可以填符号

    oper := [5]string{" ", "+", "-", "*", "/"}; // 运算符, 1表示+ ，2 表示-,3 表示* ，4 表示/
    blankI := make([]int, blankTotal+1)     // 循环变量，数组i用来表示需要填充的blankTotal个运算符 oper[blankI[j]]

    // 数字字符串转成float类型
    var arr03 []float64
    for _, val := range arr2 {
        if "" == val {
            continue // 过滤空白
        }
        i, _ := strconv.ParseFloat(val, 64) // 字符串转float64类型
        arr03 = append(arr03, i)
    }
    num := []float64{0.0}
    // 参数判断
    if (len(arr03) > 1 && len(arr03) < numTotal) {
        fmt.Printf("参数错误，-n请提供%d个数字，用空格分隔\n", numTotal)
        return
    } else if (1 == len(arr03)) {
        // 如果arg_Num只有一个数字，则4个自动填充
        for i := 0; i < numTotal; i++ {
            num = append(num, arr03[0])
        }
    } else {
        // 只需要numTotal个
        for i := 0; i < numTotal; i++ {
            num = append(num, arr03[i])
        }
    }
    //num = []float64{0.0, 8.0, 8.0, 8.0}; // 保存操作数
    result := arg_NumResult // 保存运算式的结果值

    count := 0; // 计数器，统计符合条件的方案

    for blankI[1] = 1; blankI[1] <= 4; blankI[1]++ { // 循环4种运算符，1表示+ ，2 表示-,3 表示* ，4 表示/
        if ((blankI[1] == 4) && (num[2] == 0)) { // 运算符若是/, 则第二个运算数不能为0
            continue
        }
        for blankI[2] = 1; blankI[2] <= 4; blankI[2]++ {
            if ((blankI[2] == 4) && (num[3] == 0)) {
                continue
            }
            for blankI[3] = 1; blankI[3] <= 4; blankI[3]++ {
                if ((blankI[3] == 4) && (num[numTotal] == 0)) {
                    continue
                }
                left := 0.0;                       // 保存中间结果，初始的left就是0
                right := num[1];                   // 初始的right就是num[1]，拆分成左右因为乘除法优先级高于加减
                sign := 1;                         // 累加运算时的符号
                for j := 1; j <= blankTotal; j++ { // 只有3个空白需要填符号
                    switch oper[blankI[j]] {
                    case "+":
                        left = left + float64(sign)*right; // 上一轮的该符号左右两边的数据
                        sign = 1;
                        right = num[j+1];
                        break;
                    case "-":
                        left = left + float64(sign)*right;
                        sign = -1;
                        right = num[j+1];
                        break; // 通过f=-1实现减法
                    case "*": // 乘除法优先级高于加减，因此跟left无关了
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
    if (count == 0) {
        fmt.Println("没有符合要求的方法！");
    }
    return
}
