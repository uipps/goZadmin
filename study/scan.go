package main

import "fmt"

func main() {
	var (
		name    string
		age     int
		married bool
	)
	//fmt.Scan(&name, &age, &married)
	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married) // 输入完整格式： "1:chfeng 2:18 3:true"
	//fmt.Scanf("name:%s age:%d married:%t", &name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}
