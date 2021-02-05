// 字符串反转
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\study\string\fanzhuan_string.go
package main

func ReverseStr(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

func main() {
    a := "Hello, 世界"
    println(a)
    println(ReverseStr(a))
}
