// 走迷宫
/**
cd /D F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\game\migong
 go run migong01.go

-- 报错1.txt找不到
 go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\suanfa\game\migong\migong01.go

1.txt的内容：
6 5
0 1 0 0 0
0 0 0 1 0
0 1 0 0 0
1 1 1 0 0
0 1 0 0 1
0 1 0 0 0

*/
package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    //currentPath, _ := common.GetCurrentPath();fmt.Println(currentPath); // C:/Users/cf/AppData/Local/Temp/go-build308556738/b001/exe/migong01.exe
    arr := read("1.txt")
    fmt.Println("\n the arr:")
    fmt.Println(arr)
    //os.Exit(0)
    fmt.Println("\n the arr:")
    ss := walk(arr, point{0, 0}, point{len(arr) - 1, len(arr[0]) - 1})
    fmt.Println(arr)
    fmt.Println("\n the ss:")
    fmt.Println(ss)
    for _, v := range ss {
        for _, vv := range v {
            fmt.Print(vv, "    ")
        }
        fmt.Println()
    }
}

type point struct {
    i, j int
}

var dirs = [4]point{
    {-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (p point) add(r point) point {
    return point{p.i + r.i, p.j + r.j}
}

func walk(maze [][]int, start, end point) [][]int {

    //维护一个离原点多远可以到达的二维数组
    steps := make([][]int, len(maze))

    for k, _ := range steps {
        steps[k] = make([]int, len(maze[k]))
    }

    //将头放进队列
    Q := []point{start}
    for len(Q) > 0 {
        cur := Q[0]
        Q = Q[1:]

        if cur == end {
            break
        }

        //发现四个节点
        for _, dir := range dirs {
            next := cur.add(dir)
            val, ok := next.at(maze)
            //val等于1说明撞墙了
            if !ok || val == 1 {
                continue
            }
            //走另一张图
            val, ok = next.at(steps)
            if !ok || val != 0 {
                continue
            }
            //是否回到了原点，因为maze和steps图中原点的值都是零
            if next == start {
                continue
            }

            //可以走了
            i, _ := cur.at(steps)
            steps[next.i][next.j] = i + 1
            //将该点加入队列继续找
            Q = append(Q, next)

        }

    }
    return steps

}

//是否撞墙了
func (p point) at(grid [][]int) (int, bool) {
    if p.i < 0 || p.i >= len(grid) {
        return 0, false
    }
    if p.j < 0 || p.j >= len(grid[p.i]) {
        return 0, false
    }
    //获取实现迷宫的值，如果是1的时候就撞墙了
    return grid[p.i][p.j], true
}

func read(filename string) [][]int {
    f, e := os.Open(filename)
    if e != nil {
        fmt.Println(e)
        panic(e)
    }
    defer f.Close()
    buf := make([]byte, 4096)
    n, err := f.Read(buf)
    if err != nil {
        panic(err)
    }
    str := string(buf[:n])
    split := strings.Split(str, "\n")

    //获取行和列
    //var row, col int
    ss := strings.TrimSpace(split[0])
    row, err := strconv.Atoi(strings.Split(ss, " ")[0])
    if err != nil {
        panic(err)
    }
    col, err := strconv.Atoi(strings.Split(ss, " ")[1])
    //fmt.Println(col)
    if err != nil {
        panic(err)
    }

    arr := make([][]int, row)
    for k, _ := range arr {
        sArr := strings.Split(strings.TrimSpace(split[k+1]), " ")
        arr[k] = make([]int, col)
        for kk, _ := range arr[k] {
            i, err := strconv.Atoi(sArr[kk])
            if err != nil {
                panic(err)
            }
            arr[k][kk] = i
        }
    }
    return arr
}
