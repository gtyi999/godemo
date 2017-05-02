package main

import (
    "fmt"
    //"time"
    "github.com/nsqio/go-nsq"
    "os"
    "time"
    "log"
    "sync"
    "runtime"
)
var nsqserver string = "10.10.18.156:4150"
var allcnt int64 //总数
var mutex sync.Mutex

// nsq发布消息
func Producer(docnt int64, Logname string,mychan chan bool) {
    p, err := nsq.NewProducer(nsqserver, nsq.NewConfig())                // 新建生产者
    if err != nil {
        panic(err)
    }
    logfile,err:=os.OpenFile(Logname,os.O_RDWR|os.O_CREATE,0666)
    if err!=nil{
        fmt.Printf("%s\r\n",err.Error())
        os.Exit(-1)
    }
    defer logfile.Close()
    logger:=log.New(logfile,"\r\n",log.Ldate|log.Ltime|log.Llongfile)
    fmt.Println("start...")
    timestamp := time.Now().Unix()
    tm := time.Unix(timestamp, 0)
    logger.Println("开始时间：",tm.Format("2006-01-02 03:04:05 PM"))
    for i:=0; int64(i)<docnt; i++ {
        if err := p.Publish("test", []byte("A")); err != nil {           // 发布消息
            panic(err)

        }
        //time.Sleep(time.Second*2)
    }
    mychan<-true
    timestamp = time.Now().Unix()
    tm = time.Unix(timestamp, 0)
    logger.Println("结束时间：",tm.Format("2006-01-02 03:04:05 PM"))
}

// nsq订阅消息
type ConsumerT struct{
    Info *log.Logger
    Cnt int64
    Docnt int64
    Dochan chan bool
}

func (this *ConsumerT) HandleMessage(msg *nsq.Message) error {
    //fmt.Println("handler message...")
    //fmt.Println(string(msg.Body))

    //fmt.Println("pid",os.Getpid()," now:",this.Docnt)
    //time.Sleep(time.Second*10)

    //
    //if allcnt<=0 {
    //    fmt.Println("完成数: ",this.Docnt, "allcnt:",allcnt)
    //    msg.Finish()
    //    this.Dochan<-true
    //    return nil
    //}
    this.Docnt++
    msg.Finish()
    //time.Sleep(time.Second*1)
    runtime.Gosched() //一定要加，Finish是异步发送，加这句是为了防止FIN还没发送就关闭



    mutex.Lock()
    allcnt--
    mutex.Unlock()
    if allcnt<=0 {
        fmt.Println("当前线程完成数: ",this.Docnt, "allcnt:",allcnt)
        this.Dochan<-true
    }

    fmt.Println("ack...")
    return nil
}

func Consumer(docnt int64, Logname string,mychan chan bool) {
    c, err := nsq.NewConsumer("test", "test-channel", nsq.NewConfig())   // 新建一个消费者
    if err != nil {
        panic(err)
    }
    logfile,err:=os.OpenFile(Logname,os.O_RDWR|os.O_CREATE,0666)
    if err!=nil{
        fmt.Printf("%s\r\n",err.Error())
        os.Exit(-1)
    }
    defer logfile.Close()
    logger:=log.New(logfile,"\r\n",log.Ldate|log.Ltime|log.Llongfile)

    c.AddHandler(&ConsumerT{logger,docnt,0,mychan})                                           // 添加消息处理
    if err := c.ConnectToNSQD(nsqserver); err != nil {            // 建立连接
        panic(err)
    } else {
        fmt.Println("consumer connect success...")
    }
}
// 主函数
func main() {
    var threadcnt int = 1
    var msgcnt int64 = 1
    var mychan chan bool = make(chan bool,threadcnt)

    //发送代码段
    //timestamp := time.Now().Unix()
    //tm := time.Unix(timestamp, 0)
    //fmt.Println("开始时间：",tm.Format("2006-01-02 03:04:05 PM"))
    //for i:=0; i<threadcnt; i++ {
    //    go func(docnt int64,mychan chan bool) {
    //        Producer(docnt,"publis.log",mychan)
    //    }(msgcnt/int64(threadcnt),mychan)
    //}
    //for i:=0; i<threadcnt; i++ {
    //    <-mychan
    //}
    //fmt.Println(threadcnt," 个线程执行完毕")
    //timestamp = time.Now().Unix()
    //tm = time.Unix(timestamp, 0)
    //fmt.Println("结束时间：",tm.Format("2006-01-02 03:04:05 PM"))





 //====================================================================================


    //消费代码段
    timestamp := time.Now().Unix()
    tm := time.Unix(timestamp, 0)
    fmt.Println("开始时间：",tm.Format("2006-01-02 03:04:05 PM"))
    allcnt = msgcnt
    for i:=0; i<threadcnt; i++ {
        go func(docnt int64,mychan chan bool) {
            Consumer(docnt, "subscribe.log",mychan)
        }(msgcnt / int64(threadcnt),mychan)
    }
    <-mychan
    fmt.Println(threadcnt," 个线程执行完毕")
    timestamp = time.Now().Unix()
    tm = time.Unix(timestamp, 0)
    fmt.Println("结束时间：",tm.Format("2006-01-02 03:04:05 PM"))




}
