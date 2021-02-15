// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\suanfa\shaizi\shaizi01.go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

var (
    shaiZiNum = -1
    renShu    = -1
    shaizi    []int // 定义碎片
)

func main() {
    // 接收参数
    for shaiZiNum < 0 {
        fmt.Println("请输入色子（骰子）数量（输入0退出）：")
        fmt.Scan(&shaiZiNum)
    }
    if shaiZiNum < 1 {
        fmt.Println("色子（骰子）数量最少1个")
        return
    }

    // 参数校验
    for renShu < 0 {
        fmt.Println("请输入参赛人数（输入0退出）：")
        fmt.Scan(&renShu)
    }
    if renShu < 2 {
        fmt.Println("参赛人数至少2人")
        return
    }

    rand.Seed(time.Now().UnixNano())
    shaizi = make([]int, shaiZiNum)
    // 输出每位选手投掷的筛子点数
    for i := 0; i < renShu; i++ {
        fmt.Printf("第%d位选手投掷的筛子点数为：", i+1)
        shaiziRand(shaizi, shaiZiNum)
        fmt.Println(shaizi)
    }

    return
}

// 几个色子的点数
func shaiziRand(shaizi []int, num int) {
    for i := 0; i < num; i++ {
        // 生成随机数
        r := rand.Intn(6) + 1 // 实际随机生成的数字范围[0,5]，所以要加1
        //fmt.Printf("rand-num的类型为[%T],rand-num的随机数值为:[%d]\n", r, r)
        shaizi[i] = r
    }
    return
}
