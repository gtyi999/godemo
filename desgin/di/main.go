package main

import "fmt"

type myhandler func(a int, b string)

func MethodA(name string,  handler interface{}) {
    fmt.Println("Enter MethodA:", name)
    // error handler.(myhandler)
    handler.(func(a int, b string))(3030,"zdd")// 给f注入参数
    fmt.Println("Exit MethodA:", name)
}

func MethodB(a int, b string) {
    fmt.Println(a, b)
}

func main() {
    d := MethodB
    MethodA("zddhub", d)
}