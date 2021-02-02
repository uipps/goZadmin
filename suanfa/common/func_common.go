package common

import "fmt"

func OutPrintFmt(a_pi_arr []int, xiaoshuLeng int, n int, fenzi int, fenmu int) {
	fmt.Printf("\t---第1-1000位小数---\n")

	if (1 == fenmu && fenzi == 3) {
		fmt.Printf("PI=%d.", a_pi_arr[n-1])	// 计算π
	} else {
		fmt.Printf("%d/%d=%d.", fenzi, fenmu, a_pi_arr[n-1])
	}

	// 小数部分要循环输出
	//n := 1 // 小数点开始的序号
	for i := n; i < xiaoshuLeng+n; i++ {
		if i > n && (i-n)%10 == 0 { // 每十位输入一个空格
			fmt.Print(" ")
		}
		if i > n && (i-n)%50 == 0 { // 每50位换行
			fmt.Println("")
		}
		if i > n && (i-n)%1000 == 0 { // 每1000位, 显示一个提示
			fmt.Printf("\t---显示第%d-%d位小数---\n", (i-n)+1, i-n+1000)
		}
		fmt.Printf("%d", a_pi_arr[i]) // 输出一位小数
	}
}

func Int_in_array(slice []int, val int) (int, bool) {
	for key, item := range slice {
		if item == val {
			return key, true
		}
	}
	return -1, false
}