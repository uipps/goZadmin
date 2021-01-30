//
// go run caijiage.go
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\suanfa\caijiage\caijiage.go

//    由计算机随机生成一个 1~100的整数，由用户来猜这个数。

package main

import (
    "fmt"
    "math/rand"
    "time"
)

var (
    continueDo int // 继续计算还是退出 1-继续 非1-退出
)

func main() {
    game01()

    // do-while循环
    for {
        // 检测全局变量的值
        if 1 != continueDo {
            //os.Exit(0)
            break
        }
        fmt.Println("\n\n")
        game01()
    }

    return
}

func game01() {
    scanData()

    // 继续还是退出
    fmt.Println("\n")
    fmt.Println("继续计算还是退出，输入1表示继续，输入其他数字表示退出，非1的数都退出)")
    fmt.Scan(&continueDo)

    return
}

func scanData() {
    lguess := -1
    nCount := 0

    // 生成随机数
    rand.Seed(time.Now().UnixNano())
    n := rand.Intn(100) + 1 // 实际随机生成的数字范围[0,99]
    //fmt.Printf("rand-num的类型为[%T],rand-num的随机数值为:[%d]\n", n, n)

    // 循环判断
    for lguess != n {
        // 请输入所猜数字
        fmt.Println("\n")
        fmt.Println("请输入所猜数字：")
        fmt.Scan(&lguess)
        nCount++

        // 输出调试信息
        if lguess > n {
            fmt.Printf("所猜数字太大！\n")
        } else if lguess < n {
            fmt.Printf("所猜数字太小！\n")
        }
    }

    fmt.Printf("  答对了, 共猜测了%d次。\n", nCount)

    return
}
