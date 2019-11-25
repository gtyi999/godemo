package main

import (
    "sppush/mypkg"
    "fmt"
    "reflect"
)

type MsgInfo struct{
    Method string
    Params map[string]string
    MsgId int64
}

func main() {


    home := &mypkg.Home{}

    msg := MsgInfo{"Home.AddPeople",map[string]string{"Name":"lh"},1}
    myparam := make([]reflect.Value,1)
    myparam[0]= reflect.ValueOf(msg.Params["Name"])

    v := reflect.ValueOf(home.AddPeople)
    ret := v.Call(myparam)
    fmt.Println(ret[0])
}
