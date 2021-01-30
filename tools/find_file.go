// 查找目录
// go run E:/develope/go/go_code_path/src/github.com/uipps/goZadmin/tools/find_file.go -d



package main

import (
    "fmt"
    "io/ioutil"
    //"flag"
    "os"
)

func main() {
    //srcDir := "D:\\test\\项目"
    getFileList("D:/www/tests/namespace_php")

    fmt.Println(os.Args)
    fmt.Print(os.Args)
    fmt.Println(len(os.Args))
    fmt.Println(os.Args[0])
    fmt.Println(os.Args[1])
    fmt.Println(os.Args[2])
    //pathSeparator := string(os.PathSeparator)
    ///level := 1
    //listAllFileByName(level, pathSeparator, srcDir)
}

func getFileList(path string) {
    fs,_:= ioutil.ReadDir(path)
    for _,file:=range fs{
        if file.IsDir(){
            fmt.Println(" dir: " + path + "/" + file . Name())
            getFileList(path + "/" + file.Name()+"/")
        }else{
            fmt.Println(" file: " + path  + "/" + file.Name())
        }
    }
}