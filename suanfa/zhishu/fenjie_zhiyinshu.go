// 分解质因数
// 	go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\zhishu\fenjie_zhiyinshu.go -n 407
// 	go run ~/develope/go/go_code/src/github.com/uipps/goZadmin/suanfa/zhishu/fenjie_zhiyinshu.go -n 407

package main

import (
    "flag"
    "fmt"
)

var (
    argNum01 int
    rltNum01 []int
    runNum01 = 0
)

func init() {
    flag.IntVar(&argNum01, "n", 24, "Usage: 407 2310")
}

func main() {
    flag.Parse()

    // 参数校验
    if (argNum01 < 2) {
        fmt.Printf("\n参数有误，请输入大于2的正整数\n")
        return
    }

    fmt.Printf(" 第一种方法primeNum01：\n")
    rlt := fenjiePrimeNum01(argNum01)
    fmt.Printf("\n %d 分解质因数结果为： ", argNum01)
    fmt.Println(rlt)

    fmt.Println("\n")
    fmt.Printf("\n 第二种方法primeNum02：\n")
    fenjiePrimeNum02(argNum01, 2) // 从2开始
    fmt.Printf("  运算次数 %d\n", runNum01)
    fmt.Printf("\n %d 分解质因数结果为： ", argNum01)
    fmt.Println(rltNum01)

    return
}

// 1.1 递增试探，时间复杂度O(n^1/2)
func fenjiePrimeNum01(n int) []int {
    runCounter := 0
    rlt := []int{}

    i := 2 // 分解质因数，最小从2开始
    for n > 1 {
        if (n%i == 0) {
            n = n / i
            rlt = append(rlt, i)
            // 这里不能i++ , 因为可能包含多次
        } else {
            // 数据逐渐递增
            i++
        }
        runCounter++
    }

    fmt.Printf("  运算次数 %d\n", runCounter)
    return rlt
}

// 1.2 循环嵌套-递归方式
func fenjiePrimeNum02(n int, i int) {
    if (n <= 1) {
        return
    }
    runNum01++

    if (n%i == 0) {
        n = n / i
        rltNum01 = append(rltNum01, i)
    } else {
        i++
    }
    fenjiePrimeNum02(n, i)
}
