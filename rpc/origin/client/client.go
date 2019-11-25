package main

import (
    "net/rpc"
    "github.com/prometheus/common/log"

    "fmt"
)


type Args struct {
    A,B int
}

type Quotient struct {
    Quo, Rem int
}

func main() {
    //client, err := rpc.DialHTTP("tcp","127.0.0.1:1234")
    client, err := rpc.Dial("tcp","127.0.0.1:4321")
    if err!=nil {
        log.Fatal("dial failed")
    }
    args := &Args{10,5}
    var reply int
    err = client.Call("Arith.Multiply",args,&reply)
    if err!=nil {
        log.Fatal("Multiply called failed")
    }
    fmt.Println(reply)

    //采用异步
    q := new(Quotient)
    divCall := client.Go("Arith.Divide",args,q,nil)
    replyCall := <- divCall.Done
    fmt.Println("异步返回：",replyCall.Reply)



}
