// 测试主程序，有main包和main函数
package main

import (
    "fmt"
    "github.com/uipps/goZadmin/redis_go"
    //go_redis "github.com/uipps/goZadmin/redis_go"

    "github.com/uipps/goZadmin/test01"
    "github.com/uipps/goZadmin/test02"
)

const s string = "abcde"

func init() {
    fmt.Println("init , main.main\n")
}

func main() {
    test01.Test01()
    test02.Test02()
    var str01 string = "haha"
    var int01, int02 int = 2, 3
    const C = 3e20
    fmt.Println(str01)
    fmt.Println(int64(int01), int02, C)
    fmt.Println(s)
    fmt.Println("go,test01")
    fmt.Println("1+1=", 1+1)
    fmt.Println(true)

    redis_go.Redis_go()
}


