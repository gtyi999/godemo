package publish

import (
    "github.com/gmallard/stompngo"
    "fmt"
    "time"
    "net"
    "os"
    "runtime"
    "log"
)


func Run(Cpunums int,Logname,Hoststomp,Portstomp,Queuename,Verstomp string,Trans bool,Persistent bool,Msgcnt int64,Threadnum int) {

    runtime.GOMAXPROCS(Cpunums)
    var baseconn net.Conn
    var stompconn *stompngo.Connection
    var errinfo error
    var ch stompngo.Headers

    logfile,err:=os.OpenFile(Logname,os.O_RDWR|os.O_CREATE,0666)
    if err!=nil{
        fmt.Printf("%s\r\n",err.Error())
        os.Exit(-1)
    }
    defer logfile.Close()
    logger:=log.New(logfile,"\r\n",log.Ldate|log.Ltime|log.Llongfile)
    fmt.Println("start...")

    //创建原始套接字
    hap := net.JoinHostPort(Hoststomp,Portstomp)
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
    ch.Add("accept-version",Verstomp)
    ch.Add("host",Hoststomp)
    ch.Add("heart-beat","0,0")

    stompconn,errinfo = stompngo.Connect(baseconn,ch)
    if errinfo != nil {
        fmt.Println("stompngo.Connect failed:",errinfo.Error())
    } else {
        fmt.Println("stompngo connect ok")
    }

    id := stompngo.Uuid()
    sh := stompngo.Headers{"destination", Queuename, "ack", "client", "id", id}
    if Persistent {
        sh = sh.Add("persistent", "true")
    }

    //分成几个线程
    fmt.Println("发起请求数：",Msgcnt)
    mychan := make(chan bool,Threadnum)

    reqperthread := Msgcnt / int64(Threadnum)

    timestamp := time.Now().Unix()
    tm := time.Unix(timestamp, 0)
    logger.Println("开始时间：",tm.Format("2006-01-02 03:04:05 PM"))

    for index:=0; index< Threadnum; index++ {
        go func(index int) {


            //添加事务 如果有事务，那么在发送成功后一定要commit
            //if Trans {
            //    sh = sh.Add("transaction","tx1")
            //    stompconn.Begin(sh)
            //    //fmt.Println("trans begin...:",sh)
            //} else {
            //    //sh.Delete("transaction")
            //}


            for i:=0; int64(i)<reqperthread; i++ {
                strmsg := "消息内容"


                //如果需要ACK 必须"message-id"有值
                sh.Add("message-id",strmsg)



                e := stompconn.Send(sh,strmsg)
                if e != nil {
                    fmt.Println(e.Error())
                } else {
                    //fmt.Println("send ok")
                }


            }
            //如果有事务 那么需要commit
            //if Trans {
            //    stompconn.Commit(sh)
            //    //fmt.Println("trans commit...:",sh)
            //}

            mychan <- true
        } (index)
    }

    for i:=0; i<Threadnum; i++ {
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
