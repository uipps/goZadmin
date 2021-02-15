// 奇数魔术方阵(2k+1)：方阵中的每行、每列或对角线位置的数各自相加的和均相等，
//  简捷连续填数法：1立首行中，右1上1，受阻下1，具体看下面解释
//  (1) 将1放在第一行中间一列;
//  (2) 从2开始直到n×n止各数依次按下列规则存放：每一个数存放的行比前一个数的列数加1，行数减1，
//  (3) 如果行列范围超出矩阵范围，则回绕。例如1在第1行，则2应放在最下一行，列数同样加1;
//  (4) 如果按上面规则确定的位置上已有数，或上一个数是第1行第n列时，则把下一个数放在上一个数的下面。

// go run "F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\math\moshu_fangzhen\fangzhen_2k+1.go" -b 1 -n 3
// go run "F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\math\moshu_fangzhen\fangzhen_2k+1.go" -n 5
/**

请输入矩阵行数(奇数2k+1):3
 8    1    6
 3    5    7
 4    9    2

*/

package main

import (
    "flag"
    "fmt"
    "os"
    "time"
)

var (
    argN1     int
    argBase1  int   // 从几开始的基数，默认是从1开始
    arrYiWei1 []int // 用1维数组存放数字好了，TODO 以后改成二维数组
)

func init() {
    flag.IntVar(&argN1, "n", 3, "Usage: 5 6 7")
    flag.IntVar(&argBase1, "b", 1, "Usage: 1")
}

func main() {
    flag.Parse()

    startTime := time.Now().UnixNano()
    fmt.Printf("startTime：%d, %s\n", startTime/1e3, time.Unix(0, startTime).Format("2006-01-02 15:04:05"))

    // 参数校验
    if argBase1 < 0 {
        fmt.Println("-b 参数错误，只能是大于等于0的数，提供的数值是：", argBase1)
        os.Exit(0)
    }
    // 2k+1矩阵
    fangzhen2k1(argN1)

    // 执行时间计算
    endTime := time.Now().UnixNano()
    fmt.Printf("  endTime：%d, %s\n", endTime/1e3, time.Unix(0, endTime).Format("2006-01-02 15:04:05"))
    nanoSeconds := float64(endTime-startTime) / 1e3
    fmt.Println("spendTime：", nanoSeconds)
}

func Magic1(n int) {
    /*var matrix [][]int
      var tmp_arr []int
      // 初始化一下
      for j := 0; j <= n; j++ {
          tmp_arr = append(tmp_arr, 0)
      }
      for i := 0; i <= n; i++ {
          matrix = append(matrix, tmp_arr)
          fmt.Println();
      }
      fmt.Println(matrix);

      rlt := Magic1(matrix, n)

      if 1 == rlt {
          // 输出
          for i := 1; i < n; i++ {
              for j := 1; j < n; j++ {
                  fmt.Printf("%2d   ", matrix[i][j]);
              }
              fmt.Println();
          }
      } else {
          fmt.Println("生成方阵失败！可能是输入的阶数不正确！");
      }
    */
}

// 生成方阵的函数
func fangzhen2k1(n int) int {
    if (n%2 == 0 || n < 3) {
        return 0 // 偶数返回
    }
    //const MAXSIZE = 99
    //matrix := [MAXSIZE][MAXSIZE]int{} // 用二维数组也行
    // 二维切片初始化，二维数组初始化
    sliceLen := n+6 // 容量要大一点，因为第三行用a[3]表示
    //matrix := make([][]int, sliceLen)
    //for i := 0; i < sliceLen; i++ {
    //    matrix[i] = make([]int, sliceLen)
    //}
    //matrix := [][]int 错误的语法
    defaultV := -1000   // 默认值可以用0，也可以自定义一个值
    var matrix [][]int
    for i := 0; i < sliceLen; i++ {
        var tmp_s []int
        for j := 0; j < sliceLen; j++ {
            tmp_s = append(tmp_s, defaultV)
        }
        matrix = append(matrix, tmp_s)
    }
    //fmt.Println(matrix)

    i := 1 // 行号，从1开始，1表示第一行,不是从0开始；列号也是从1开始，1表示第一列

    j := (n + 1) / 2 // 行中间列号，
    matrix[i][j] = argBase1 // 将1放在第一行中间一列

    k := argBase1
    for k = argBase1 + 1; k <= n*n+argBase1-1; k++ {
        tmpi := i   // 记录原位置，占用的时候需要进行下移操作（行号+1）
        tmpj := j
        if (j == n) { // 1列就是第一列，为j==n不能++就越界，所以从头开始
            j = 1
        } else {
            j++
        }
        if (i == 1) { // i==0,i--就会越界，因此用最大行号替代
            i = n
        } else {
            i--
        }
        // 新位置被占用，那么原位置下移一位即可
        if (matrix[i][j] != defaultV) {
            j = tmpj
            i = tmpi + 1
            // tmpi可能为n，i可能会越界，实际上不会发生
            //if i > n {
            //    i = i - n   // 也不用判断新位置是否被占用的情况了
            //}
        }
        matrix[i][j] = k
    }

    // 输出
    for i := 1; i <= n; i++ {
        for j := 1; j <= n; j++ {
            fmt.Printf("%2d   ", matrix[i][j])
        }
        fmt.Println()
    }
    // 对角线的和
    sum := 0;
    for i := 1; i <= n; i++ {// 统计对角线的和
        sum += matrix[i][i];
    }
    fmt.Printf("\n各行、列、对角线的和为：%d\n", sum);

    return 1
}
