package main

import (
    "fmt"
    "time"
)
type newField string

func (f newField) print() {
    //fmt.Println(f)
    fmt.Printf("\n print %s \n", f)
}

func main() {
    data := []newField{"one", "two", "thred"}
    for _, v := range data {
        v.print() //print one two three
        //go v.print() //print one two three
        //go func(){fmt.Printf(" | goroutine, %s | \n",v)}() // print three three three
        fmt.Println(v)
    }
    time.Sleep(3 * time.Second)
}

/*

func main() {

    //创建一个通道
    ch := make(chan int)
    //开启一个goroutine
    go func() {
        //无限循环
        for {
            //往通道里发送数据
            ch <- 1
            //睡眠一秒
            time.Sleep(time.Second)
        }
    }()
    //无限循环从通道中读取数据
    for i := range ch {
        fmt.Println(i)
    }

}
*/