// 正则匹配
//  go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\study\regexp\regexp01.go -s 134.5tff877
//  go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\study\regexp\regexp01.go -s 134.5877
//
package main

import (
    "bytes"
    "flag"
    "fmt"
    "regexp"
)

var (
    argString01 string
)

func init() {
    flag.StringVar(&argString01, "s", "123456", "Usage: 1233")
}

func Match(str1 string, pattern string) {
    //result, _ := regexp.MatchString(pattern, str)
    //r := regexp.MustCompile(pattern)
    fmt.Printf("regexp.Match: ")
    fmt.Println(regexp.Match(pattern, []byte(str1)))
}

func FindString(str1 string, pattern string) {
    reg := regexp.MustCompile(pattern)
    rlt := reg.FindString(str1)
    fmt.Printf("FindString: %T - %q\n", rlt, rlt)
}

func FindAll(str1 string, pattern string) {
    reg := regexp.MustCompile(pattern)
    fmt.Printf("FindAll：%q\n", reg.FindAll([]byte(str1), -1))
}

func FindAllString(str1 string, pattern string) {
    reg := regexp.MustCompile(pattern)
    fmt.Printf("findAllString：%q\n", reg.FindAllString(str1, -1))
}

func FindIndex(str1 string, pattern string) {
    reg := regexp.MustCompile(pattern)
    fmt.Printf("FindIndex: ")
    fmt.Println(reg.FindIndex([]byte(str1)))
}

func FindStringIndex(str1 string, pattern string) {
    reg := regexp.MustCompile(pattern)
    fmt.Printf("FindStringIndex: ")
    fmt.Println(reg.FindStringIndex(str1))
}

func FindReaderIndex(str1 string, pattern string) {
    r := bytes.NewReader([]byte(str1))
    reg := regexp.MustCompile(pattern)
    fmt.Printf("FindReaderIndex: ")
    fmt.Println(reg.FindReaderIndex(r))
}

func FindSubmatch(str1 string, pattern string) {
    reg := regexp.MustCompile(pattern)
    fmt.Printf("FindSubmatch: %q\n", reg.FindSubmatch([]byte(str1)))
}

func FindAllSubmatch(str1 string, pattern string) {
    reg := regexp.MustCompile(pattern)
    fmt.Printf("FindAllSubmatch: %q\n", reg.FindAllSubmatch([]byte(str1), -1))
}

func SubexpNames(str1 string, pattern string) {
    reg := regexp.MustCompile(pattern)
    fmt.Printf("SubexpNames: %q\n", reg.SubexpNames())
}

