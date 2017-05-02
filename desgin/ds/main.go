package main

import (
    "fmt"
    "reflect"
)

type pp func() (interface{},error)

func step1() (interface{}, error){
    return "step1 call",nil
}

func step2() (interface{}, error){
    return "step1 call",nil
}

func step3() (interface{}, error){
    return "step1 call",nil
}
func step4() (interface{}, error){
    return "step1 call",nil
}

func main() {
    myhandler := make(map[string]interface{},10)


    myhandler["step1"] = step1

    myhandler["step2"] = step2

    myhandler["step3"] = step3

    myhandler["step4"] = step4

    for key,val := range myhandler {
        //fmt.Println("call : ",key)

        fmt.Println("key: ",key,"type: ",reflect.TypeOf(val))
        ret :=val.(func() (interface {}, error))
        strout,e:=ret()
        if e!=nil {
            fmt.Println(e.Error())
        } else {
            fmt.Println(strout)
        }
    }
}
