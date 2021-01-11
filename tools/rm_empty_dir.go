// 删除空的目录
// 	go run rm_empty_dir.go -p "E:/BaiduNetdiskDownload/aa/aa_bb"

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	sourcePath1 string // 原路径
)

func init() {
	flag.StringVar(&sourcePath1, "p", "", "Usage: ")
}

func main() {
	flag.Parse()

	sourcePath1 = strings.Trim(sourcePath1, " ")             // 去掉左边空白字符
	sourcePath1 = strings.TrimRight(sourcePath1, "\\/")      // 暂不允许直接放到/目录下，去除/不影响原路径
	sourcePath1 = strings.ReplaceAll(sourcePath1, "\\", "/") // 路径统一用/分隔

	if "" == sourcePath1 || !FileOrPathExists2(sourcePath1) {
		fmt.Printf("source path %s is empty or not exist! \n", sourcePath1)
	} else {
		transPath2(sourcePath1)
	}
}

func transPath2(source_path string) {
	//fmt.Println(source_path)
	//fmt.Println(strings.ReplaceAll(filepath.Dir(source_path), "\\", "/"))
	//fmt.Println(filepath.Dir(source_path))
	//os.Exit(0)

	files, err := ioutil.ReadDir(source_path)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		//fmt.Printf("\n ------ path %s ,files :", source_path)
		//fmt.Printf("\n  ------  the varchar 'files' Type is: %T ,  v: %v ----- \n", files, files)
		//fmt.Println(len(files))
		//fmt.Println(cap(files))

		// 检查目录是否为空，非空目录不能删除，只删除空目录
		if 0 == len(files) {
			// 删除空目录
			err2 := os.RemoveAll(source_path)
			if err2 != nil {
				fmt.Println(" Remove path %s Error: %s", source_path, err2)
			} else if sourcePath1 != source_path {
				// 反过来再重新检查一遍上级目录，因为上级可能就只包含一个空的下级目录，当空下级目录删除了，上级目录也需要检查一下，然后删除
				new_path := filepath.Dir(source_path)               // 获取目录的上一级目录, 类似于PHP中dirname — 返回路径中的目录部分
				transPath2(strings.ReplaceAll(new_path, "\\", "/")) // 字符串替换，\ ==> /

				// TODO 还是可能报错： open E:/BaiduNetdiskDownload/aa/aa_bb/empty_dir: The system cannot find the file specified.
				// 		目录被删除导致的报错
			}
		} else {
			for _, v := range files {
				if !v.IsDir() {
					fmt.Println(" ------ file: " + source_path + "/" + v.Name())
					continue
				}
				transPath2(source_path + "/" + v.Name())
			}
		}
	}
}

func IsEmptyDir(source_path string) {
	files, err := ioutil.ReadDir(source_path)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("\n ------ path %s ,files :", source_path)
		fmt.Println(files)

		//for _, v := range files {
		//	if (!v.IsDir()) {
		//		continue
		//	}
		//
		//	// 检查目录是否为空，非空目录不能删除，只删除空目录
		//	fmt.Println(source_path + "/" + v.Name())
		//	transPath2(source_path + "/" + v.Name())
		//	fmt.Println(v.Name())
		//}
	}
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
