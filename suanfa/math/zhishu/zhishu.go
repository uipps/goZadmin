// 输出1000（n可变值）以内的所有质数。并输出计算次数进行对比
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\math\zhishu\zhishu.go -n 1000
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\math\zhishu\zhishu.go -n 1000000000 -m 100

// go run ~/develope/go/go_code/src/github.com/uipps/goZadmin/suanfa/math/zhishu/zhishu.go -n 1000
// go run ~/develope/go/go_code/src/github.com/uipps/goZadmin/suanfa/math/zhishu/zhishu.go -n 1000000000 -m 100
package main

import (
    "flag"
    "fmt"
    "time"

    "github.com/uipps/goZadmin/suanfa/common/jinzhiToAny"
)

var (
    argN int
    argM int
)

func init() {
    flag.IntVar(&argN, "n", 100, "Usage: 10 100 1000 10000 ")
    flag.IntVar(&argM, "m", -1, "Usage: 10000 1000000000 ")
}

func main() {
    flag.Parse()

    startTime := time.Now().UnixNano()
    //fmt.Println("startTime：", startTime)
    fmt.Printf("startTime：%d, %s\n", startTime/1e3, time.Unix(0, startTime).Format("2006-01-02 15:04:05"))

    if (argM > 0) {
        fmt.Printf(" 第二种方法primeNum02：\n")
        primeYiNei02(argN, argM)

        // 执行时间计算
        endTime := time.Now().UnixNano()
        fmt.Printf("  第二种方法endTime：%d, %s ；也是下面方面的开始时间\n", endTime/1e3, time.Unix(0, endTime).Format("2006-01-02 15:04:05"))
        nanoSeconds := float64(endTime-startTime) / 1e3
        seconds := nanoSeconds / 1e6
        fmt.Printf("第二种方法spendTime：%.2f s , %.2f ns", seconds, nanoSeconds)

        fmt.Println("\n")

        fmt.Printf(" 第一种方法primeNum01：\n")
        primeYiNei01(argN, argM)
        endTime2 := time.Now().UnixNano()
        fmt.Printf("  第一种方法endTime：%d, %s\n", endTime2/1e3, time.Unix(0, endTime2).Format("2006-01-02 15:04:05"))
        nanoSeconds = float64(endTime2-endTime) / 1e3
        seconds = nanoSeconds / 1e6
        fmt.Printf("第一种方法spendTime：%.2f s , %.2f ns", seconds, nanoSeconds)

    } else {

        fmt.Printf(" 第一种方法primeNum01：\n")
        primeNum01(argN)

        fmt.Println("\n")

        fmt.Printf(" 第二种方法primeNum02：\n")
        primeNum02(argN)
    }

    fmt.Println("\n")

    // 执行时间计算
    endTime := time.Now().UnixNano()
    fmt.Printf("  总endTime：%d, %s\n", endTime/1e3, time.Unix(0, endTime).Format("2006-01-02 15:04:05"))
    nanoSeconds := float64(endTime-startTime) / 1e3
    seconds := nanoSeconds / 1e6
    fmt.Printf("总spendTime：%.2f s , %.2f ns", seconds, nanoSeconds)
    return
}

// 1.1 质数的定义进行遍历，时间复杂度最大的
func primeNum01(n int) int {
    runCounter := 0
    counter := 0
    i := 0

    for no := 2; no <= n; no++ {
        for i = 2; i < no; i++ {
            runCounter++
            if 0 == no%i {
                break
            }
        }
        if no == i {
            counter++
            // 直到最后也没有找出
            fmt.Printf("%8d, %8o, %8X, %20b \n", no, no, no, no)
            //fmt.Printf("Line %4d: %8d, %8o, %8X, %20b \n", counter, no, no, no, no)
            //fmt.Printf("%d\n", no)
        }
    }

    fmt.Printf("运算次数 %d\n", runCounter)
    fmt.Printf("%d 以内的质数个数 %d\n", n, counter)
    return 0
}

