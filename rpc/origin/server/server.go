package main

import (
    "net/rpc"
    //"net/http"
    "fmt"
    "net"
)

type Args struct {
    A,B int
}

type Quotient struct {
    Quo, Rem int
}

type Arith int

func (this *Arith)Multiply(args *Args, reply *int) error{
    *reply = args.A * args.B
    return nil
}

func (this *Arith)Divide(args *Args, quo *Quotient) error{
    if args.B == 0 {
        fmt.Println("divide by zero")
    }
    quo.Quo = args.A/ args.B
    quo.Rem = args.A% args.B
    return nil
}

func main() {
    a := new(Arith)

    ////采用HTTP暴露接口
    //rpc.Register(a)
    //rpc.HandleHTTP()
    //l,e := net.Listen("tcp",":1234")
    //if e!=nil {
    //    fmt.Println("err: ",e.Error())
    //}
    //go http.Serve(l,nil)

    //采用socket暴露接口
    rpc.Register(a)
    l,e := net.Listen("tcp",":4321")
    if e!=nil {
        fmt.Println("err: ",e.Error())
        return
    }
    rpc.Accept(l)



    select {}
}