func main() {
    flag.Parse()

    //pattern := "(-)?\\d(\\d+)?(\\.)?(\\d+)?" //反斜杠要转义
    pattern := `(-)?\d(\d+)?(\.)?(\d+)?`

    Match(argString01, pattern)
    FindString(argString01, pattern)
    FindAll(argString01, pattern)
    FindAllString(argString01, pattern)
    FindIndex(argString01, pattern)
    FindStringIndex(argString01, pattern)
    FindReaderIndex(argString01, pattern)
    fmt.Println("\n")
    FindSubmatch(argString01, pattern)      // ["134.5" "" "34" "." "5"] , 原字符串：134.5tff877
    FindAllSubmatch(argString01, pattern)   // [["134.5" "" "34" "." "5"] ["877" "" "77" "" ""]] , 原字符串：134.5tff877
    SubexpNames(argString01, pattern)


    return
    //是否匹配字符串
    // .匹配任意一个字符 ，*匹配零个或多个 ，优先匹配更多(贪婪)
    match, _ := regexp.MatchString("H(.*)d!", "Hello World!")
    fmt.Println(match) //true
    //或
    match, _ = regexp.Match("H(.*)d!", []byte("Hello World!"))
    fmt.Println(match) //true
    //或通过`Compile`来使用一个优化过的正则对象
    r, _ := regexp.Compile("H(.*)d!")
    fmt.Println(r.MatchString("Hello World!")) //true

    // 这个方法返回匹配的子串
    fmt.Println(r.FindString("Hello World! world")) //Hello World!
    //同上
    fmt.Println(string(r.Find([]byte("Hello World!")))) //Hello World!

    // 这个方法查找第一次匹配的索引
    // 的起始索引和结束索引，而不是匹配的字符串
    fmt.Println(r.FindStringIndex("Hello World! world")) //[0 12]

    // 这个方法返回全局匹配的字符串和局部匹配的字符，匹配最大的子字符串一次。
    // 它和r.FindAllStringSubmatch("Hello World! world"，1) 等价。  比如
    // 这里会返回匹配`H(.*)d!`的字符串
    // 和匹配`(.*)`的字符串
    fmt.Println(r.FindStringSubmatch("Hello World! world")) //[Hello World! ello Worl]

    // 和上面的方法一样，不同的是返回全局匹配和局部匹配的
    // 起始索引和结束索引
    fmt.Println(r.FindStringSubmatchIndex("Hello World! world")) //[0 12 1 10]
    // 这个方法返回所有正则匹配的字符，不仅仅是第一个
    fmt.Println(r.FindAllString("Hello World! Held! world", -1)) //[Hello World! Held!]

    // 这个方法返回所有全局匹配和局部匹配的字符串起始索引,只匹配最大的串
    // 和结束索引
    fmt.Println(r.FindAllStringSubmatchIndex("Hello World! world", -1))       //[[0 12 1 10]]
    fmt.Println(r.FindAllStringSubmatchIndex("Hello World! Held! world", -1)) //[[0 18 1 16]]

    // 为这个方法提供一个正整数参数来限制匹配数量
    res, _ := regexp.Compile("H([a-z]+)d!")
    fmt.Println(res.FindAllString("Hello World! Held! Hellowrld! world", 2)) //[Held! Hellowrld!]

    fmt.Println(r.FindAllString("Hello World! Held! world", 2)) //[Hello World! Held!]
    //注意上面两个不同，第二参数是一最大子串为单位计算。

    // regexp包也可以用来将字符串的一部分替换为其他的值
    fmt.Println(r.ReplaceAllString("Hello World! Held! world", "html")) //html world

    // `Func`变量可以让你将所有匹配的字符串都经过该函数处理
    // 转变为所需要的值
    in := []byte("Hello World! Held! world")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))

    // 在 b 中查找 reg 中编译好的正则表达式，并返回第一个匹配的位置
    // {起始位置, 结束位置}
    b := bytes.NewReader([]byte("Hello World!"))
    reg := regexp.MustCompile(`\w+`)
    fmt.Println(reg.FindReaderIndex(b)) //[0 5]

    // 在 字符串 中查找 r 中编译好的正则表达式，并返回所有匹配的位置
    // {{起始位置, 结束位置}, {起始位置, 结束位置}, ...}
    // 只查找前 n 个匹配项，如果 n < 0，则查找所有匹配项

    fmt.Println(r.FindAllIndex([]byte("Hello World!"), -1)) //[[0 12]]
    //同上
    fmt.Println(r.FindAllStringIndex("Hello World!", -1)) //[[0 12]]

    // 在 s 中查找 re 中编译好的正则表达式，并返回所有匹配的内容
    // 同时返回子表达式匹配的内容
    // {
    //     {完整匹配项, 子匹配项, 子匹配项, ...},
    //     {完整匹配项, 子匹配项, 子匹配项, ...},
    //     ...
    // }
    // 只查找前 n 个匹配项，如果 n < 0，则查找所有匹配项
    reg = regexp.MustCompile(`(\w)(\w)+`)                      //[[Hello H o] [World W d]]
    fmt.Println(reg.FindAllStringSubmatch("Hello World!", -1)) //[[Hello H o] [World W d]]

    // 将 template 的内容经过处理后，追加到 dst 的尾部。
    // template 中要有 $1、$2、${name1}、${name2} 这样的“分组引用符”
    // match 是由 FindSubmatchIndex 方法返回的结果，里面存放了各个分组的位置信息
    // 如果 template 中有“分组引用符”，则以 match 为标准，
    // 在 src 中取出相应的子串，替换掉 template 中的 $1、$2 等引用符号。
    reg = regexp.MustCompile(`(\w+),(\w+)`)
    src := []byte("Golang,World!")           // 源文本
    dst := []byte("Say: ")                   // 目标文本
    template := []byte("Hello $1, Hello $2") // 模板
    m := reg.FindSubmatchIndex(src)          // 解析源文本
    // 填写模板，并将模板追加到目标文本中
    fmt.Printf("%q", reg.Expand(dst, template, src, m))
    // "Say: Hello Golang, Hello World"

    // LiteralPrefix 返回所有匹配项都共同拥有的前缀（去除可变元素）
    // prefix：共同拥有的前缀
    // complete：如果 prefix 就是正则表达式本身，则返回 true，否则返回 false
    reg = regexp.MustCompile(`Hello[\w\s]+`)
    fmt.Println(reg.LiteralPrefix())
    // Hello false
    reg = regexp.MustCompile(`Hello`)
    fmt.Println(reg.LiteralPrefix())
    // Hello true

    text := `Hello World! hello world`
    // 正则标记“非贪婪模式”(?U)
    reg = regexp.MustCompile(`(?U)H[\w\s]+o`)
    fmt.Printf("%q\n", reg.FindString(text)) // Hello
    // 切换到“贪婪模式”
    reg.Longest()
    fmt.Printf("%q\n", reg.FindString(text)) // Hello Wo

    // 统计正则表达式中的分组个数（不包括“非捕获的分组”）
    fmt.Println(r.NumSubexp()) //1

    //返回 r 中的“正则表达式”字符串
    fmt.Printf("%s\n", r.String())

    // 在 字符串 中搜索匹配项，并以匹配项为分割符，将 字符串 分割成多个子串
    // 最多分割出 n 个子串，第 n 个子串不再进行分割
    // 如果 n < 0，则分割所有子串
    // 返回分割后的子串列表
    fmt.Printf("%q\n", r.Split("Hello World! Helld! hello", -1)) //["" " hello"]

    // 在 字符串 中搜索匹配项，并替换为 repl 指定的内容
    // 如果 rep 中有“分组引用符”（$1、$name），则将“分组引用符”当普通字符处理
    // 全部替换，并返回替换后的结果
    s := "Hello World, hello!"
    reg = regexp.MustCompile(`(Hell|h)o`)
    rep := "${1}"
    fmt.Printf("%q\n", reg.ReplaceAllLiteralString(s, rep)) //"${1} World, hello!"

    // 在 字符串 中搜索匹配项，然后将匹配的内容经过 repl 处理后，替换 字符串 中的匹配项
    // 如果 repb 的返回值中有“分组引用符”（$1、$name），则将“分组引用符”当普通字符处理
    // 全部替换，并返回替换后的结果
    ss := []byte("Hello World!")
    reg = regexp.MustCompile("(H)ello")
    repb := []byte("$0$1")
    fmt.Printf("%s\n", reg.ReplaceAll(ss, repb))
    // HelloH World!

    fmt.Printf("%s\n", reg.ReplaceAllFunc(ss,
        func(b []byte) []byte {
            rst := []byte{}
            rst = append(rst, b...)
            rst = append(rst, "$1"...)
            return rst
        }))
    // Hello$1 World!

}
