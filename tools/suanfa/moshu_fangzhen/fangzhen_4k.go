// 4k方阵填数：方阵中的每行、每列或对角线位置的数各自相加的和均相等，
//  双向翻转法：将位于方阵中间部分的数据进行一次纵向、一次横向的翻转，即可。
//  (1) 将数字按从左向右、从上到下的顺序填入方阵;
//  (2) 将中间部分半数行的数字左右翻转
//  (3) 然后，将中间半数列的数字上下翻转，就完成了。

// 参考： https://www.cnblogs.com/heisaijuzhen/articles/4324474.html

// go run "F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\suanfa\moshu_fangzhen\fangzhen_4k.go" -b 1 -n 8
// go run "F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\suanfa\moshu_fangzhen\fangzhen_4k.go" -n 4
/**

请输入矩阵行数(奇数):3
 1   14   15    4
 8   11   10    5
12    7    6    9
13    2    3   16

*/

package main

import (
    "flag"
    "fmt"
    "os"
    "time"
)

var (
    argN3    int
    argBase3 int // 从几开始的基数，默认是从1开始、
)

func init() {
    flag.IntVar(&argN3, "n", 3, "Usage: 5 6 7")
    flag.IntVar(&argBase3, "b", 1, "Usage: 1")
}

func main() {
    flag.Parse()

    startTime := time.Now().UnixNano()
    fmt.Printf("startTime：%d, %s\n", startTime/1e3, time.Unix(0, startTime).Format("2006-01-02 15:04:05"))

    // 参数校验
    if argBase3 < 0 {
        fmt.Println("-b 参数错误，只能是大于等于0的数，提供的数值是：", argBase3)
        os.Exit(0)
    }
    // 4k矩阵
    fangzhen4k(argN3)

    // 执行时间计算
    endTime := time.Now().UnixNano()
    fmt.Printf("  endTime：%d, %s\n", endTime/1e3, time.Unix(0, endTime).Format("2006-01-02 15:04:05"))
    nanoSeconds := float64(endTime-startTime) / 1e3
    fmt.Println("spendTime：", nanoSeconds)
}

// 4k生成方阵的函数
func fangzhen4k(n int) int {
    if (n%4 != 0 || n < 4) {
        fmt.Println("\n请输入4的倍数，如:4 8 12等！\n");
        return 0 // 最小是4
    }

    // 初始化二维切片
    sliceLen := n
    matrix := make([][]int, sliceLen)
    for i := 0; i < sliceLen; i++ {
        matrix[i] = make([]int, sliceLen)
    }
    // 初始填入数字，并且进行水平翻转
    kBase := argBase3
    for i := 0; i < n; i++ {
        if (i <= n/4-1 || i > n/4-1+n/2) { //上下部分按顺序填入数字
            for j := 0; j < n; j++ {
                matrix[i][j] = kBase
                kBase++
            }
        } else {
            for j := n - 1; j >= 0; j-- { // 中间半数行按倒序填入数字, j从大到小
                matrix[i][j] = kBase
                kBase++
            }
        }
    }

    // 中间半数列翻转
    for j := n / 4; j < (n/4 + n/2); j++ { //将中间半数的列进行上下翻转
        for i := 0; i < n/2; i++ {
            matrix[i][j], matrix[n-i-1][j] = matrix[n-i-1][j], matrix[i][j] //交换数据
        }
    }

    // 输出
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            fmt.Printf("%2d   ", matrix[i][j])
        }
        fmt.Println()
    }
    // 对角线的和
    sum := 0;
    for i := 0; i < n; i++ { // 统计对角线的和
        sum += matrix[i][i];
    }
    fmt.Printf("\n各行、列、对角线的和为：%d\n", sum);
    return 1
}
