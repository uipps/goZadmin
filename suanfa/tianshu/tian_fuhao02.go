/**

填运算符
  具体的题目是：在下面的算式中适当地添加“＋、－、×、÷”运算符，使等式成立，可使用括号。
2 2 2=6
3 3 3=6
4 4 4=6
5 5 5=6
6 6 6=6
7 7 7=6
8 8 8=6
9 9 9=6

go run tian_fuhao02.go -n "8 8 8" -r 6
go run tian_fuhao02.go -n 8 -r 6
go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\tianshu\tian_fuhao02.go -n "8 8 8" -r 6
go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\tianshu\tian_fuhao02.go -n 8 -r 6

*/

package main

import (
    "flag"
    "fmt"
    "strconv"
    "strings"
)

var (
    arg3Num02      string  // 3个数字，用空格分隔
    argNumResult02 float64 // 结果
    arr02          []float64
)

func init() {
    flag.StringVar(&arg3Num02, "n", "8 8 8", "Usage: 8 8 8") // 左边的多个数字
    flag.Float64Var(&argNumResult02, "r", 6, "Usage: 8")     // 右边结果数字
}

func main() {
    flag.Parse()

    tianfuhao02(arg3Num02, argNumResult02)
}

func tianfuhao02(arg_3Num string, arg_NumResult float64) {
    // 解析字符串
    //fmt.Println(arg_3Num)
    arr2 := strings.Split(arg_3Num, " ") // 重置
    //fmt.Println(arr2)

    numTotal := 3       // 左边数字个数，3个
    blankTotal := 2     // 左边数字之间的空白数，3个数字2个空白可以填符号

    // 数字字符串转成float类型
    for _, val := range arr2 {
        if "" == val {
            continue // 过滤空白
        }
        i, _ := strconv.ParseFloat(val, 64) // 字符串转float64类型
        arr02 = append(arr02, i)
    }
    num := []float64{0.0}
    // 参数判断
    if (len(arr02) > 1 && len(arr02) < numTotal) {
        fmt.Println("参数错误，-n请提供3个数字，用空格分隔")
        return
    } else if (1 == len(arr02)) {
        // 如果arg_3Num只有一个数字，则3个自动填充
        for i := 0; i < numTotal; i++ {
            num = append(num, arr02[0])
        }
    } else {
        // 只需要3个
        for i := 0; i < numTotal; i++ {
            num = append(num, arr02[i])
        }
    }
    //num = []float64{0.0, 8.0, 8.0, 8.0}; // 保存操作数
    result := arg_NumResult // 保存运算式的结果值

    count := 0;              // 计数器，统计符合条件的方案

    i := []int{0, 1, 1}                   // 循环变量，数组i用来表示4个运算符 oper[i[j]]
    oper := [5]string{" ", "+", "-", "*", "/"}; // 运算符, 1表示+ ，2 表示-,3 表示* ，4 表示/

    for i[1] = 1; i[1] <= 4; i[1]++ {           // 循环4种运算符，1表示+ ，2 表示-,3 表示* ，4 表示/
        if ((i[1] == 4) && (num[2] == 0)) {     // 运算符若是/, 则第二个运算数不能为0
            continue
        }
        for i[2] = 1; i[2] <= 4; i[2]++ {
            if ((i[2] == 4) && (num[numTotal] == 0)) {
                continue
            }
            left := 0.0;         // 保存中间结果
            right := num[1];
            sign := 1;           // 累加运算时的符号
            for j := 1; j <= blankTotal; j++ {   // 只有2个空白需要填符号
                switch oper[i[j]] {
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
                    fmt.Printf("%.f%s", num[j], oper[i[j]]);
                }

                fmt.Printf("%.f=%.f\n", num[numTotal], result);
            }
        }
    }
    if (count == 0) {
        fmt.Println("没有符合要求的方法！");
    }
    return
}
