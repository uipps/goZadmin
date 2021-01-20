// 阶乘 5*4*3*2*1
// 	go run jiecheng.go -n 5
// 	go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\suanfa\math\jiecheng.go -n 5
package main

import (
    "flag"
    "fmt"
)

var (
    argN int64
)

func init() {
    flag.Int64Var(&argN, "n", 0, "Usage: 5 6 7")
}

func main() {
    flag.Parse()

    var i int64

    for i = 1; i <= argN; i++ {
        fmt.Printf("%2d! = %d\n", i, jiecheng(i))
    }
}

func jiecheng(i int64) int64 {
    if i <= 1 {
        return i
    }
    return i * jiecheng(i-1)
}
