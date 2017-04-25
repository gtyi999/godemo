//activemq接收消息示例代码
package main

import (
    "fmt"
    "net"
    "github.com/gmallard/stompngo"
)

func main() {

    var baseconn net.Conn
    var stompconn *stompngo.Connection
    var errinfo error
    var ch stompngo.Headers
    var msgdata stompngo.MessageData

    fmt.Println("start...")

    //创建原始套接字
    hap := net.JoinHostPort("192.168.10.166","61613")
    baseconn, errinfo = net.Dial("tcp", hap)
    if errinfo != nil {
        fmt.Println("errinfo:",errinfo.Error())
    } else {
        fmt.Println("net.Dial ok")
    }
    //创建连接头
    ch = stompngo.Headers{}
    ch.Add("login","guest") //账号
    ch.Add("passcode","guest") //密码
    ch.Add("accept-version",stompngo.SPL_10)
    ch.Add("host","192.168.10.166")
    ch.Add("heart-beat","0,0")

    stompconn,errinfo = stompngo.Connect(baseconn,ch)
    if errinfo != nil {
        fmt.Println("stompngo.Connect failed:",errinfo.Error())
    } else {
        fmt.Println("stompngo connect ok")
    }


    sh := stompngo.Headers{"destination", "/queue/shaoyong-test", "ack", "auto"}
    //sh = sh.Add("id", id) 1.1 1.2版本需要，1.0版本不需要

    msgchan, e := stompconn.Subscribe(sh)
    if e != nil {
        fmt.Println("订阅操作失败")
    }

    for {
        select {
            case msgdata= <-msgchan:
            case msgdata= <-stompconn.MessageData:
                fmt.Println("conn.Session:",stompconn.Session())
        }
        if msgdata.Error != nil {
            fmt.Println("conn.Session:",stompconn.Session()," err:",msgdata.Error.Error())
        } else {
            wh := msgdata.Message.Headers
            for j := 0; j < len(wh)-1; j += 2 {
                fmt.Printf("%stag:%s connsess:%s Header:%s:%s\n", stompconn.Session(), wh[j], wh[j+1])
            }
            //最大不能超过1024长度
            fmt.Println("msgdata:",string(msgdata.Message.Body))
        }
    }
    // stomp 1.0版本
    sbh := stompngo.Headers{}
    id := stompngo.Uuid()
    sbh.Add("id",id)
    sbh.Add("destination", "/queue/shaoyong-test")
    e = stompconn.Unsubscribe(sbh)
    if e != nil {
        fmt.Println("err:",e.Error())
    }
    e = stompconn.Disconnect(stompngo.Headers{})
    if e != nil {
        fmt.Println("err:",e.Error())
    }
    e = baseconn.Close()
    if e!=nil {
        fmt.Println("err:",e.Error())
    }


}

