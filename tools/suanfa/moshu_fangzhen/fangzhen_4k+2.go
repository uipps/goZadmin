// 4k+2方阵：方阵中的每行、每列或对角线位置的数各自相加的和均相等，
//  井字调整法：
//  (1) 将数字按从左向右、从上到下的顺序填入方阵,然后在第k+1、3k+2行及列做井字标记
//  (2) 将标记内部的行做水平翻转，再将列做垂直翻转（井字上的数字不动）
//  (3) 将井字分隔线的行及第k+2行两侧的数字左右对调，两行中央的数字上下对调，左边列的数字除交叉点外进行垂直翻转。
//  (4) 将井字分隔线列的纵向中央的数字除第2k+1列外左右对调，井字横行左方的第一个数字上下对调，上横线中央的数字水平翻转。

// go run "F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\suanfa\moshu_fangzhen\fangzhen_4k+2.go" -b 2 -n 6
// go run "F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\suanfa\moshu_fangzhen\fangzhen_4k+2.go" -n 6
/**
请输入矩阵行数(4k+2):6
 1   32   33   34    5    6
30    8   28   27   11    7
13   20   22   21   17   18
24   23   16   15   14   19
12   26    9   10   29   25
31    2    3    4   35   36

*/

package main

import (
    "flag"
    "fmt"
    "os"
    "time"
)

var (
    argN2     int
    argBase2  int   // 从几开始的基数，默认是从1开始
)

func init() {
    flag.IntVar(&argN2, "n", 3, "Usage: 5 6 7")
    flag.IntVar(&argBase2, "b", 1, "Usage: 1")
}

func main() {
    flag.Parse()

    startTime := time.Now().UnixNano()
    fmt.Printf("startTime：%d, %s\n", startTime/1e3, time.Unix(0, startTime).Format("2006-01-02 15:04:05"))

    // 参数校验
    if argBase2 < 0 {
        fmt.Println("-b 参数错误，只能是大于等于0的数，提供的数值是：", argBase2)
        os.Exit(0)
    }
    // 4k+2矩阵
    fangzhen4k2(argN2)

    // 执行时间计算
    endTime := time.Now().UnixNano()
    fmt.Printf("  endTime：%d, %s\n", endTime/1e3, time.Unix(0, endTime).Format("2006-01-02 15:04:05"))
    nanoSeconds := float64(endTime-startTime) / 1e3
    fmt.Println("spendTime：", nanoSeconds)
}

// 4k+2生成方阵的函数
func fangzhen4k2(n int) int {
    if (n%4 != 2 || n < 6) {
        return 0 // 最小是6
    }

    // 二维切片初始化，二维数组初始化
    magic := make([][]int, n)
    for i := 0; i < n; i++ {
       magic[i] = make([]int, n)
    }

    sn := (n - 2) / 4  // sn=1

    //1. 第一步,按顺序填入数字
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            magic[i][j] = i * n + j + 1;
        }
    }


    //2. 第二步，将井字左右部分进行翻转
    for i := sn + 1; i <= 3 * sn; i++ {
        for j := 0; j <= n / 2 - 1; j++ {
            if j != sn {
                magic[i][j], magic[i][n - 1 - j] =  magic[i][n - 1 - j], magic[i][j]
            }
        }
    }
    //   2.2 第二步，井字上下部分进行翻转
    for i := sn + 1; i <= 3 * sn; i++ {
        for j := 0; j <= n / 2 - 1; j++ {
            if (j != sn) {
                magic[j][i], magic[n - 1 - j][i] = magic[n - 1 - j][i], magic[j][i]
            }
        }
    }

    //3. 第三步，将井字分割线之横列及第k+2列两侧数字对调
    for i := 0; i <= sn - 1; i++ {
        magic[sn][i], magic[sn][n - 1 - i] = magic[sn][n - 1 - i], magic[sn][i]
        magic[n - 1 - sn][i], magic[n - 1 - sn][n - 1 - i] = magic[n - 1 - sn][n - 1 - i],  magic[n - 1 - sn][i]
        magic[sn + 1][i], magic[sn + 1][n - 1 - i] = magic[sn + 1][n - 1 - i],magic[sn + 1][i]
    }
    //  3.2 两横列中央数字上下对调
    for i := sn + 1; i <= n - 1 - sn - 1; i++ {
        magic[sn][i], magic[n - 1 - sn][i] = magic[n - 1 - sn][i], magic[sn][i]
    }
    //   3.3 左变列除交点外的数字垂直翻转
    for i := 0; i <= n / 2 - 1; i++ {
        if (i != sn) {
            magic[i][sn], magic[n - 1 - i][sn] = magic[n - 1 - i][sn],magic[i][sn]
        }
    }

    //4. 第四步，将井字横线左侧数字上下对调
    magic[sn][0], magic[n - 1 - sn][0] = magic[n - 1 - sn][0], magic[sn][0]

    //井字横线中央翻转
    for i := sn + 1; i <= n / 2 - 1; i++ {
        magic[sn][i], magic[sn][n - 1 - i] = magic[sn][n - 1 - i], magic[sn][i]
    }
    //井字分割线纵向中央数字除第2K+1列外左右对调
    for i := sn + 1; i <= n - 1 - sn - 1; i++ {
        if (i != n / 2 - 1) {
            magic[i][sn], magic[i][n - 1 - sn] = magic[i][n - 1 - sn], magic[i][sn]
        }
    }

    // 输出
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            fmt.Printf("%2d   ", magic[i][j])
        }
        fmt.Println()
    }
    // 对角线的和
    sum := 0;
    for i := 0; i < n; i++ { // 统计对角线的和
        sum += magic[i][i];
    }
    fmt.Printf("\n各行、列、对角线的和为：%d\n", sum);

    return 1
}
