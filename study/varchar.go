// 变量声明及作用域
package main

import (
    "fmt"
    "runtime"
)

var x = 10
var y = 10    // 全局变量可以不使用

func main() {
    fmt.Println("FuncName=", y)

    funcName, file, line, ok := runtime.Caller(0) // 获取行号等信息
    if (ok) {
        fmt.Println("FuncName=" + runtime.FuncForPC(funcName).Name())
    }

    fmt.Printf("x:%v, line:%s , file:%s , funcName:%v, ok:%v\n", x, line, file, funcName, ok)
    x := 1
    fmt.Println(x)

    {
        fmt.Println(x)
        x := 2
        fmt.Println(x)
    }
    fmt.Println(x)

    // 交换两个变量的值
    a, b := 11, 21
    fmt.Printf("a: %d , b: %d \n", a, b)
    a, b = b, a
    fmt.Printf("a: %d , b: %d \n", a, b)
}
