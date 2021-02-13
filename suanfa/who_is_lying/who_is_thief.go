/**
雷米警长正在盘问一宗盗窃案的5个嫌疑人，他们当中只有3个人说的是真话，根据他们的说辞，你能猜出谁是小偷吗？
A：D是小偷
B：我是无辜的
C：E不是小偷
D：A说的全是谎话
E：B说的全是真话

--
解：(小偷一个还是多个？假设小偷数量不确定，将全部情况打印出来，最后在其中挑选只有一个小偷的，因为“一宗盗窃案”，题意应该就是一个小偷)
    真话为1，假话为0
    小偷为1，非小偷为0
    5人分别用a-e表示



-- 程序运行
go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\who_is_lying\who_is_thief.go

-- 运行结果
真假话情况(1真话，0假话)：
  A: 0,
  B: 1,
  C: 0,
  D: 1,
  E: 1; 小偷是E
真假话情况(1真话，0假话)：
  A: 1,
  B: 1,
  C: 0,
  D: 0,
  E: 1; 小偷是DE
*/

package main

import "fmt"

func main() {
    whoIsThief()
}

func whoIsThief() {
    lThief := "" // 小偷可能是多个，没有明确说只有一个小偷，虽然“一宗盗窃案”，暂且当多个小偷，看看有哪些情况


    for a := 0; a <= 1; a++ {
        for b := 0; b <= 1; b++ {
            for c := 0; c <= 1; c++ {
                for d := 0; d <= 1; d++ {
                    for e := 0; e <= 1; e++ {
                        // 他们当中只有3个人说的是真话
                        if (a+b+c+d+e == 3) {
                            lThief = ""         // 清除上一轮
                            // E：B说的全是真话；B：我是无辜的；E和B都为1，或者E和B都是0
                            if (e + b ==2) {
                                // e,b都是1
                                //lThief = ""
                            } else if (e + b == 0 ) {
                                lThief += "B"
                            } else {
                                // e + b ==1不成立
                                continue
                            }

                            // A：D是小偷;D：A说的全是谎话; 所以a+d == 1；如果A说的是真话a==1，那么D是小偷
                            if (a+d != 1) {
                                continue
                            }
                            if (a == 1) {
                                lThief += "D"
                            }

                            if (c == 0) {
                                lThief += "E"
                                //fmt.Printf("E不是小偷 \n")
                            }
                            fmt.Printf("真假话情况(1真话，0假话)：\n  A: %d, \n  B: %d, \n  C: %d, \n  D: %d, \n  E: %d; 小偷是%s \n", a, b, c, d, e, lThief)
                        }
                    }
                }
            }
        }
    }
}
