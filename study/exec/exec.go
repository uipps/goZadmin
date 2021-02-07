// 执行外部命令
// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\study\exec\exec.go

package main

import (
    "bufio"
    "fmt"
    "io"
    "os/exec"
    "runtime"
)

func main() {
    command_str := ""

    // exec PHP
    num1 := "1001"
    num2 := "9876"
    command_str = fmt.Sprintf("echo bcmul('%s', '%s');", num1, num2)
    fmt.Println(command_str)
    execPhp(command_str)

    command_str = "dir"
    sysType := runtime.GOOS
    if ("windows" == sysType) {
        execWin01(command_str)
        rlt := CombinedExecWin(command_str)
        fmt.Println(rlt)
        StdoutPipeWin(command_str)
    } else {
        exec01(command_str)
        rlt := CombinedExec(command_str)
        fmt.Println(rlt)
        //StdoutPipe(command_str)
    }

    //
    LookPath("curl")
}

// exec PHP code, unix和windows均可
func execPhp(cmd_str string) {
    cmd := exec.Command("php", "-r", cmd_str) // /usr/bin/php -- unix ; windows -- php
    bytes, err := cmd.Output()
    if err != nil {
        fmt.Println(err)
    }
    resp := string(bytes)
    fmt.Println(resp)
}

// windows下报错：2021/02/07 06:12:59 exec: "/bin/bash": file does not exist
//  windows用下面win版本
func exec01(cmd_str string) {
    cmd := exec.Command("/bin/bash", "-c", cmd_str)
    //c := exec.Command("cmd", "/C", cmd_str)   // 此处是windows版本
    bytes, err := cmd.Output()
    if err != nil {
        fmt.Println(err)
    }
    resp := string(bytes)
    fmt.Println(resp)
}
// windows版本
func execWin01(cmd_str string) {
    // 此处是windows版本
    cmd := exec.Command("cmd", "/C", cmd_str)
    bytes, err := cmd.Output()
    if err != nil {
        fmt.Println(err)
    }
    resp := string(bytes)
    fmt.Println(resp)
}

func CombinedExec(cmd_str string) string {
    c := exec.Command("bash", "-c", cmd_str)
    output, err := c.CombinedOutput()
    if err != nil {
        fmt.Println(err)
    }
    return string(output)
}
// shell标准输出的逐行实时进行处理
func StdoutPipeWin(cmd_str string) bool {
    cmd := exec.Command("cmd", "/C", cmd_str)

    //显示运行的命令
    fmt.Println(cmd.Args)
    //StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        fmt.Println(err)
        return false
    }

    cmd.Start()
    //创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
    reader := bufio.NewReader(stdout)

    //实时循环读取输出流中的一行内容
    for {
        line, err2 := reader.ReadString('\n')
        if err2 != nil || io.EOF == err2 {
            break
        }
        fmt.Println(line)
    }

    //阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
    cmd.Wait()
    return true
}
func CombinedExecWin(cmd_str string) string {
    c := exec.Command("cmd", "/C", cmd_str)
    output, err := c.CombinedOutput()
    if err != nil {
        fmt.Println(err)
    }
    return string(output)
}

// unix和windows下均可
func LookPath(str1 string) {
    f, err := exec.LookPath("curl")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(f) //  /bin/ls
}
// 非阻塞方式(不需要执行结果) , 适用于调用自己写的程序(服务器死循环，且不需要返回结果的), 后面加&符号的
//  不需要执行命令的结果与成功与否，执行命令马上就返回
/*func exec_shell_no_result(command string) {
    //处理启动参数，通过空格分离 如：setsid /home/luojing/gotest/src/test_main/iwatch/test/while_little &
    command_name_and_args := strings.FieldsFunc(command, splite_command)
    cmd := exec.Command(commandName, params...)
    //开始执行c包含的命令，但并不会等待该命令完成即返回
    cmd.Start()
    if err != nil {
        fmt.Printf("%v: exec command:%v error:%v\n", get_time(), command, err)
    }
    fmt.Printf("Waiting for command:%v to finish...\n", command)
    //阻塞等待fork出的子进程执行的结果，和cmd.Start()配合使用[不等待回收资源，会导致fork出执行shell命令的子进程变为僵尸进程]
    err = cmd.Wait()
    if err != nil {
        fmt.Printf("%v: Command finished with error: %v\n", get_time(), err)
    }
    return
}*/