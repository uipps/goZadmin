// 汉诺塔：
//  把A塔上编号从小号到大号的圆盘从A塔通过B辅助塔移动到C塔上去，要求大盘不能在小盘上面
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\hanoi\hanoi1.go -n 3
// go run ~/develope/go/go_code/src/github.com/uipps/goZadmin/suanfa/hanoi/hanoi1.go -n 3 -a A -b B -c C
//
// 不难发现规律：1个圆盘的次数 2的1次方减1
//
//　　　　　　　2个圆盘的次数 2的2次方减1
//
//                         3个圆盘的次数 2的3次方减1
//
//                         。  。   。    。   。
//
//                         n个圆盘的次数 2的n次方减1
//
// 故：移动次数为：2^n - 1
//     调用方法的栈机制:（特点：先进后出）

package main

import (
    "flag"
    "fmt"
)

var (
    hanoiNum       int    // 盘数量
    hanoiCount     int    // 移动的次数
    hanoiTowBegin  string // 起始塔
    hanoiTowFuzhu  string // 辅助塔
    hanoiTowTarget string // 目标塔
)

func init() {
    flag.IntVar(&hanoiNum, "n", 3, "Usage: 3 4 5")
    flag.StringVar(&hanoiTowBegin, "a", "A", "Usage: A B C")
    flag.StringVar(&hanoiTowFuzhu, "b", "B", "Usage: A B C")
    flag.StringVar(&hanoiTowTarget, "c", "C", "Usage: A B C")
}

func main() {
    flag.Parse()

    a := hanoiTowBegin
    b := hanoiTowFuzhu
    c := hanoiTowTarget

    // 参数校验 a,b,c必须都不相同，否则报错
    if a == b || a == c || b == c {
        fmt.Printf("三个柱子必须不同!\n")
        return
    }

    fmt.Println("******************************************************************************************")
    fmt.Println("这是汉诺塔问题（把A塔上编号从小号到大号的圆盘从A塔通过B辅助塔移动到C塔上去）")
    fmt.Println("    要求移动过程中，大盘不能放在小盘上面；从上到下编号依次为： 1 2 3 ...... n-1 n")
    fmt.Println("******************************************************************************************")
    fmt.Printf("圆盘的个数为：%d\n\n", hanoiNum)
    hanoiCount = 0
    TowersOfHanoi1(hanoiNum, a, b, c)
    fmt.Printf("\n>>移动了 %d 次，把A上的圆盘都移动到了C上\n", hanoiCount)

    return
}

func TowersOfHanoi1(num int, a string, b string, c string) {
    if 1 == num {
        printMove(num, a, c)
        //hanoiCount++
        //fmt.Printf("第%4d 次移动 : 把 %3d 号圆盘从 %s ---> %s  \n", hanoiCount, num, a, c);
    } else {
        TowersOfHanoi1(num-1, a, c, b)

        printMove(num, a, c)
        //hanoiCount++
        //fmt.Printf("第%4d 次移动 : 把 %3d 号圆盘从 %s ---> %s  \n", hanoiCount, num, a, c);
        TowersOfHanoi1(num-1, b, a, c)
    }
    return
}

func printMove(num int, a string, c string) {
    hanoiCount++
    fmt.Printf("第%4d 次移动 : 把 %3d 号圆盘从 %s ---> %s  \n", hanoiCount, num, a, c)
}
