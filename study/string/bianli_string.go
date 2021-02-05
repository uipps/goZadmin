/*
golang中，用反引号定义多行字符串

go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\study\string\bianli_string.go

*/
package main

import "fmt"

func main() {
    //rune01()
    //str_bianli01()
    str_bianli02()
    //str_bianli03()
}

/*
方法一  格式化打印
'H''世''界'

方法二  转化输出格式
H
世
界
*/
func str_bianli03() {
    str := "H世界"
    fmt.Println("方法一  格式化打印")
    for _, ch1 := range str {
        fmt.Printf("%q",ch1) //单引号围绕的字符字面值，由go语法安全的转义
    }

    fmt.Println("\n\n方法二  转化输出格式")
    for _, ch2 := range str {
        fmt.Println(string(ch2))
    }
}

/*
Utf-8遍历：
72
228
184
150
231
149
140

Unicode遍历：
72
19990
30028
*/
func str_bianli02() {
    str := "H世界"
    fmt.Println("Utf-8遍历：")
    for i := 0; i < len(str); i++ { // 长度7,
        ch := str[i]
        fmt.Println(ch)
    }

    fmt.Println("\nUnicode遍历：")
    for _, ch1 := range str {   // 输出3个数字
        fmt.Println(ch1)
    }
}

// 字符串
func str_bianli01() {
    //v6 := "床前明月光,\n疑似地上霜.\n举头望明月,\n低头思故乡.\n"  // 汉字索引，间隔3，因为汉字占用3个字节；emoji表情没有测试
    /*
    0--床--int32
    3--前--int32
    6--明--int32
    9--月--int32
    12--光--int32
    15--,--int32
    16--
    --int32
    */
    v6 := "abc12356.\n"     // 单字符，其索引是连续的
    fmt.Println(v6)
    for k, v := range v6 {
        fmt.Printf("%d--%c--%T\n", k, v, v)
    }
}

func rune01() {

    // 反引号定义多行字符串
    v4 := `
床前明月光,
疑似地上霜.
举头望明月,
低头思故乡.
    `
    //v5 := "床前明月光,\n疑似地上霜.\n举头望明月,\n低头思故乡.\n"

    v6 := []rune(v4)
    fmt.Println(v4)
    fmt.Println(v6)

    v7 := "前"
    for k, v := range v6 {
        if string(v) == v7 {
            fmt.Printf("找到字符---\"%s\",\n其索引为%d\n", v7, k)
            fmt.Printf("%d--%c--%T\n", k, v, v)
        }
    }
}
