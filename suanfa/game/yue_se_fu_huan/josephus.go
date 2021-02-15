// 约瑟夫环：
//  41个人围成一个圆圈，由第一个人开始报数，报到3的人就必须自杀，再由下一个人开始报数，直到所有人都自杀为止。
//      约瑟夫和他的朋友安排在16和31的位置，最后就剩下他们两人，不符合游戏规则，所以逃过了该死亡游戏
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\suanfa\josephus.go -m 3 -n 41

package main

import (
    "flag"
    "fmt"
)

var (
    josephusPerentNum int // 参与的总人数
    josephusM         int // 数到3出列
)

func init() {
    flag.IntVar(&josephusPerentNum, "n", 41, "Usage: 41 42 43 44")
    flag.IntVar(&josephusM, "m", 3, "Usage: 2 3 4")
}

func main() {
    flag.Parse()

    josephus01(josephusPerentNum, josephusM)
}

func josephus01(total_num int, m int) {
    var josephusMan []int

    // 初始化切片
    for i := 0; i < total_num; i++ {
        josephusMan = append(josephusMan, 0)
    }

    pos := -1
    i := 0
    for count := 1; count <= total_num; count++ {
        for {
            pos = (pos + 1) % total_num //  环状处理
            if 0 == josephusMan[pos] {
                i++
            }
            if (i == m) {
                i = 0
                break
            }
        }
        josephusMan[pos] = count
    }

    fmt.Printf("\n约瑟夫排列（最初位置-约瑟夫环位置）：\n") // 输出排列位置
    for i := 0; i < total_num; i++ {
        fmt.Printf("%d-%d  ", i+1, josephusMan[i])
        if (i != 0 && 0 == i%10) {
            fmt.Println("")
        }
    }

    fmt.Printf("\n\n准备剩下的人数？")
    alive := 2 // 剩下的人数
    fmt.Printf("这%d人初始位置应排在以下序号：\n", alive)

    alive = total_num - alive // 需要跳过的编号
    for i := 0; i < total_num; i++ {
        if (josephusMan[i] > alive) {
            fmt.Printf("初始序号:%d, 约瑟夫环序号：%d\n", i+1, josephusMan[i])
        }
    }
    fmt.Println("")

    return
}
