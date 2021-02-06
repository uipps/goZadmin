package common

import (
    "fmt"
    "regexp"
    "strconv"
    "strings"
)

func OutPrintFmt(a_pi_arr []int, xiaoshuLeng int, n int, fenzi int, fenmu int) {
    fmt.Printf("\t---第1-1000位小数---\n")

    if 1 == fenmu && fenzi == 3 {
        fmt.Printf("PI=%d.", a_pi_arr[n-1]) // 计算π
    } else if 1 == fenmu && fenzi == 2 {
        fmt.Printf("e=%d.", a_pi_arr[n-1]) // 计算e
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

//
func LongNumFmtOutPrint(num_arr []int, xiaoshuLeng int, n int, fenzi int, fenmu int) {
    fmt.Printf("\t---第1-1000位小数---\n")

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
        fmt.Printf("%d", num_arr[i]) // 输出一位小数
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

// 判断字符串是否数字，包括小数点。	float64的最大正数为: 1.797693134862315708145274237317043567981e+308；
// 								float64的最小非负数: 4.940656458412465441765687928682213723651e-324
// TODO ：超出了float64范围的数字就会出错
func Is_numeric_float(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}

// 判断是否数字（101.45fdfd判断）
func Is_numeric(str1 string) bool {
    //pattern := "(-)?\\d(\\d+)?(\\.)?(\\d+)?" //反斜杠要转义
    pattern := `(-)?\d(\d+)?(\.)?(\d+)?`
    reg := regexp.MustCompile(pattern)
    str_new := reg.FindString(str1)
    if (str_new == str1) {
        return true
    }
    return false
}

// 是否数字 参考： https://www.cnblogs.com/syyong/p/8940902.html
func IsNumeric(val interface{}) bool {
    switch val.(type) {
    case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
    case float32, float64, complex64, complex128:
        return true
    case string:
        str := val.(string)
        if str == "" {
            return false
        }
        // Trim any whitespace
        str = strings.Trim(str, " \\t\\n\\r\\v\\f")
        if str[0] == '-' || str[0] == '+' {
            if len(str) == 1 {
                return false
            }
            str = str[1:]
        }
        // hex
        if len(str) > 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X') {
            for _, h := range str[2:] {
                if !((h >= '0' && h <= '9') || (h >= 'a' && h <= 'f') || (h >= 'A' && h <= 'F')) {
                    return false
                }
            }
            return true
        }
        // 0-9,Point,Scientific
        p, s, l := 0, 0, len(str)
        for i, v := range str {
            if v == '.' { // Point
                if p > 0 || s > 0 || i+1 == l {
                    return false
                }
                p = i
            } else if v == 'e' || v == 'E' { // Scientific
                if i == 0 || s > 0 || i+1 == l {
                    return false
                }
                s = i
            } else if v < '0' || v > '9' {
                return false
            }
        }
        return true
    }

    return false
}

// 字符串反转
func ReverseStr(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

func ReverseNum(r string) string {
    j := len(r) - 1
    c := []int32(r)
    for _, item := range r {
        c[j] = item
        j--
    }
    return string(c)
}
