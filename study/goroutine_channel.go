//
package main

import "fmt"

func main() {
    oneSimpleTest()

}

func oneSimpleTest()  {
    ch := make(chan string)

    go func() { ch <- "ping" }()

    msg := <-ch
    fmt.Println(msg)
}