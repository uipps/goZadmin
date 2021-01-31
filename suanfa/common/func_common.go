package common

import "fmt"

func OutPrintFmt(a_pi_arr []int, xiaoshuLeng int) {
	fmt.Printf("\t---第1-1000位小数---\n")
	fmt.Printf("PI=%d.", a_pi_arr[1])
	// 小数部分要循环输出
	for i := 2; i < xiaoshuLeng; i++ {
		if i > 2 && (i-2)%10 == 0 { // 每十位输入一个空格
			fmt.Print(" ")
		}
		if i > 2 && (i-2)%50 == 0 { // 每50位换行
			fmt.Println("")
		}
		if i > 2 && (i-2)%1000 == 0 { // 每1000位, 显示一个提示
			fmt.Printf("\t---显示第%d-%d位小数---\n", (i-2)/1000*1000+1, ((i-2)/1000+1)*1000)
		}
		fmt.Printf("%d", a_pi_arr[i]) // 输出一位小数
	}
}
