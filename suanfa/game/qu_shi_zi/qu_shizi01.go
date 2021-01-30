/**
 取石子游戏 (必负局)

一、规则1 （参考： https://www.cnblogs.com/heisaijuzhen/articles/4324474.html）
有n堆石子，每堆有若干石子，数量不一定相同，两人(游戏者与计算机)轮流从任一堆中拿走任意数量的石子，拿走最后一个石子者为胜利方。
所谓“必负局”，是指把剩余的每一堆的数目都转化成二进制的数，然后把它们相加，进行不进位的加法（也就是异或运算）,
即0+0=0、1+0=1,0+1=1、1+1=0（不进位），如果所得和是0（多个0）,那么此局势称为“必负局”。

二、规则2 (参考： https://blog.csdn.net/csy981848153/article/details/9005248)
题目：现有21根火柴，两人轮流取，每人每次可以取走1至4根，不可多取，也不能不取，谁取最后一根火柴谁输。
请编写一个程序进行人机对弈，要求人先取，计算机后取；计算机一方为“常胜将军”。

*/

//     go run qu_shizi01.go -n 21
//     go run /Users/cf/develope/go/go_code/src/github.com/uipps/goZadmin/suanfa/game/qu_shi_zi/qu_shizi01.go -n 21
//     go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\game\qu_shi_zi\qu_shizi01.go -n 22

package main

import (
    "flag"
    "fmt"
)

var (
    num      int
    heap     []int //保存各堆石子的数量 unsigned int *heap;
    argTotal int
)

func init() {
    flag.IntVar(&argTotal, "n", 21, "Usage: 21") // 总数
}

func main() {
    flag.Parse()

    guize01()
}

func Output() { // 显示各堆石子的状态
    fmt.Printf("各堆剩下的石子数量:\n")
    for i := 1; i <= num; i++ {
        fmt.Printf("第%2d堆剩下：%d \n", i, heap[i])
    }
}

func remain() int { //统计剩下的石子总数
    s := 0
    for i := 1; i <= num; i++ {
        s += heap[i]
    }
    return s
}

func xorall() int { //按位异或的结果
    s := 0
    for i := 1; i <= num; i++ {
        s ^= heap[i]
    }
    return s
}

// 规则1
func guize01() {
    h, t := 0, 0

    fmt.Printf(" 游戏规则是： 有n堆石子，每堆有若干石子，数量不一定相同，两人(游戏者与计算机)轮流从任一堆中拿走任意数量的石子，拿走最后一个石子者为胜利方。\n\n")

    fmt.Printf("输入石子的堆数：")
    fmt.Scan(&num)
    if (num < 2) {
        fmt.Printf("至少应该有2堆石子!\n")
        return
    }

    // 初始化切片
    for i := 0; i <= num; i++ {
        heap = append(heap, 0)
    }

    for i := 1; i <= num; i++ {
        fmt.Printf("输入第%d堆石子的数量：", i)
        fmt.Scan(&heap[i])
    }

    for remain() > 0 { //剩余石子大于0
        if (xorall() == 0) {
            for i := 1; i <= num; i++ { //从一堆石子中取一粒
                if (heap[i] > 0) {
                    fmt.Printf("\n计算机从第%2d堆中拿1粒。\n", i)
                    heap[i]--
                    break
                }
            }
        } else {
            for i := 1; i <= num; i++ {
                s := heap[i] - (xorall() ^ heap[i]) //计算要取的石子数量
                if (s > 0) {
                    fmt.Printf("\n计算机从第%2d堆中拿%d粒。\n", i, s)
                    heap[i] ^= xorall()
                    break
                }
            }
        }
        if (remain() == 0) {
            fmt.Printf("\n计算机胜")
            break
        }
        Output() //显示剩余的石堆情况

        for {
            fmt.Printf("\n输入你的选择（堆 数量）：")
            fmt.Scanf("%d %d", &h, &t)
            if ((h >= 1) && (h <= num) && (heap[h] >= t)) {
                heap[h] -= t
                break
            } else {
                fmt.Printf("\n输入数据出错，重新输入！\n")
            }
        }
        if (remain() == 0) {
            fmt.Printf("\n恭喜你获胜")
            break
        }
    }

    return
}
