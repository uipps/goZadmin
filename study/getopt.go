package main

import (
	"fmt"
	"flag"
)

var (
	arg string
	arg2 string
	arg3 int
)

func init()  {
	flag.StringVar(&arg, "uFlags", "logoff", "Usage: shutdown logoff reboot")
	flag.StringVar(&arg2, "m", "", "1,2,3")
	flag.IntVar(&arg3, "n", "", "1,2,3")
}

func main() {

	flag.Parse()

	fmt.Println(arg)
	fmt.Println(arg2)
	fmt.Println(arg3)
	//fmt.Println(reflect.TypeOf(arg))
}
