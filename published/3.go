//activemq发布消息示例代码
package main

import (
    "fmt"
    "net"
    "github.com/gmallard/stompngo"
    "strconv"
)

func main() {

    var baseconn net.Conn
    var stompconn *stompngo.Connection
    var errinfo error
    var ch stompngo.Headers

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

    id := stompngo.Uuid()
    sh := stompngo.Headers{"destination", "/queue/shaoyong-test", "ack", "client", "id", id}
    sh = sh.Add("persistent", "true")

    for i:=1; i<1000000; i++ {
        strmsg := "消息内容" + strconv.Itoa(i)
        e := stompconn.Send(sh,strmsg)
        if e != nil {
            fmt.Println(e.Error())
        } else {
            fmt.Println("send ok")
        }
    }

    errinfo = stompconn.Disconnect(stompngo.Headers{})
    if errinfo != nil {
        fmt.Println("base connection err:",errinfo.Error())
    }

    errinfo = baseconn.Close()
    if errinfo != nil {
        fmt.Println("base connect close err:",errinfo.Error())
    }

}