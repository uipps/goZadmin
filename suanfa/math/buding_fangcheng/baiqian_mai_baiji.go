/*
10.不定方程问题(百钱买白鸡)

公鸡5文钱1只，母鸡3文钱1只，小鸡3只1文钱，要求用100文钱买100只鸡，求公鸡，母鸡和小鸡各应该买多少只？
x+y+z=100;
5x+3y+z/3=100

参考： https://www.cnblogs.com/heisaijuzhen/articles/4324474.html


// go run buding_fangcheng.go
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\math\buding_fangcheng\baiqian_mai_baiji.go
// go run ~/develope/go/go_code/src/github.com/uipps/goZadmin/suanfa/math/buding_fangcheng/baiqian_mai_baiji.go


-- 几种可能结果
公鸡:0,母鸡：25,小鸡:75
公鸡:4,母鸡：18,小鸡:78
公鸡:8,母鸡：11,小鸡:81
公鸡:12,母鸡：4,小鸡:84

*/
package main

import "fmt"

func main() {
	baiqianMaiBaiJi()
}

func baiqianMaiBaiJi() {
    for x := 0; x <= 20; x++ {
        for y := 0; y <= 33; y++ {
            z := 100 - x - y;
            if (z % 3 == 0 && x * 5 + y * 3 + z / 3 == 100) {
                fmt.Printf("公鸡:%d, 母鸡：%d, 小鸡:%d\n", x, y, z);
            }
        }
    }
    return
}
