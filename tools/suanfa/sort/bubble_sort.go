// 冒泡排序
// go run bubble_sort.go -a "4 93 84 85 80 37 81 93 27 12 1"
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\suanfa\sort\bubble_sort.go -a "69 65 90 37 92 6 1"

package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var (
	arr1 []int
	str1 string
)

func init() {
	flag.StringVar(&str1, "a", "69 65 90 37 92 6", "Usage: 69 65 90 37 1 92 6")
}

func main() {
	flag.Parse()

	// 解析字符串
	fmt.Println(str1)
	arr2 := strings.Split(str1, " ") // 重置
	fmt.Println(arr2)

	// 数字字符串转成int类型
	for _, val := range arr2 {
		if "" == val {
			continue // 过滤空白
		}
		i, _ := strconv.Atoi(val) // 字符串转int类型
		arr1 = append(arr1, i)
	}

	// 冒泡排序
	bubbleSort(arr1)
	fmt.Println(arr1)
}

// 冒泡排序
func bubbleSort(arr []int) {
	n := len(arr)
	flag := 0

	for i := 0; i < n; i++ {
		/*// 这个第一位总是最小，时间复杂度O(n^2)，从左边第一个开始，每次都能找到其i+i后最小的，并放到i位置上，也就是找最小放左边
		for j := i + 1; j < n; j++ {
		    if (arr[i] > arr[j]) {
		        arr[i], arr[j] = arr[j], arr[i] // 交换
		    }
		}*/
		// 时间复杂度O(n^2)
		/*for j := 1; j < n; j++ {
		    if (arr[j-1] > arr[j]) {
		        arr[j-1], arr[j] = arr[j], arr[j-1] // 交换
		    }
		}*/

		// 时间复杂度更低O(n*logn)，就是从第一项开始，每项都跟其后的 数据比较，这个冒泡排序找最大的一个往右放
		for j := 1; j < n-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1] // 交换
				flag = 1                            // 设置交换标记
			}
		}
		fmt.Printf("第%2d遍：", i+1)
		fmt.Println(arr)
		if 0 == flag { //  没有交换，退出循环，后面不用做了，减少了多余的判断
			break
		}
	}
}
