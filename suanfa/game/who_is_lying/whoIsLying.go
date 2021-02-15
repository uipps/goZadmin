/**
50.谁在说谎

题目：张三说李四在说谎，李四说王五在说谎，王五说张三和李四都在说谎。现在问：这三人中到底谁说的是真话，谁说的是假话？

参考： https://blog.csdn.net/csy981848153/article/details/7626784


go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\game\who_is_lying\whoIsLying.go

解：设张三为A，李四为B，王五为C，说真话为1，说假话为0．

　　(1)若A＝1，即张三说真话．由于张三说：“李四在说谎”，可推知B＝0．

　　而李四说：“王五在说谎”，但B＝0，李四说假话，则王五说真话C＝1；

　　由于王五说：“张三和李四都在说谎”，可知A＝0，B＝0与A＝1相矛盾．

　　故A＝1时问题无解．

　　(2)若张三说假话，即A＝0．

　　由于张三说：“李四在说谎”，可知李四说真话，即B＝1；

　　李四说：“王五在说谎”，知C＝0．

　　由于王五说：“张三和李四都说谎”，而C＝0，可得A＝1，B＝1或A＝0，B＝1，或A＝1，B＝0

　　则C＝0符合要求．

　　结论：张三、王五说假话，李四说真话．

　　思路解析：该问题看起来复杂，但若把它变为数学问题，则清晰多了．


-- 程序运行结果
张三说谎
李四没说谎
王五说谎

*/

package main

import "fmt"

func main() {
    whoIsLying()
}

func whoIsLying() {
    tem := 0
    lying := [2]string{"说谎", "没说谎"} // 0-说假话、说谎；1-真话、没说谎

    for a := 1; a >= 0; a-- {
        for b := 1; b >= 0; b-- {
            for c := 1; c >= 0; c-- {
                // 张三a说李四b在说谎; 李四b说王五c在说谎; 王五说张三和李四都在说谎
                // ( (a==1 && b+c == 0)||(a ==0 && b+c>=1) )  // (a==1 && b+c == 0) 不成立
                if (a+b == 1) && (b+c == 1) && (a == 0 && b+c >= 1) {
                    fmt.Printf("张三%s\n", lying[a])
                    fmt.Printf("李四%s\n", lying[b])
                    fmt.Printf("王五%s\n", lying[c])
                    tem = 1
                }
            }
        }
    }

    if 0 == tem {
        fmt.Printf("三个人都说谎了！")
    }
}
