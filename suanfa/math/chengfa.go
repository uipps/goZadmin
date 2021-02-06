// 计算乘法，n*m, 任意大数或小数，任意小数点位
//
// go run chengfa.go -n 123 -m 987 -l 100
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\math\chengfa.go -n 201 -m 134
// go run ~/develope/go/go_code/src/github.com/uipps/goZadmin/suanfa/math/chengfa.go -n 201 -m 134

// 2个非常大的数字，310位，超出了float64范围
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\math\chengfa.go -n 8888888888999999999977777777776666666666555555555588888888889999999999777777777766666666665555555555888888888899999999997777777777666666666655555555558888888888999999999977777777776666666666555555555588888888889999999999777777777766666666665555555555888888888899999999997777777777666666666655555555559999999999 -m 8888888888999999999977777777776666666666555555555588888888889999999999777777777766666666665555555555888888888899999999997777777777666666666655555555558888888888999999999977777777776666666666555555555588888888889999999999777777777766666666665555555555888888888899999999997777777777666666666655555555559999999999

package main

import (
    "flag"
    "fmt"
    "github.com/syyongx/php2go"
    "github.com/uipps/goZadmin/suanfa/common"
    "strconv"
    "strings"
    "time"
)

var (
    chengshu1 string // 乘数1
    chengshu2 string // 乘数2
)

func init() {
    flag.StringVar(&chengshu1, "n", "123456", "Usage: 1")
    flag.StringVar(&chengshu2, "m", "87654321", "Usage: 7 13")
}

func forexample01() {
    var a, b, c float64
    a = 1.69
    b = 1.7
    c = a * b
    fmt.Println(c) // 2.8729999999999998

    a = 1.69 * 100
    b = 1.7 * 10
    c = a * b / (100 * 10)

    fmt.Println(c) // 2.873， 正确结果
}

func main() {
    flag.Parse()

    // 参数校验，是否数字(包括小数), 超过了float64范围的数字，直接退出了
    if (!common.Is_numeric(chengshu1) || !common.Is_numeric(chengshu2)) {
       fmt.Println("乘数必须都是数字!")
       return
    }

    startTime := time.Now().UnixNano()
    fmt.Printf("startTime：%d, %s\n", startTime/1e3, time.Unix(0, startTime).Format("2006-01-02 15:04:05"))

    fmt.Printf("\n  计算 %s x %s 的值，第一种方法(任意位，负数或小数均可)，结果为：\n", chengshu1, chengshu2)
    rlt := chengFa01(chengshu1, chengshu2)
    fmt.Println(rlt)
    fmt.Println("\n")

    // 如果数字太长，考虑格式化输出
    //common.OutPrintFmt(k_arr1, xiaoshuLeng, 1, fenzi_orig, fenmu)
    fmt.Printf("\n  计算 %s x %s 的值，第二种方法(不适用小数和负数)，结果为：\n", chengshu1, chengshu2)
    rlt = chengFa02(chengshu1, chengshu2)
    fmt.Println(rlt)
    fmt.Println("\n")

    fmt.Printf("\n  使用PHP的bcmul方法，结果为：\n")
    command_str := fmt.Sprintf("echo bcmul('%s', '%s');", chengshu1, chengshu2)
    rlt = common.ExecPHP(command_str)
    //fmt.Printf("%T , '%s' , \n", rlt, rlt)
    fmt.Println(rlt)
    fmt.Println("\n")

    fmt.Printf("\n  计算 %s x %s 的值，直接转float64，计算结果为：\n", chengshu1, chengshu2)
    rlt_float64 := chengFa04(chengshu1, chengshu2)
    fmt.Println(rlt_float64)
    fmt.Printf("%f , 非科学计数法，需手动改位数.100f这样的. \n", rlt_float64)
    fmt.Println("\n")


    // 如果数字太长，考虑格式化输出
    //common.OutPrintFmt(k_arr1, xiaoshuLeng, 1, fenzi_orig, fenmu)
    //fmt.Printf("\n  计算 %s x %s 的值，第三种方法(计算错误，废弃此方法)\n", chengshu1, chengshu2)
    //rlt = chengFa03(chengshu1, chengshu2)
    //fmt.Println(rlt)
    //fmt.Println("\n")

    // 执行时间计算
    endTime := time.Now().UnixNano()
    fmt.Printf("  endTime：%d, %s\n", endTime/1e3, time.Unix(0, endTime).Format("2006-01-02 15:04:05"))
    nanoSeconds := float64(endTime-startTime) / 1e3
    fmt.Println("spendTime：", nanoSeconds)
}

// 转化为float64位进行计算
func chengFa04(chengshuA, chengshuB string) (result float64) {
    chengA, _ := strconv.ParseFloat(chengshuA, 64)
    chengB, _ := strconv.ParseFloat(chengshuB, 64)
    return  chengA * chengB
}

