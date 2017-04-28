//activemq接收消息示例代码
package main

import (
    "fmt"
    "net"
    "github.com/gmallard/stompngo"
    "os"
    "log"
    "time"
)
var perthread int64 = 1000000/10
var mychan chan bool
func main() {
    mychan = make(chan bool,10)
    for i:=0; i<10; i++ {
        go test()
    }
    for i:=0; i<10; i++ {
        <-mychan
    }
    fmt.Println("完成")
}

func test() {

    var baseconn net.Conn
    var stompconn *stompngo.Connection
    var errinfo error
    var ch stompngo.Headers
    var msgdata stompngo.MessageData

    logfile,err:=os.OpenFile("test.log",os.O_RDWR|os.O_CREATE,0666)
    if err!=nil{
        fmt.Printf("%s\r\n",err.Error())
        os.Exit(-1)
    }
    defer logfile.Close()
    logger:=log.New(logfile,"\r\n",log.Ldate|log.Ltime|log.Llongfile)

    fmt.Println("start...")

    //创建原始套接字
    hap := net.JoinHostPort("192.168.100.123","61613")
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
    ch.Add("host","192.168.100.123")
    ch.Add("heart-beat","0,0")
    stompconn,errinfo = stompngo.Connect(baseconn,ch)
    if errinfo != nil {
        fmt.Println("stompngo.Connect failed:",errinfo.Error())
    } else {
        fmt.Println("stompngo connect ok")
    }
    sh := stompngo.Headers{"destination", "/queue/liuhui_test", "ack", "client"}
    //sh = sh.Add("id", id) 1.1 1.2版本需要，1.0版本不需要
    msgchan, e := stompconn.Subscribe(sh)
    if e != nil {
        fmt.Println("订阅操作失败")
    }
    timestamp := time.Now().Unix()
    tm := time.Unix(timestamp, 0)
    logger.Println("开始时间：",tm.Format("2006-01-02 03:04:05 PM"))

    icnt :=0
    for i:=0; int64(i)<perthread; i++{
        select {
        case msgdata= <-msgchan:
        case msgdata= <-stompconn.MessageData:
            //fmt.Println("conn.Session:",stompconn.Session())
        }
        if msgdata.Error != nil {
            //fmt.Println("conn.Session:",stompconn.Session()," err:",msgdata.Error.Error())
        } else {
            wh := msgdata.Message.Headers
            ah := stompngo.Headers{}
            ah = ah.Add("message-id", wh.Value("message-id"))
            if cv, ok := wh.Contains(stompngo.HK_RECEIPT); ok {
                ah = ah.Add(stompngo.HK_RECEIPT, cv)
            }
            e := stompconn.Ack(ah)
            if e != nil {
                fmt.Println("err info:",e.Error())
            } else {
                //fmt.Println("ack success")
                icnt++
            }
            //time.Sleep(10*time.Second)
        }
    }
    timestamp = time.Now().Unix()
    tm = time.Unix(timestamp, 0)
    fmt.Println("结束时间：",tm.Format("2006-01-02 03:04:05 PM"),"总笔数:",icnt)
    logger.Println("结束时间：",tm.Format("2006-01-02 03:04:05 PM"))





    //e = stompconn.Unsubscribe(sbh)
    //if e != nil {
    //    fmt.Println("err:",e.Error())
    //}
    e = stompconn.Disconnect(stompngo.Headers{})
    if e != nil {
        fmt.Println("err:",e.Error())
    }
    e = baseconn.Close()
    if e!=nil {
        fmt.Println("err:",e.Error())
    }

    mychan<-true
}

