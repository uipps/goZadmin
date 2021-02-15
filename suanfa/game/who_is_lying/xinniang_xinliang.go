/*
48. 新娘和新郎
题目：三对情侣参加婚礼，三个新郞为A、B、C，三个新娘为X、Y、Z。有人不知道谁和谁结婚，于是询问了六位新人中的三位，
但听到的回答是这样的：A说他将和X结婚；X说她的未婚夫是C；C说他将和Z结婚。这人听后知道他们在开玩笑，全是假话。请编程找出谁将和谁结婚。

参考： https://blog.csdn.net/csy981848153/article/details/7626779

--
解：
  ABC分别用123表示三个新郎，并且ABC的值不能相同
  XYZ分别表示新娘，XYZ的取值范围就是123，分别对应新郎的值即为匹配


-- 程序运行
go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\game\who_is_lying\xinniang_xinliang.go

-- 程序运行结果
新娘X和新郎B是一对情侣！
新娘Y和新郎C是一对情侣！
新娘Z和新郎A是一对情侣！

*/

package main

import "fmt"

func main() {
    xinniangXinlang()
}

func xinniangXinlang() {
    //A,B,C := 1,2,3  // 用1，2，3表示三个不同的新郎，
    //xinlang := map[int]string {1:"A", 2:"B", 3:"C"}   // 也可以用map
    //xinlang := [...]string {" ", "A", "B", "C"} // 也可以用数组
    xinlang := []string{" ", "A", "B", "C"} // 也可以用切片

    x, y, z := 0, 0, 0 // 新娘设定初值，找出xyz哪个跟1，2，3相等就是对应的新娘

    for x = 1; x <= 3; x++ {
        for y = 1; y <= 3; y++ {
            for z = 1; z <= 3; z++ {
                if x != 1 && x != 3 && z != 3 && x != y && y != z && x != z {
                    fmt.Printf("新娘x和新郎%s是一对情侣！\n"+
                        "新娘y和新郎%s是一对情侣！\n"+
                        "新娘z和新郎%s是一对情侣！\n", xinlang[x], xinlang[y], xinlang[z])
                }
            }
        }
    }
}
