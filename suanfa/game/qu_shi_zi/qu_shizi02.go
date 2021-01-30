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
//     go run /Users/cf/develope/go/go_code/src/github.com/uipps/goZadmin/suanfa/game/qu_shi_zi/qu_shizi02.go -n 21
//     go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\game\qu_shi_zi\qu_shizi02.go -n 22

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

    fmt.Printf(" ------ 游戏规则是： 两人轮流取，每人每次可以取走1至4根，不可多取，也不能不取，谁取最后一根火柴谁输。\n\n")
    guize02(argTotal02)

    // 命令行输入参数
    /*var huochai_num int
    fmt.Printf("  请输入火柴数量(默认21)：")
    _, err :=fmt.Scan(&huochai_num)
    if err != nil {
        fmt.Printf( " %s == %s " , os.Stderr, err)
        return
    }
    fmt.Printf("  s:%s , v:%v, d:%d,c:%c,q:%q, p:%p", huochai_num,huochai_num,huochai_num, huochai_num,huochai_num,huochai_num)
    if (huochai_num < 1 || huochai_num > 40) {
        fmt.Printf("\nyour choice is %d! too big or too small，use default 21\n", huochai_num)
        huochai_num = argTotal02
    }
    guize02(huochai_num)
    */
}

// 规则2：找到规律，原来是4+1一个循环组。 参考： https://blog.csdn.net/csy981848153/article/details/9005248
func guize02(sum int) {
    //sum := 21       //tick num, 5k+1，则后手（电脑）赢；否则看对手是否失误
    computer_win := true //电脑获胜

    fmt.Printf(" 火柴数量是：%d \n\n", sum)

    countN := 0
    //game run
    for sum > 0  {
        countN++

        fmt.Printf("\n第%3d轮, now the number of ticks is: %d\n", countN, sum)

        // 剩余1个提前退出
        if sum == 1 {
            computer_win = true
            fmt.Printf("剩余1根，您不得不拿!\n")
            break
        }

        // 接收用户输入
        fmt.Printf("第%3d轮, turn to your choice(range:1~4): ", countN)
        for {
            fmt.Scan(&input)
            if (input < 1 || input > 4 || input > sum ) {   // 用户拿走的不能大于sum数量
                fmt.Printf("your choice is out of the range(1~4)，请重新输入!\n")
            } else {
                break
            }
        }
        sum -= input

        // 没有剩余，则电脑赢
        if (0 == sum) {
            fmt.Printf("您拿走了全部，抱歉，您输了!\n")
            computer_win = true
            break
        } else if (1 == sum) {
            fmt.Printf("剩余1根，电脑不得不拿!\n")
            computer_win = false
            break
        }

        // 电脑获取多少张？电脑为了赢，所以让剩下的数量为5k+1
        computer_nazou := 1
        yushu := sum % 5
        if (1 == yushu) {
            // 此时对方会赢，因此就拿走1个或者随机，等待对方失误
            computer_nazou = 1
        } else {
            computer_nazou = yushu - 1
            if (computer_nazou < 0) {   // yushu = 0整除的情况
                computer_nazou += 5
            }
        }

        fmt.Printf("第%3d轮, computer choices the num is: %d\n", countN, computer_nazou)    // 电脑总是拿走5-input数量
        sum -= computer_nazou
    }

    fmt.Println("\n")
    // 判断谁输谁赢
    if computer_win {
        fmt.Printf(" computer is the winner!")
    } else {
        fmt.Printf(" you are the winner! computer is loser")
    }

    return
}
