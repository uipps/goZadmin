
package main

import (
	"fmt"
	"flag"
	"reflect"
	)

var (
	arg string
	arg2 string
)



var parvalue = [...]int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10}

//parvalue := [...]int{100, 50, 20, 10, 5, 2, 1}
func init()  {
	flag.StringVar(&arg, "uFlags", "logoff", "Usage: shutdown logoff reboot")
	flag.StringVar(&arg2, "m", "", "1,2,3")
}

func main() {
	//arr3 := […]int{10,20,30,40,50}

	flag.Parse()

	fmt.Println(arg)
	fmt.Println(arg2)
	fmt.Println(reflect.TypeOf(arg))
	zhaolinqian()
}

// 找零，贪婪算法
func zhaolinqian() {
	//parvalue := [7]int{100, 50, 20, 10, 5, 2, 1}
	//arr3 := […]int{10,20,30,40,50}
	parvalue2 := [...]int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10}

	//fmt.Println(parvalue)
	fmt.Println(parvalue2)
	//
	//fmt.Println(parvalue[0])
	//return  1;
}