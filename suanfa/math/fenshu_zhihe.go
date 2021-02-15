/**
题目：求这样的四个自然数p,q,r,s(p<=q<=r<=s)，使得以下等式成立：
1/p+1/q+1/r+1/s=1

参考： https://blog.csdn.net/csy981848153/article/details/7625938

// 分析一下：
  1) p<=q<=r<=s , 所以1/p>=1/q>=1/r>=1/s , 得：1/p+1/q+1/r+1/s <= 4/p ,  所以 4/p>=1 ，因此p<=4，并且p不能为1，所以p>=2且p<=4
  2) p<=q<=r<=s , 所以1/p>=1/q>=1/r>=1/s , 得：1/p+1/q+1/r+1/s >= 4/s ,  所以 4/s<=1 ，因此s>=4  (这项好像没多大用处)

  3) 因p>=2且p<=4，所以 1/2 <= 1/q+1/r+1/s <= 3/4,
     因1/p>=1/q>=1/r>=1/s，所以 1/q+1/r+1/s <= 3/q,  结合上式，得：3/q >= 1/2，因此q<=6，并且q<=p, q不能为2，所以q最小值是3.
    可以推导q的范围为： 3到6； 3<=q<=6

  4) 同理, 所以 1/2 - 1/3 <= 1/r+1/s <= 3/4 - 1/6  ===> 1/6 <= 1/r+1/s <= 7/12
     1/r+1/s <= 2/r , 所以 2/r >= 1/6, 因此r<=12, r最大也是4，
     可知r的范围4~12， 4<=r<=12

  5) 1/6 - 1/4 <= 1/s <= 7/12 - 1/12  ===> -1/12 <= 1/s <= 1/2 , 这么推导，只能得出 s>=2, 不能得出s的上限。
     但是我们知道，当p，q，r都取最小的时候(1/x才最大)，s才能最大.
     p,q,r都是正整数，p的最小值是2，p最小3，则r最小比6大(1/2+1/3+1/6 = 1)，因此r最小是7；所以s最大为 1/s = 1-1/2-1/3-1/7=1/42
     可知s的范围 2<=s<=42

用程序计算一下：

go run math/fenshu_zhihe.go
go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\math\fenshu_zhihe.go

-- 所需范围4,6,12,1000
p: 2 q: 3 r: 7 s: 42
p: 2 q: 3 r: 8 s: 24
p: 2 q: 3 r: 9 s: 18
p: 2 q: 3 r: 10 s: 15
p: 2 q: 3 r: 12 s: 12
p: 2 q: 4 r: 5 s: 20
p: 2 q: 4 r: 6 s: 12
p: 2 q: 4 r: 8 s: 8
p: 2 q: 5 r: 5 s: 10
p: 2 q: 6 r: 6 s: 6
p: 3 q: 3 r: 4 s: 12
p: 3 q: 3 r: 6 s: 6
p: 3 q: 4 r: 4 s: 6
p: 4 q: 4 r: 4 s: 4

*/

package main

import "fmt"

func main() {
	fenshu_zihe()
}

func fenshu_zihe() {
	count :=0

	for  p := 2; p <= 4; p++ {
		for  q := p; q <= 6; q++ {
			for r := q; r <= 12; r++ {
				for s := r; s <= 1000; s++ {
					count++
					if (q * r * s + p * r * s + p * q * s + p * q * r == p * q * r * s) {
						fmt.Printf("p: %d q: %d r: %d s: %d\n", p,q,r,s)
					}
				}
			}
		}
	}
	fmt.Printf("counter: %d", count)
	return
}
