// n最大支持76进制，最小2进制。
// num数值必须是正整数
package jinzhiToAny

import (
    "math"
    "strconv"
    "strings"
)

var tenToAny map[int]string = map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i", 19: "j", 20: "k", 21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q", 27: "r", 28: "s", 29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y", 35: "z", 36: ":", 37: ";", 38: "<", 39: "=", 40: ">", 41: "?", 42: "@", 43: "[", 44: "]", 45: "^", 46: "_", 47: "{", 48: "|", 49: "}", 50: "A", 51: "B", 52: "C", 53: "D", 54: "E", 55: "F", 56: "G", 57: "H", 58: "I", 59: "J", 60: "K", 61: "L", 62: "M", 63: "N", 64: "O", 65: "P", 66: "Q", 67: "R", 68: "S", 69: "T", 70: "U", 71: "V", 72: "W", 73: "X", 74: "Y", 75: "Z"}

// func main() {
//     fmt.Println(DecimalToAny(9999, 76))
//     fmt.Println(AnyToDecimal("1F[", 76))
// }

// 10进制转任意进制
func DecimalToAny(num, n int) string {
    // 参数校验
    if n < 2 {
        // 没有1进制，否则会一直循环；也没有0进制，被除数为0会异常。
        return strconv.Itoa(num)
    }
    if 0 == num {
        //fmt.Printf(" string(num) , %T, %v, %s\n", string(num),string(num),strconv.Itoa(num))
        return strconv.Itoa(num) // 不能用string(num)，string(num) == ""
    }
    if num < 0 {
        num = -1 * num // 强制转成正数处理
    }

    new_num_str := ""
    var remainder int
    var remainder_string string
    for num != 0 {
        remainder = num % n
        if 76 > remainder && remainder > 9 {
            remainder_string = tenToAny[remainder]
        } else {
            remainder_string = strconv.Itoa(remainder)
        }
        new_num_str = remainder_string + new_num_str
        num = num / n
    }
    return new_num_str
}

func D十进制转换(num int, n int) string {
    // 参数校验
    if n < 2 {
        // 没有1进制，否则会一直循环；也没有0进制，被除数为0会异常。
        return strconv.Itoa(num)
    }
    if 0 == num {
        //fmt.Printf(" string(num) , %T, %v, %s\n", string(num),string(num),strconv.Itoa(num))
        return strconv.Itoa(num) // 不能用string(num)，string(num) == ""
    }
    if num < 0 {
        num = -1 * num // 强制转成正数处理
    }

    new_num_str := ""
    var remainder int
    var remainder_string string
    for num != 0 {
        remainder = num % n
        if 76 > remainder && remainder > 9 {
            remainder_string = tenToAny[remainder]
        } else {
            remainder_string = strconv.Itoa(remainder) // int转字符串
        }
        new_num_str = remainder_string + new_num_str
        num = num / n
    }
    return new_num_str
}

// 采用递归方法进行实现，十进制转任意进制
func DecimalToAnyDigui(new_num_str *string, num int, n int) {
    // 参数校验
    if n < 2 {
        // 没有1进制，否则会一直循环；也没有0进制，被除数为0会异常。
        return
    }
    if num < 0 {
        num = -1 * num // 强制转成正数处理
    }

    y := num / n           // 商
    remainder_string := "" // 10进制以上数值，需要寻找对应的单字母表示，例如10在16进制中就是a字母
    remainder := num % n   // 余数

    if 76 > remainder && remainder > 9 {
        remainder_string = tenToAny[remainder]
    } else {
        remainder_string = strconv.Itoa(remainder)
    }
    *new_num_str = remainder_string + *new_num_str

    if 0 != y {
        // 商为0，终止递归
        DecimalToAnyDigui(new_num_str, y, n)
    }

    return
}

// map根据value找key
func findkey(in string) int {
    result := -1
    for k, v := range tenToAny {
        if in == v {
            result = k
        }
    }
    return result
}

// 任意进制转10进制
func AnyToDecimal(num string, n int) int {
    new_num := 0.0
    nNum := len(strings.Split(num, "")) - 1
    for _, value := range strings.Split(num, "") {
        tmp := float64(findkey(value))
        if tmp != -1 {
            new_num = new_num + tmp*math.Pow(float64(n), float64(nNum))
            nNum = nNum - 1
        } else {
            break
        }
    }
    return int(new_num)
}
