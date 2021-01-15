// 将某文件或文件夹删除
//  go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\tools\rm_file_dir.go -p "F:/develope/go/go_code_path/src/github.com/rubyhan1314/go_foundation" -n ".idea"
//  go run rm_file_dir.go -p "F:/develope/go/go_code_path/src/github.com/rubyhan1314/go_foundation" -n ".idea"

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	sourcePath2 string // 原路径
	fileName2   string // 替换成的字符串，
)

func init() {
	flag.StringVar(&sourcePath2, "p", "", "Usage: ")
	flag.StringVar(&fileName2, "n", "", "Usage: ")
}

func main() {
	flag.Parse()

	sourcePath2 = strings.TrimRight(sourcePath2, "\\/ ") // 原路径
	//fmt.Printf("--%s--", fileName2)
	//os.Exit(0)

	// 参数校验, 目标目录不存在或不是目录
	if "" == sourcePath2 || !FileOrPathExists2(sourcePath2) {
		fmt.Printf("source path %s is empty or not exist! \n", sourcePath2)
	} else {
		transPathRm(sourcePath2)
	}

	return
}

func transPathRm(source_path string) {
	// 读取文件、目录列表
	files, err := ioutil.ReadDir(source_path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, v := range files {
		fmt.Printf("--- %v --- \n --- %q --- \n --- %p --- \n", v, v, v)
		fmt.Println("\n\n\n\n")

		// 删除文件或目录
		if (v.Name() == fileName2) {
			if v.IsDir() {
				err = os.RemoveAll(source_path + "/" + v.Name())
			} else {
				err = os.Remove(source_path + "/" + v.Name())
			}
			if err != nil {
				// 删除失败
				fmt.Println("删除失败")
			} else {
				// 删除成功
				fmt.Println("删除成功")
			}
		} else if v.IsDir() {
			transPathRm(source_path + "/" + v.Name())
		}
	}

	return
}

// 判断文件或目录是否存在
func FileOrPathExists2(path string) bool {
	_, err := os.Stat(path)
	if err == nil { //文件或者目录存在
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
