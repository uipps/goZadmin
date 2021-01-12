// 将某文件或某一类型的文件，移动到指定的目录下
//		例如，将分散的mp4视频文件，转移到统一的目录下
// 	go run move_file.go -p "E:/BaiduNetdiskDownload/aa/aa_bb" -t "E:/qianfeng_golang_mp4/" -T ".mp4"
//  go run move_file.go -p "H:/qianfeng_golang_mp4" -t "H:/qianfeng_golang_mp4" -M "fileNameReplaceStr" -S "千锋Go语言教程：/千锋GO语言教程：" -R ""
//  go run move_file.go -p "H:/qianfeng_golang_mp4" -t "H:/qianfeng_golang_mp4" -M "fileNameReplaceStr" -S "千锋Go语言教程："

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var (
	sourcePath string // 原路径
	toPath     string // 目标路径
	isSon      int    // 是否递归子目录
	fileType   string // 需要移动的文件类型
	methodFunc string // 执行方法
	strOrig    string // 被替换的字符串，用“/”分隔
	replaceStr string // 替换成的字符串，
)

func init() {
	flag.StringVar(&sourcePath, "p", "", "Usage: ")
	flag.StringVar(&toPath, "t", "", "Usage: ")
	flag.StringVar(&fileType, "T", ".mp4", "Usage: 1 0")
	flag.StringVar(&methodFunc, "M", "", "fileNameReplaceStr")
	flag.StringVar(&strOrig, "S", "", "需要被替换的字符串，多个用“/”分隔")
	flag.StringVar(&replaceStr, "R", "", "替换成的字符串，暂时仅仅支持一条")
	flag.IntVar(&isSon, "i", 1, "Usage: 1 0")
}

func main() {
	flag.Parse()

	//println(toPath)
	toPath = strings.TrimRight(toPath, "\\/ ")         // 暂不允许直接放到/目录下，去除/不影响原路径
	sourcePath = strings.TrimRight(sourcePath, "\\/ ") // 暂不允许直接放到/目录下，去除/不影响原路径
	//fmt.Printf("--%s--", toPath)
	//os.Exit(0)

	// 参数校验, 目标目录不存在或不是目录
	if "" == toPath || !FileOrPathExists(toPath) {
		fmt.Printf("target path %s is empty or not exist! \n", toPath)
		return
	}
	//

	if "" == sourcePath || !FileOrPathExists(sourcePath) {
		fmt.Printf("source path %s is empty or not exist! \n", sourcePath)
	} else {
		if ("fileNameReplaceStr" == methodFunc) {
			fileNameReplaceStr(sourcePath)
		} else {
			transPath(sourcePath)
		}

	}

	return
}

func transPath(source_path string) {
	fileExt := ""
	files, err := ioutil.ReadDir(source_path)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, v := range files {
			//fmt.Printf("--- %v --- \n --- %q --- \n --- %p --- \n", v, v, v)
			//fmt.Println("\n\n\n\n")
			if v.IsDir() { //	&& 1 == isSon
				// fmt.Println(source_path + "/" + v.Name())
				transPath(source_path + "/" + v.Name())
			} else {
				fileExt = path.Ext(v.Name()) // 后缀
				if fileType == strings.ToLower(fileExt) {
					fmt.Println(source_path + "/" + v.Name())
					FilePutContents("D:/temp/go_write_test.txt", source_path+"/"+v.Name()+"\r\n")

					// 重命名文件，即剪切到新位置，必须相同的盘符，否则：The system cannot move the file to a different disk drive.
					err := os.Rename(source_path+"/"+v.Name(), toPath+"/"+v.Name())
					if err != nil {
						//如果重命名文件失败,则输出错误 file rename Error!
						fmt.Println("file rename Error!")
						//打印错误详细信息
						fmt.Printf("%s", err)
					} else {
						//如果文件重命名成功,则输出 file rename OK!
						fmt.Println("file rename OK!")
					}
				}
			}
		}
	}
}

//
func tOne(s_path string, s_file string) {

}

// 判断文件或目录是否存在
func FileOrPathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil { //文件或者目录存在
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

//使用ioutil.WriteFile方式写入文件,是将[]byte内容写入文件,如果content字符串中没有换行符的话，默认就不会有换行符
func FilePutContents(name string, content string) {
	data := []byte(content)
	if ioutil.WriteFile(name, data, 0775) == nil {
		//fmt.Println("写入文件成功:", content)
	}
}

func fileNameReplaceStr(source_path string) {
	files, err := ioutil.ReadDir(source_path)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, v := range files {
			//fmt.Printf("--- %v --- \n --- %q --- \n --- %p --- \n", v, v, v)
			//fmt.Println("\n\n\n\n")
			if v.IsDir() { //	&& 1 == isSon
				// fmt.Println(source_path + "/" + v.Name())
				transPath(source_path + "/" + v.Name())
			} else {
				fmt.Println(source_path + "/" + v.Name())
				FilePutContents("D:/temp/go_write_test.txt", source_path+"/"+v.Name()+"\r\n")

				// 重命名文件，即剪切到新位置，必须相同的盘符，否则：The system cannot move the file to a different disk drive.
				tmpSplit := strings.Split(strOrig, "/")
				new_file := v.Name()
				for _, str01 := range tmpSplit {
					//fmt.Printf("%T, %s \n", str01,str01)
					new_file = strings.ReplaceAll(new_file, str01, replaceStr)
				}

				fmt.Printf(" -- new file name is %s \n", new_file)
				err := os.Rename(source_path+"/"+v.Name(), toPath+"/"+new_file)
				if err != nil {
					//如果重命名文件失败,则输出错误 file rename Error!
					fmt.Println("file rename Error!")
					//打印错误详细信息
					fmt.Printf("%s", err)
				} else {
					//如果文件重命名成功,则输出 file rename OK!
					fmt.Println("file rename OK!")
				}

			}
		}
	}
}
