/**

一、 推理题1：
https://www.zybang.com/question/dbd62d25cfcec11e707886ca31be7a09.html
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


二、 推理题2：
https://blog.csdn.net/qq_29720657/article/details/78435177

警察抓住了A、B、C、D四名盗窃嫌疑犯，其中只有一人是小偷。在审问时，
A说：“我不是小偷”；
B说：“C是小偷”；
C说：“小偷肯定是D”；
D说：“C在冤枉好人”。
现在已经知道这四人中有三人说的是真话，一人说的是假话。请问到底谁是小偷？

-- 输出结果
真假话情况(1真话，0假话)：
  A: 1,
  B: 1,
  C: 0,
  D: 1; 小偷是C
真假话情况(1真话，0假话)：
  A: 1,
  B: 1,
  C: 1,
  D: 0; 小偷是CD
因为小偷只有一个，因此采纳“小偷是C”

*/

package main

import (
    "fmt"
    "github.com/uipps/goZadmin/suanfa/common"
)

func main() {
    //whoIsThief01()
    whoIsThief01_2()
    //whoIsThief02()
    //whoIsThief02_2()
}

func whoIsThief01() {
    lThief := "" // 小偷可能是多个，没有明确说只有一个小偷，虽然“一宗盗窃案”，暂且当多个小偷，看看有哪些情况

    for a := 0; a <= 1; a++ {
        for b := 0; b <= 1; b++ {
            for c := 0; c <= 1; c++ {
                for d := 0; d <= 1; d++ {
                    for e := 0; e <= 1; e++ {
                        // 他们当中只有3个人说的是真话
                        if (a+b+c+d+e == 3) {
                            lThief = "" // 清除上一轮
                            // E：B说的全是真话；B：我是无辜的；E和B都为1，或者E和B都是0
                            if (e+b == 2) {
                                // e,b都是1
                                //lThief = ""
                            } else if (e+b == 0) {
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

func whoIsThief02() {
    lThief := "" // 小偷可能是多个，没有明确说只有一个小偷，虽然“一宗盗窃案”，暂且当多个小偷，看看有哪些情况

    for a := 0; a <= 1; a++ {
        for b := 0; b <= 1; b++ {
            for c := 0; c <= 1; c++ {
                for d := 0; d <= 1; d++ {
                    // 他们当中只有3个人说的是真话
                    if (a+b+c+d == 3) {
                        lThief = "" // 清除上一轮

                        //c和d互咬只能有一个人对，c+d==1
                        if (c+d != 1) {
                            continue
                        }

                        if (a == 0) {
                            lThief += "A"
                        }

                        if (b == 1) {
                            lThief += "C"
                        }

                        if (c == 1) {
                            lThief += "D"
                        }

                        fmt.Printf("真假话情况(1真话，0假话)：\n  A: %d, \n  B: %d, \n  C: %d, \n  D: %d; 小偷是%s \n", a, b, c, d, lThief)
                    }
                }
            }
        }
    }
}

// 小偷只有一个的情况，可以逐个假定哪位是小偷进行一重循环，无需多重循环
func whoIsThief02_2() {
    // 对ABCD进行循环
    for i := 'A'; i <= 'D'; i++ {
        A := common.Btoi(i != 'A'); // A说：“我不是小偷”
        B := common.Btoi(i == 'C'); // B说：“C是小偷”；
        C := common.Btoi(i == 'D'); // C说：“小偷肯定是D”；
        D := common.Btoi(i != 'D'); // D说：“C在冤枉好人”。
        if (A+B+C+D == 3) {         // 四人中有三人说的是真话
            fmt.Printf("小偷是：%c \n", i); // 打印谁是小偷
        }
    }
}

// 假设问题1中小偷只有一个，可以逐个假定哪位是小偷进行一重循环，无需多重循环
func whoIsThief01_2() {
    // 对ABCDE进行循环
    for i := 'A'; i <= 'E'; i++ {
        A := common.Btoi(i == 'D'); // A说：“D是小偷”
        B := common.Btoi(i != 'B'); // B说：“我是无辜的”；
        C := common.Btoi(i != 'E'); // C说：“E不是小偷”；
        D := common.Btoi(i != 'D'); // D说：“A说的全是谎话”。
        E := common.Btoi(i != 'B'); // E说：“B说的全是真话”。
        if (A+B+C+D+E == 3) {       // 五人中有三人说的是真话
            fmt.Printf("小偷是：%c \n", i); // 打印
        }
    }
}
