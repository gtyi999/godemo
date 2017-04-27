//activemq发布消息示例代码
package main

import (
    "fmt"
    "net"
    "github.com/gmallard/stompngo"
    "strconv"
    "time"
    "os"
    "log"
    "runtime"
)

func main() {

    runtime.GOMAXPROCS(runtime.NumCPU() * 2)
    var baseconn net.Conn
    var stompconn *stompngo.Connection
    var errinfo error
    var ch stompngo.Headers

    logfile,err:=os.OpenFile("test.log",os.O_RDWR|os.O_CREATE,0666)
    if err!=nil{
        fmt.Printf("%s\r\n",err.Error())
        os.Exit(-1)
    }
    defer logfile.Close()
    logger:=log.New(logfile,"\r\n",log.Ldate|log.Ltime|log.Llongfile)
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

    //分成几个线程
    var msgCnt int64 = 1000000
    fmt.Println(msgCnt)
    var threadCnt int = 10
    mychan := make(chan bool,threadCnt)

    reqperthread := msgCnt / int64(threadCnt)

    timestamp := time.Now().Unix()
    tm := time.Unix(timestamp, 0)
    logger.Println("开始时间：",tm.Format("2006-01-02 03:04:05 PM"))

    for index:=0; index< threadCnt; index++ {
        go func(index int) {
            for i:=0; int64(i)<reqperthread; i++ {
                strmsg := "消息内容" + strconv.Itoa(i)
                //如果需要ACK 必须"message-id"有值
                sh.Add("message-id",strmsg)
                e := stompconn.Send(sh,strmsg)
                if e != nil {
                    fmt.Println(e.Error())
                } else {
                    //fmt.Println("send ok")
                }
            }
            mychan <- true
        } (index)
    }

    for i:=0; i<threadCnt; i++ {
        <-mychan
    }
    timestamp = time.Now().Unix()
    tm = time.Unix(timestamp, 0)
    logger.Println("结束时间：",tm.Format("2006-01-02 03:04:05 PM"))
    fmt.Println("发送完毕。。。。")


    errinfo = stompconn.Disconnect(stompngo.Headers{})
    if errinfo != nil {
        fmt.Println("base connection err:",errinfo.Error())
    }

    errinfo = baseconn.Close()
    if errinfo != nil {
        fmt.Println("base connect close err:",errinfo.Error())
    }

}