// 任意位数乘法, 包括小数相乘
func chengFa01(chengshuA, chengshuB string) (result string) {
    // 1. 确定正负值，顺便去掉负号
    plus_minus_1 := php2go.Strpos(chengshuA, "-", 0)
    plus_minus_2 := php2go.Strpos(chengshuB, "-", 0)
    if (-1 == plus_minus_1) {
        plus_minus_1 = 1 // 不包含，则为正数
    } else {
        // 含有-负号
        plus_minus_1 = -1
        chengshuA = strings.Replace(chengshuA, "-", "", -1) // 去掉负号
    }
    if (-1 == plus_minus_2) {
        plus_minus_2 = 1 // 不包含，则小数位的长度重置为0
    } else {
        // 含有-负号
        plus_minus_2 = -1
        chengshuB = strings.Replace(chengshuB, "-", "", -1) // 去掉负号
    }

    // 2. 确定小数位数，顺便去掉小数点
    xiaoshuLen1 := php2go.Strpos(chengshuA, ".", 0) // 返回-1表示不包含小数点
    xiaoshuLen2 := php2go.Strpos(chengshuB, ".", 0)
    if (-1 == xiaoshuLen1) {
        xiaoshuLen1 = 0 // 不包含，则小数位的长度重置为0
    } else {
        // 含有小数点，记录小数位长度，然后删除小数点
        xiaoshuLen1 = len(chengshuA) - xiaoshuLen1 - 1      // 字符串总长度减去(整数和小数点位数)，假如小数点后全为0也不影响
        chengshuA = strings.Replace(chengshuA, ".", "", -1) // 去掉小数点
    }
    if (-1 == xiaoshuLen2) {
        xiaoshuLen2 = 0 // 不包含，则小数位的长度重置为0
    } else {
        // 含有小数点，记录小数位长度，然后删除小数点
        xiaoshuLen2 = len(chengshuB) - xiaoshuLen2 - 1
        chengshuB = strings.Replace(chengshuB, ".", "", -1) // 去掉小数点
    }

    // 3. 为方便进位，需将数位反转，由原来的从高到低位 反转成 从低到高位
    chengshuA_r := common.ReverseStr(chengshuA)
    chengshuB_r := common.ReverseStr(chengshuB)

    // 4. 大数相乘，不考虑小数点
    a_length := len(chengshuA_r)
    b_length := len(chengshuB_r)
    c := make([]byte, a_length+b_length) // 存放结果的数组，初始每位都是0
    for j := 0; j < b_length; j++ {
        for i := 0; i < a_length; i++ {
            if (0 == chengshuB_r[j]-'0' || 0 == chengshuA_r[i]-'0') {
                continue
            }
            ji := (chengshuB_r[j] - '0') * (chengshuA_r[i] - '0') // 可能进位了
            //fmt.Printf("%d, %c, %T | %d, %c, %T, chengji: %d\n", chengshuB_r[j], chengshuB_r[j], chengshuB_r[j], chengshuA_r[i], chengshuA_r[i], chengshuA_r[i], ji)

            c[i+j] = c[i+j] + ji%10 // 这里是10进制
            c[i+j+1] = c[i+j+1] + ji/10
        }
    }

    // 5. 继续进位处理，对每项进行进位处理，plus保存上一次的进位数目, 顺便对结果加小数点
    xiaoshu_l := xiaoshuLen1 + xiaoshuLen2
    var plus byte = 0
    for i := 0; i < len(c); i++ {
        temp := c[i] + plus
        plus = 0
        if temp > 9 {
            plus = temp / 10
            result += string(temp - plus*10 + '0')
        } else {
            result += string(temp + '0')
        }
        if (xiaoshu_l > 0 && i == xiaoshu_l-1) {
            result += "." // 添加小数点
        }
    }
    // 6. 删除前导0，这里是先删除右侧0，因为还是反转状态
    result = php2go.Rtrim(result, "0")
    result = common.ReverseStr(result) // 将字符串反转

    // 7. 第一位如果是小数点，则首位补0
    xiaoshu_l = php2go.Strpos(result, ".", 0)
    if (0 == xiaoshu_l) {
        result = "0" + result // 小数点前面补0
    }

    // 8. 加符号：+-号
    if (plus_minus_1*plus_minus_2 < 0) {
        result = "-" + result // 加-负号
    }
    return result
}

// 此方法对于 999 * 999 就不准了 , 见 https://github.com/mnhkahn/go_code/blob/master/largenumberx.go
// 于是在原有基础上做了一些修改，使之正确
// 相乘的过程中，需要将字符和整数进行转换，通过a[i] - '0'和temp + '0'就能实现。
// 进位在最后一并进行。通过变量plus保存上一次的进位数目。
func chengFa02(a, b string) (reuslt string) {
    a = common.ReverseStr(a)
    b = common.ReverseStr(b)
    c := make([]byte, len(a)+len(b))

    for i := 0; i < len(a); i++ {
        for j := 0; j < len(b); j++ {
            ji := (a[i] - '0') * (b[j] - '0')
            c[i+j] += ji % 10
            c[i+j+1] += ji / 10 // 将进位累加
        }
    }
    //fmt.Println(c)

    /* // uint16换成uint8则可能超出范围导致最终结果不准
       d := make([]uint16, len(a)+len(b))
       for i := 0; i < len(a); i++ {
          for j := 0; j < len(b); j++ {
              ji := uint16(a[i] - '0') *uint16(b[j] - '0')
              d[i+j] += ji // 这里是10进制  这里做了一些修改
              //d[i+j] += ji%10
              //d[i+j+1] += ji/10   // 将进位累加
          }
       }
       fmt.Println(d)*/

    var plus byte = 0
    for i := 0; i < len(c); i++ {
        temp := c[i] + plus
        plus = 0
        if temp > 9 {
            plus = temp / 10
            reuslt += string(temp - plus*10 + '0')
        } else {
            reuslt += string(temp + '0')
        }
    }
    return common.ReverseStr(reuslt)
}

// 参考： https://www.cnblogs.com/PasserByOne/p/12019885.html
// 运行结果错误，废弃
//func chengFa03(str1, str2 string) (result string) {}