// 1.2 除的时候，只需要no开方以下即可，不需要尝试到no-1; 17只需要尝试4以下的，大于4的不需要尝试
func primeNum02(n int) int {
    runCounter := 0
    counter := 0
    i := 0
    isPrime := 1 // 质数
    jinzhiTo3 := ""
    jinzhiTo5 := ""
    jinzhiTo7 := ""

    // 2也是质数
    no := 2
    counter++
    //jinzhiTo7 = jinzhiToAny.DecimalToAny(no, 7)
    fmt.Printf("%8d, \n", no)
    //fmt.Printf("%8d, %8s, %8o, %8X, %20b \n", no, jinzhiTo7, no, no, no)
    //fmt.Printf("%8d, %8o, %8X, %20b \n", no, no, no, no)

    // 大于2的偶数不可能，所以只在奇数中找
    for no = 3; no <= n; no += 2 {
        isPrime = 1
        for i = 2; i*i <= no; i++ {
            runCounter++
            if 0 == no%i {
                isPrime = 0 // 合数
                break
            }
        }
        if 1 == isPrime {
            counter++
            // 直到最后也没有找出
            jinzhiTo3 = jinzhiToAny.DecimalToAny(no, 3)
            jinzhiTo5 = jinzhiToAny.DecimalToAny(no, 5)
            jinzhiTo7 = jinzhiToAny.D十进制转换(no, 7)
            fmt.Printf("%8d, %8s, %8s, %8s, %8o, %8X, %20b \n", no, jinzhiTo3, jinzhiTo5, jinzhiTo7, no, no, no)
            //fmt.Printf("Line %4d: %8d, %8o, %8X, %20b \n", counter, no, no, no, no)
            //fmt.Printf("%d\n", no)
        }
    }

    fmt.Printf("运算次数 %d\n", runCounter)
    fmt.Printf("%d 以内的质数个数 %d\n", n, counter)
    return 0
}

// 1.1_2 （从大到小遍历小于n的m个数）;质数的定义进行遍历，时间复杂度最大的
func primeYiNei01(n, m int) int {
    if (m <= 0) {
        return primeNum01(n)
    }
    if (m > n) {
        m = n     // 最大也就是把n全部遍历
    }

    runCounter := 0
    counter := 0
    i := 0

    for no := n; no > n - m; no-- {
        if (no < 2) {
            continue   // no递减可能是1，1不是质数
        }
        for i = 2; i < no; i++ {
            runCounter++
            if 0 == no%i {
                break
            }
        }
        // 直到最后也没有找出，或2也是质数
        if no == i {
            counter++
            fmt.Printf("%8d, %8o, %8X, %20b \n", no, no, no, no)
        }
    }

    fmt.Printf("运算次数 %d\n", runCounter)
    fmt.Printf("%d 以内（从大到小遍历%d次）的质数个数: %d\n", n, m, counter)
    return 0
}

// 1.2_2 （从大到小遍历小于n的m个数）;除的时候，只需要no开方以下即可，不需要尝试到no-1; 17只需要尝试4以下的，大于4的不需要尝试
func primeYiNei02(n, m int) int {
    if (m <= 0) {
        return primeNum01(n)
    }
    if (m > n) {
        m = n     // 最大也就是把n全部遍历
    }

    runCounter := 0
    counter := 0
    i := 0
    isPrime := 1 // 质数
    jinzhiTo3 := ""
    jinzhiTo5 := ""
    jinzhiTo7 := ""

    // 大于2的偶数不可能，所以只在奇数中找
    for no := n; no > n - m; no-- {
        if (no < 2) {
            continue   // no递减可能是1，1不是质数
        }
        isPrime = 1
        for i = 2; i*i <= no; i++ {
            runCounter++
            if 0 == no%i {
                isPrime = 0 // 合数
                break
            }
        }
        if 1 == isPrime {
            counter++
            // 直到最后也没有找出
            jinzhiTo3 = jinzhiToAny.DecimalToAny(no, 3)
            jinzhiTo5 = jinzhiToAny.DecimalToAny(no, 5)
            jinzhiTo7 = jinzhiToAny.D十进制转换(no, 7)
            fmt.Printf("%8d, %8s, %8s, %8s, %8o, %8X, %30b \n", no, jinzhiTo3, jinzhiTo5, jinzhiTo7, no, no, no)
        }
    }

    fmt.Printf("运算次数 %d\n", runCounter)
    fmt.Printf("%d 以内（从大到小遍历%d次）的质数个数: %d\n", n, m, counter)
    return 0
}
