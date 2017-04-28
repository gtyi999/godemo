package main

import (
    "github.com/gmallard/stompngo"
    "net"
    "fmt"
    "strconv"
)

func main() {
    var stompconn *stompngo.Connection = nil
    var baseconn net.Conn
    host := "192.168.100.123"
    port := "61613"

    baseconn, e := net.Dial("tcp",net.JoinHostPort(host, port))
    if e!=nil {
        fmt.Println("error: ",e.Error())
    } else {
        fmt.Println("socket connect success")
    }
    header := stompngo.Headers{}
    header.Add("login","guest")
    header.Add("passcode","guest")
    //使用stomp 1.2协议
    header.Add("accept-version","1.2")
    header.Add("host","/")
    header.Add("heart-beat", "0,0")
    stompconn, err := stompngo.Connect(baseconn, header)
    if err != nil {
        fmt.Println("连接失败:",err)
    } else {
        fmt.Println("stompconn connect success")
    }
    //time.Sleep(time.Second*100)
    // 必须客户端响应才可以删除MQ队列数据
    f := stompngo.Headers{"destination", "/queue/liuhui_test", "ack", "client"}
    f.Add("persistent", "true")
    msgdata := "msg data "
    fmt.Println("开始发送...")
    for i:=0; i<10000; i++ {
        msgdata = msgdata + strconv.Itoa(i)
        err := stompconn.Send(f,msgdata)
        if err!=nil {
            fmt.Println("err: ",err.Error())
        } else {
            fmt.Println("send data success")
        }
    }

    err = stompconn.Disconnect(stompngo.Headers{})
    if err!=nil {
        fmt.Println("断开失败")
    }
    err = baseconn.Close()
    if err!=nil {
        fmt.Println("socket 断开失败")
    }
}
