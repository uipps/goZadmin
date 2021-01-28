/**
 取石子游戏 （常胜将军）

一、规则1 （参考： https://www.cnblogs.com/heisaijuzhen/articles/4324474.html）
有n堆石子，每堆有若干石子，数量不一定相同，两人(游戏者与计算机)轮流从任一堆中拿走任意数量的石子，最后把石子全部拿走者为胜利方。
所谓“必负局”，是指把剩余的每一堆的数目都转化成二进制的数，然后把它们相加，进行不进位的加法（也就是异或运算）,
即0+0=0、1+0=1,0+1=1、1+1=0（不进位），如果所得和是0（多个0）,那么此局势称为“必负局”。


二、规则2 (参考： https://blog.csdn.net/csy981848153/article/details/9005248)
题目：现有21根火柴，两人轮流取，每人每次可以取走1至4根，不可多取，也不能不取，谁取最后一根火柴谁输。
请编写一个程序进行人机对弈，要求人先取，计算机后取；计算机一方为“常胜将军”。

这个相对容易一点，就是找规律。



*/

//     go run qu_shizi.go -n 21
//     go run /Users/cf/develope/go/go_code/src/github.com/uipps/goZadmin/suanfa/game/qu_shi_zi/qu_shizi.go -n 21
//     go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\game\qu_shi_zi\qu_shizi.go -n 22

package main

import (
    "flag"
    "fmt"
)

var (
    input int
    argTotal02 int
)

func init() {
    flag.IntVar(&argTotal02, "n", 21, "Usage: 21")  // 总数
}

func main() {
    flag.Parse()

    guize02()
}

// 规则2：找到规律，原来是4+1一个循环组。 参考： https://blog.csdn.net/csy981848153/article/details/9005248
func guize02() {
    sum := 21;       //tick num, 5k+1，后手（电脑）赢；否则 TODO
    //input = 0;       //for input
    bBreak := false; //if user input the wrong num is true

    fmt.Printf(" 游戏规则是： 两人轮流取，每人每次可以取走1至4根，不可多取，也不能不取，谁取最后一根火柴谁输。\n\n")

    //game run
    for sum > 4 {
        fmt.Printf("now the number of ticks is: %d\n", sum)

        // 接收用户输入
        fmt.Println("turn to your choice(range:1~4): ")
        fmt.Scan(&input)

        if (input < 1 || input > 4) {
            fmt.Printf("your choice is out of the range(1~4)!\n")
            bBreak = true;
            break;
        }
        fmt.Printf("computer choices the num is: %d\n", 5-input)    // 电脑总是拿走5-input数量
        sum -= 5;
    }
    if bBreak {
        fmt.Printf("because you input the wrong num,", input)
    } else {
        fmt.Printf("the num of tick is %d, less than 4,", sum)
    }
    fmt.Printf(" computer is the winner!")

    return
}
